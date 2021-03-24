package insertsort

import (
	"gods/sort"
	"testing"
)

func TestInsertSort_Sort(t *testing.T) {
	var testCase = []struct {
		in       []interface{}
		expected []int
	}{
		{[]interface{}{1, 4, 2, 3}, []int{1, 2, 3, 4}},
		{[]interface{}{1, 1, 3, 2}, []int{1, 1, 2, 3}},
		{[]interface{}{1}, []int{1}},
	}

	insertSort := InsertSort{}
	equal := func(a []interface{}, b []int) bool {
		if len(a) != len(b) {
			return false
		}

		for i := 0; i < len(a); i++ {
			if a[i].(int) != b[i] {
				return false
			}
		}

		return true
	}

	for _, c := range testCase {
		res := insertSort.Sort(sort.NewSorter(c.in, sort.IntComparator))
		if !equal(res.Values, c.expected) {
			t.Fatalf("values = %v, expected %v", res.Values, c.expected)
		}
	}
}
