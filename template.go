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

func (d *Decorator) Template(text string) (*Template, error) {
	ss := series.New()

	if err := ss.Parse(text); err != nil {
		return nil, err
	}

	return &Template{ss: ss, dec: d}, nil
}

func (t *Template) String() string {
	return t.ss.String()
}

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
