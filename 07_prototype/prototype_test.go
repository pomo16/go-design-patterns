package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	p1 := new(PrototypeStruct)
	p2 := p1.Clone()
	fmt.Println(p1 == p2) //false
}
