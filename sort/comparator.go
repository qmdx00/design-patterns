package sort

type Comparator func(x, y interface{}) bool

func IntComparator(x, y interface{}) bool {
	a, b := x.(int), y.(int)
	return a < b
}
