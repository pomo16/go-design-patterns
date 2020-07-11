package interpreter

import (
	"strconv"
	"strings"
)

/*
	interpreter 解释器模式：以整数加减计算解析器为例
*/

//Node 解释器接口，需要被解析的对象都要实现解析器接口
type Node interface {
	Interpret() int
}

//ValNode 数值
type ValNode struct {
	val int
}

//Interpret 数值解析
func (n *ValNode) Interpret() int {
	return n.val
}

//AddNode 加法解析器
type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

//MinNode 减法解析器
type MinNode struct {
	left, right Node
}

func (n *MinNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

//Calculator 计算器
type Calculator struct {
	exp   []string
	index int
	prev  Node
}

//Parse 计算器解析
func (c *Calculator) Parse(exp string) {
	c.exp = strings.Split(exp, " ")

	for {
		if c.index >= len(c.exp) {
			return
		}
		switch c.exp[c.index] {
		case "+":
			c.prev = c.newAddNode()
		case "-":
			c.prev = c.newMinNode()
		default:
			c.prev = c.newValNode()
		}
	}
}

func (c *Calculator) newAddNode() Node {
	c.index++
	return &AddNode{
		left:  c.prev,
		right: c.newValNode(),
	}
}

func (c *Calculator) newMinNode() Node {
	c.index++
	return &MinNode{
		left:  c.prev,
		right: c.newValNode(),
	}
}

func (c *Calculator) newValNode() Node {
	v, _ := strconv.Atoi(c.exp[c.index])
	c.index++
	return &ValNode{
		val: v,
	}
}

func (c *Calculator) Result() Node {
	return c.prev
}
