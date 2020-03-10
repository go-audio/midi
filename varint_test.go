package midi

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncodeDecodeVarint(t *testing.T) {
	testCases := []struct {
		input  []byte
		output uint32
	}{
		0: {[]byte{0x7F}, 127},
		1: {[]byte{0x81, 0x00}, 128},
		2: {[]byte{0xC0, 0x00}, 8192},
		3: {[]byte{0xFF, 0x7F}, 16383},
		4: {[]byte{0x81, 0x80, 0x00}, 16384},
		5: {[]byte{0xFF, 0xFF, 0x7F}, 2097151},
		6: {[]byte{0x81, 0x80, 0x80, 0x00}, 2097152},
		7: {[]byte{0xC0, 0x80, 0x80, 0x00}, 134217728},
		8: {[]byte{0xFF, 0xFF, 0xFF, 0x7F}, 268435455},

		9:  {[]byte{0x82, 0x00}, 256},
		10: {[]byte{0x81, 0x10}, 144},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v", tc.input), func(t *testing.T) {
			if o, k := DecodeVarint(tc.input); o != tc.output || k != len(tc.input) {
				t.Fatalf("expected %d len %d\ngot\n%d len %d\n", tc.output, len(tc.input), o, k)
			}
			if encoded := EncodeVarint(tc.output); bytes.Compare(encoded, tc.input) != 0 {
				t.Fatalf("%d - expected %#v\ngot\n%#v\n", tc.output, tc.input, encoded)
			}
		})
	}
}

func TestDecodeVarint(t *testing.T) {
	tests := []struct {
		name  string
		buf   []byte
		wantX uint32
		wantN int
	}{
		{name: "empty decoder", buf: []byte{}, wantX: 0, wantN: 0},
		{name: "fall through", buf: []byte{0x81}, wantX: 1, wantN: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotN := DecodeVarint(tt.buf)
			if gotX != tt.wantX {
				t.Errorf("DecodeVarint() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotN != tt.wantN {
				t.Errorf("DecodeVarint() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
