package abstractfactory

import "fmt"

/*
	AbstractFactory 抽象工厂模式：生成对象存在关联，若无关联则退化成工厂方法模式
*/

//FactoryInterface 工厂接口
type FactoryInterface interface {
	CreateProduct1() ProductInterface // 创建产品1
	CreateProduct2() ProductInterface // 创建产品2
}

//ProductInterface 产品接口
type ProductInterface interface {
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

type Factory1 struct {}

func (Factory1) CreateProduct1() ProductInterface {
	return Factory1Product1{}
}

func (Factory1) CreateProduct2() ProductInterface {
	return Factory1Product2{}
}

type Factory2 struct {}

func (Factory2) CreateProduct1() ProductInterface {
	return Factory1Product1{}
}

func (Factory2) CreateProduct2() ProductInterface {
	return Factory1Product2{}
}
