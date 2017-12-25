package midi

import (
	"reflect"
	"testing"
)

func Test_TickPosition(t *testing.T) {
	tests := []struct {
		name     string
		absTicks uint64
		ppq      uint32
		want     Position
	}{
		{name: "One quantized", absTicks: 0, ppq: 96, want: Position{Bar: 0, Beat: 0, Div: 0, Ticks: 0}},
		{name: "One unquantized", absTicks: 35, ppq: 96, want: Position{Bar: 0, Beat: 0, Div: 1, Ticks: 11}},
		{name: "75", absTicks: 75, ppq: 96, want: Position{Bar: 0, Beat: 0, Div: 3, Ticks: 3}},
		{name: "ppq", absTicks: 96, ppq: 96, want: Position{Bar: 0, Beat: 1, Div: 0, Ticks: 0}},
		{name: "250", absTicks: 250, ppq: 96, want: Position{Bar: 0, Beat: 2, Div: 2, Ticks: 10}},
		{name: "high ppq", absTicks: 1250, ppq: 480, want: Position{Bar: 0, Beat: 2, Div: 2, Ticks: 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TickPosition(tt.absTicks, tt.ppq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TickPosition() = %v, want %v", got, tt.want)
			}
			ev := &Event{AbsTicks: tt.absTicks}
			if got := ev.Position(tt.ppq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Event.Position() = %v, want %v", got, tt.want)
			}
		})
	}
	var ev *Event
	if got := ev.Position(96); !reflect.DeepEqual(got, Position{}) {
		t.Errorf("Event.Position() of a nil element expected a blank position but got %v", got)
	}
}

func TestPosition_ToTicks(t *testing.T) {
	tests := []struct {
		name     string
		absTicks uint64
		ppq      uint32
		pos      Position
	}{
		{name: "One quantized", absTicks: 0, ppq: 96, pos: Position{Bar: 0, Beat: 0, Div: 0, Ticks: 0}},
		{name: "One unquantized", absTicks: 35, ppq: 96, pos: Position{Bar: 0, Beat: 0, Div: 1, Ticks: 11}},
		{name: "75", absTicks: 75, ppq: 96, pos: Position{Bar: 0, Beat: 0, Div: 3, Ticks: 3}},
		{name: "ppq", absTicks: 96, ppq: 96, pos: Position{Bar: 0, Beat: 1, Div: 0, Ticks: 0}},
		{name: "250", absTicks: 250, ppq: 96, pos: Position{Bar: 0, Beat: 2, Div: 2, Ticks: 10}},
		{name: "high ppq", absTicks: 1250, ppq: 480, pos: Position{Bar: 0, Beat: 2, Div: 2, Ticks: 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pos.ToTicks(tt.ppq); !reflect.DeepEqual(got, tt.absTicks) {
				t.Errorf("ToTicks() = %v, want %v", got, tt.absTicks)
			}
		})
	}
}
