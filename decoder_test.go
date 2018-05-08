package midi

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestVarint(t *testing.T) {
	expecations := []struct {
		dec   uint32
		bytes []byte
	}{
		{0, []byte{0}},
		{42, []byte{0x2a}},
		{4610, []byte{0xa4, 0x02}},
	}

	for _, exp := range expecations {
		conv := EncodeVarint(exp.dec)
		if bytes.Compare(conv, exp.bytes) != 0 {
			t.Fatalf("%d was converted to %#v didn't match %#v\n", exp.dec, conv, exp.bytes)
		}
	}

	for _, exp := range expecations {
		conv, _ := DecodeVarint(exp.bytes)
		if conv != exp.dec {
			t.Fatalf("%#v was converted to %d didn't match %d\n", exp.bytes, conv, exp.dec)
		}
	}
}

func TestParsingFile(t *testing.T) {
	expectations := []struct {
		path                string
		format              uint16
		numTracks           uint16
		ticksPerQuarterNote uint16
		timeFormat          timeFormat
		trackNames          []string
		bpms                []int
	}{
		{"fixtures/elise.mid", 1, 4, 960, MetricalTF, []string{"Track 0", "F\xfcr Elise", "http://www.forelise.com/", "Composed by Ludwig van Beethoven"}, []int{69, 0, 0, 0}},
		{"fixtures/elise1track.mid", 0, 1, 480, MetricalTF, []string{"F"}, []int{69}},
		{"fixtures/bossa.mid", 0, 1, 96, MetricalTF, []string{"bossa 1"}, []int{0}},
		{"fixtures/closedHat.mid", 0, 1, 96, MetricalTF, []string{"01 4th Hat Closed Side"}, []int{0}},
	}

	for _, exp := range expectations {
		path, _ := filepath.Abs(exp.path)
		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		p := NewDecoder(f)
		if err := p.Parse(); err != nil {
			t.Fatal(err)
		}

		if p.Format != exp.format {
			t.Fatalf("%s of %s didn't match %v, got %v", "format", exp.path, exp.format, p.Format)
		}
		if p.NumTracks != exp.numTracks {
			t.Fatalf("%s of %s didn't match %v, got %v", "numTracks", exp.path, exp.numTracks, p.NumTracks)
		}
		if p.TicksPerQuarterNote != exp.ticksPerQuarterNote {
			t.Fatalf("%s of %s didn't match %v, got %v", "ticksPerQuarterNote", exp.path, exp.ticksPerQuarterNote, p.TicksPerQuarterNote)
		}
		if p.TimeFormat != exp.timeFormat {
			t.Fatalf("%s of %s didn't match %v, got %v", "format", exp.path, exp.timeFormat, p.TimeFormat)
		}

		if len(p.Tracks) == 0 {
			t.Fatal("Tracks not parsed")
		}
		for i, tr := range p.Tracks {
			t.Run(fmt.Sprintf("Track %s", tr.Name()), func(t *testing.T) {
				if tName := tr.Name(); tName != exp.trackNames[i] {
					t.Fatalf("expected name of track %d to be %s but got %s -> (%q)", i, exp.trackNames[i], tName, tName)
				}
				if bpm := tr.Tempo(); bpm != exp.bpms[i] {
					t.Fatalf("expected tempo of track %d to be %d but got %d", i, exp.bpms[i], bpm)
				}
			})
		}

	}
}

