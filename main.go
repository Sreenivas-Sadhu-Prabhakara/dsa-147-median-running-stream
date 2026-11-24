package main

import (
	"fmt"
	"sort"
)

// Median maintains the running median of a stream by keeping every value in
// a sorted slice: each Add is a binary-searched insertion, each Median a
// direct index. The two-heap approach (see dsa-072) trades this O(n) insert
// for O(log n); this version is simpler and faster for small streams.
type Median struct {
	sorted []int
}

// Add inserts v, keeping the slice sorted.
func (m *Median) Add(v int) {
	i := sort.SearchInts(m.sorted, v)
	m.sorted = append(m.sorted, 0)
	copy(m.sorted[i+1:], m.sorted[i:])
	m.sorted[i] = v
}

// Median returns the current median; ok is false while empty.
func (m *Median) Median() (float64, bool) {
	n := len(m.sorted)
	if n == 0 {
		return 0, false
	}
	if n%2 == 1 {
		return float64(m.sorted[n/2]), true
	}
	return (float64(m.sorted[n/2-1]) + float64(m.sorted[n/2])) / 2, true
}

func main() {
	var m Median
	for _, v := range []int{5, 15, 1, 3} {
		m.Add(v)
		med, _ := m.Median()
		fmt.Println(med)
	}
}
