package simplefactory

import "fmt"

/*
	SimpleFactory 简单工厂
*/

//Factory 工厂结构体
type Factory struct{}

//Product 产品接口，可理解为工厂生产的产品模板，可扩展多个方法
type Product interface {
	create()
}

// Product1
type Product1 struct{}

func (Product1) create() {
	fmt.Println("this is Product1")
}

// Product2
type Product2 struct{}

func (Product2) create() {
	fmt.Println("this is Product2")
}

//NewProduct 工厂生产方法
func (f Factory) NewProduct(t int) Product {
	switch t {
	case 1:
		return Product1{}
	case 2:
		return Product2{}
	default:
		return nil
	}
}
