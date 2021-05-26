package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// LinkedList 带头节点的单链表实现
// head 链表指向头节点的指针
// size 链表中元素个数
type LinkedList struct {
	head *node
	size int
}

// node 节点定义
// data 节点数据
// next 下一个节点指针
type node struct {
	data interface{}
	next *node
}

/**
 * NewList 初始化指定元素的带头节点的单链表，返回链表的引用
 */
func NewList(values ...interface{}) *LinkedList {
	list := &LinkedList{head: &node{}}
	for idx, item := range values {
		_ = list.Insert(idx, item)
	}
	return list
}

func (l *LinkedList) Insert(idx int, value interface{}) error {
	if idx < 0 || idx > l.size {
		return errors.New(fmt.Sprintf("{ LinkList [ Insert Error ]: invalid position: %v, it should in [0, %d] }", idx, l.size))
	}

	pos := l.head
	for i := idx; i > 0; i-- {
		pos = pos.next
	}

	s := &node{data: value}
	s.next = pos.next
	pos.next = s
	l.size++

	return nil
}

func (l *LinkedList) Delete(idx int) (interface{}, error) {
	if idx < 0 || idx >= l.size {
		return nil, errors.New(fmt.Sprintf("{ LinkList [ Delete Error ]: invalid position: %v, it should in [0, %d) }", idx, l.size))
	}

	pos := l.head
	for i := idx; i > 0; i-- {
		pos = pos.next
	}

	s := pos.next
	pos.next = s.next
	s.next = nil
	l.size--

	return s.data, nil
}

func (l *LinkedList) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= l.size {
		return nil, errors.New(fmt.Sprintf("{ LinkList [ Get Error ]: Out of boundary, it should in [0, %d) }", l.size))
	}

	pos := l.head
	for i := idx; i >= 0; i-- {
		pos = pos.next
	}

	return pos.data, nil
}

func (l *LinkedList) Clear() {
	l.head.next = nil
	l.size = 0
}

func (l *LinkedList) String() string {
	str := fmt.Sprintf("{ LinkedList [ head: (%v), size: %d ] }: ", l.head.next.data, l.size)

	values := make([]string, 0)
	l.Traverse(func(value interface{}) {
		values = append(values, fmt.Sprintf("(%v)", value))
	})
	str += strings.Join(values, " -> ")

	return str
}

func (l *LinkedList) Traverse(visit func(interface{})) {
	for item := l.head.next; item != nil; item = item.next {
		visit(item.data)
	}
}

func (l *LinkedList) Size() int {
	return l.size
}

