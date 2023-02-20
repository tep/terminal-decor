// Copyright © 2023 Timothy E. Peoples

package termcolor

import (
	"fmt"
	"strings"
)

type segmentStack struct {
	active []segment
	*Term
}

func (t *Term) stack() *segmentStack {
	return &segmentStack{Term: t}
}

func (ss *segmentStack) String() string {
	parts := make([]string, len(ss.active))
	for i, s := range ss.active {
		parts[i] = s.String()
	}
	return strings.Join(parts, " ")
}

func (ss *segmentStack) start(sgmt segment) string {
	orig := ss.active
	ss.active = nil
	out := ss.enterCode(sgmt)

	for _, as := range orig {
		if as.stype == sgmt.stype {
			// There's already an active attr of this type so we'll skip sending
			// the enter code iff it's the exact same attr. Otherwize, this one
			// will simply override the previous one.
			if as.equal(&sgmt) {
				out = ""
			}
		} else {
			ss.active = append(ss.active, as)
		}
	}

	// ...but, no matter what, we'll append this segment to the end of the
	// active list.
	ss.active = append(ss.active, sgmt)

	if DEBUG {
		msg := out
		if msg == "" {
			msg = "<already-active>"
		}
		fmt.Printf("%-30q", msg)
	}

	return out
}

func (ss *segmentStack) stop(sgmt segment) string {
	x := ss.lastIndexType(sgmt.stype)
	if x == -1 {
		// sgmt.stype isn't actually active so there's nothing to do here.
		return ""
	}

	// remove the *last* instance of sgmt.stype from ss.active
	rem := ss.active[x+1:]
	ss.active = append(ss.active[:x], rem...)

	// add the exit code for this sgmt...
	out := ss.exitCode(sgmt)

	// ...but, if we've turned *everything* off, we need to re-add
	// everything that's still active.
	if out == ss.sgr0 {
		for _, as := range ss.active {
			out += ss.enterCode(as)
		}
	}

	if DEBUG {
		fmt.Printf("%-30q", out)
	}

	return out
}

func (ss *segmentStack) lastIndexType(st segType) int {
	x := -1
	for i, as := range ss.active {
		if as.stype == st {
			x = i
		}
	}
	return x
}
