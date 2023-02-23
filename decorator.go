// Copyright © 2023 Timothy E. Peoples

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

func New() (*Decorator, error) {
	ti, err := terminfo.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	return newDecorator(ti), nil
}

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
