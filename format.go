// Copyright © 2023 Timothy E. Peoples

package termcolor

import (
	"fmt"
	"strconv"
)

var DEBUG bool

func (t *Term) Format(mesg string) (string, error) {
	segments, err := segmentize(mesg)
	if err != nil {
		return "", err
	}

	var out string
	stack := t.stack()

	for _, sgmt := range segments {
		if DEBUG {
			fmt.Printf(">> %-25s", sgmt.String())
		}

		switch {
		case sgmt.stype == stText:
			out += sgmt.text
		case sgmt.start:
			out += stack.start(sgmt)
		default:
			out += stack.stop(sgmt)
		}

		if DEBUG {
			var pad string
			if sgmt.stype == stText {
				pad = fmt.Sprintf("%30s", "")
			}
			fmt.Printf("%s[%s]\n", pad, stack)
		}
	}

	return out, nil

}

func (t *Term) enterCode(sgmt segment) string {
	switch sgmt.stype {
	case stFGcolor:
		return t.fgColor(sgmt.text)
	case stBGcolor:
		return t.bgColor(sgmt.text)
	default:
		return t.enter[sgmt.stype]
	}
}

func (t *Term) exitCode(sgmt segment) string {
	stype := sgmt.stype
	switch stype {
	case stFGcolor, stBGcolor:
		return t.sgr0
	default:
		return t.exit[stype]
	}
}

func (t *Term) fgColor(s string) string { return lookup(s, t.fg) }
func (t *Term) bgColor(s string) string { return lookup(s, t.bg) }

func lookup(color string, codes []string) string {
	if n, ok := NumberOK(color); ok {
		return codes[n]
	}

	if n, err := strconv.Atoi(color); err == nil {
		return codes[n]
	}

	return fmt.Sprintf("<!color:%s>", color)
}
