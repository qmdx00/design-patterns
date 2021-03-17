package linkedlist

import (
	"fmt"
	"gods/common"
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList(1, 2, 3, 4, 5, 6)
	t.Log(list.String())
}

func TestLinkedList_Insert(t *testing.T) {
	list := NewList("aaa", "bbb", "ccc")

	index := 3
	if err := list.Insert(index, "hello"); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(list)
	}
}

func TestLinkedList_Delete(t *testing.T) {
	list := NewList(1, 2, 3)

	index := 2
	if value, err := list.Delete(index); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf("delete %v at pos %d", value, index)
	}
}

func TestLinkedList_Get(t *testing.T) {
	list := NewList("hello", "world")

	index := 1
	if value, err := list.Get(index); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log(list)
		t.Logf(common.C.PrintMsg(common.Green, fmt.Sprintf("Get %v at pos %d", value, index)))
	}
}
