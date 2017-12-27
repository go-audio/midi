package midi

import (
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
		name  string
		track *Track
		want  string
	}{
		{name: "nil", track: nil, want: ""},
		{name: "no events", track: &Track{Events: []*Event{}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.track.Name(); got != tt.want {
				t.Errorf("Track.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
