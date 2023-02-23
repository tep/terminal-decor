// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"testing"

	"toolman.org/terminal/decor/internal/item"
)

func TestDecorator(t *testing.T) {
	item.ResetID()
	d, err := Load("xterm-256color")
	if err != nil {
		t.Fatal(err)
	}

	// d.debug = 1

	var (
		input = "@F(Grey37)[ABC:@I123@i]@f"
		want  = "\x1b[38;5;59m[ABC:\x1b[3m123\x1b[23m]\x1b(B\x1b[m"
	)

	if got, err := d.Format(input); err != nil || got != want {
		t.Errorf("Format(%q) == (%q, %v); Wanted (%q, nil)", input, got, err, want)
	} else {
		t.Logf("Got: %q", got)
	}
}

func TestTemplate(t *testing.T) {
	item.ResetID()
	d, err := Load("xterm-256color")
	if err != nil {
		t.Fatal(err)
	}

	// d.debug = 2
	input := "@F(Grey37)[${Glyph}:@I${Key}@i]@f"
	want := `001:START:FGCOLOR:"Grey37" + 002:TEXT:"[" + 003:VAR:"Glyph" + 004:TEXT:":" + 005:START:ITALIC + 006:VAR:"Key" + 007:STOP:ITALIC + 008:TEXT:"]" + 009:STOP:FGCOLOR`

	var tmpl *Template
	if tmpl, err = d.Template(input); err != nil || tmpl.String() != want {
		t.Fatalf("d.Template(%q) == (%q, %v);\nWanted (%q, nil)", input, tmpl, err, want)
	}

	vars := map[string]string{"Glyph": "@F<Orange1>Ж@f", "Key": "@F{204}ABC@f"}
	want = "\x1b[38;5;59m[\x1b[38;5;214mЖ\x1b[38;5;59m:\x1b[3m\x1b[38;5;204mABC\x1b[38;5;59m]\x1b(B\x1b[m"

	if got, err := tmpl.Format(vars); err != nil || got != want {
		t.Errorf("tmpl.Format(%#v) == (%q, %v); Wanted (%q, nil)", vars, got, err, want)
	} else {
		t.Logf("OUT: %q", got)
	}
}
