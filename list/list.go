package list

import (
	"cmp"
	"iter"
	"slices"
)

const minSize = 256

// ordered list with an underlying generic slice of elements
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
