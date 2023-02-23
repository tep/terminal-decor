// Copyright © 2023 Timothy E. Peoples

package series

import (
	"testing"

	"toolman.org/terminal/decor/internal/item"
)

func TestSegment(t *testing.T) {
	s := New()

	want := Build(
		item.FGColorItem("Grey37"),
		item.TextItem("["),
		item.VarItem("Glyph"),
		item.TextItem(":"),
		item.AttrItem('I'),
		item.VarItem("Key"),
		item.AttrItem('i'),
		item.TextItem("]"),
		item.AttrItem('f'),
	)

	input := "@F(Grey37)[${Glyph}:@I${Key}@i]@f"
	if err := s.Parse(input); err != nil {
		t.Errorf("s.Parse(%q) error: %v", input, err)
	} else if !s.Equal(want) {
		t.Errorf("s.Parse(%q) -> >>%s<< Wanted >>%s<<", input, s, want)
	} else {
		t.Logf("OK: s.Parse(%q) -> >>%s<<", input, s)
	}
}
