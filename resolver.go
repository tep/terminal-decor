// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"fmt"
	"strings"

	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

func (d *Decorator) resolve(name string, values map[string]string) *series.Series {
	return d.resolver().resolve(name, values)
}

// type void struct{} // XXX is this needed?

type resolver struct {
	refs map[string]string
	*Decorator
}

func (d *Decorator) resolver() *resolver {
	return &resolver{make(map[string]string), d}
}

func (r *resolver) checkRefs(name string) error {
	ref, in := r.refs[name]
	if !in {
		return nil
	}

	refs := []string{ref}
	circle := []string{ref, name}

	for ref != name {
		if ref = r.refs[ref]; ref != name {
			refs = append([]string{ref}, refs...)
		}
	}

	circle = append(circle, refs...)

	return &crefError{circle}
}

type crefError struct {
	circle []string
}

func (cre *crefError) Error() string {
	return fmt.Sprintf("circular reference: %s", strings.Join(cre.circle, "->"))
}

func (r *resolver) resolve(name string, values map[string]string) *series.Series {
	r.debugf(1, "  resolving: %q", name)

	ss := series.New()
	val, ok := values[name]
	if !ok {
		return ss.Append(item.ErrItemf("<undef:%s>", name))
	}

	r.debugf(1, "      value: %q", val)

	if val == "" {
		return ss.Append(item.TextItem(""))
	}

	if err := ss.Parse(val); err != nil {
		return ss.Append(item.ErrItem(err))
	}

	out := series.New()

	var attr bool
	for itm := ss.Front().Clone(); itm != nil; itm = itm.Next() {
		r.debugf(2, "  ## %s", itm)
		if itm.Type == item.VAR {
			if err := r.checkRefs(itm.Text); err != nil {
				out.Append(item.ErrItem(err))
			} else {
				r.refs[itm.Text] = name
				defer delete(r.refs, itm.Text)
				out.AppendList(r.resolve(itm.Text, values))
			}
			continue
		}

		switch {
		case itm.IsAttrOn():
			attr = true
		case itm.IsAttrOff():
			attr = false
		}

		out.Append(itm.Detach())
	}

	if attr {
		out.Append(item.AttrItem('b'))
	}

	return out
}
