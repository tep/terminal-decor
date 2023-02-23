// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

func (d *Decorator) optimize(input *series.Series) *series.Series {
	output := series.New()
	active := series.New()

	d.debugf(1, "optimizing %d segments: %v", input.Len(), input.SegmentIDs())

	for itm := input.Front(); itm != nil; itm = itm.Next() {
		if d.debug > 1 {
			d.debugf(2, "OS: %s", itm)
			d.debugf(3, "    Output: %v", output.SegmentIDs())
			d.debugf(3, "    Active: %v", active.SegmentIDs())
		}

		switch itm.Action {
		case item.START:
			if prev := output.Back(); prev != nil && prev.Type == itm.Type && prev.Action == item.STOP {
				d.debugf(2, "    START immediately preceded by STOP: removing previous %s from output", prev)
				output.RemoveBack()
			}
			active.MoveToBackOrAppend(itm.Clone())

		case item.STOP:
			// 1. if prev.Type == itm.Type && prev.Action == START:
			// 		a. Remove prev from both 'active' AND output
			// 		b. continue
			if prev := output.Back(); prev != nil && prev.Type == itm.Type && prev.Action == item.START {
				d.debugf(2, "    STOP immediately preceded by START: undoing")
				active.RemoveBack()
				d.debugf(3, "    'active' done")
				output.RemoveBack()
				d.debugf(3, "    'output' done")
				continue
			}

			// 2. Find the final instance of itm.Type in 'active' and remove it
			d.debugf(2, "    Updating 'active' to remove most recent %s", itm.Type)
			rem := active.RemoveLast(itm.Type)
			d.debugf(3, "    Removed %s", rem)
		}

		output.Append(itm.Clone())

		if d.isAllOff(itm) && active.Len() > 0 {
			// If the above has turned everything off (i.e. 'sgr0') then
			// we'll need to turn all of the 'active' stuff back on.
			d.debugf(2, "    inserting active list after %q: %v", itm, active.SegmentIDs())
			d.debugf(3, "      before: %v", input.SegmentIDs())
			input.InsertAfterList(itm, active)
			d.debugf(3, "       after: %v", input.SegmentIDs())
		}
	}

	return output
}
