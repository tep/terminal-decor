// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

type Template struct {
	ss  *series.Series
	dec *Decorator
}

// Template is similar to Decorator's Format method but instead of returning
// a formatted string, a parsed *Template is returned. The returned *Template
// then has a Format method that accepts a map of variable names to values
// that will be expanded to emit a formatted string.
func (d *Decorator) Template(text string) (*Template, error) {
	ss := series.New()

	if err := ss.Parse(text); err != nil {
		return nil, err
	}

	return &Template{ss: ss, dec: d}, nil
}

// Format expands the given map of variable name to value for the items parsed
// when creating the reciever. Values may themselves contain decor attributes
// and/or references to other variables -- however, circular references are
// not allowed.
func (t *Template) Format(values map[string]string) (string, error) {
	ss := series.New()

	t.dec.debugf(1, "template: formatting %d segments", t.ss.Len())
	t.dec.debugf(2, "          using values: %#v", values)

	for itm := t.ss.Front(); itm != nil; itm = itm.Next() {
		t.dec.debugf(2, "    %s", itm)
		if itm.Type == item.VAR {
			ss.AppendList(t.dec.resolve(itm.Text, values))
		} else {
			ss.Append(itm.Detach())
		}
	}

	return t.dec.format(t.dec.optimize(ss))
}

/*
XXX I'm not sure I like this 'Blend' method defined below.

Instead of blending/merging templates ad-hoc with a method like this -- the
'Template' method should also accept a 'name' parameter to be used as a key
into a template cache held within the Decorator... and, since those templates
are now cached, there's no need to return a *Template object for later use.

With the above, Series.Parse would then support several 'sigils' for declaring
different things:

    @ Attributes
      For bold, underline, colors, etc...

    $ Environment Variables
      The '$' sigil is currently used for "Decor" Variables

    = Decor Variables
      As currently provided by the map[string]string argument to Template.Format
      (Note: these are currently declared with '$')

    & Template Names
      Names of cached templates

The end result would be Decor Text similar to the following:

    frame: "@B@F{Grey35}[={Glyph}:={Name}]@f@b"
    line:  "&frame:${PWD}"

...of course, there's a chance none of this is really useful at all so, think
about it before putting a lot of work into it.

---

func (t *Template) Blend(others map[string]*Template) *Template {
	ss := series.New()

	for itm := t.ss.Front(); itm != nil; itm = itm.Next() {
		if itm.Type == item.VAR {
			if ot := others[itm.Text]; ot != nil {
				ss.AppendList(ot.ss)
				continue
			}
		}

		ss.Append(itm.Detach())
	}

	return &Template{ss: ss, dec: t.dec}
}
*/
