package quicksort

import (
	"gods/sort"
)

type QuickSort struct {
}

func (q *QuickSort) Sort(s sort.Sorter) sort.Sorter {
	if s.Len() <= 1 {
		return s
	}
	qSort(s, 0, s.Len()-1)
	return s
}

func qSort(s sort.Sorter, l, h int) {
	if l < h {
		p := partition(s, l, h)
		qSort(s, l, p-1)
		qSort(s, p+1, h)
	}
}

func partition(s sort.Sorter, l, h int) int {
	pivot := l
	index := l + 1

	for i := index; i <= h; i++ {
		if s.Less(s.Values[i], s.Values[pivot]) {
			s.Swap(i, index)
			index++
		}
	}
	s.Swap(pivot, index-1)
	return index - 1
}

func partition2(s sort.Sorter, l, h int) int {
	pivot := s.Values[l]

	for l < h {
		// 从右往左找比pivot小的元素
		for l < h && !s.Less(s.Values[h], pivot) {
			h--
		}
		s.Values[l] = s.Values[h]

		// 从左往右找比pivot大的元素
		for l < h && !s.More(s.Values[l], pivot) {
			l++
		}
		s.Values[h] = s.Values[l]
	}

	s.Values[l] = pivot
	return l
}
