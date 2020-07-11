package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	c := &Calculator{}
	c.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := c.Result().Interpret()
	correct := 1
	fmt.Println(res == correct)
}
