package grid

// Res is the resolution of the grid
type Res string

// StepsInBeat returns the number of steps to fill a beat
func (g Res) StepsInBeat() uint64 {
	switch g {
	case One8:
		return 2
	case One16:
		return 4
	case One32:
		return 8
	case One64:
		return 16
	}
	return 1
}

const (
	One4  Res = "1/4"
	One8  Res = "1/8"
	One16 Res = "1/16"
	One32 Res = "1/32"
	One64 Res = "1/64"
	// TODO: add triplets
)

// StepSize returns the size of a step in ticks given its grid resolution and ppqn
func (g Res) StepSize(ppqn uint16) uint64 {
	switch g {
	case One4:
		return uint64(ppqn)
	case One8:
		return uint64(ppqn / 2)
	case One16:
		return uint64(ppqn / 4)
	case One32:
		return uint64(ppqn / 8)
	case One64:
		return uint64(ppqn / 16)
	}
	return 0
}
