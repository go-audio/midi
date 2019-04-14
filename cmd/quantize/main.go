package main

import (
	"github.com/go-audio/midi"
	"github.com/go-audio/midi/transform"
	"log"
	"os"
	"path/filepath"
)

func main() {
	r, err := os.Open(filepath.Join(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	dec := midi.NewDecoder(r)
	if err := dec.Parse(); err != nil {
		log.Fatal(err)
	}
	events := dec.Tracks[0].AbsoluteEvents()
	quantizedEvents := transform.One16thQuantizer.Quantize(events, dec.TicksPerQuarterNote)
	f, err := os.Create("quantized.mid")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	enc := midi.NewEncoder(f, dec.Format, dec.TicksPerQuarterNote)
	quantizedEvents.ToMIDITrack(enc)
	enc.Write()
	log.Println("Quantized.mid saved")
}
