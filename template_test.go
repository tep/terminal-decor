// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"testing"

	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
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
	want := "\x1b[38;5;59m[\x1b[38;5;214mЖ\x1b[38;5;59m:\x1b[3m\x1b[38;5;204mABC\x1b[38;5;59m]\x1b(B\x1b[m"

	if got := tmpl.Expand(vars); got != want {
		t.Errorf("tmpl.Format(%#v) == %q; Wanted %q", vars, got, want)
	} else {
		t.Logf("OUT: %q", got)
	}
}

func (t *Template) equal(os *series.Series) bool {
	return t.ss.Equal(os)
}
