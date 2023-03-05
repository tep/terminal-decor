// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"testing"
)

func TestStrip(t *testing.T) {
	input := "@F(Grey37)[ABC:@I123@i]@f"
	want := "[ABC:123]"

	if got, err := Strip(input); err != nil || got != want {
		t.Errorf("Strip(%q) == (%q, %v); Wanted (%q, %v)", input, got, err, want, nil)
	}
}
