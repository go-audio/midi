package midi

// ChordDefinition defines chords by name and by defining the interval between
// adjacent notes on the chromatic scale (half steps) to create the chord.
type ChordDefinition struct {
	Name      string
	Abbrev    string
	HalfSteps []int
}

var (
	// ChordDefs contain all the known chord definitions key by name.
	ChordDefs = map[string]ChordDefinition{
		"Major": {
			Name: "Major", Abbrev: "maj",
			HalfSteps: []int{4, 3},
		},
		"Minor": {
			Name: "Minor", Abbrev: "min",
			HalfSteps: []int{3, 4},
		},
		"Suspended 2nd": {
			Name: "Suspended 2nd", Abbrev: "sus2",
			HalfSteps: []int{2, 5},
		},
		"Suspended 4th": {
			Name: "Suspended 4th", Abbrev: "sus4",
			HalfSteps: []int{5, 2},
		},
		"Major Flat 5th": {
			Name: "Major Flat 5th", Abbrev: "majb5",
			HalfSteps: []int{4, 2},
		},
		"Diminished": {
			Name: "Diminished", Abbrev: "mb5",
			HalfSteps: []int{3, 3},
		},
		"Augmented": {
			Name: "Augmented", Abbrev: "aug",
			HalfSteps: []int{4, 4},
		},
		"Augmented Suspended 4th": {
			Name: "Augmented Suspended 4th", Abbrev: "augsus4",
			HalfSteps: []int{5, 3},
		},
		// 6th
		"Triton": {
			Name: "Triton", Abbrev: "tri",
			HalfSteps: []int{3, 3, 3},
		},
		"Sixth": {
			Name: "Sixth", Abbrev: "6", HalfSteps: []int{4, 3, 2},
		},
		"Sixth Suspended 4th": {
			Name: "Sixth Suspended 4th", Abbrev: "6sus4", HalfSteps: []int{5, 2, 2},
		},
		"Sixth add 9th": {
			Name: "Sixth add 9th", Abbrev: "6add9", HalfSteps: []int{4, 3, 2, 5},
		},
		"Minor Sixth": {
			Name: "Minor Sixth", Abbrev: "m6", HalfSteps: []int{3, 4, 2},
		},
		"Minor Sixth 9th": {
			Name: "Minor Sixth 9th", Abbrev: "m6add9", HalfSteps: []int{3, 4, 2, 5},
		},
		// 7th
		"Seventh": {
			Name: "Seventh", Abbrev: "7", HalfSteps: []int{4, 3, 3},
		},
		"Seventh Suspended 4th": {
			Name: "Seventh Suspended 4th", Abbrev: "7sus4", HalfSteps: []int{5, 2, 3},
		},
		"Seventh Sharp 5th": {
			Name: "Seventh Sharp 5th", Abbrev: "7#5", HalfSteps: []int{4, 4, 2},
		},
		"Seventh Flat 5th": {
			Name: "Seventh Flat 5th", Abbrev: "7b5", HalfSteps: []int{4, 2, 4},
		},
		"Seventh Sharp 9th": {
			Name: "Seventh Sharp 9th", Abbrev: "7#9", HalfSteps: []int{4, 3, 3, 5},
		},
		"Seventh Flat 9th": {
			Name: "Seventh Flat 9th", Abbrev: "7b9", HalfSteps: []int{4, 3, 3, 3},
		},
		"Seventh Sharp 5th Sharp 9th": {
			Name: "Seventh Sharp 5th Sharp 9th", Abbrev: "7#5#9", HalfSteps: []int{4, 4, 2, 5},
		},
		"Seventh Sharp 5th Flat 9th": {
			Name: "Seventh Sharp 5th Flat 9th", Abbrev: "7#5b9", HalfSteps: []int{4, 4, 2, 3},
		},
		"Seventh Flat 5th Flat 9th": {
			Name: "Seventh Flat 5th Flat 9th", Abbrev: "7b5b9", HalfSteps: []int{4, 2, 4, 3},
		},
		"Seventh add 11th": {
			Name: "Seventh add 11th", Abbrev: "7add11", HalfSteps: []int{4, 3, 3, 7},
		},
		"Seventh add 13th": {
			Name: "Seventh add 13th", Abbrev: "7add13", HalfSteps: []int{4, 3, 3, 11},
		},
		"Seventh Sharp 11th": {
			Name: "Seventh Sharp 11th", Abbrev: "7#11", HalfSteps: []int{4, 3, 3, 8},
		},
		"Major Seventh": {
			Name: "Major Seventh", Abbrev: "Maj7", HalfSteps: []int{4, 3, 4},
		},
		"Major Seventh Flat 5th": {
			Name: "Major Seventh Flat 5th", Abbrev: "Maj7b5", HalfSteps: []int{4, 2, 5},
		},
		"Major Seventh Sharp 5th": {
			Name: "Major Seventh Sharp 5th", Abbrev: "Maj7#5", HalfSteps: []int{4, 4, 3},
		},
		"Major Seventh Sharp 11th": {
			Name: "Major Seventh Sharp 11th", Abbrev: "Maj7#11", HalfSteps: []int{4, 3, 4, 7},
		},
		"Major Seventh add 13th": {
			Name: "Major Seventh add 13th", Abbrev: "Maj7add13", HalfSteps: []int{4, 3, 4, 10},
		},
		"Minor Seventh": {
			Name: "Minor Seventh", Abbrev: "m7", HalfSteps: []int{3, 4, 3},
		},
		"Minor Seventh Flat 5th": {
			Name: "Minor Seventh Flat 5th", Abbrev: "m7b5", HalfSteps: []int{3, 3, 4},
		},
		"Minor Seventh Flat 9th": {
			Name: "Minor Seventh Flat 9th", Abbrev: "m7b9", HalfSteps: []int{3, 4, 3, 3},
		},
		"Minor Seventh add 11th": {
			Name: "Minor Seventh add 11th", Abbrev: "m7add11", HalfSteps: []int{3, 4, 3, 7},
		},
		"Minor Seventh add 13th": {
			Name: "Minor Seventh add 13th", Abbrev: "m7add13", HalfSteps: []int{3, 4, 3, 11},
		},
		"Minor Major Seventh": {
			Name: "Minor Major Seventh", Abbrev: "m-Maj7", HalfSteps: []int{3, 4, 4},
		},
		"Minor Major Seventh add 11th": {
			Name: "Minor Major Seventh add 11th", Abbrev: "m-Maj7add11", HalfSteps: []int{3, 4, 4, 6},
		},
		"Minor Major Seventh add 13th": {
			Name: "Minor Major Seventh add 13th", Abbrev: "m-Maj7add13", HalfSteps: []int{3, 4, 4, 10},
		},
		// 9th
		"Ninth": {
			Name: "Ninth", Abbrev: "9", HalfSteps: []int{4, 3, 3, 4},
		},
		"Ninth Suspended 4th": {
			Name: "Ninth Suspended 4th", Abbrev: "9sus4", HalfSteps: []int{5, 2, 3, 4},
		},
		"Ninth Add": {
			Name: "Ninth Add", Abbrev: "add9", HalfSteps: []int{4, 3, 7},
		},
		"Ninth Sharp 5th": {
			Name: "Ninth Sharp 5th", Abbrev: "9#5", HalfSteps: []int{4, 4, 2, 4},
		},
		"Ninth Flat 5th": {
			Name: "Ninth Flat 5th", Abbrev: "9b5", HalfSteps: []int{4, 2, 4, 4},
		},
		"Ninth Sharp 11th": {
			Name: "Ninth Sharp 11th", Abbrev: "9#11", HalfSteps: []int{4, 3, 3, 4, 4},
		},
		"Ninth Flat 13th": {
			Name: "Ninth Flat 13th", Abbrev: "9b13", HalfSteps: []int{4, 3, 3, 4, 6},
		},
		"Major Ninth": {
			Name: "Major Ninth", Abbrev: "Maj9", HalfSteps: []int{4, 3, 4, 3},
		},
		"Major Ninth Suspended 4th": {
			Name: "Major Ninth Suspended 4th", Abbrev: "Maj9sus4", HalfSteps: []int{5, 2, 4, 3},
		},
		"Major Ninth Sharp 5th": {
			Name: "Major Ninth Sharp 5th", Abbrev: "Maj9#5", HalfSteps: []int{4, 4, 3, 3},
		},
		"Major Ninth Sharp 11th": {
			Name: "Major Ninth Sharp 11th", Abbrev: "Maj9#11", HalfSteps: []int{4, 3, 4, 3, 4},
		},
		"Minor Ninth": {
			Name: "Minor Ninth", Abbrev: "m9", HalfSteps: []int{3, 4, 3, 4},
		},
		"Minor add 9th": {
			Name: "Minor add 9th", Abbrev: "madd9", HalfSteps: []int{3, 4, 7},
		},
		"Minor Ninth Flat 5th": {
			Name: "Minor Ninth Flat 5th", Abbrev: "m9b5", HalfSteps: []int{3, 3, 4, 4},
		},
		"Minor Major Ninth": {
			Name: "Minor Major Ninth", Abbrev: "m9-Maj7", HalfSteps: []int{3, 4, 4, 3},
		},
		// 11th
		"Eleventh": {
			Name: "Eleventh", Abbrev: "11", HalfSteps: []int{4, 3, 3, 4, 3},
		},
		"Eleventh Flat 9th": {
			Name: "Eleventh Flat 9th", Abbrev: "11b9", HalfSteps: []int{4, 3, 3, 3, 4},
		},
		"Major Eleventh": {
			Name: "Major Eleventh", Abbrev: "Maj11", HalfSteps: []int{4, 3, 4, 3, 3},
		},
		"Minor Eleventh": {
			Name: "Minor Eleventh", Abbrev: "m11", HalfSteps: []int{3, 4, 3, 4, 3},
		},
		"Minor Major Eleventh": {
			Name: "Minor Major Eleventh", Abbrev: "m-Maj11", HalfSteps: []int{3, 4, 4, 3, 3},
		},
		// 13th

		"Thirteenth": {
			Name: "Thirteenth", Abbrev: "13", HalfSteps: []int{4, 3, 3, 4, 7},
		},
		"Thirteenth Sharp 9th": {
			Name: "Thirteenth Sharp 9th", Abbrev: "13#9", HalfSteps: []int{4, 3, 3, 5, 6},
		},
		"Thirteenth Flat 9th": {
			Name: "Thirteenth Flat 9th", Abbrev: "13b9", HalfSteps: []int{4, 3, 3, 3, 8},
		},
		"Thirteenth Flat 5th Flat 9th": {
			Name: "Thirteenth Flat 5th Flat 9th", Abbrev: "13b5b9", HalfSteps: []int{4, 2, 4, 3, 8},
		},
		"Major Thirteenth": {
			Name: "Major Thirteenth", Abbrev: "Maj13", HalfSteps: []int{4, 3, 4, 3, 7},
		},
		"Minor Thirteenth": {
			Name: "Minor Thirteenth", Abbrev: "m13", HalfSteps: []int{3, 4, 3, 4, 7},
		},
		"Minor Major Thirteenth": {
			Name: "Minor Major Thirteenth", Abbrev: "m-Maj13", HalfSteps: []int{3, 4, 4, 3, 7},
		},
	}
)
