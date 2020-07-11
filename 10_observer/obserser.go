package observer

import "fmt"

/*
	Observer 观察者模式
*/

//IObserver 抽象观察者
type IObserver interface {
	Notify() // 当被观察对象有更改的时候，出发观察者的Notify() 方法
}

//ISubject 抽象被观察者
type ISubject interface {
	AddObservers(observers ...IObserver) // 添加观察者
	NotifyObservers()                    // 通知观察者
}

//Observer 实现观察者
type Observer struct{}

func (o *Observer) Notify() {
	fmt.Println("已经触发了观察者")
}

//Subject 实现被观察者
type Subject struct {
	observers []IObserver
}

func (s *Subject) AddObservers(observers ...IObserver) {
	s.observers = append(s.observers, observers...)
}

func (s *Subject) NotifyObservers() {
	for k := range s.observers {
		s.observers[k].Notify() // 触发观察者
	}
}
