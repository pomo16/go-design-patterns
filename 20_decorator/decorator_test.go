package decorator

import "testing"

func TestDecorator(t *testing.T) {
	a := new(DecorateA)
	b := new(DecorateB)
	b.Base.DecorateFun(a)
	b.Do()
}
