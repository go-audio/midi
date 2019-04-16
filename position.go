package midi

import (
	"fmt"
)

// Position indicates where we are at in the project.
type Position struct {
	// Bar is the bar/measure of position (index 0)
	Bar uint64
	// Beat is the beat position within the bar, index 0, the max value depends
	// of the time signature
	Beat uint32
	// Div is the division position within the beat
	Div uint32
	// Ticks are the leftover ticks
	Ticks uint32
}

func (p Position) String() string {
	return fmt.Sprintf("Bar: %d, Beat: %d, Div: %d, Ticks: %d", p.Bar, p.Beat, p.Div, p.Ticks)
}

// ToTicks converts a position to an absolute tick number.
func (p Position) ToTicks(ppq uint16) uint64 {
	ppqUint64 := uint64(ppq)
	barLen := 4 * ppqUint64
	divLen := ppqUint64 / 4

	ticks := uint64(p.Bar) * barLen
	ticks += (uint64(p.Beat) * ppqUint64)
	ticks += uint64(p.Div) * divLen
	ticks += uint64(p.Ticks)

	return ticks
}

// TickPosition returns the position of the passed tick
func TickPosition(tick uint64, ppq uint16) Position {
	// TODO: support more than 4/4 time signature
	if tick == 0 {
		return Position{}
	}
	barLen := 4 * uint64(ppq)
	divLen := uint32(ppq) / 4

	p := Position{Bar: (tick / barLen)}
	leftOver := tick % uint64(barLen)
	p.Beat = uint32(leftOver / uint64(ppq))

	if rest := uint32(leftOver) % uint32(ppq); rest != 0 {
		p.Div = rest / divLen
		p.Ticks = rest % divLen
	}
	return p
}

// Position returns the start position of the event
// in index zero.
func (e *Event) Position(ppq uint16) Position {
	if e == nil {
		return Position{}
	}
	return TickPosition(e.AbsTicks, ppq)
}
