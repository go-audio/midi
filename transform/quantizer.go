package transform

import (
	"github.com/go-audio/midi"
	"github.com/go-audio/midi/grid"
	"sort"
)

// One8thQuantizer is a default 1/8th grid 100% quantizer for the start of the
// notes.
var One8thQuantizer = Quantizer{
	GridRes:           grid.One8,
	QuantizationLevel: 1.0,
	Start:             true,
}

// One16thQuantizer is a default 1/16th grid 100% quantizer for the start of the
// notes.
var One16thQuantizer = Quantizer{
	GridRes:           grid.One16,
	QuantizationLevel: 1.0,
	Start:             true,
}

// One32thQuantizer is a default 1/32th grid 100% quantizer for the start of the
// notes.
var One32thQuantizer = Quantizer{
	GridRes:           grid.One32,
	QuantizationLevel: 1.0,
	Start:             true,
}

// Quantizer's job is to modify a list of absolute events to align them as per
// the settings.
// See https://en.wikipedia.org/wiki/Quantization_(music)
type Quantizer struct {
	// GridRes is the grid resolution used to quantize notes
	GridRes grid.Res
	// QuantizationLevel is the amount of quantization to apply from 0.0 to 1.0
	// 0.0 means no quantization, 1.0 means exact snap, 0.5 is half way from
	// where the note current is and where it should snap.
	QuantizationLevel float64
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
	cc := events.Copy()

	// distance between "grid lines"
	stepSize := int(q.GridRes.StepSize(ppq))
	halfStep := int(stepSize / 2)

	// enforce a valid value for the quantization level
	if q.QuantizationLevel < 0 {
		q.QuantizationLevel = 0
	} else if q.QuantizationLevel > 1.0 {
		q.QuantizationLevel = 1.0
	}
	// if the quantization level is set to zero, we don't need to quantize and
	// can return the copy right away
	if q.QuantizationLevel == 0 {
		return cc
	}

	var snapPoint int
	var delta int
	var fullStepPos int

	if q.Start {
		for i, ev := range cc {
			// snap start point
			if remainder := ev.Start % stepSize; remainder != 0 {
				// decide if the note should be moved sooner or later depending
				// on how close it is from the nearby steps
				fullStepPos = (ev.Start / stepSize * stepSize)
				if remainder >= halfStep {
					if q.QuantizationLevel == 1.0 {
						cc[i].Start = fullStepPos + stepSize
					} else {
						// @ 100%
						snapPoint = fullStepPos + stepSize
						// apply the quantization level to the distance to snap
						delta = int(float64(snapPoint-cc[i].Start) * q.QuantizationLevel)
						cc[i].Start = cc[i].Start + delta
					}
				} else {
					if q.QuantizationLevel == 1.0 {
						cc[i].Start = fullStepPos
					} else {

						delta = int(float64(cc[i].Start-fullStepPos) * q.QuantizationLevel)
						cc[i].Start = cc[i].Start - delta
					}
				}
			}
		}
	}
	if q.End {
		for i, ev := range cc {
			// TODO: apply the quantization level
			// snap end point by adjusting the duration
			if remainder := ev.Duration % int(stepSize); remainder != 0 {
				if remainder >= halfStep {
					cc[i].Duration = (ev.Duration / stepSize * stepSize) + stepSize
				} else {
					cc[i].Duration = (ev.Duration / stepSize) * stepSize
				}
			}
		}
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
