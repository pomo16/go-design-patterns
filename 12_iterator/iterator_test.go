package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	//创建容器并初始化
	c := &Aggregate{
		container: []int{1, 2, 3, 4},
	}

	//获取迭代器
	iterator := c.Iterator()

	for {
		// 打印当前数据
		fmt.Println(iterator.Current())
		// 如果有下一个元素，则将游标移动到下一个元素
		// 否则跳出循环，迭代结束
		if iterator.HasNext() {
			iterator.Next()
		} else {
			break
		}
	}

}
