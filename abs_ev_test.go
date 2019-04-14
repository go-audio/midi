package midi

import (
	"github.com/mattetti/filebuffer"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestAbsEvents_ToMIDITrack(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
	}{
		{name: "unquantized track",
			fixture: "unquantized2bars.mid",
		},
	}
	for _, tt := range tests {
		r, err := os.Open(filepath.Join("fixtures", tt.fixture))
		if err != nil {
			t.Fatal(err)
		}
		defer r.Close()
		dec := NewDecoder(r)
		if err := dec.Parse(); err != nil {
			t.Fatal(err)
		}
		buff := filebuffer.New(nil)
		enc := NewEncoder(buff, dec.Format, dec.TicksPerQuarterNote)
		evs := dec.Tracks[0].AbsoluteEvents()
		t.Run(tt.name, func(t *testing.T) {
			encodedTrack := evs.ToMIDITrack(enc)
			encodedEvs := encodedTrack.AbsoluteEvents()
			if !reflect.DeepEqual(evs, encodedEvs) {
				t.Errorf("AbsEvents.ToMIDITrack() = %v, want %v", evs, encodedEvs)
			}
		})
	}
}
