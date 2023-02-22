// Copyright © 2023 Timothy E. Peoples

package termcolor

import "fmt"

type Template struct {
	segments []segment
	term     *Term
}

func (t *Term) Template(text string) (*Template, error) {
	segs, err := t.segmentize(text)
	if err != nil {
		return nil, err
	}

	return &Template{segments: segs, term: t}, nil
}

func (t *Template) Format(values map[string]string) (string, error) {
	var list []segment

	for _, seg := range t.segments {
		if seg.stype != stVariable {
			list = append(list, seg)
			continue
		}

		vs := t.term.resolve(seg.text, values)
		list = append(list, vs...)
	}

	return t.term.format(t.term.optimize(list))
}

func (t *Term) resolve(name string, values map[string]string) []segment {
	val, ok := values[name]
	if !ok {
		return []segment{textSegment(fmt.Sprintf("<undef:%s>", name))}
	}

	if val == "" {
		return []segment{textSegment("")}
	}

	segs, err := t.segmentize(val)
	if err != nil {
		return []segment{errSegment(err)}
	}

	var out []segment

	var attr bool
	for _, s := range segs {
		if s.stype == stVariable {
			vs := t.resolve(s.text, values)
			out = append(out, vs...)
			continue
		}

		switch {
		case s.isAttrOn():
			attr = true
		case s.isAttrOff():
			attr = false
		}

		out = append(out, s)
	}

	if attr {
		out = append(out, newSegment('b'))
	}

	return out
}
