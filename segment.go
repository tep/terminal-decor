// Copyright © 2023 Timothy E. Peoples

package termcolor

import (
	"errors"
	"fmt"
	"strings"
)

const (
	stText segType = iota
	stBold
	stUnderline
	stItalic
	stFGcolor
	stBGcolor
	stVariable
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
	default:
		return "<unknown>"
	}
}

type segment struct {
	stype segType
	text  string
	start bool
}

func (s segment) String() string {
	if s.stype == stText {
		return fmt.Sprintf("TEXT:%q", s.text)
	}

	var parts []string
	ss := "STOP"
	if s.start {
		ss = "START"
	}

	parts = append(parts, ss, s.stype.String())

	if s.text != "" {
		parts = append(parts, fmt.Sprintf("%q", s.text))
	}

	return strings.Join(parts, ":")
}

func textSegment(s string) segment {
	return segment{text: s}
}

func newSegment(c byte) segment {
	switch c {
	case 'B':
		return segment{stype: stBold, start: true}
	case 'F':
		return segment{stype: stFGcolor, start: true}
	case 'I':
		return segment{stype: stItalic, start: true}
	case 'K':
		return segment{stype: stBGcolor, start: true}
	case 'U':
		return segment{stype: stUnderline, start: true}
	case 'b':
		return segment{stype: stBold}
	case 'f':
		return segment{stype: stFGcolor}
	case 'i':
		return segment{stype: stItalic}
	case 'k':
		return segment{stype: stBGcolor}
	case 'u':
		return segment{stype: stUnderline}
	default:
		return segment{}
	}
}

func (ss *segment) equal(os *segment) bool {
	switch {
	case ss == nil && os == nil:
		return true
	case ss == nil || os == nil:
		return false
	default:
		return ss.stype == os.stype && ss.start == os.start && ss.text == os.text
	}
}

func segmentize(s string) ([]segment, error) {
	var out []segment
	var text string

	for s != "" {
		i := strings.IndexRune(s, '@')
		if i == -1 {
			text += s
			break
		}

		text += s[:i]

		if i == len(s)-1 {
			return nil, errors.New("@ not allowed at end of string")
		}

		c := s[i+1]
		s = s[i+2:]

		if c == '@' {
			text += "@"
			continue
		}

		if text != "" {
			out = append(out, textSegment(text))
			text = ""
		}

		seg := newSegment(c)

		if c != 'F' && c != 'K' {
			out = append(out, seg)
			continue
		}

		if s == "" {
			return nil, fmt.Errorf("malformed attribute: %c", c)
		}

		var cc byte
		switch s[0] {
		case '{':
			cc = '}'
		case '(':
			cc = ')'
		case '[':
			cc = ']'
		case '<':
			cc = '>'
		default:
			cc = s[0]
		}

		s = s[1:]

		j := strings.IndexByte(s, cc)
		if j == -1 {
			return nil, fmt.Errorf("unterminated attribute %q at pos %d", c, i)
		}

		seg.text = s[:j]
		out = append(out, seg)

		s = s[j+1:]
	}

	if text != "" {
		out = append(out, textSegment(text))
	}

	return out, nil
}
