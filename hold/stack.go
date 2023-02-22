// Copyright © 2023 Timothy E. Peoples

package termcolor

// import (
// 	"fmt"
// 	"strings"
// )
//
// type segmentStack struct {
// 	active []segment
// 	*Term
// }
//
// func (t *Term) stack() *segmentStack {
// 	return &segmentStack{Term: t}
// }
//
// func (ss *segmentStack) String() string {
// 	if ss == nil {
// 		return "<nil>"
// 	}
//
// 	if len(ss.active) == 0 {
// 		return "<empty>"
// 	}
//
// 	parts := make([]string, len(ss.active))
// 	for i, s := range ss.active {
// 		parts[i] = s.String()
// 	}
// 	return strings.Join(parts, " ")
// }
//
// func (ss *segmentStack) start(sgmt segment) string {
// 	var newlist []segment
//
// 	out := ss.enterCode(sgmt)
//
// 	// Here we iterate over the receiver's current list of active attributes
// 	// comparing each to our 'sgmt' argument. Those that are not the same as
// 	// 'sgmt' are copied into a new "active" list. If any of them are equal to
// 	// 'sgmt', we'll disable sending its "enter" code.
// 	//
// 	// tl/dr; If 'sgmt' is already active, disable sending its "enter" code
// 	//        and move it to the end of the active list.
// 	for _, as := range ss.active {
// 		if as.equal(&sgmt) {
// 			out = ""
// 		} else {
// 			newlist = append(newlist, as)
// 		}
// 	}
//
// 	// 'sgmt' *always* goes to the end of the "active" list.
// 	newlist = append(newlist, sgmt)
//
// 	ss.active = newlist
//
// 	if DEBUG {
// 		msg := out
// 		if msg == "" {
// 			msg = "<already-active>"
// 		}
// 		fmt.Printf("++ %q\n", msg)
// 	}
//
// 	return out
// }
//
// func (ss *segmentStack) stop(sgmt segment) string {
// 	x := ss.lastIndexType(sgmt.stype)
// 	if x == -1 {
// 		// sgmt.stype isn't actually active so there's nothing to do here.
// 		return ""
// 	}
//
// 	// remove the *last* instance of sgmt.stype from ss.active
// 	rem := ss.active[x+1:]
// 	ss.active = append(ss.active[:x], rem...)
//
// 	// add the exit code for this sgmt...
// 	out := ss.exitCode(sgmt)
//
// 	// ...but, if we've turned *everything* off, we need to re-add
// 	// everything that's still active.
// 	if out == ss.sgr0 {
// 		for _, as := range ss.active {
// 			out += ss.enterCode(as)
// 		}
// 	}
//
// 	if DEBUG {
// 		fmt.Printf("-- %q\n", out)
// 	}
//
// 	return out
// }
//
// func (ss *segmentStack) lastIndexType(st segType) int {
// 	x := -1
// 	for i, as := range ss.active {
// 		if as.stype == st {
// 			x = i
// 		}
// 	}
// 	return x
// }
