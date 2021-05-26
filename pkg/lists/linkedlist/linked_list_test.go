package linkedlist

import (
	"fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList(1, 2, 3, 4)
	t.Log(list.String())
}

func TestLinkedList_Insert(t *testing.T) {
	list := NewList("aaa", "bbb", "ccc")
	t.Log(list)

	index := 2
	value := "hello"
	if err := list.Insert(index, value); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf("Insert %v at pos %d", value, index)
		t.Log(list)
	}
}

func TestLinkedList_Delete(t *testing.T) {
	list := NewList(1, 2, 3)
	t.Log(list)

	index := 1
	if value, err := list.Delete(index); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf("Delete %v at pos %d", value, index)
		t.Log(list)
	}
}

func TestLinkedList_Get(t *testing.T) {
	list := NewList("hello", "world", "something")

	index := 1
	if value, err := list.Get(index); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(list)
		t.Logf(fmt.Sprintf("Get %v at pos %d", value, index))
	}
}
