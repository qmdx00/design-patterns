package main

import (
	"fmt"
	"github.com/qmdx00/gods/pkg/lists/linkedlist"
)

func main() {
	l := linkedlist.NewList(1, 2, 3)
	fmt.Println(l)

	if v, err := l.Get(3); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v)
	}

	if err := l.Insert(3, 4); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(l)
	}

	if _, err := l.Delete(2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(l)
	}

	l.Traverse(func(i interface{}) {
		fmt.Println(i)
	})

	l.Clear()
}
