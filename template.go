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
// then has an Expand method that accepts a map of variable names to values
// that will be expanded to emit a formatted string. If the given text cannot
// be parsed as decor notation, nil and an error are returned.
func (d *Decorator) Template(text string) (*Template, error) {
	ss := series.New()

	if err := ss.Parse(text); err != nil {
		return nil, err
	}

	return &Template{ss: ss, dec: d}, nil
}

// Expand will expand the given map of variable names to values for the items
// parsed when creating the Template receiver. Values may themselves contain
// decor attributes and/or references to other variables -- however, circular
// references are not allowed and will be rendered as errors in the output
// string.
func (t *Template) Expand(values map[string]string) string {
	ss := series.New()

	t.dec.debugf(1, "template: formatting %d items", t.ss.Len())
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

---

XXX I'm not sure I like this 'Blend' method defined above
    (which is why it's commented out).

Instead of blending/merging templates ad-hoc with a method like this -- the
'Template' method chould also accept a 'name' parameter which would be used as
a key into a template cache held within the Decorator... and, since those
templates are now cached, there's no need to return a *Template object for
later use; Decorator would have an 'Expand' method taking a template name and
ssmap.

With the above, Series.Parse could then support several 'sigils' for declaring
a variety of different things:

    @ Attributes
      For bold, underline, colors, etc...

    $ Environment Variables
      The '$' sigil is currently used for "Decor" Variables

    = Decor Variables
      As currently provided by the ssmap argument to Template.Format
      (Note: these are currently declared with '$')

    & Template Names
      Names of cached templates

The end result would be Decor Text similar to the following:

    frame: "@B@F{Grey35}[={Glyph}:={Name}]@f@b"
    line:  "&{frame}:${PWD}"

...of course, there's a chance none of this is really useful at all so, think
about it before putting a lot of work into it.
*/
