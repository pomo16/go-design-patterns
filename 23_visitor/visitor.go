package visitor

import "fmt"

/*
	visitor 访问者模式
*/

//IVisitor 访问者接口
type IVisitor interface {
	Visit() // 访问者的访问方法
}

//ProductionVisitor 生产环境访问者，实现IVisitor
type ProductionVisitor struct{}

func (v ProductionVisitor) Visit() {
	fmt.Println("这是生产环境")
}

//TestingVisitor 测试环境访问者，实现IVisitor
type TestingVisitor struct{}

func (t TestingVisitor) Visit() {
	fmt.Println("这是测试环境")
}

//IElement 元素接口
type IElement interface {
	Accept(visitor IVisitor)
}

//Element 实现元素接口
type Element struct{}

func (el Element) Accept(visitor IVisitor) {
	visitor.Visit()
}

type EnvExample struct {
	Element
}

func (e EnvExample) Print(visitor IVisitor) {
	e.Element.Accept(visitor)
}
