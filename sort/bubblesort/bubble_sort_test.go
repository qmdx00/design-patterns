package bubblesort

import (
	"gods/sort"
	"testing"
)

func TestBubbleSort_Sort(t *testing.T) {
	var testCase = []struct {
		in       []interface{}
		expected []int
	}{
		{[]interface{}{1, 4, 2, 3}, []int{1, 2, 3, 4}},
		{[]interface{}{1, 1, 3, 2}, []int{1, 1, 2, 3}},
		{[]interface{}{1}, []int{1}},
	}

	bubbleSort := BubbleSort{}
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
		res := bubbleSort.Sort(sort.NewSorter(c.in, sort.IntComparator))
		if !equal(res.Values, c.expected) {
			t.Fatalf("values = %v, expected %v", res.Values, c.expected)
		}
	}
}


func BenchmarkBubbleSort_Sort(b *testing.B) {
	bubbleSort := BubbleSort{}

	for i := 0; i < b.N; i++ {
		bubbleSort.Sort(sort.NewSorter([]interface{}{1, 3, 5, 6, 9, 8, 4, 7, 2}, sort.IntComparator))
	}
}
