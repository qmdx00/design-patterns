package sequencelist

import (
	"errors"
	"fmt"
	"gods/common"
	"strings"
)

const (
	InitCap    = 10 // 初始化容量
	GrowthSize = 1  // 扩容 2^GrowthSize 倍
)

// SequenceList 顺序表实现
// items 元素列表
// size 元素个数
// cap 容量大小
type SequenceList struct {
	items []element
	size  int
	cap   int
}

// element 顺序表元素
// data 数据内容
type element struct {
	data interface{}
}

func NewList(values ...interface{}) *SequenceList {
	list := &SequenceList{
		items: make([]element, InitCap, InitCap),
		size:  0,
		cap:   InitCap,
	}
	for idx, value := range values {
		_ = list.Insert(idx, value)
	}
	return list
}

func (l *SequenceList) Insert(idx int, value interface{}) error {
	if idx < 0 || idx > l.size {
		return errors.New(common.C.PrintMsg(common.Red,
			fmt.Sprintf("{ SequenceList [ Insert Error ]: invalid position: %v, it should in [0, %d] }", idx, l.size)))
	}

	// 扩容
	if l.size+1 >= l.cap {
		l.growth()
	}

	// 插入点之后的所有元素后移一位
	for i := l.size + 1; i > idx; i-- {
		l.items[i] = l.items[i-1]
	}

	// 插入元素
	l.items[idx] = element{data: value}
	l.size++

	return nil
}

func (l *SequenceList) Delete(idx int) (interface{}, error) {
	if idx < 0 || idx >= l.size {
		return nil, errors.New(common.C.PrintMsg(common.Red,
			fmt.Sprintf("{ SequenceList [ Delete Error ]: invalid position: %v, it should in [0, %d) }", idx, l.size)))
	}

	item := l.items[idx]
	for i := idx; i < l.size; i++ {
		l.items[i] = l.items[i+1]
	}
	l.size--

	return item.data, nil
}

func (l *SequenceList) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= l.size {
		return nil, errors.New(common.C.PrintMsg(common.Red,
			fmt.Sprintf("{ SequenceList [ Get Error ]: Out of boundary, it should in [0, %d) }", l.size)))
	}

	return l.items[idx].data, nil
}

func (l *SequenceList) Traverse(visit func(interface{})) {
	if l.size > 0 {
		for i := 0; i < l.size; i++ {
			visit(l.items[i].data)
		}
	}
}

func (l *SequenceList) Clear() {
	l.items = make([]element, InitCap, InitCap)
	l.size = 0
	l.cap = InitCap
}

func (l SequenceList) String() string {
	str := fmt.Sprintf("{ SequenceList [ size: %d, cap: %d ] }: ", l.size, l.cap)

	values := make([]string, 0)
	l.Traverse(func(value interface{}) {
		values = append(values, fmt.Sprintf("(%v)", value))
	})

	str += strings.Join(values, ", ")

	return common.C.PrintMsg(common.Green, str)
}

func (l *SequenceList) growth() {
	l.cap <<= GrowthSize
	newItems := make([]element, l.cap, l.cap)
	copy(newItems, l.items)
	l.items = newItems
}
