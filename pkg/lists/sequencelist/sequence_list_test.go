package sequencelist

import (
	"fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList("hello", "world")
	t.Log(list)
}

func TestSequenceList_Insert(t *testing.T) {
	list := NewList(1, 2, 3)
	t.Log(list)

	index := 2
	value := "hello"
	if err := list.Insert(index, value); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf(fmt.Sprintf("Insert %v at pos %d", value, index))
		t.Log(list)
	}
}

func TestSequenceList_Delete(t *testing.T) {
	list := NewList("a", "b", "c", "d")
	t.Log(list)

	index := 2
	if value, err := list.Delete(index); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf(fmt.Sprintf("Delete %v at pos %d", value, index))
		t.Log(list)
	}
}

func TestSequenceList_Get(t *testing.T) {
	list := NewList("a", "b", "c")

	index := 1
	if value, err := list.Get(index); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf(fmt.Sprintf("Get %v at pos %d", value, index))
		t.Log(list)
	}
}
