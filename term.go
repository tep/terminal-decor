// Copyright © 2023 Timothy E. Peoples

package termcolor

import "github.com/xo/terminfo"

type Term struct {
	ti    *terminfo.Terminfo
	sgr0  string
	enter map[segType]string
	exit  map[segType]string
	fg    []string
	bg    []string
}

func New() (*Term, error) {
	ti, err := terminfo.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	t := &Term{
		ti: ti,

		sgr0: ti.Printf(terminfo.ExitAttributeMode),

		enter: map[segType]string{
			stBold:      ti.Printf(terminfo.EnterBoldMode),
			stItalic:    ti.Printf(terminfo.EnterItalicsMode),
			stUnderline: ti.Printf(terminfo.EnterUnderlineMode),
		},

		exit: map[segType]string{
			stBold:      ti.Printf(terminfo.ExitAttributeMode),
			stItalic:    ti.Printf(terminfo.ExitItalicsMode),
			stUnderline: ti.Printf(terminfo.ExitUnderlineMode),
		},

		fg: make([]string, len(names)),
		bg: make([]string, len(names)),
	}

	for n := range names {
		t.fg[n] = ti.Printf(terminfo.SetAForeground, n)
		t.bg[n] = ti.Printf(terminfo.SetABackground, n)
	}

	return t, nil
}
