package midi

import (
	"fmt"
	"strings"
)

// ChordDefinition defines chords by name and by defining the interval between
// adjacent notes on the chromatic scale (half steps) to create the chord.
type ChordDefinition struct {
	// Root isn't usually defined until the definition is used
	Root string
	// Name is the English name of the chord
	Name string
	// Abbrev is the English abbreviation for the chord
	Abbrev string
	// HalfSteps are the number of half steps between the chord's notes
	HalfSteps []uint
}

// WithRoot returns a copy of the chord definition with the chord root set.
func (cd *ChordDefinition) WithRoot(root string) *ChordDefinition {
	return &ChordDefinition{
		Root:      root,
		Name:      cd.Name,
		Abbrev:    cd.Abbrev,
		HalfSteps: cd.HalfSteps,
	}
}

// RootInt returns the note number (0-11).
// -1 is returned if the root isn't set.
func (cd *ChordDefinition) RootInt() int {
	if cd == nil || len(cd.Root) < 1 {
		return -1
	}
	return KeyInt(cd.Root, 0) % 12
}

func (cd *ChordDefinition) String() string {
	if len(cd.Root) > 0 {
		return fmt.Sprintf("%s %s", strings.ToUpper(cd.Root), cd.Name)
	}
	return cd.Name
}

