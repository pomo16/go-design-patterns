package prototype

/*
	Prototype 原型模式：相当于深拷贝方法
*/

//Cloneable 原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type PrototypeStruct struct {
	name string
}

//Clone 实现克隆方法（深拷贝）
func (p *PrototypeStruct) Clone() *PrototypeStruct {
	res := *p
	return &res
}