func TestDecoder_parseTrack(t *testing.T) {
	toEvTime := func(e *Event) evTime {
		return evTime{Note: e.Note, timeDelta: e.TimeDelta, absTicks: e.AbsTicks}
	}
	tests := []struct {
		name     string
		path     string
		trackIDX int
		events   []evTime
	}{
		{name: "unquantized", path: "fixtures/unquantized.mid", events: []evTime{
			{Note: 0, timeDelta: 0, absTicks: 0},
			{Note: 0, timeDelta: 0, absTicks: 0},
			{Note: 0, timeDelta: 0, absTicks: 0},
			{Note: 36, timeDelta: 0, absTicks: 0},
			{Note: 36, timeDelta: 58, absTicks: 58},
			{Note: 38, timeDelta: 17, absTicks: 75},
			{Note: 38, timeDelta: 20, absTicks: 95},
			{Note: 38, timeDelta: 0, absTicks: 95},
			{Note: 38, timeDelta: 48, absTicks: 143},
			{Note: 38, timeDelta: 0, absTicks: 143},
			{Note: 36, timeDelta: 48, absTicks: 191},
			{Note: 38, timeDelta: 31, absTicks: 222},
			{Note: 36, timeDelta: 28, absTicks: 250},
			{Note: 38, timeDelta: 16, absTicks: 266},
			{Note: 38, timeDelta: 21, absTicks: 287},
			{Note: 38, timeDelta: 0, absTicks: 287},
			{Note: 38, timeDelta: 48, absTicks: 335},
			{Note: 38, timeDelta: 0, absTicks: 335},
			{Note: 36, timeDelta: 48, absTicks: 383},
			{Note: 38, timeDelta: 9, absTicks: 392},
			{Note: 36, timeDelta: 42, absTicks: 434},
			{Note: 38, timeDelta: 25, absTicks: 459},
			{Note: 38, timeDelta: 20, absTicks: 479},
			{Note: 38, timeDelta: 0, absTicks: 479},
			{Note: 38, timeDelta: 48, absTicks: 527},
			{Note: 38, timeDelta: 0, absTicks: 527},
			{Note: 36, timeDelta: 48, absTicks: 575},
			{Note: 38, timeDelta: 9, absTicks: 584},
			{Note: 36, timeDelta: 50, absTicks: 634},
			{Note: 38, timeDelta: 16, absTicks: 650},
			{Note: 38, timeDelta: 21, absTicks: 671},
			{Note: 38, timeDelta: 0, absTicks: 671},
			{Note: 38, timeDelta: 48, absTicks: 719},
			{Note: 38, timeDelta: 0, absTicks: 719},
			{Note: 36, timeDelta: 48, absTicks: 767},
			{Note: 38, timeDelta: 9, absTicks: 776},
			{Note: 36, timeDelta: 43, absTicks: 819},
			{Note: 38, timeDelta: 23, absTicks: 842},
			{Note: 38, timeDelta: 21, absTicks: 863},
			{Note: 38, timeDelta: 0, absTicks: 863},
			{Note: 38, timeDelta: 48, absTicks: 911},
			{Note: 38, timeDelta: 0, absTicks: 911},
			{Note: 36, timeDelta: 48, absTicks: 959},
			{Note: 38, timeDelta: 8, absTicks: 967},
			{Note: 36, timeDelta: 57, absTicks: 1024},
			{Note: 38, timeDelta: 10, absTicks: 1034},
			{Note: 38, timeDelta: 21, absTicks: 1055},
			{Note: 38, timeDelta: 0, absTicks: 1055},
			{Note: 38, timeDelta: 48, absTicks: 1103},
			{Note: 38, timeDelta: 0, absTicks: 1103},
			{Note: 36, timeDelta: 48, absTicks: 1151},
			{Note: 38, timeDelta: 9, absTicks: 1160},
			{Note: 36, timeDelta: 43, absTicks: 1203},
			{Note: 38, timeDelta: 23, absTicks: 1226},
			{Note: 38, timeDelta: 21, absTicks: 1247},
			{Note: 38, timeDelta: 0, absTicks: 1247},
			{Note: 38, timeDelta: 48, absTicks: 1295},
			{Note: 38, timeDelta: 0, absTicks: 1295},
			{Note: 36, timeDelta: 48, absTicks: 1343},
			{Note: 38, timeDelta: 9, absTicks: 1352},
			{Note: 36, timeDelta: 50, absTicks: 1402},
			{Note: 38, timeDelta: 16, absTicks: 1418},
			{Note: 38, timeDelta: 21, absTicks: 1439},
			{Note: 38, timeDelta: 0, absTicks: 1439},
			{Note: 38, timeDelta: 48, absTicks: 1487},
			{Note: 38, timeDelta: 0, absTicks: 1487},
			{Note: 38, timeDelta: 49, absTicks: 1536},
			{Note: 0, timeDelta: 0, absTicks: 1536}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, _ := filepath.Abs(tt.path)
			f, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			p := NewDecoder(f)
			if err := p.Parse(); err != nil {
				t.Fatal(err)
			}
			track := p.Tracks[tt.trackIDX]
			for i, e := range track.Events {
				if evt := toEvTime(e); !reflect.DeepEqual(evt, tt.events[i]) {
					t.Fatalf("[%d] Expected %+v, got %+v", i, tt.events[i], evt)
				}
			}
		})
	}
}

type evTime struct {
	Note      uint8
	timeDelta uint32
	absTicks  uint64
}
