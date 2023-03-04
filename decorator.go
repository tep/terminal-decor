// Copyright © 2023 Timothy E. Peoples

/*
Package decor provides facilities for decorating a string of characters
with display attributes for the current (or specified) terminal type. This
is done using a notation inspired by (but slightly different from) Zsh
prompt formatting.  Supported attributes are bold, italic, and/or underlined
characters as well as 256-color support for foreground and background
colors.

As a simple example, the numbered markings in the decor notated string argument
to the Format method here:

	// error handling elided
	d, _ := decor.New()
	s, _ := d.Format("@B@F{44}@Iuser@i@F{Orchid1}@@@F{Green3}host@f@b")
	//                ^ ^     ^     ^ ^          ^ ^            ^ ^
	//                1 2     3     4 5          6 7            8 9

...each have the following meaning.

	#1. Start Bold Text
	#2. Start Foreground Color #44 (DarkTurquoise)
	#3. Start Italics Text
	#4. End Italics Text
	#5. Start Foreground Color Orchid1 (#213)
	#6. A Literal '@' Character
	#7. Begin Foreground Color Green3 (#40)
	#8. End Foreground Color
	#9. End Bold Text

Therefore, the value assigned to s would be the string "user@host" - formatted
for the current terminal - in all bold text with "user" displayed in italics
and foreground color DarkTurquoise, the '@' character with color Orchid1, and
"host" shown with color Green3.

Or more specifically, for TERM="xterm-256color", the value of s would be:

	"\x1b[1m\x1b[38;5;44m\x1b[3muser\x1b[23m\x1b[38;5;213m@\x1b[38;5;40mhost\x1b(B\x1b[m\x1b(B\x1b[m"

# Attributes

All attribute designators begin with an '@' sign (for "ATtribute");
a literal '@' sign is specified as two consecutive characters (e.g. "@@").
The '@' sign is immediately followed by one of several capital letters
to begin the attribute or a lower case letter to end that attribute.
The full list of supported attribute designators is as follows:

	@B (@b) - Start (stop) boldface mode
	@I (@i) - Start (stop) italics mode
	@U (@u) - Start (stop) underline mode
	@F (@f) - Start (stop) specified foreground color
	@K (@k) - Start (stop) specified background color

# Color Designations

The start-color designators (@F and @K) are then followed by a color name
or number wrapped in braces (such as "@F{DodgerBlue}"). Note that the braces
surrounding the color name (or number) are not limited to '{' and '}'; these
can be any matching pair of brace-like characters (i.e. "<color>", "[color]"
or "(color)") -or- any character at all, such that the beginning and ending
characters agree (e.g. "@F+color+").

See package toolman.org/terminal/decor/color for a complete list of supported
color names.

# Templates

In addition to simple string decoration, this package also supports variable
expansion through templates. The example above could be modified to instead
create a template like so:

	t, _ := d.Template("@B@F{44}@I${UserName}@i@F{Orchid1}@@@F{Green3}${HostName}@f@b")

	// ...later followed by...

	s, _ := t.Expand(map[string]string{
		"UserName": "user",
		"HostName": "host@F(Grey42).${Domain}@f",
		"Domain":   "example.com",
	})

Here, the template's Expand method is provided a map of variable names and
values for "${UserName}", "${HostName}", and "${Domain}" which are expanded
to format the template for the current terminal.

Note that values for template variable may reference other variables (as above
where ${HostName} references ${Domain}) and can themselves contain attribute
designations (as ${HostName} uses a difference color for ${Domain}). Special
care is taken to restore attributes after a variable expansion that were in
effect before the expansion started and unnecessary or redundant attributes are
optimized away.

For example, if a template specifies a foreground color of "SpringGreen" and
a variable changes the foreground color to "DarkCyan", the foreground color
will be restored to "SpringGreen" once the variable has been expanded.
*/
package decor

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xo/terminfo"

	"toolman.org/terminal/decor/color"
	"toolman.org/terminal/decor/internal/colors"
	"toolman.org/terminal/decor/internal/item"
)

type level int32

type Decorator struct {
	term  string
	sgr0  string
	enter map[item.Type]string
	exit  map[item.Type]string
	fg    []string
	bg    []string
	debug int
}

// New returns a new *Decorator for the terminal type specified by the $TERM
// environment variable, or nil and an error if a new *Decorator cannot be
// created.
func New() (*Decorator, error) {
	ti, err := terminfo.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	return newDecorator(os.Getenv("TERM"), ti), nil
}

// Load returns a new *Decorator for the specified terminal type (ignoring the
// current environment) or nil and an error if a new *Decorator cannot be
// created.
func Load(term string) (*Decorator, error) {
	ti, err := terminfo.Load(term)
	if err != nil {
		return nil, err
	}

	return newDecorator(term, ti), nil
}

// Term returns the terminal type used to create the receiver.
func (d *Decorator) Term() string {
	if d != nil {
		return d.term
	}
	return ""
}

func newDecorator(term string, ti *terminfo.Terminfo) *Decorator {
	d := &Decorator{
		term: term,
		sgr0: ti.Printf(terminfo.ExitAttributeMode),

		enter: map[item.Type]string{
			item.BOLD:      ti.Printf(terminfo.EnterBoldMode),
			item.ITALIC:    ti.Printf(terminfo.EnterItalicsMode),
			item.UNDERLINE: ti.Printf(terminfo.EnterUnderlineMode),
		},

		exit: map[item.Type]string{
			item.BOLD:      ti.Printf(terminfo.ExitAttributeMode),
			item.ITALIC:    ti.Printf(terminfo.ExitItalicsMode),
			item.UNDERLINE: ti.Printf(terminfo.ExitUnderlineMode),
		},

		fg: make([]string, len(colors.Names)),
		bg: make([]string, len(colors.Names)),
	}

	for n := range colors.Names {
		d.fg[n] = ti.Printf(terminfo.SetAForeground, n)
		d.bg[n] = ti.Printf(terminfo.SetABackground, n)
	}

	return d
}

func (d *Decorator) enterCode(itm *item.Item) string {
	if d == nil || itm == nil {
		return ""
	}

	switch itm.Type {
	case item.FGCOLOR:
		return d.fgColor(itm.Text)
	case item.BGCOLOR:
		return d.bgColor(itm.Text)
	default:
		return d.enter[itm.Type]
	}
}

func (d *Decorator) exitCode(itm *item.Item) string {
	if d == nil || itm == nil {
		return ""
	}

	switch itm.Type {
	case item.FGCOLOR, item.BGCOLOR:
		return d.sgr0
	default:
		return d.exit[itm.Type]
	}
}

func (d *Decorator) isAllOff(itm *item.Item) bool {
	if d != nil && itm != nil && itm.Action == item.STOP {
		return d.exitCode(itm) == d.sgr0
	}

	return false
}

func (d *Decorator) fgColor(s string) string { return lookup(s, d.fg) }
func (d *Decorator) bgColor(s string) string { return lookup(s, d.bg) }

func lookup(clr string, codes []string) string {
	if n := color.Number(clr); n >= 0 {
		return codes[n]
	}

	if n, err := strconv.Atoi(clr); err == nil {
		return codes[n]
	}

	return fmt.Sprintf("<!color:%s>", clr)
}

func (d *Decorator) debugf(level int, msg string, args ...any) {
	if d == nil || d.debug < level {
		return
	}

	if c := msg[len(msg)-1]; c != '\n' {
		msg += "\n"
	}

	fmt.Fprintf(os.Stderr, msg, args...)
}
