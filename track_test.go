package midi

import (
	"os"
	"reflect"
	"testing"
)

func TestTrack_Add(t *testing.T) {
	type args struct {
		beatDelta float64
		e         *Event
	}
	tests := []struct {
		name         string
		ticksPerBeat uint16
		events       []*Event
		args         args
	}{
		{name: "nil event",
			ticksPerBeat: 0,
			events:       []*Event{},
			args:         args{beatDelta: 0, e: nil},
		},
		{name: "0 ticks per beat",
			ticksPerBeat: 0,
			events:       []*Event{},
			args:         args{beatDelta: 0, e: &Event{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Track{
				Events:       tt.events,
				ticksPerBeat: tt.ticksPerBeat,
			}
			tr.Add(tt.args.beatDelta, tt.args.e)
		})
	}
}

func TestTrack_AddAfterDelta(t *testing.T) {
	type args struct {
		ticks uint32
		e     *Event
	}
	tests := []struct {
		name     string
		events   []*Event
		args     args
		expected *Event
	}{
		{name: "nil event",
			events:   []*Event{},
			args:     args{ticks: 0, e: nil},
			expected: &Event{TimeDelta: 0},
		},
		{name: "next beat",
			events:   []*Event{},
			args:     args{ticks: 96, e: &Event{Copyright: "Foo"}},
			expected: &Event{TimeDelta: 96, Copyright: "Foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Track{
				Events: tt.events,
			}
			tr.AddAfterDelta(tt.args.ticks, tt.args.e)

			if len(tr.Events) > 0 {
				if last := tr.Events[len(tr.Events)-1]; !reflect.DeepEqual(last, tt.expected) {
					t.Errorf("Expected the last event to be %#v but got %#v", tt.expected, last)
				}
			}
		})
	}
}

