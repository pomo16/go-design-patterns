package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder := new(ProductBuilder)
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	fmt.Println(res)
}
