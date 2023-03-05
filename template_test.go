// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"testing"

	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

const (
	setaf59  = "\x1b[38;5;59m"
	setaf214 = "\x1b[38;5;214m"
)

func TestTemplate(t *testing.T) {
	d, err := xterm256Decorator()
	if err != nil {
		t.Fatal(err)
	}

	input := "@F(Grey37)[${Glyph}:@I${Key}@i]@f"

	tmplWant := series.Build(
		item.FGColorItem("Grey37"),
		item.TextItem("["),
		item.VarItem("Glyph"),
		item.TextItem(":"),
		item.StartItem(item.ITALIC),
		item.VarItem("Key"),
		item.StopItem(item.ITALIC),
		item.TextItem("]"),
		item.StopItem(item.FGCOLOR),
	)

	var tmpl *Template
	if tmpl, err = d.Template(input); err != nil || !tmpl.equal(tmplWant) {
		t.Fatalf("d.Template(%q) == (%q, %v);\nWanted (%s, nil)", input, tmpl.ss, err, tmplWant)
	}

	vars := map[string]string{"Glyph": "@F<Orange1>Ж@f", "Key": "@F{204}ABC@f"}
	want := "" +
		xt_setaf59 + "[" +
		xt_setaf214 + "Ж" +
		xt_setaf59 + ":" +
		xt_sitm + xt_setaf204 + "ABC" +
		xt_setaf59 + xt_ritm + "]" +
		xt_defFG

	if got := tmpl.Expand(vars); got != want {
		t.Errorf("tmpl.Format(%#v):\n   Got: %q\nWanted: %q", vars, decodeAttrString(got), decodeAttrString(want))
	} else {
		t.Logf("OUT: %s (%s)", got, decodeAttrString(got))
	}
}

func (t *Template) equal(os *series.Series) bool {
	return t.ss.Equal(os)
}
