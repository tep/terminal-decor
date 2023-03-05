// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	d, err := xterm256Decorator()
	if err != nil {
		t.Fatal(err)
	}

	// d.debug = 1

	var (
		input = "@F(Grey37)[ABC:@I123@i]@f"
		want  = "\x1b[38;5;59m[ABC:\x1b[3m123\x1b[23m]\x1b[39m"
	)

	if got, err := d.Format(input); err != nil || got != want {
		t.Errorf("Format(%q) == (%q, %v); Wanted (%q, nil)", input, got, err, want)
	} else {
		t.Logf("Got: %q", got)
	}
}
