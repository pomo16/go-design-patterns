package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	var health int
	ctx := NewContext(health)
	for health = -11; health < 10; health += 10 {
		ctx.SetHealth(health)
		ctx.State.Comment()
		ctx.State.Post()
		ctx.State.View()
		fmt.Println()
	}
}
