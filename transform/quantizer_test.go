package transform

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/go-audio/midi"
	"github.com/go-audio/midi/grid"
)

func TestQuantizer_Quantize(t *testing.T) {
	type fields struct {
		GridRes           grid.Res
		QuantizationLevel float64
		Start             bool
		End               bool
		MoveEndOnStartQ   bool
	}
	tests := []struct {
		name    string
		fields  fields
		fixture string
		want    string // fixture path
	}{
		{
			name:    "Quantize start of unquantized content at 1/8 100%",
			fixture: "../fixtures/unquantized2bars.mid",
			fields: fields{
				GridRes:           grid.One8,
				QuantizationLevel: 1.0,
				Start:             true,
			},
			want: "../fixtures/unquantized2bars-quantized1_8.mid",
		},
		{
			name:    "Quantize start of unquantized content at 1/16 100%",
			fixture: "../fixtures/unquantized2bars.mid",
			fields: fields{
				GridRes:           grid.One16,
				QuantizationLevel: 1.0,
				Start:             true,
			},
			want: "../fixtures/unquantized2bars-quantized1_16.mid",
		},
		{
			name:    "Quantize start of unquantized content at 1/32 100%",
			fixture: "../fixtures/unquantized2bars.mid",
			fields: fields{
				GridRes:           grid.One32,
				QuantizationLevel: 1.0,
				Start:             true,
			},
			want: "../fixtures/unquantized2bars-quantized1_32.mid",
		},
		{
			name:    "Quantize start of other unquantized content at 1/32 100%",
			fixture: "../fixtures/unquantized.mid",
			fields: fields{
				GridRes:           grid.One32,
				QuantizationLevel: 1.0,
				Start:             true,
			},
			want: "../fixtures/unquantized_quantized-1_32.mid",
		},
		{
			name:    "Quantize start and end of unquantized content at 1/32 100%",
			fixture: "../fixtures/unquantized.mid",
			fields: fields{
				GridRes:           grid.One32,
				QuantizationLevel: 1.0,
				Start:             true,
				End:               true,
			},
			want: "../fixtures/unquantized_quantized-1_32_start_end.mid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := os.Open(filepath.Join(tt.fixture))
			if err != nil {
				t.Fatal(err)
			}
			dec := midi.NewDecoder(r)
			if err = dec.Parse(); err != nil {
				t.Fatal(err)
			}
			events := dec.Tracks[0].AbsoluteEvents()
			r.Close()
			// expected output
			r, err = os.Open(filepath.Join(tt.want))
			if err != nil {
				t.Fatal(err)
			}
			dec = midi.NewDecoder(r)
			if err = dec.Parse(); err != nil {
				t.Fatal(err)
			}
			want := dec.Tracks[0].AbsoluteEvents()
			r.Close()

			q := Quantizer{
				GridRes:           tt.fields.GridRes,
				QuantizationLevel: tt.fields.QuantizationLevel,
				Start:             tt.fields.Start,
				End:               tt.fields.End,
				MoveEndOnStartQ:   tt.fields.MoveEndOnStartQ,
			}
			got := q.Quantize(events, dec.TicksPerQuarterNote)
			if len(got) != len(want) {
				t.Fatalf("expected %d events but got %d", len(want), len(got))
			}
			for i, ev := range got {
				if !reflect.DeepEqual(ev, want[i]) {
					t.Errorf("[%d] expected\t%+v\ngot\t\t%+v\nOriginal event: %+v", i, want[i], ev, events[i])
				}
			}
		})
	}
}
