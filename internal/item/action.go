// Copyright © 2023 Timothy E. Peoples

package item

type Action int

const (
	NONE Action = iota
	START
	STOP
)

func (a Action) String() string {
	switch a {
	case NONE:
		return ""
	case START:
		return "START"
	case STOP:
		return "STOP"
	default:
		return "<UNKNOWN>"
	}
}
