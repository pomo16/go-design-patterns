package simplefactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	//创建一个工厂类，在应用中可以将这个工厂类实例作为一个全局变量
	factory := new(Factory)

	//在工厂类中传入不同的参数，获取不同的实例
	p1 := factory.NewProduct(1)
	p1.create() // output:   this is product 1

	p2 := factory.NewProduct(2)
	p2.create() // output:   this is product 2
}
