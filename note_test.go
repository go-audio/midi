package midi

import (
	"fmt"
	"testing"
)

var epsilon = 0.00000001

func floatEquals(a, b float64) bool {
	if (a-b) < epsilon && (b-a) < epsilon {
		return true
	}
	return false
}

func TestKeyToInt(t *testing.T) {
	testCases := []struct {
		key    string
		octave int
		n      int
	}{
		{"C", 0, 24},
		{"C#", 0, 25},
		{"Db", 0, 25},
		{"D", 0, 26},
		{"E", 0, 28},
		{"F", 0, 29},
		{"G", 0, 31},
		{"A", 0, 33},
		{"B", 0, 35},
		{"C", 3, 60},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s %d", tc.key, tc.octave), func(t *testing.T) {
			if o := KeyInt(tc.key, tc.octave); o != tc.n {
				t.Fatalf("expected %s %d -> %d\ngot\n%d\n", tc.key, tc.octave, tc.n, o)
			}
		})
	}
}

func TestKeyFreq(t *testing.T) {
	testCases := []struct {
		key    string
		octave int
		freq   float64
	}{
		{"C", 0, 32.70319566257483},
		{"D", 0, 36.70809598967595},
		{"E", 0, 41.20344461410875},
		{"F", 0, 43.653528929125486},
		{"G", 0, 48.99942949771867},
		{"A", 0, 55.00},
		{"B", 0, 61.7354126570155},
		{"A", 3, 440.00},
		{"D#", 3, 311.12698372208087},
		{"C", 3, 261.6255653005986},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s %d", tc.key, tc.octave), func(t *testing.T) {
			if o := KeyFreq(tc.key, tc.octave); !floatEquals(o, tc.freq) {
				t.Fatalf("expected %s %d -> %v\ngot\n%v\n", tc.key, tc.octave, tc.freq, o)
			}
		})
	}
}

func TestFreqToNote(t *testing.T) {
	testCases := []struct {
		note int
		freq float64
	}{
		{KeyInt("C", 0), 32.70319566257483},
		{KeyInt("D", 0), 36.70809598967595},
		{KeyInt("E", 0), 41.20344461410875},
		{KeyInt("F", 0), 43.653528929125486},
		{KeyInt("G", 0), 48.99942949771867},
		{KeyInt("A", 0), 55.00},
		{KeyInt("B", 0), 61.7354126570155},
		{KeyInt("A", 3), 440.00},
		{KeyInt("D#", 3), 311.12698372208087},
		{KeyInt("C", 3), 261.6255653005986},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%2.f", tc.freq), func(t *testing.T) {
			if note := FreqToNote(tc.freq); note != tc.note {
				t.Fatalf("expected freq %v -> %v\ngot\n%v\n", tc.freq, tc.note, note)
			}
		})
	}
}

func TestNoteOctave(t *testing.T) {
	tests := []struct {
		name string
		note int
		want int
	}{
		{name: "A", note: KeyInt("A", -2), want: -2},
		{name: "B", note: KeyInt("B", -1), want: -1},
		{name: "C", note: KeyInt("C", 0), want: 0},
		{name: "D", note: KeyInt("D", 1), want: 1},
		{name: "E", note: KeyInt("E", 2), want: 2},
		{name: "F", note: KeyInt("F", 3), want: 3},
		{name: "G", note: KeyInt("G", 4), want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoteOctave(tt.note); got != tt.want {
				t.Errorf("NoteOctave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteToName(t *testing.T) {
	tests := []struct {
		name string
		note int
		want string
	}{
		{name: "A", note: KeyInt("A", -2), want: "A-2"},
		{name: "B", note: KeyInt("B", -1), want: "B-1"},
		{name: "C", note: KeyInt("C", 0), want: "C0"},
		{name: "D", note: KeyInt("D", 1), want: "D1"},
		{name: "E", note: KeyInt("E", 2), want: "E2"},
		{name: "F", note: KeyInt("F", 3), want: "F3"},
		{name: "G", note: KeyInt("G", 4), want: "G4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoteToName(tt.note); got != tt.want {
				t.Errorf("NoteToName(%d) = %v, want %v", tt.note, got, tt.want)
			}
		})
	}
}
