// Copyright © 2023 Timothy E. Peoples

package item

import (
	"container/list"
	"fmt"
	"strings"
	"sync/atomic"
)

var idSequence uint64

func nextID() uint64 { return atomic.AddUint64(&idSequence, 1) }

func ResetID() { atomic.StoreUint64(&idSequence, 0) }

type Item struct {
	ID      uint64
	Type    Type
	Action  Action
	Text    string
	element *list.Element
}

func AttrItem(c byte) *Item {
	switch c {
	case 'B':
		return StartItem(BOLD)
	case 'F':
		return StartItem(FGCOLOR)
	case 'I':
		return StartItem(ITALIC)
	case 'K':
		return StartItem(BGCOLOR)
	case 'U':
		return StartItem(UNDERLINE)
	case 'b':
		return StopItem(BOLD)
	case 'f':
		return StopItem(FGCOLOR)
	case 'i':
		return StopItem(ITALIC)
	case 'k':
		return StopItem(BGCOLOR)
	case 'u':
		return StopItem(UNDERLINE)
	default:
		return &Item{ID: nextID()}
	}
}

func TextItem(text string) *Item { return newItem(TEXT, NONE, text) }
func VarItem(name string) *Item  { return newItem(VAR, NONE, name) }
func ErrItem(err error) *Item    { return newItem(ERROR, NONE, fmt.Sprintf("<err:%v>", err)) }

func FGColorItem(color string) *Item {
	itm := StartItem(FGCOLOR)
	itm.Text = color
	return itm
}

func BGColorItem(color string) *Item {
	itm := StartItem(FGCOLOR)
	itm.Text = color
	return itm
}

func ErrItemf(msg string, args ...any) *Item { return ErrItem(fmt.Errorf(msg, args...)) }

func StartItem(stype Type) *Item { return newItem(stype, START, "") }
func StopItem(stype Type) *Item  { return newItem(stype, STOP, "") }

func newItem(stype Type, action Action, text string) *Item {
	s := &Item{
		ID:     nextID(),
		Type:   stype,
		Action: action,
		Text:   text,
	}

	// s.element = &list.Element{Value: s}

	return s
}

func (i *Item) Element() *list.Element {
	if i == nil {
		return nil
	}
	return i.element
}

func (i *Item) Bind(e *list.Element) bool {
	if i != nil && e != nil {
		if es, ok := e.Value.(*Item); ok && es == i {
			i.element = e
			return true
		}
	}
	return false
}

func (i *Item) Next() *Item {
	if i != nil && i.element != nil {
		if n := i.element.Next(); n != nil {
			if v, ok := n.Value.(*Item); ok {
				return v
			}
		}
	}

	return nil
}

func (i *Item) Prev() *Item {
	if i != nil && i.element != nil {
		if p := i.element.Prev(); p != nil {
			if v, ok := p.Value.(*Item); ok {
				return v
			}
		}
	}

	return nil
}

// Clone returns a copy of its reciever that is still attached to its
// containing SegmentList. Use the Detach method for a detached copy.
func (i *Item) Clone() *Item {
	if i == nil {
		return nil
	}

	clone := *i
	return &clone
}

func (i *Item) Detach() *Item {
	d := i.Clone()
	if d != nil {
		d.element = nil
	}
	return d
}

func (i *Item) String() string {
	if i == nil {
		return "<nil>"
	}

	parts := []string{fmt.Sprintf("%03d", i.ID)}

	if i.Type > attrs {
		parts = append(parts, fmt.Sprintf("%s", i.Action))
	}

	parts = append(parts, fmt.Sprintf("%s", i.Type))

	if i.Text != "" {
		parts = append(parts, fmt.Sprintf("%q", i.Text))
	}

	return strings.Join(parts, ":")
}

func (i *Item) Equal(o *Item) bool {
	switch {
	case i == nil && o == nil:
		return true
	case i == nil || o == nil:
		return false
	default:
		return i.Type == o.Type && i.Action == o.Action && i.Text == o.Text
	}
}

func (i *Item) IsAttrOn() bool {
	if i == nil {
		return false
	}
	return i.Action == START
}

func (i *Item) IsAttrOff() bool {
	if i == nil || i.Action != STOP {
		return false
	}

	switch i.Type {
	case BOLD, FGCOLOR, BGCOLOR:
		return true
	default:
		return false
	}
}
