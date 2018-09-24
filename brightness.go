package main

import "fmt"

// Brightness is an int that is constrained between 0 and max
type Brightness struct {
	value, max int
}

// Percent is the brightness ratio to the max (expressed where 100 is 100%)
func (b Brightness) Percent() int {
	return int(100 * b.value / b.max)
}

// Value is the raw brightness (note that this range probably varies for different displays)
func (b Brightness) Value() int {
	return b.value
}

// Increment (or decrement with a negative delta) by a percent (expressed where 100 is 100%)
func (b *Brightness) Increment(delta int) {
	b.setPercent(b.Percent() + delta)
}

func (b Brightness) String() string {
	return fmt.Sprintf("Brightness: %v%% (%v/%v)", b.Percent(), b.value, b.max)
}

func (b *Brightness) setPercent(p int) {
	b.setValue(b.max * p / 100)
}

func (b *Brightness) setValue(v int) {
	switch {
	case v < 0:
		b.value = 0
	case v > b.max:
		b.value = b.max
	default:
		b.value = v
	}
}
