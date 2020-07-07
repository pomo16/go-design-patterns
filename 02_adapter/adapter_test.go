package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	fmt.Println(res)
}
