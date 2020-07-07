package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	var f FactoryInterface  // 特意以这种方式声明，更好的体会抽象工厂模式的好处
	f = new(Factory1)
	b := f.CreateProduct1()
	b.create()
}
