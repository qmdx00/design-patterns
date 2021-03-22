package sort

import (
	"fmt"
	"strconv"
	"strings"
)

type Sorter struct {
	Values     []interface{}
	comparator Comparator
}

func NewSorter(values []interface{}, comparator Comparator) Sorter {
	return Sorter{
		Values:     values,
		comparator: comparator,
	}
}

func (s *Sorter) Len() int {
	return len(s.Values)
}

func (s *Sorter) Less(x, y interface{}) bool {
	return s.comparator(x, y)
}

func (s *Sorter) More(x, y interface{}) bool {
	return s.comparator(y, x)
}

func (s *Sorter) Swap(x, y int) {
	s.Values[x], s.Values[y] = s.Values[y], s.Values[x]
}

func (s Sorter) String() string {
	values := make([]string, 0)

	for _, value := range s.Values {
		switch value.(type) {
		case int:
			values = append(values, strconv.Itoa(value.(int)))
		case string:
			values = append(values, value.(string))
		default:
			return fmt.Sprintf("{ Sorter [ Values: %v ] }", s.Values)
		}
	}

	return fmt.Sprintf("{ Sorter [ Values: (%v) ] }", strings.Join(values, ", "))
}
