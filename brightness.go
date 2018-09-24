package main

import "fmt"

// Brightness is an int64 that is constrained between 0 and max
type Brightness struct {
	value, max int64
}

// Percent is the brightness ratio to the max (expressed where 100 is 100%)
func (b Brightness) Percent() int64 {
	return int64(100 * b.value / b.max)
}

// Value is the raw brightness (note that this range probably varies for different displays)
func (b Brightness) Value() int64 {
	return b.value
}

// Increment (or decrement with a negative delta) by a percent (expressed where 100 is 100%)
func (b *Brightness) Increment(delta int64) {
	b.setPercent(b.Percent() + delta)
}

func (b Brightness) String() string {
	return fmt.Sprintf("Brightness: %v%% (%v/%v)", b.Percent(), b.value, b.max)
}

func (b *Brightness) setPercent(p int64) {
	b.setValue(b.max * p / 100)
}

func (b *Brightness) setValue(v int64) {
	switch {
	case v < 0:
		b.value = 0
	case v > b.max:
		b.value = b.max
	default:
		b.value = v
	}
}
