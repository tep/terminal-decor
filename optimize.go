// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"toolman.org/terminal/decor/internal/item"
	"toolman.org/terminal/decor/internal/series"
)

func (d *Decorator) optimize(input *series.Series) *series.Series {
	output := series.New()
	active := series.New()

	d.debugf(1, "optimizing %d items: %v", input.Len(), input.ItemIDs())
	if d.debug > 1 {
		d.debugf(2, "    %s", input.String())
	}

	var restorePoints seriesStack

	for itm := input.Front(); itm != nil; itm = itm.Next() {
		if d.debug > 1 {
			d.debugf(2, "ITEM: %s", itm)
			d.debugf(3, "     Input: %v", input.ItemIDs())
			d.debugf(3, "    Output: %v", output.ItemIDs())
			d.debugf(3, "    Active: %v", active.ItemIDs())
		}

		switch itm.Action {
		case item.START:
			if itm.Type == item.SAVE {
				d.debugf(2, "    Saving Active: %v", active.ItemIDs())
				restorePoints.push(active)
				continue
			}

			if prev := output.Back(); prev != nil && prev.Type == itm.Type && prev.Action == item.STOP {
				d.debugf(2, "    START immediately preceded by STOP: removing previous %s from output", prev)
				output.RemoveBack()
			}
			active.MoveToBackOrAppend(itm.Clone())

		case item.STOP:
			if itm.Type == item.SAVE {
				if rp := restorePoints.pop(); rp != nil {
					active = rp
					if active.Len() == 0 {
						d.debugf(2, "    Restoring Empty Active")
					} else {
						d.debugf(2, "    Restoring Active: %v", rp.ItemIDs())
						input.InsertAfterList(itm, active)
						p := itm.Prev()
						input.Remove(itm)
						itm = p
					}
				}
				continue
			}

			// 1. if prev.Type == itm.Type && prev.Action == START --> Remove prev from output
			if prev := output.Back(); prev != nil && prev.Type == itm.Type && prev.Action == item.START {
				d.debugf(2, "    STOP (%s) immediately preceded by START (%s): undoing", itm, prev)
				output.RemoveBack()
				d.debugf(3, "    'output' done")
			}

			// 2. Find the final instance of itm.Type in 'active' and remove it
			d.debugf(2, "    Updating 'active' and 'output' to remove most recent START:%s", itm.Type)
			rem := active.RemoveLast(itm.Type)
			d.debugf(3, "    Removed %s from active", rem)
		}

		output.Append(itm.Clone())

		if d.isAllOff(itm) && active.Len() > 0 {
			// If the above has turned everything off (i.e. 'sgr0') then
			// we'll need to turn all of the 'active' stuff back on.
			d.debugf(2, "    inserting active list after %q: %v", itm, active.ItemIDs())
			d.debugf(3, "      before: %v", input.ItemIDs())
			input.InsertAfterList(itm, active)
			d.debugf(3, "       after: %v", input.ItemIDs())
		}
	}

	return output
}

type seriesStack struct {
	stack []*series.Series
}

func (ss *seriesStack) push(s *series.Series) {
	ss.stack = append([]*series.Series{s.Clone()}, ss.stack...)
}

func (ss *seriesStack) pop() *series.Series {
	if len(ss.stack) == 0 {
		return nil
	}

	s := ss.stack[0]
	ss.stack = ss.stack[1:]

	return s
}
