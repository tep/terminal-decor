// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

func (d *Decorator) resolve(name string, values map[string]string) *series.Series {
	d.debugf(1, "  resolving: %q", name)

	ss := series.New()
	val, ok := values[name]
	if !ok {
		return ss.Append(item.ErrItemf("<undef:%s>", name))
	}

	d.debugf(1, "      value: %q", val)

	if val == "" {
		return ss.Append(item.TextItem(""))
	}

	if err := ss.Parse(val); err != nil {
		return ss.Append(item.ErrItem(err))
	}

	out := series.New()

	var attr bool
	for itm := ss.Front().Clone(); itm != nil; itm = itm.Next() {
		d.debugf(2, "  ## %s", itm)
		if itm.Type == item.VAR {
			out.AppendList(d.resolve(itm.Text, values))
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
