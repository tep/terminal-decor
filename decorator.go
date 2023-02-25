// Copyright © 2023 Timothy E. Peoples

// Package decor provides facilities for decorating a strings of characters
// with display attributes for the current (or specified) terminal type. This
// is done using a notation inspired by (but slightly different from) Zsh
// prompt formatting.  Supported attributes are bold, italic, and/or underlined
// characters as well as 256-color support for foreground and background
// colors.
//
// As a simple example, the decor notated string provided to the Format method
// below:
//
//	// error handling elided
//	d, _ := decor.New()
//	s, _ := d.Format("@B@F{44}@Iuser@i@F{250}@@@F{21}host@f@b")
//	//                ^ ^     ^     ^ ^      ^ ^         ^ ^
//	//                1 2     3     4 5      6 7         8 9
//
// ...has the following meaning.
//
//	#1. Start Bold Text
//	#2. Start Foreground Color #44
//	#3. Start Italics Text
//	#4. End Italics Text
//	#5. Start Foreground Color #250
//	#6. A Literal '@' Character
//	#7. Begin Foreground Color #21
//	#8. End Foreground Color
//	#9. End Bold Text
//
// Therefore, the value of s would be the string "user@host" - formatted for
// the current terminal - in all bold text with "user" displayed in italics
// with foreground color 44, the '@' character with color 250, and "host" shown
// with color 21.
//
// All attribute designators begin with an '@' sign (for "attribute");
// a literal '@' sign is specified as two consecutive characters (e.g. "@@").
// The '@' sign is immediately followed by one of several capital letters
// to begin the attribute or a lower case letter to end that attribute.
// The full list of supported attribute designators is as follows:
//
//	@B (@b) - Start (stop) boldface mode
//	@I (@i) - Start (stop) italics mode
//	@U (@u) - Start (stop) underline mode
//	@F (@f) - Start (stop) specified foreground color
//	@K (@k) - Start (stop) specified background color
//
// The start-color designators (@F and @K) are then followed by a color name
// or number wrapped in braces (such as "@F{DodgerBlue}"). Note that the braces
// surrounding the color name (or number) are not limited to '{' and '}'; these
// can be any matching pair of brace-like characters (i.e. "<color>", "[color]"
// or "(color)") -or- any character at all, such that the beginning and ending
// characters agree (e.g. "@F+color+").
//
// In addition to simple string decoration, this package also supports variable
// expansion through templates. The above decor text could be altered to create
// a template with the string:
//
//	t, _ := d.Template("@B@F{44}@I${UserName}@i@F{250}@@@F{21}${HostName}@f@b")
//
//	// ...later followed by...
//
//	s, _ := t.Format(map[string]string{"UserName": "user", "HostName": "host"})
//
// Here, the variables "${UserName}" and "${HostName}" would be expanded by
// formatting the template with the provided map of variable names to values.
//
// Note that variable values may reference other variables and can themselves
// contain attribute designations. Special care is taken to restore attributes
// after a variable expansion that were in effect before the expansion started.
// For example, if a template specifies a foreground color of "SpringGreen" and
// a variable changes the foreground color to "DarkCyan", the foreground color
// will be restored to "SpringGreen" once the variable has been expanded.
package decor

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xo/terminfo"

	"toolman.org/terminal/decor/internal/item"
)

type level int32

type Decorator struct {
	sgr0  string
	enter map[item.Type]string
	exit  map[item.Type]string
	fg    []string
	bg    []string
	debug int
}

// New returns a new *Decorator for the terminal type specified by the $TERM
// environment variable, or nil and an error if a new *Decorator connot be
// created.
func New() (*Decorator, error) {
	ti, err := terminfo.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	return newDecorator(ti), nil
}

// Load returns a new *Decorator for the specified terminal type (ignoring the
// current environment) or nil and an error if a new *Decorator connot be
// created.
func Load(term string) (*Decorator, error) {
	ti, err := terminfo.Load(term)
	if err != nil {
		return nil, err
	}

	return newDecorator(ti), nil
}

func newDecorator(ti *terminfo.Terminfo) *Decorator {
	d := &Decorator{
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

		fg: make([]string, len(names)),
		bg: make([]string, len(names)),
	}

	for n := range names {
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

func lookup(color string, codes []string) string {
	if n, ok := NumberOK(color); ok {
		return codes[n]
	}

	if n, err := strconv.Atoi(color); err == nil {
		return codes[n]
	}

	return fmt.Sprintf("<!color:%s>", color)
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
