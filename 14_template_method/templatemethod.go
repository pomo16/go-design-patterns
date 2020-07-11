package templatemethod

import "fmt"

/*
	templatemethod 模板方法模式
*/

//AbstractClass 实体抽象接口
type AbstractClass interface {
	DoSomeThing(name string)
}

//template 模板
type template struct {
	templateMethod
	name string
}

//templateMethod 模板方法
type templateMethod interface {
	method1()
	method2()
}

//newTemplate 创建模板
func newTemplate(tm templateMethod) *template {
	return &template{
		templateMethod: tm,
	}
}

//DoSomeThing 实体方法
func (t *template) DoSomeThing(name string) {
	t.name = name
	fmt.Println("prepare")
	t.templateMethod.method1()
	t.templateMethod.method2()
	fmt.Println("finish")
}

//method2 模板方法默认实现
func (t *template) method2() {
	fmt.Println("Default method2")
}

//ConcreteClass1 实体类1，实现了所有模板方法
type ConcreteClass1 struct {
	*template
}

func NewConcreteClass1() AbstractClass {
	concreteClass := &ConcreteClass1{}
	template := newTemplate(concreteClass)
	concreteClass.template = template
	return concreteClass
}

func (d *ConcreteClass1) method1() {
	fmt.Printf("ConcreteClass1 %s method1\n", d.name)
}

func (d *ConcreteClass1) method2() {
	fmt.Printf("ConcreteClass1 %s method2\n", d.name)
}

//ConcreteClass2 实体类2，实现部分模板方法
type ConcreteClass2 struct {
	*template
}

func NewConcreteClass2() AbstractClass {
	concreteClass := &ConcreteClass2{}
	template := newTemplate(concreteClass)
	concreteClass.template = template
	return concreteClass
}

func (d *ConcreteClass2) method1() {
	fmt.Printf("ConcreteClass2 %s method1\n", d.name)
}
