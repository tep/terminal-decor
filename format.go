// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"fmt"

	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

// Format converts the given decor-notated text into a string ready for
// display on the receiver's associated terminal type, or the empty string
// and an error if it cannot do so. Note that if text contains any decor
// variable references they will be ignored in the resultant output.
// Create a Template and use its Expand method to resolve decor variables.
//
// See the package documentation for more details on decor notation.
func (d *Decorator) Format(text string) (string, error) {
	ss := series.New()

	if err := ss.Parse(text); err != nil {
		return "", err
	}

	return d.format(ss), nil
}

// Formatf is a wrapper around Format providing a Printf like interface.
func (d *Decorator) Formatf(msg string, args ...any) (string, error) {
	return d.Format(fmt.Sprintf(msg, args...))
}

func (d *Decorator) format(ss *series.Series) string {
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

	return out
}
