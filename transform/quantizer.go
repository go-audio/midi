package transform

import (
	"github.com/go-audio/midi"
	"github.com/go-audio/midi/grid"
	"sort"
)

// Quantizer's job is to modify a list of absolute events to align them as per
// the settings.
// See https://en.wikipedia.org/wiki/Quantization_(music)
type Quantizer struct {
	// GridRes is the grid resolution used to quantize notes
	GridRes grid.Res
	// QuantizationLevel is the amount of quantization to apply from 0.0 to 1.0
	// 0.0 means no quantization, 1.0 means exact snap, 0.5 is half way from
	// where the note current is and where it should snap.
	QuantizationLevel float32
	// Start indicates if the quantizer should quantize the start of an event (note on)
	Start bool
	// End indicates if the quantizer should quantize the end of an event (note off)
	End bool
	// MoveEndOnStartQ indicates that we want to keep the duration when quantizing the start of an event instead of shortening/expanding the duration due to the quantization. Note that this option will be applied before the end quantization is evaluated.
	MoveEndOnStartQ bool
}

// Quantize creates a copy of the passed events and quantizees the copy as per
// the quantizer settings.
func (q Quantizer) Quantize(events midi.AbsEvents, ppq uint16) midi.AbsEvents {
	cc := make(midi.AbsEvents, len(events))
	copy(cc, events)

	// distance between "grid lines"
	stepSize := int(q.GridRes.StepSize(ppq))
	halfStep := int(stepSize / 2)

	for i, ev := range cc {
		if remainder := ev.Start % int(stepSize); remainder != 0 {
			if remainder >= halfStep {
				cc[i].Start = (ev.Start / stepSize * stepSize) + stepSize
			} else {
				cc[i].Start = (ev.Start / stepSize) * stepSize
			}
		}
		// find the closest snap point
		// apply the quantization level to start if start
		// adjust the length of the note if MoveEndOnStartQ
	}

	// sort the events, first ones first
	sort.Slice(cc, func(i, j int) bool {
		if cc[i].Start < cc[j].Start {
			return true
		}
		if cc[i].Start > cc[j].Start {
			return false
		}
		// if both items start at the same time, use the duration to sort
		if cc[i].Duration < cc[j].Duration {
			return true
		}
		if cc[i].Duration > cc[j].Duration {
			return false
		}
		// if both items start at the same time, have the same duration, we sort by note
		if cc[i].MIDINote < cc[j].MIDINote {
			return true
		}
		return false
	})

	return cc
}
