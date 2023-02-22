// Copyright © 2023 Timothy E. Peoples

package termcolor

import (
	"fmt"
	"strings"
	"sync/atomic"
)

const (
	stText segType = iota
	stError
	stVariable
	__st_attrs__
	stBold
	stUnderline
	stItalic
	stFGcolor
	stBGcolor
)

type segType int

func (t segType) String() string {
	switch t {
	case stText:
		return "TEXT"
	case stBold:
		return "BOLD"
	case stUnderline:
		return "UNDERLINE"
	case stItalic:
		return "ITALIC"
	case stFGcolor:
		return "FGCOLOR"
	case stBGcolor:
		return "BGCOLOR"
	case stVariable:
		return "VAR"
	case stError:
		return "ERR"
	default:
		return "<unknown>"
	}
}

var segmentID uint64

func nextID() uint64 {
	return atomic.AddUint64(&segmentID, 1)
}

type segment struct {
	stype  segType
	text   string
	action action
	id     uint64
}

type action int

const (
	aNone action = iota
	aStart
	aStop
)

func (a action) String() string {
	switch a {
	case aNone:
		return ""
	case aStart:
		return "START"
	case aStop:
		return "STOP"
	default:
		return fmt.Sprintf("<err:%d>", a)
	}
}

func textSegment(s string) segment {
	return segment{text: s, id: nextID()}
}

func varSegment(name string) segment {
	return segment{stype: stVariable, text: name, id: nextID()}
}

func errSegment(err error) segment {
	return segment{stype: stError, text: fmt.Sprintf("<err:%v>", err), id: nextID()}
}

func startSegment(stype segType) segment {
	return segment{stype: stype, action: aStart, id: nextID()}
}

func stopSegment(stype segType) segment {
	return segment{stype: stype, action: aStop, id: nextID()}
}

func newSegment(c byte) segment {
	switch c {
	case 'B':
		return startSegment(stBold)
	case 'F':
		return startSegment(stFGcolor)
	case 'I':
		return startSegment(stItalic)
	case 'K':
		return startSegment(stBGcolor)
	case 'U':
		return startSegment(stUnderline)
	case 'b':
		return stopSegment(stBold)
	case 'f':
		return stopSegment(stFGcolor)
	case 'i':
		return stopSegment(stItalic)
	case 'k':
		return stopSegment(stBGcolor)
	case 'u':
		return stopSegment(stUnderline)
	default:
		return segment{id: nextID()}
	}
}

func (s segment) String() string {
	if s.stype < __st_attrs__ {
		return fmt.Sprintf("%03d:%s:%q", s.id, s.stype, s.text)
	}

	parts := []string{fmt.Sprintf("%03d", s.id), s.action.String(), s.stype.String()}

	if s.text != "" {
		parts = append(parts, fmt.Sprintf("%q", s.text))
	}

	return strings.Join(parts, ":")
}

func (ss *segment) equal(os *segment) bool {
	switch {
	case ss == nil && os == nil:
		return true
	case ss == nil || os == nil:
		return false
	default:
		return ss.stype == os.stype && ss.action == os.action && ss.text == os.text
	}
}

func (ss *segment) isAttrOn() bool {
	if ss != nil {
		return ss.action == aStart
	}
	return false
}

func (ss *segment) isAttrOff() bool {
	if ss == nil || ss.action == aStart {
		return false
	}

	switch ss.stype {
	case stBold, stFGcolor, stBGcolor:
		return true
	default:
		return false
	}
}

func (t *Term) isAllOff(sgmt *segment) bool {
	if t == nil || sgmt == nil || sgmt.action != aStop {
		return false
	}

	xcode := t.exitCode(*sgmt)

	if DEBUG {
		fmt.Printf("   %-17s [ALLOFF=%v]\n", sgmt, xcode == t.sgr0)
	}

	return xcode == t.sgr0
}

