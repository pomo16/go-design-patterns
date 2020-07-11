package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	var ctx Context
	ctx = Context{Strategy: BubbleSortStrategy{}}
	ctx.Exec()

	ctx = Context{Strategy: MergeSortStrategy{}}
	ctx.Exec()
}
