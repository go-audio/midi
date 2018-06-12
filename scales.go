package midi

// ScaleDefinition defines a scale by giving it a name and the spacing between adjacent notes on the chromatic scale.
type ScaleDefinition struct {
	// Popular indicates that the scale is commonly used
	Popular   bool
	Name      ScaleName
	HalfSteps []int
}

// ScaleDefinitions is a type representing slice of scale definitions
type ScaleDefinitions []ScaleDefinition

// Popular filter down to only return the popular scales found
func (def ScaleDefinitions) Popular() ScaleDefinitions {
	out := ScaleDefinitions{}
	for _, scale := range def {
		if scale.Popular {
			out = append(out, scale)
		}
	}
	return out
}

// ScaleName is the English name of the scale
type ScaleName string

const (
	MajorScale           ScaleName = "Major"
	HarmonicMinorScale   ScaleName = "Harmonic Minor"
	MelodicMinorScale    ScaleName = "Melodic Minor"
	WholeToneScale       ScaleName = "Whole Tone"
	DiminishedScale      ScaleName = "Diminished"
	MajorPentatonicScale ScaleName = "Major Pentatonic"
	MinorPentatonicScale ScaleName = "Minor Pentatonic"
	JapInSenScale        ScaleName = "Jap In Sen"
	MajorBebopScale      ScaleName = "Major Bebop"
	DominantBebopScale   ScaleName = "Dominant Bebop"
	BluesScale           ScaleName = "Blues"
	ArabicScale          ScaleName = "Arabic"
	EnigmaticScale       ScaleName = "Enigmatic"
	NeapolitanScale      ScaleName = "Neapolitan"
	NeapolitanMinorScale ScaleName = "Neapolitan Minor"
	HungarianMinorScale  ScaleName = "Hungarian Minor"
	DorianScale          ScaleName = "Dorian"
	PhrygianScale        ScaleName = "Phrygian"
	LydianScale          ScaleName = "Lydian"
	MixolydianScale      ScaleName = "Mixolydian"
	AeolianScale         ScaleName = "Aeolian"
	LocrianScale         ScaleName = "Locrian"
)

var (
	// ScaleDefs list all known scales
	ScaleDefs = map[ScaleName]ScaleDefinition{
		MajorScale:           {Name: MajorScale, HalfSteps: []int{2, 2, 1, 2, 2, 2}, Popular: true},
		HarmonicMinorScale:   {Name: HarmonicMinorScale, HalfSteps: []int{2, 1, 2, 2, 1, 3}},
		MelodicMinorScale:    {Name: MelodicMinorScale, HalfSteps: []int{2, 1, 2, 2, 2, 2}},
		WholeToneScale:       {Name: WholeToneScale, HalfSteps: []int{2, 2, 2, 2, 2}},
		DiminishedScale:      {Name: DiminishedScale, HalfSteps: []int{2, 1, 2, 1, 2, 1, 2}},
		MajorPentatonicScale: {Name: MajorPentatonicScale, HalfSteps: []int{2, 2, 3, 2}},
		MinorPentatonicScale: {Name: MinorPentatonicScale, HalfSteps: []int{3, 2, 2, 3}, Popular: true},
		DorianScale:          {Name: DorianScale, HalfSteps: []int{2, 1, 2, 2, 2, 1}, Popular: true},
		//
		JapInSenScale:        {Name: JapInSenScale, HalfSteps: []int{1, 4, 2, 3}},
		MajorBebopScale:      {Name: MajorBebopScale, HalfSteps: []int{2, 2, 1, 2, 1, 1, 2}},
		DominantBebopScale:   {Name: DominantBebopScale, HalfSteps: []int{2, 2, 1, 2, 2, 1, 1}},
		BluesScale:           {Name: BluesScale, HalfSteps: []int{3, 2, 1, 1, 3}, Popular: true},
		ArabicScale:          {Name: ArabicScale, HalfSteps: []int{1, 3, 1, 2, 1, 3}},
		EnigmaticScale:       {Name: EnigmaticScale, HalfSteps: []int{1, 3, 2, 2, 2, 1}},
		NeapolitanScale:      {Name: NeapolitanScale, HalfSteps: []int{1, 2, 2, 2, 2, 2}},
		NeapolitanMinorScale: {Name: NeapolitanMinorScale, HalfSteps: []int{1, 2, 2, 2, 1, 3}},
		HungarianMinorScale:  {Name: HungarianMinorScale, HalfSteps: []int{2, 1, 3, 1, 1, 3}},
		PhrygianScale:        {Name: PhrygianScale, HalfSteps: []int{1, 2, 2, 2, 1, 2}},
		LydianScale:          {Name: LydianScale, HalfSteps: []int{2, 2, 2, 1, 2, 2}},
		MixolydianScale:      {Name: MixolydianScale, HalfSteps: []int{2, 2, 1, 2, 2, 1}, Popular: true},
		AeolianScale:         {Name: AeolianScale, HalfSteps: []int{2, 1, 2, 2, 1, 2}, Popular: true}, // AKA natural minor scale
		LocrianScale:         {Name: LocrianScale, HalfSteps: []int{1, 2, 2, 1, 2, 2}},
	}
)

// ScaleNotes returns the notes in the scale. The return data contains the
// note numbers (0-11) and the English musical notes
func ScaleNotes(tonic string, scale ScaleName) ([]int, []string) {
	k := KeyInt(tonic, 0) % 12
	scaleKeys := []int{k}
	for _, hs := range ScaleDefs[scale].HalfSteps {
		k += hs
		scaleKeys = append(scaleKeys, k%12)
	}
	notes := []string{}
	for _, k := range scaleKeys {
		notes = append(notes, Notes[k%12])
	}
	return scaleKeys, notes
}
