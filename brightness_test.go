package main

import "testing"

func TestPercent(t *testing.T) {
	brightness := Brightness{value: 1, max: 5}
	percent := brightness.Percent()
	if percent != 20 {
		t.Errorf("Percent was incorrect, got: %d, want: %d.", percent, 20)
	}
}

func TestIncrement(t *testing.T) {
	tables := []struct {
		value int
		max   int
		inc   int
		want  int
	}{
		{0, 10, 0, 0},
		{1, 10, 10, 2},
		{1, 10, -10, 0},
		{0, 10, -10, 0},
		{0, 10, 999, 10},
	}

	for _, table := range tables {
		brightness := Brightness{value: table.value, max: table.max}
		brightness.Increment(table.inc)
		result := brightness.Value()
		if result != table.want {
			t.Errorf("Increment %d from %d/%d result was incorrect, got: %d, want: %d.", table.inc, table.value, table.max, result, table.want)
		}
	}
}
