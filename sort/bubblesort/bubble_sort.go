package bubblesort

import (
	"gods/sort"
)

type BubbleSort struct {
}

func (*BubbleSort) Sort(s sort.Sorter) sort.Sorter {
	for i := 0; i < s.Len(); i++ {

		flag := false
		for j := 0; j < s.Len()-i-1; j++ {
			if s.More(s.Values[j], s.Values[j+1]) {
				s.Swap(j, j+1)
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	return s
}
