package insertsort

import "gods/sort"

type InsertSort struct {
}

func (*InsertSort) Sort(s sort.Sorter) sort.Sorter {
	if s.Len() <= 1 {
		return s
	}

	for i := 1; i < s.Len(); i++ {
		value := s.Values[i]
		j := i - 1

		for j >= 0 {
			if s.More(s.Values[j], value) {
				s.Values[j+1] = s.Values[j]
				j--
			} else {
				break
			}
		}
		s.Values[j+1] = value
	}

	return s
}
