package strategy

import "fmt"

/*
	strategy 策略模式
*/

//IStrategy 策略接口
type IStrategy interface {
	SortList() // 对列表进行排序
}

//BubbleSortStrategy 定义策略，这里定义了冒泡排序和归并排序两种策略
type BubbleSortStrategy struct {}

func ( b BubbleSortStrategy) SortList()  {
	fmt.Println("这是冒泡排序")
}

type MergeSortStrategy struct {}

func (m MergeSortStrategy) SortList()  {
	fmt.Println("这是归并排序")
}

//Context 定义上下文
type Context struct {
	Strategy IStrategy  // 上下文中指定的策略
}

func (c Context) Exec() {
	c.Strategy.SortList()
}
