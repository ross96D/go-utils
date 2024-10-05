package list

import (
	"cmp"
	"iter"
	"slices"
)

const minSize = 256

type Ordered interface {
	// on x.Compare(y) 0 is x == y, less than 0 is x < y, more than 0 is x > y
	Compare(any) int
}

type ISortedList[T Ordered] struct {
	values []T
}

func (l *ISortedList[T]) Elem(idx uint) *T {
	return &l.values[uint(idx)]
}

func (l *ISortedList[T]) Iter(p ...uint) iter.Seq2[int, *T] {
	start := uint(0)
	end := uint(len(l.values))
	if len(p) >= 1 {
		start = p[0]
	}
	if len(p) >= 2 {
		start = p[1]
	}
	return func(yield func(int, *T) bool) {
		for i := start; i < end; i++ {
			if !yield(int(i), &l.values[i]) {
				break
			}
		}
	}
}

func (l *ISortedList[T]) ensureInit() {
	if l.values == nil {
		l.values = make([]T, 0, minSize)
	}
}

func (l *ISortedList[T]) Search(v T) (int, bool) {
	n := len(l.values)
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if l.values[h].Compare(v) < 0 {
			i = h + 1 // preserves x[i-1] < target
		} else {
			j = h // preserves x[j] >= target
		}
	}
	// i == j, x[i-1] < target, and x[j] (= x[i]) >= target  =>  answer is i.
	return i, i < n && (l.values[i].Compare(v) == 0)
}

func (l *ISortedList[T]) append(v T) {
	l.ensureInit()
	pos := l.pos(v)
	l.values = append(l.values, v)
	copy(l.values[pos+1:], l.values[pos:len(l.values)-1])
	l.values[pos] = v
}

func (l *ISortedList[T]) Append(v ...T) {
	l.ensureInit()
	for _, i := range v {
		l.append(i)
	}
}

func (l *ISortedList[T]) pos(v T) int {
	if len(l.values) == 0 {
		return 0
	}

	// there is room for optimizations here if we take in account the minium effort
	// to move the values to the right
	low := 0
	high := len(l.values) - 1
	var pos int

	for low <= high {
		pos = (high + low) / 2

		if v.Compare(l.values[pos]) < 0 {
			high = pos - 1
		} else if v.Compare(l.values[pos]) > 0 {
			low = pos + 1
		} else {
			break
		}
	}

	switch pos {
	case 0, len(l.values) - 1:
		if v.Compare(l.values[pos]) < 0 {
			return pos
		} else {
			return pos + 1
		}
	default:
		return pos
	}
}

// ordered list with primitives only
type SortedList[T cmp.Ordered] struct {
	values []T
}

func (l *SortedList[T]) Elem(idx uint) T {
	return l.values[uint(idx)]
}

func (l *SortedList[T]) Iter(p ...uint) iter.Seq2[int, T] {
	start := uint(0)
	end := uint(len(l.values))
	if len(p) >= 1 {
		start = p[0]
	}
	if len(p) >= 2 {
		start = p[1]
	}
	return func(yield func(int, T) bool) {
		for i := start; i < end; i++ {
			if !yield(int(i), l.values[i]) {
				break
			}
		}
	}
}

func (l *SortedList[T]) ensureInit() {
	if l.values == nil {
		l.values = make([]T, 0, minSize)
	}
}

func (l *SortedList[T]) Search(v T) (int, bool) {
	return slices.BinarySearch(l.values, v)
}

func (l *SortedList[T]) append(v T) {
	l.ensureInit()
	pos := l.pos(v)
	l.values = append(l.values, v)
	copy(l.values[pos+1:], l.values[pos:len(l.values)-1])
	l.values[pos] = v
}

func (l *SortedList[T]) Append(v ...T) {
	l.ensureInit()
	for _, i := range v {
		l.append(i)
	}
}

func (l *SortedList[T]) pos(v T) int {
	if len(l.values) == 0 {
		return 0
	}

	// there is room for optimizations here if we take in account the minium effort
	// to move the values to the right
	low := 0
	high := len(l.values) - 1
	var pos int

	for low <= high {
		pos = (high + low) / 2

		if v < l.values[pos] {
			high = pos - 1
		} else if v > l.values[pos] {
			low = pos + 1
		} else {
			break
		}
	}

	switch pos {
	case 0, len(l.values) - 1:
		if v < l.values[pos] {
			return pos
		} else {
			return pos + 1
		}
	default:
		return pos
	}
}
