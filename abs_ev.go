package midi

import (
	"sort"
)

// AbsEv is the representation of an absolute event
type AbsEv struct {
	Start    int // in ticks
	Duration int // in ticks
	Vel      int // 0 - 127
	MIDINote int
}

// End  returns the calculated end position (in ticks) of the event
func (ev *AbsEv) End() int {
	if ev == nil {
		return 0
	}
	return int(ev.Start) + int(ev.Duration)
}

// AbsEvents is a collection of absolute events
type AbsEvents []*AbsEv

// Copy returns a deep copy of the absolute event points
func (evs AbsEvents) Copy() AbsEvents {
	cc := make(AbsEvents, len(evs))
	for i, ev := range evs {
		cc[i] = &AbsEv{
			Start:    ev.Start,
			Duration: ev.Duration,
			Vel:      ev.Vel,
			MIDINote: ev.MIDINote,
		}
	}
	return cc
}

// ToMIDITrack adds a MIDI track to the encoder, make sure that the encoder uses
// the same PPQ as the events.
func (evs AbsEvents) ToMIDITrack(e *Encoder) *Track {
	tr := e.NewTrack()
	// convert our *AbsEv-s to *Event-s
	events := []*Event{}
	for _, ev := range evs {
		on := NoteOn(0, ev.MIDINote, ev.Vel)
		on.AbsTicks = uint64(ev.Start)
		events = append(events, on)
		off := NoteOff(0, ev.MIDINote)
		off.AbsTicks = uint64(ev.End())
		events = append(events, off)
	}
	// sort notes by starting time
	sort.Slice(events, func(i, j int) bool {
		// if the 2 events happen at the same time, the note off should have the priority.
		if events[i].AbsTicks == events[j].AbsTicks {
			return events[i].MsgType < events[j].MsgType
		}
		return events[i].AbsTicks < events[j].AbsTicks
	})

	// add to the actual track
	var lastTick uint64
	for _, e := range events {
		delta := e.AbsTicks - lastTick
		tr.AddAfterDelta(uint32(delta), e)
		lastTick = e.AbsTicks
	}

	return tr
}
