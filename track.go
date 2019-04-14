package midi

import (
	"bytes"
	"sort"
	"strings"
)

// Track
// <Track Chunk> = <chunk type><length><MTrk event>
// <MTrk event> = <delta-time><event>
//
type Track struct {
	Size         uint32
	Events       []*Event
	_name        string
	ticksPerBeat uint16
}

// Add schedules the passed event after x beats (relative to the previous event)
func (t *Track) Add(beatDelta float64, e *Event) {
	if t == nil || e == nil {
		return
	}
	if t.ticksPerBeat == 0 {
		t.ticksPerBeat = 96
	}
	e.TimeDelta = uint32(beatDelta * float64(t.ticksPerBeat))
	t.Events = append(t.Events, e)
	t.Size += uint32(len(EncodeVarint(e.TimeDelta))) + e.Size()
}

// AddAfterDelta schedules the passed event after x beats (relative to the previous event)
func (t *Track) AddAfterDelta(ticks uint32, e *Event) {
	if t == nil || e == nil {
		return
	}
	if t.ticksPerBeat == 0 {
		t.ticksPerBeat = 96
	}
	e.TimeDelta = ticks
	t.Events = append(t.Events, e)
	t.Size += uint32(len(EncodeVarint(e.TimeDelta))) + e.Size()
}

// Tempo returns the tempo of the track if set, 0 otherwise
func (t *Track) Tempo() int {
	if t == nil {
		return 0
	}
	tempoEvType := MetaByteMap["Tempo"]
	for _, ev := range t.Events {
		if ev.Cmd == tempoEvType {
			return int(ev.Bpm)
		}
	}
	return 0
}

// SetName sets the name on the track so it can be encoded when the track is serialized.
func (t *Track) SetName(name string) *Track {
	if t == nil {
		return t
	}
	t._name = name
	return t
}

// Name returns the name of the track if provided
func (t *Track) Name() string {
	if t == nil {
		return ""
	}
	if t._name != "" {
		return t._name
	}
	nameEvType := MetaByteMap["Sequence/Track name"]
	for _, ev := range t.Events {
		if ev.Cmd == nameEvType {
			// trim spaces and null bytes
			t._name = strings.TrimRight(strings.TrimSpace(ev.SeqTrackName), "\x00")
			return t._name
		}
	}
	return ""
}

// ChunkData converts the track and its events into a binary byte slice (chunk
// header included). If endTrack is set to true, the end track metadata will be
// added if not already present.
func (t *Track) ChunkData(endTrack bool) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	// time signature
	// TODO: don't have 4/4 36, 8 hardcoded
	buff.Write([]byte{0x00, 0xFF, 0x58, 0x04, 0x04, 0x02, 0x24, 0x08})
	// name event if name set
	if name := t.Name(); len(name) > 0 {
		t.Events = append([]*Event{TrackName(name)}, t.Events...)
	}

	if endTrack {
		if l := len(t.Events); l > 0 {
			if t.Events[l-1].Cmd != MetaByteMap["End of Track"] {
				t.Add(0, EndOfTrack())
			}
		}
	}
	for _, e := range t.Events {
		if _, err := buff.Write(e.Encode()); err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

// AbsoluteEvents converts a midi track into a list of absolute events. The
// events are sorted by start time, duration (shorter notes first) and finally
// order of notes (lower notes first)
func (t *Track) AbsoluteEvents() AbsEvents {
	totalDuration := uint32(0) // in ticks

	absEvs := map[int][]*AbsEv{}
	curEvsStart := map[string]*Event{}

	for _, ev := range t.Events {
		totalDuration += ev.TimeDelta
		pitch := int(ev.Note)
		n := NoteToName(pitch)
		// fmt.Printf("%s %s @ %.2f beats\n", n, EventMap[ev.MsgType], float64(totalDuration))

		if _, ok := absEvs[pitch]; !ok {
			absEvs[pitch] = []*AbsEv{}
		}
		if _, ok := curEvsStart[n]; !ok {
			curEvsStart[n] = nil
		}
		switch ev.MsgType {
		case EventByteMap["NoteOn"]:
			if curEvsStart[n] != nil {
				// end previous note (weird but sure)
				start := uint32(curEvsStart[n].AbsTicks)
				absEvs[pitch] = append(absEvs[pitch], &AbsEv{
					MIDINote: pitch,
					Start:    int(curEvsStart[n].AbsTicks),
					Duration: int(totalDuration - start),
					Vel:      int(ev.Velocity),
				},
				)
			}
			curEvsStart[n] = ev
		case EventByteMap["NoteOff"]:
			absEvs[pitch] = append(absEvs[pitch],
				&AbsEv{
					MIDINote: pitch,
					Start:    int(curEvsStart[n].AbsTicks),
					Duration: int(ev.AbsTicks) - int(curEvsStart[n].AbsTicks),
					Vel:      int(curEvsStart[n].Velocity),
				})
			curEvsStart[n] = nil
		}
	}
	events := []*AbsEv{}
	for _, ev := range absEvs {
		events = append(events, ev...)
	}

	// sort the events, first ones first
	sort.Slice(events, func(i, j int) bool {
		if events[i].Start < events[j].Start {
			return true
		}
		if events[i].Start > events[j].Start {
			return false
		}
		// if both items start at the same time, use the duration to sort
		if events[i].Duration < events[j].Duration {
			return true
		}
		if events[i].Duration > events[j].Duration {
			return false
		}
		// if both items start at the same time, have the same duration, we sort by note
		if events[i].MIDINote < events[j].MIDINote {
			return true
		}
		return false
	})
	return events
}
