package midi_test

import (
	"reflect"
	"testing"

	"github.com/go-audio/midi"
)

func TestScaleNotes(t *testing.T) {
	tests := []struct {
		name      string
		tonic     string
		scale     midi.ScaleName
		wantKeys  []int
		wantNames []string
	}{
		{
			name: "C Major", tonic: "c", scale: midi.MajorScale,
			wantKeys:  []int{24, 26, 28, 29, 31, 33, 35},
			wantNames: []string{`C`, `D`, `E`, `F`, `G`, `A`, `B`},
		},
		{
			name: "C melodic Minor", tonic: "C", scale: midi.MelodicMinorScale,
			wantKeys:  []int{24, 26, 27, 29, 31, 33, 35},
			wantNames: []string{`C`, `D`, `D#`, `F`, `G`, `A`, `B`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, notes := midi.ScaleNotes(tt.tonic, tt.scale)
			if !reflect.DeepEqual(keys, tt.wantKeys) {
				t.Errorf("ScaleNotes() keys = %v, want %v", keys, tt.wantKeys)
			}
			if !reflect.DeepEqual(notes, tt.wantNames) {
				t.Errorf("ScaleNotes() notes = %v, want %v", notes, tt.wantNames)
			}
		})
	}
}
