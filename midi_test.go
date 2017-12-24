package midi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUint24(t *testing.T) {
	tests := []struct {
		n    uint32
		want []byte
	}{
		{n: 42, want: []byte{0x0, 0x0, 0x2a}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := Uint24(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint24() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
