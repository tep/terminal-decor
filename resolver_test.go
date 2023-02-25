// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"testing"

	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

type resolverTestcase struct {
	label   string
	varname string
	want    *series.Series
}

func TestResolver(t *testing.T) {
	dec, err := xterm256Decorator()
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"A": "foo-${C}-bar",
		"B": "123-${D}-789",
		"C": "abc-${E}-xyz",
		"D": "klm-${F}-nop",
		"E": "987-654-321",
		"F": "xyz-${B}-abc",
	}

	cases := []resolverTestcase{
		{"ok", "A", series.Build(
			item.TextItem("foo-"),
			item.TextItem("abc-"),
			item.TextItem("987-654-321"),
			item.TextItem("-xyz"),
			item.TextItem("-bar"),
		)},

		{"circular-ref", "B", series.Build(
			item.TextItem("123-"),
			item.TextItem("klm-"),
			item.TextItem("xyz-"),
			item.TextItem("123-"),
			item.ErrItem(&crefError{[]string{"B", "D", "F", "B"}}),
			item.TextItem("-789"),
			item.TextItem("-abc"),
			item.TextItem("-nop"),
			item.TextItem("-789"),
		)},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			if got := dec.resolve(tc.varname, vars); !got.Equal(tc.want) {
				t.Errorf("dec.resolve(%q, ...) == (%s)\nWanted (%s)", tc.varname, got, tc.want)
			}
		})
	}
}
