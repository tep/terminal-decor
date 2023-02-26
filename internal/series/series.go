// Copyright © 2023 Timothy E. Peoples

package series

import (
	"container/list"
	"fmt"
	"os"
	"strings"

	"toolman.org/terminal/decor/internal/item"
)

type Series struct {
	clist *list.List
}

func New() *Series {
	return &Series{clist: list.New()}
}

func Build(items ...*item.Item) *Series {
	s := New()

	for _, itm := range items {
		s.Append(itm.Detach())
	}

	return s
}

func (s *Series) Len() int {
	if s == nil || s.clist == nil {
		return 0
	}

	return s.clist.Len()
}

func (s *Series) Equal(o *Series) bool {
	slen, olen := s.Len(), o.Len()

	switch {
	case slen == 0 && olen == 0:
		return true
	case slen == 0 || olen == 0:
		return false
	case slen != olen:
		return false
	}

	for si, oi := s.Front(), o.Front(); si != nil && oi != nil; si, oi = si.Next(), oi.Next() {
		if !si.Equal(oi) {
			return false
		}
	}

	return true
}

func (s *Series) String() string {
	if s.Len() == 0 {
		return ""
	}

	var out []string

	for itm := s.Front(); itm != nil; itm = itm.Next() {
		out = append(out, strings.Replace(itm.String(), "+", "++", -1))
	}

	return strings.Join(out, " + ")
}

func (s *Series) Front() *item.Item {
	if s != nil && s.clist != nil {
		if e := s.clist.Front(); e != nil {
			if itm, ok := e.Value.(*item.Item); ok {
				return itm
			}
		}
	}

	return nil
}

func (s *Series) Back() *item.Item {
	if s != nil && s.clist != nil {
		if e := s.clist.Back(); e != nil {
			if itm, ok := e.Value.(*item.Item); ok {
				return itm
			}
		}
	}

	return nil
}

func (s *Series) Append(itm *item.Item) *Series {
	if s == nil || s.clist == nil || itm == nil {
		return nil
	}

	e := s.clist.PushBack(itm)
	if !itm.Bind(e) {
		fmt.Fprintf(os.Stderr, "Item %q failed to bind", itm)
	}
	return s
}

func (s *Series) AppendList(other *Series) *Series {
	for itm := other.Front(); itm != nil; itm = itm.Next() {
		s.Append(itm.Detach())
	}
	return s
}

func (s *Series) Clone() *Series {
	if s.Len() == 0 {
		return nil
	}

	ns := New()

	return ns.AppendList(s)
}

func (s *Series) ItemIDs() []string {
	if s == nil || s.Len() == 0 {
		return nil
	}

	ids := make([]string, s.Len())

	for it, itm := 0, s.Front(); itm != nil; it, itm = it+1, itm.Next() {
		ids[it] = fmt.Sprintf("%03d", itm.ID)
	}

	return ids
}

// MoveToBackOrAppend examines the existing elements in its receiver's list
// and, if one is found that is equal to the given Item, it is moved to
// the back of the list. Otherwise, itm is appended to the reciever's list.
func (s *Series) MoveToBackOrAppend(itm *item.Item) *Series {
	for it := s.Front(); it != nil; it = it.Next() {
		if it.Equal(itm) {
			s.clist.MoveToBack(it.Element())
			return s
		}
	}

	return s.Append(itm)
}

// RemoveLast iterates over the receiver's list from back to front until it
// finds an Item of the given type, removes it, then returns the removed
// *Item. If no Item is found for the give Type, a nil pointer is
// returned.
func (s *Series) RemoveLast(itype item.Type) *item.Item {
	if s == nil {
		return nil
	}

	for it := s.Back(); s != nil; it = it.Prev() {
		if it.Type == itype {
			s.clist.Remove(it.Element())
			return it
		}
	}

	return nil
}

// RemoveBack removes the element at the back of the reciever's list
// and returns it. If this list is empty a nil pointer is returned.
func (s *Series) RemoveBack() *item.Item {
	if s == nil {
		return nil
	}

	back := s.Back()
	if back == nil {
		return nil
	}

	s.clist.Remove(back.Element())

	return back
}

func (s *Series) InsertAfterList(itm *item.Item, os *Series) {
	ele := itm.Element()
	for it := os.Front(); it != nil; it = it.Next() {
		ic := it.Clone()
		ele = s.clist.InsertAfter(ic, ele)
		ic.Bind(ele)
	}
}
