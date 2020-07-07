package factorymethod

import "testing"

func TestFactoryMethod(t *testing.T) {
	var p Product
	factory1 := new(Factory1)
	p = factory1.Generate(1)
	p.create()
	p = factory1.Generate(2)
	p.create()

	factory2 := new(Factory2)
	factory2.Generate(1).create()
	factory2.Generate(2).create()
}
