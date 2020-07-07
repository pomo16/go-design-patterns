package factorymethod

import "fmt"

/*
	FactoryMethod 工厂方法模式
*/

//Factory 工厂接口
type Factory interface {
	Generate() Product
}

//Product 产品接口
type Product interface {
	create()
}

//Factory1Product1
type Factory1Product1 struct{}

func (Factory1Product1) create() {
	fmt.Println("this is Factory1's Product1")
}

//Factory1Product2
type Factory1Product2 struct{}

func (Factory1Product2) create() {
	fmt.Println("this is Factory1's Product2")
}

//Factory2Product1
type Factory2Product1 struct{}

func (Factory2Product1) create() {
	fmt.Println("this is Factory2's Product1")
}

//Factory2Product2
type Factory2Product2 struct{}

func (Factory2Product2) create() {
	fmt.Println("this is Factory2's Product2")
}

//Factory1
type Factory1 struct{}

func (f Factory1) Generate(t int) Product {
	switch t {
	case 1:
		return Factory1Product1{}
	case 2:
		return Factory1Product2{}
	default:
		return nil
	}
}

//Factory2
type Factory2 struct{}

func (f Factory2) Generate(t int) Product {
	switch t {
	case 1:
		return Factory2Product1{}
	case 2:
		return Factory2Product2{}
	default:
		return nil
	}
}
