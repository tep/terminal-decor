// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

func (d *Decorator) Format(text string) (string, error) {
	ss := series.New()

	if err := ss.Parse(text); err != nil {
		return "", err
	}

	return d.format(ss)
}

func (d *Decorator) format(ss *series.Series) (string, error) {
	var out string

	for itm := ss.Front(); itm != nil; itm = itm.Next() {
		d.debugf(1, ">> %s", itm)

		switch itm.Action {
		case item.START:
			code := d.enterCode(itm)
			d.debugf(1, "++ %q", code)
			out += code
		case item.STOP:
			code := d.exitCode(itm)
			d.debugf(1, "-- %q", code)
			out += code
		default:
			out += itm.Text
		}
	}

	return out, nil
}
