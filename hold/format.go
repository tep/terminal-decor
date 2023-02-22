// Copyright © 2023 Timothy E. Peoples

package termcolor

import (
	"fmt"
	"strconv"
)

var DEBUG bool

func (t *Term) Format(text string) (string, error) {
	segments, err := t.segmentize(text)
	if err != nil {
		return "", err
	}

	return t.format(segments)
}

func (t *Term) format(segments []segment) (string, error) {
	var out string

	for _, sgmt := range segments {
		if DEBUG {
			fmt.Printf(">> %s\n", sgmt.String())
		}

		switch sgmt.action {
		case aStart:
			code := t.enterCode(sgmt)
			if DEBUG {
				fmt.Printf("++ %q\n", code)
			}
			out += code
		case aStop:
			code := t.exitCode(sgmt)
			if DEBUG {
				fmt.Printf("-- %q\n", code)
			}
			out += code
		default:
			out += sgmt.text
		}
	}

	return out, nil
}

// func (t *Term) _x_format(segments []segment) (string, error) {
// 	var out string
// 	stack := t.stack()
//
// 	segCount := len(segments)
// 	for i := 0; i < segCount; i++ {
// 		sgmt := segments[i]
// 		var next segment
// 		if i < segCount-1 {
// 			next = segments[i+1]
// 		}
//
// 		if DEBUG {
// 			fmt.Printf(">> %s\n", sgmt.String())
// 		}
//
// 		// if this segment enables an attribute and the next segment immediately
// 		// disables the same attribute, let's just skip the both of them.
// 		if sgmt.action == aStart && sgmt.stype == next.stype && next.action == aStop {
// 			if DEBUG {
// 				fmt.Printf("## SKIP: %s -> %s\n", sgmt, next)
// 			}
// 			i++
// 			continue
// 		}
//
// 		switch {
// 		case sgmt.stype == stText:
// 			out += sgmt.text
// 		case sgmt.action == aStart:
// 			out += stack.start(sgmt)
// 		case sgmt.action == aStop:
// 			out += stack.stop(sgmt)
// 		}
//
// 		if DEBUG {
// 			fmt.Printf("== [%s]\n", stack)
// 		}
// 	}
//
// 	return out, nil
//
// }

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
