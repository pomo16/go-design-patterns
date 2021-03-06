package visitor

import "testing"

func TestVisitor(t *testing.T) {
	e := new(Element)
	e.Accept(new(ProductionVisitor)) // output: 这是生产环境
	e.Accept(new(TestingVisitor))    // output: 这是测试环境

	m := new(EnvExample)
	m.Print(new(ProductionVisitor)) // output: 这是生产环境
	m.Print(new(TestingVisitor))    // output: 这是测试环境
}
