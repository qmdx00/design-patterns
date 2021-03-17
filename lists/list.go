package lists

// List 线性表顶层接口，定义线性表的公共函数
type List interface {
	Insert(int, interface{}) error   // 在指定索引处插入指定元素
	Delete(int) (interface{}, error) // 删除指定索引处的元素并返回
	Get(int) (interface{}, error)    // 获取指定索引处的元素并返回
	Traverse(func(interface{}))      // 遍历链表并用指定函数处理每个元素
	Clear()                          // 清空链表
}