func (t *Term) segmentize(input string) ([]segment, error) {
	var output []segment
	var bufstr string

	for input != "" {
		i := strings.IndexAny(input, "@$")
		if i == -1 {
			bufstr += input
			break
		}

		sigil := input[i]

		bufstr += input[:i]

		if i == len(input)-1 {
			return nil, fmt.Errorf("sigil %q not allowed at end of string", sigil)
		}

		c := input[i+1]
		input = input[i+2:]

		if c == sigil {
			bufstr += string(sigil)
			continue
		}

		if bufstr != "" {
			output = append(output, textSegment(bufstr))
			bufstr = ""
		}

		if sigil == '$' {
			if c != '{' {
				return nil, fmt.Errorf("malformed variable reference at pos %d", i)
			}

			j := strings.IndexByte(input, '}')
			if j == -1 {
				return nil, fmt.Errorf("unterminated variable reference at pod %d", i)
			}

			output = append(output, varSegment(input[:j]))
			input = input[j+1:]
			continue
		}

		seg := newSegment(c)

		if c != 'F' && c != 'K' {
			output = append(output, seg)
			continue
		}

		if input == "" {
			return nil, fmt.Errorf("malformed attribute: %c", c)
		}

		var cc byte
		switch input[0] {
		case '{':
			cc = '}'
		case '(':
			cc = ')'
		case '[':
			cc = ']'
		case '<':
			cc = '>'
		default:
			cc = input[0]
		}

		input = input[1:]

		j := strings.IndexByte(input, cc)
		if j == -1 {
			return nil, fmt.Errorf("unterminated attribute %q at pos %d", c, i)
		}

		seg.text = input[:j]
		output = append(output, seg)

		input = input[j+1:]
	}

	if bufstr != "" {
		output = append(output, textSegment(bufstr))
	}

	if DEBUG {
		fmt.Println("\nSEGMENTIZED")
		for _, s := range output {
			fmt.Println("  ", s)
		}
		fmt.Println("--\n")
	}

	return output, nil
}

func (t *Term) optimize(segments []segment) []segment {
	count := len(segments)
	if count == 0 {
		return nil
	}

	fmt.Printf("\nOPTIMIZING %d segments=%v\n", count, segments)

	var active, out []segment

	var prev, this *segment
	for i, s := range segments {
		prev = this
		this = &s

		fmt.Printf("i=%-3d this=%-30s active=%v\n", i, this, active)
		fmt.Printf("  segments: %v\n", segments)

		switch this.action {
		case aStart:
			var nact []segment
			for _, as := range active {
				if !as.equal(this) {
					nact = append(nact, as)
				}
			}
			nact = append(nact, *this)
			active = nact

		case aStop:
			if prev != nil && prev.stype == this.stype && prev.action == aStart {
				fmt.Printf("      optimizing %s for %s\n", prev, this) // XXX
				// splice 'prev' from OUT!
				fmt.Printf("          before out=%v\n", out) // XXX
				out = out[:len(out)-1]
				fmt.Printf("           after out=%v\n", out) // XXX
				// ...and remove prev from "active"
				fmt.Printf("          before active=%v\n", active) // XXX
				active = active[:len(active)-1]
				fmt.Printf("           after active=%v\n", active) // XXX
				i--
				continue
			}

			// find index of last instance of this.stype in 'active' and remove it
			for j := len(active) - 1; j >= 0; j-- {
				if active[j].stype == this.stype {
					fmt.Printf("      for %d, removing %d from active\n", this.id, active[j].id) // XXX
					rem := active[j+1:]
					active = active[:j]
					active = append(active, rem...)
					break
				}
			}

			// ...but, if 'this' turned everything off, we'll need to turn the
			// 'active' stuff back on.
			if t.isAllOff(this) && len(active) > 0 && i < count-1 {
				out = append(out, active...)
				fmt.Printf("      re-adding: %v\n", active)
				fmt.Printf("          before segments=%v\n", segments) // XXX
				rem := make([]segment, len(segments[i+1:]))
				copy(rem, segments[i+1:])
				segments = append(segments[:i], active...)
				segments = append(segments, rem...)
				fmt.Printf("           after segments=%v\n", segments) // XXX
				fmt.Printf("i=%-3d out=%v\n", i, out)
				continue
			}
		}

		out = append(out, *this)
		fmt.Printf("i=%-3d out=%v\n", i, out)
	}

	fmt.Println("--\n")

	return out
}
