// Copyright © 2023 Timothy E. Peoples

package series

import (
	"fmt"
	"strings"

	"toolman.org/terminal/decor/internal/item"
)

func (s *Series) Parse(input string) error {
	var bufstr string

	for input != "" {
		i := strings.IndexAny(input, "@$")
		if i == -1 {
			bufstr += input
			break
		}

		sigil := input[i]

		if i == len(input)-1 {
			return fmt.Errorf("sigil %q not allowed at end of string", sigil)
		}

		bufstr += input[:i]
		c := input[i+1]
		input = input[i+2:]

		if c == sigil {
			bufstr += string(sigil)
			continue
		}

		if bufstr != "" {
			s.Append(item.TextItem(bufstr))
			bufstr = ""
		}

		if sigil == '$' {
			if c != '{' {
				return fmt.Errorf("malformed variable reference at pos %d", i)
			}

			j := strings.IndexByte(input, '}')
			if j == -1 {
				return fmt.Errorf("unterminated variable reference at pod %d", i)
			}

			s.Append(item.VarItem(input[:j]))
			input = input[j+1:]
			continue
		}

		sgmt := item.AttrItem(c)

		if c != 'F' && c != 'K' {
			s.Append(sgmt)
			continue
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
			return fmt.Errorf("unterminated attribute %q at pos %d", c, i)
		}

		sgmt.Text = input[:j]
		s.Append(sgmt)

		input = input[j+1:]
	}

	if bufstr != "" {
		s.Append(item.TextItem(bufstr))
	}

	return nil
}
