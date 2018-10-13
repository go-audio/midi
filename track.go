package midi

import (
	"bytes"
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
		data, err := e.Encode()
		if err != nil {
			return nil, err
		}
		if _, err := buff.Write(data); err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}
