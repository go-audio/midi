package midi

import (
	"testing"
)

func TestChordDefinition_WithRoot(t *testing.T) {
	type fields struct {
		Root      string
		Name      string
		Abbrev    string
		HalfSteps []uint
	}
	type args struct {
		root string
	}
	tests := []struct {
		name   string
		root   string
		fields fields
		args   args
		want   *ChordDefinition
	}{
		{
			name: "Basic",
			root: "c",
			fields: fields{
				Name:      ChordDefs[0].Name,
				Abbrev:    ChordDefs[0].Abbrev,
				HalfSteps: ChordDefs[0].HalfSteps,
			},
			want: &ChordDefinition{
				Root:      "C",
				Name:      ChordDefs[0].Name,
				Abbrev:    ChordDefs[0].Abbrev,
				HalfSteps: ChordDefs[0].HalfSteps,
			},
		},
		{
			name: "make sure we don't cache",
			root: "d",
			fields: fields{
				Name:      ChordDefs[0].Name,
				Abbrev:    ChordDefs[0].Abbrev,
				HalfSteps: ChordDefs[0].HalfSteps,
			},
			want: &ChordDefinition{
				Root:      "D",
				Name:      ChordDefs[0].Name,
				Abbrev:    ChordDefs[0].Abbrev,
				HalfSteps: ChordDefs[0].HalfSteps,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cd := &ChordDefinition{
				Root:      tt.fields.Root,
				Name:      tt.fields.Name,
				Abbrev:    tt.fields.Abbrev,
				HalfSteps: tt.fields.HalfSteps,
			}
			if got := cd.WithRoot(tt.root); got.String() != tt.want.String() {
				t.Errorf("ChordDefinition.WithRoot() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
