package decorator

import "fmt"

/*
	decorator 装饰模式
*/

//IDecorate 抽象装饰器
type IDecorate interface {
	Do()
}

//Decorate 装饰器实现
type Decorate struct {
	decorate IDecorate
}

//DecorateFun 装饰方法
func (c *Decorate) DecorateFun(i IDecorate) {
	c.decorate = i
}

//Do 被修饰方法
func (c *Decorate) Do() {
	if c.decorate != nil {
		c.decorate.Do()
	}
}

//DecorateA 具体修饰器
type DecorateA struct {
	Base Decorate
}

func (c *DecorateA) Do() {
	fmt.Println("run A decorate")
	c.Base.Do()
}

//DecorateB 具体修饰器B
type DecorateB struct {
	Base Decorate
}

func (c *DecorateB) Do() {
	fmt.Println("run B decorate")
	c.Base.Do()
}
