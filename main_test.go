package main

import "testing"

func TestMedian(t *testing.T) {
	var m Median
	if _, ok := m.Median(); ok {
		t.Error("empty stream has no median")
	}
	steps := []struct {
		add  int
		want float64
	}{
		{5, 5},
		{15, 10},
		{1, 5},
		{3, 4},
		{2, 3},
		{8, 4},
		{7, 5},
	}
	for _, s := range steps {
		m.Add(s.add)
		got, ok := m.Median()
		if !ok || got != s.want {
			t.Errorf("after Add(%d): median = %v, want %v", s.add, got, s.want)
		}
	}
}

func TestDuplicatesAndOrder(t *testing.T) {
	var m Median
	for _, v := range []int{2, 2, 2, 2} {
		m.Add(v)
	}
	if got, _ := m.Median(); got != 2 {
		t.Errorf("median of all 2s = %v", got)
	}
	// Descending inserts still keep the slice sorted.
	var d Median
	for v := 5; v >= 1; v-- {
		d.Add(v)
	}
	if got, _ := d.Median(); got != 3 {
		t.Errorf("median of 1..5 = %v, want 3", got)
	}
}