func TestTrack_Tempo(t *testing.T) {
	tests := []struct {
		name  string
		track *Track
		want  int
	}{
		{name: "nil", track: nil, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.track.Tempo(); got != tt.want {
				t.Errorf("Track.Tempo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrack_Name(t *testing.T) {
	tests := []struct {
		name      string
		track     *Track
		trackName string
		want      string
	}{
		{name: "nil", track: nil, want: ""},
		{name: "no events", track: &Track{Events: []*Event{}}, want: ""},
		{name: "basic", track: &Track{Events: []*Event{TrackName("basic")}}, want: "basic"},
		{name: "Using the setter", track: &Track{}, trackName: "Extra", want: "Extra"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.trackName) > 0 {
				tt.track.SetName(tt.trackName)
			}
			if got := tt.track.Name(); got != tt.want {
				t.Errorf("Track.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrack_AbsoluteEvents(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
		want    []*AbsEv
	}{
		{
			name:    "unquantized track",
			fixture: "fixtures/unquantized2bars.mid",
			want: []*AbsEv{
				0: {Start: 0, Duration: 74, Vel: 78, MIDINote: 60},
				1: {Start: 0, Duration: 77, Vel: 88, MIDINote: 48},
				2: {Start: 0, Duration: 365, Vel: 82, MIDINote: 64},
				3: {Start: 0, Duration: 366, Vel: 84, MIDINote: 72},
				4: {Start: 0, Duration: 367, Vel: 84, MIDINote: 67},
				//
				5: {Start: 87, Duration: 81, Vel: 78, MIDINote: 55},
				6: {Start: 183, Duration: 90, Vel: 53, MIDINote: 60},
				7: {Start: 184, Duration: 80, Vel: 76, MIDINote: 52},
				8: {Start: 285, Duration: 79, Vel: 86, MIDINote: 55},
				//
				9:  {Start: 382, Duration: 278, Vel: 88, MIDINote: 72},
				10: {Start: 382, Duration: 280, Vel: 78, MIDINote: 64},
				11: {Start: 382, Duration: 281, Vel: 86, MIDINote: 67},
				//
				12: {Start: 383, Duration: 83, Vel: 82, MIDINote: 48},
				13: {Start: 383, Duration: 88, Vel: 66, MIDINote: 60},
				//
				14: {Start: 484, Duration: 78, Vel: 80, MIDINote: 52},
				15: {Start: 484, Duration: 79, Vel: 60, MIDINote: 60},
				//
				16: {Start: 581, Duration: 83, Vel: 80, MIDINote: 55},
				17: {Start: 583, Duration: 64, Vel: 62, MIDINote: 48},
				//
				18: {Start: 676, Duration: 74, Vel: 78, MIDINote: 64},
				19: {Start: 676, Duration: 75, Vel: 58, MIDINote: 60},
				20: {Start: 676, Duration: 77, Vel: 80, MIDINote: 67},
				21: {Start: 676, Duration: 77, Vel: 88, MIDINote: 72},
				//
				22: {Start: 677, Duration: 47, Vel: 70, MIDINote: 52},
			},
		},
		{
			name:    "c major scale",
			fixture: "fixtures/c-maj-scale.mid",
			want: []*AbsEv{
				0: {Start: 0, Duration: 479, Vel: 80, MIDINote: 48},
				1: {Start: 480, Duration: 479, Vel: 80, MIDINote: 50},
				2: {Start: 960, Duration: 479, Vel: 80, MIDINote: 52},
				3: {Start: 1440, Duration: 479, Vel: 80, MIDINote: 53},
				4: {Start: 1920, Duration: 479, Vel: 80, MIDINote: 55},
				5: {Start: 2400, Duration: 479, Vel: 80, MIDINote: 57},
				6: {Start: 2880, Duration: 479, Vel: 80, MIDINote: 59},
				7: {Start: 3360, Duration: 479, Vel: 80, MIDINote: 60},
			},
		},
	}
	for _, tt := range tests {
		r, err := os.Open(tt.fixture)
		if err != nil {
			t.Fatal(err)
		}
		defer r.Close()
		dec := NewDecoder(r)
		if err := dec.Parse(); err != nil {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			got := dec.Tracks[0].AbsoluteEvents()
			if len(got) != len(tt.want) {
				t.Fatalf("Expected %d events, but got %d", len(tt.want), len(got))
			}
			for i, ev := range got {
				if !reflect.DeepEqual(ev, tt.want[i]) {
					t.Errorf("Expected event %d to be %+v but got %+v", i, tt.want[i], ev)
				}
			}
		})
	}
}

func TestTrack_AbsoluteEvents_MultiTrack(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
		want    [][]*AbsEv
	}{
		{
			name:    "example format 0",
			fixture: "fixtures/example-format0.mid",
			want: [][]*AbsEv{
				0: {
					// It's unclear whether different channels in format 0 should end up
					// as different tracks or not.
					0: &AbsEv{Start: 0, Duration: 384, Vel: 96, MIDINote: 48},
					1: &AbsEv{Start: 0, Duration: 384, Vel: 96, MIDINote: 60},
					2: &AbsEv{Start: 96, Duration: 288, Vel: 64, MIDINote: 67},
					3: &AbsEv{Start: 192, Duration: 192, Vel: 32, MIDINote: 76},
				},
			},
		},
		{
			name:    "example format 1",
			fixture: "fixtures/example-format1.mid",
			want: [][]*AbsEv{
				0: {},
				1: {
					0: &AbsEv{Start: 192, Duration: 192, Vel: 32, MIDINote: 76},
				},
				2: {
					0: &AbsEv{Start: 96, Duration: 288, Vel: 64, MIDINote: 67},
				},
				3: {
					0: &AbsEv{Start: 0, Duration: 384, Vel: 96, MIDINote: 48},
					1: &AbsEv{Start: 0, Duration: 384, Vel: 96, MIDINote: 60},
				},
			},
		},
		{
			name:    "SATB multitrack",
			fixture: "fixtures/satb-basic.mid",
			want: [][]*AbsEv{
				0: {
					0: &AbsEv{Start: 0, Duration: 479, Vel: 80, MIDINote: 60},
					1: &AbsEv{Start: 480, Duration: 479, Vel: 80, MIDINote: 62},
					2: &AbsEv{Start: 960, Duration: 479, Vel: 80, MIDINote: 64},
					3: &AbsEv{Start: 1440, Duration: 479, Vel: 80, MIDINote: 65},
				},
				1: {
					0: &AbsEv{Start: 0, Duration: 479, Vel: 80, MIDINote: 69},
					1: &AbsEv{Start: 480, Duration: 479, Vel: 80, MIDINote: 71},
					2: &AbsEv{Start: 960, Duration: 479, Vel: 80, MIDINote: 72},
					3: &AbsEv{Start: 1440, Duration: 479, Vel: 80, MIDINote: 74},
				},
				2: {
					0: &AbsEv{Start: 0, Duration: 479, Vel: 80, MIDINote: 65},
					1: &AbsEv{Start: 0, Duration: 479, Vel: 80, MIDINote: 71},
					2: &AbsEv{Start: 480, Duration: 479, Vel: 80, MIDINote: 64},
					3: &AbsEv{Start: 480, Duration: 479, Vel: 80, MIDINote: 67},
					4: &AbsEv{Start: 960, Duration: 479, Vel: 80, MIDINote: 60},
					5: &AbsEv{Start: 960, Duration: 479, Vel: 80, MIDINote: 65},
					6: &AbsEv{Start: 1440, Duration: 479, Vel: 80, MIDINote: 60},
					7: &AbsEv{Start: 1440, Duration: 479, Vel: 80, MIDINote: 64},
				},
				3: {
					0: &AbsEv{Start: 0, Duration: 239, Vel: 80, MIDINote: 48},
					1: &AbsEv{Start: 240, Duration: 239, Vel: 80, MIDINote: 50},
					2: &AbsEv{Start: 480, Duration: 239, Vel: 80, MIDINote: 52},
					3: &AbsEv{Start: 720, Duration: 239, Vel: 80, MIDINote: 53},
					4: &AbsEv{Start: 960, Duration: 239, Vel: 80, MIDINote: 55},
					5: &AbsEv{Start: 1200, Duration: 239, Vel: 80, MIDINote: 52},
					6: &AbsEv{Start: 1440, Duration: 239, Vel: 80, MIDINote: 53},
					7: &AbsEv{Start: 1680, Duration: 239, Vel: 80, MIDINote: 57},
				},
			},
		},
	}
	for _, tt := range tests {
		r, err := os.Open(tt.fixture)
		if err != nil {
			t.Fatal(err)
		}
		defer r.Close()
		dec := NewDecoder(r)
		if err := dec.Parse(); err != nil {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			for trackNo, want := range tt.want {
				got := dec.Tracks[trackNo].AbsoluteEvents()
				if len(got) != len(want) {
					t.Fatalf("track #%d: expected %d events, but got %d", trackNo+1, len(want), len(got))
				}
				for i, ev := range got {
					if !reflect.DeepEqual(ev, want[i]) {
						t.Errorf("track #%d: expected event %d to be %+v but got %+v", trackNo+1, i, want[i], ev)
					}
				}
			}

		})
	}
}