var (
	// ChordDefs are the most populate chord definitions
	ChordDefs = []*ChordDefinition{
		{
			Name: "Major", Abbrev: "maj",
			HalfSteps: []uint{4, 3},
		},
		{
			Name: "Minor", Abbrev: "min",
			HalfSteps: []uint{3, 4},
		},
		{
			Name: "Diminished", Abbrev: "mb5",
			HalfSteps: []uint{3, 3},
		},
		{
			Name: "Augmented", Abbrev: "aug",
			HalfSteps: []uint{4, 4},
		},
		{
			// 2 note chord
			Name: "Fifth", Abbrev: "5",
			HalfSteps: []uint{7},
		},
		{
			Name: "Minor Seventh", Abbrev: "m7", HalfSteps: []uint{3, 4, 3},
		},
	}
	// OtherChordDefs are less common chord definitions
	OtherChordDefs = []*ChordDefinition{
		{
			Name: "Suspended 2nd", Abbrev: "sus2",
			HalfSteps: []uint{2, 5},
		},
		{
			Name: "Suspended 4th", Abbrev: "sus4",
			HalfSteps: []uint{5, 2},
		},
		{
			Name: "Major Flat 5th", Abbrev: "majb5",
			HalfSteps: []uint{4, 2},
		},
		{
			Name: "Augmented Suspended 4th", Abbrev: "augsus4",
			HalfSteps: []uint{5, 3},
		},
		// 6th
		{
			Name: "Triton", Abbrev: "tri",
			HalfSteps: []uint{3, 3, 3},
		},
		{
			Name: "Sixth", Abbrev: "6", HalfSteps: []uint{4, 3, 2},
		},
		{
			Name: "Sixth Suspended 4th", Abbrev: "6sus4", HalfSteps: []uint{5, 2, 2},
		},
		{
			Name: "Sixth add 9th", Abbrev: "6add9", HalfSteps: []uint{4, 3, 2, 5},
		},
		{
			Name: "Minor Sixth", Abbrev: "m6", HalfSteps: []uint{3, 4, 2},
		},
		{
			Name: "Minor Sixth 9th", Abbrev: "m6add9", HalfSteps: []uint{3, 4, 2, 5},
		},
		// 7th
		{
			Name: "Seventh", Abbrev: "7", HalfSteps: []uint{4, 3, 3},
		},
		{
			Name: "Seventh Suspended 4th", Abbrev: "7sus4", HalfSteps: []uint{5, 2, 3},
		},
		{
			Name: "Seventh Sharp 5th", Abbrev: "7#5", HalfSteps: []uint{4, 4, 2},
		},
		{
			Name: "Seventh Flat 5th", Abbrev: "7b5", HalfSteps: []uint{4, 2, 4},
		},
		{
			Name: "Seventh Sharp 9th", Abbrev: "7#9", HalfSteps: []uint{4, 3, 3, 5},
		},
		{
			Name: "Seventh Flat 9th", Abbrev: "7b9", HalfSteps: []uint{4, 3, 3, 3},
		},
		{
			Name: "Seventh Sharp 5th Sharp 9th", Abbrev: "7#5#9", HalfSteps: []uint{4, 4, 2, 5},
		},
		{
			Name: "Seventh Sharp 5th Flat 9th", Abbrev: "7#5b9", HalfSteps: []uint{4, 4, 2, 3},
		},
		{
			Name: "Seventh Flat 5th Flat 9th", Abbrev: "7b5b9", HalfSteps: []uint{4, 2, 4, 3},
		},
		{
			Name: "Seventh add 11th", Abbrev: "7add11", HalfSteps: []uint{4, 3, 3, 7},
		},
		{
			Name: "Seventh add 13th", Abbrev: "7add13", HalfSteps: []uint{4, 3, 3, 11},
		},
		{
			Name: "Seventh Sharp 11th", Abbrev: "7#11", HalfSteps: []uint{4, 3, 3, 8},
		},
		{
			Name: "Major Seventh", Abbrev: "Maj7", HalfSteps: []uint{4, 3, 4},
		},
		{
			Name: "Major Seventh Flat 5th", Abbrev: "Maj7b5", HalfSteps: []uint{4, 2, 5},
		},
		{
			Name: "Major Seventh Sharp 5th", Abbrev: "Maj7#5", HalfSteps: []uint{4, 4, 3},
		},
		{
			Name: "Major Seventh Sharp 11th", Abbrev: "Maj7#11", HalfSteps: []uint{4, 3, 4, 7},
		},
		{
			Name: "Major Seventh add 13th", Abbrev: "Maj7add13", HalfSteps: []uint{4, 3, 4, 10},
		},
		{
			Name: "Minor Seventh Flat 5th", Abbrev: "m7b5", HalfSteps: []uint{3, 3, 4},
		},
		{
			Name: "Minor Seventh Flat 9th", Abbrev: "m7b9", HalfSteps: []uint{3, 4, 3, 3},
		},
		{
			Name: "Minor Seventh add 11th", Abbrev: "m7add11", HalfSteps: []uint{3, 4, 3, 7},
		},
		{
			Name: "Minor Seventh add 13th", Abbrev: "m7add13", HalfSteps: []uint{3, 4, 3, 11},
		},
		{
			Name: "Minor Major Seventh", Abbrev: "m-Maj7", HalfSteps: []uint{3, 4, 4},
		},
		{
			Name: "Minor Major Seventh add 11th", Abbrev: "m-Maj7add11", HalfSteps: []uint{3, 4, 4, 6},
		},
		{
			Name: "Minor Major Seventh add 13th", Abbrev: "m-Maj7add13", HalfSteps: []uint{3, 4, 4, 10},
		},
		// 9th
		{
			Name: "Ninth", Abbrev: "9", HalfSteps: []uint{4, 3, 3, 4},
		},
		{
			Name: "Ninth Suspended 4th", Abbrev: "9sus4", HalfSteps: []uint{5, 2, 3, 4},
		},
		{
			Name: "Ninth Add", Abbrev: "add9", HalfSteps: []uint{4, 3, 7},
		},
		{
			Name: "Ninth Sharp 5th", Abbrev: "9#5", HalfSteps: []uint{4, 4, 2, 4},
		},
		{
			Name: "Ninth Flat 5th", Abbrev: "9b5", HalfSteps: []uint{4, 2, 4, 4},
		},
		{
			Name: "Ninth Sharp 11th", Abbrev: "9#11", HalfSteps: []uint{4, 3, 3, 4, 4},
		},
		{
			Name: "Ninth Flat 13th", Abbrev: "9b13", HalfSteps: []uint{4, 3, 3, 4, 6},
		},
		{
			Name: "Major Ninth", Abbrev: "Maj9", HalfSteps: []uint{4, 3, 4, 3},
		},
		{
			Name: "Major Ninth Suspended 4th", Abbrev: "Maj9sus4", HalfSteps: []uint{5, 2, 4, 3},
		},
		{
			Name: "Major Ninth Sharp 5th", Abbrev: "Maj9#5", HalfSteps: []uint{4, 4, 3, 3},
		},
		{
			Name: "Major Ninth Sharp 11th", Abbrev: "Maj9#11", HalfSteps: []uint{4, 3, 4, 3, 4},
		},
		{
			Name: "Minor Ninth", Abbrev: "m9", HalfSteps: []uint{3, 4, 3, 4},
		},
		{
			Name: "Minor add 9th", Abbrev: "madd9", HalfSteps: []uint{3, 4, 7},
		},
		{
			Name: "Minor Ninth Flat 5th", Abbrev: "m9b5", HalfSteps: []uint{3, 3, 4, 4},
		},
		{
			Name: "Minor Major Ninth", Abbrev: "m9-Maj7", HalfSteps: []uint{3, 4, 4, 3},
		},
		// 11th
		{
			Name: "Eleventh", Abbrev: "11", HalfSteps: []uint{4, 3, 3, 4, 3},
		},
		{
			Name: "Eleventh Flat 9th", Abbrev: "11b9", HalfSteps: []uint{4, 3, 3, 3, 4},
		},
		{
			Name: "Major Eleventh", Abbrev: "Maj11", HalfSteps: []uint{4, 3, 4, 3, 3},
		},
		{
			Name: "Minor Eleventh", Abbrev: "m11", HalfSteps: []uint{3, 4, 3, 4, 3},
		},
		{
			Name: "Minor Major Eleventh", Abbrev: "m-Maj11", HalfSteps: []uint{3, 4, 4, 3, 3},
		},
		// 13th

		{
			Name: "Thirteenth", Abbrev: "13", HalfSteps: []uint{4, 3, 3, 4, 7},
		},
		{
			Name: "Thirteenth Sharp 9th", Abbrev: "13#9", HalfSteps: []uint{4, 3, 3, 5, 6},
		},
		{
			Name: "Thirteenth Flat 9th", Abbrev: "13b9", HalfSteps: []uint{4, 3, 3, 3, 8},
		},
		{
			Name: "Thirteenth Flat 5th Flat 9th", Abbrev: "13b5b9", HalfSteps: []uint{4, 2, 4, 3, 8},
		},
		{
			Name: "Major Thirteenth", Abbrev: "Maj13", HalfSteps: []uint{4, 3, 4, 3, 7},
		},
		{
			Name: "Minor Thirteenth", Abbrev: "m13", HalfSteps: []uint{3, 4, 3, 4, 7},
		},
		{
			Name: "Minor Major Thirteenth", Abbrev: "m-Maj13", HalfSteps: []uint{3, 4, 4, 3, 7},
		},
	}
)
