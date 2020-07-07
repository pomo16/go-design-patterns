package singleton

import (
	"fmt"
	"sync"
	"testing"
)

const count = 100

func TestSingleton(t *testing.T) {
	e1 := GetInstance()
	e2 := GetInstance()
	fmt.Println(e1 == e2)
}

//并行测试
func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(count)
	instances := [count]*Singleton{}
	for i := 0; i < count; i++ {
		go func(index int) {
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < count; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
