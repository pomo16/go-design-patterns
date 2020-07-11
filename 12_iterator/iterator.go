package iterator

/*
	Iterator 迭代器模式
*/

//IAggregate 容器接口
type IAggregate interface {
	Iterator() IIterator
}

//IIterator 迭代器接口
type IIterator interface {
	HasNext() bool
	Current() int
	Next() bool
}

//Aggregate 具体容器实现
type Aggregate struct {
	container []int // 容器中装载 int 型容器
}

//Iterator 创建一个迭代器，并让迭代器中的容器指针指向当前对象
func (a *Aggregate) Iterator() IIterator {
	i := new(Iterator)
	i.aggregate = a
	return i
}

//Iterator 迭代器实现
type Iterator struct {
	cursor    int        // 当前游标
	aggregate *Aggregate // 对应的容器指针
}

//HasNext 判断是否迭代到最后，如果没有，则返回true
func (i *Iterator) HasNext() bool {
	if i.cursor+1 < len(i.aggregate.container) {
		return true
	}
	return false
}

//Current 获取当前迭代元素（从容器中取出当前游标对应的元素）
func (i *Iterator) Current() int {
	return i.aggregate.container[i.cursor]
}

//Next 将游标指向下一个元素
func (i *Iterator) Next() bool {
	if i.cursor < len(i.aggregate.container) {
		i.cursor++
		return true
	}
	return false
}
