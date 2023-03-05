// Copyright © 2023 Timothy E. Peoples

package item

const (
	EMPTY Type = iota
	TEXT
	ERROR
	VAR
	SAVE

	attrs

	BOLD
	UNDERLINE
	ITALIC
	FGCOLOR
	BGCOLOR
)

type Type int

func (t Type) String() string {
	switch t {
	case TEXT:
		return "TEXT"
	case ERROR:
		return "ERROR"
	case VAR:
		return "VAR"
	case SAVE:
		return "SAVE"
	case BOLD:
		return "BOLD"
	case UNDERLINE:
		return "UNDERLINE"
	case ITALIC:
		return "ITALIC"
	case FGCOLOR:
		return "FGCOLOR"
	case BGCOLOR:
		return "BGCOLOR"
	default:
		return "<UNKNOWN>"
	}
}
