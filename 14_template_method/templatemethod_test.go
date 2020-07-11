package templatemethod

import "testing"

func TestTemplateMethod(t *testing.T) {
	var concrete1 AbstractClass = NewConcreteClass1()
	concrete1.DoSomeThing("c1")

	concrete2 := NewConcreteClass2()
	concrete2.DoSomeThing("c2")

	//Output:
	//prepare
	//ConcreteClass1 c1 method1
	//ConcreteClass1 c1 method2
	//finish
	//prepare
	//ConcreteClass2 c2 method1
	//Default method2
	//finish
}
