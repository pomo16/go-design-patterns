package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}
	res := sub.Do()
	fmt.Println(res)
}
