package observer

import "testing"

func TestObserver(t *testing.T) {
	// 创建被观察者
	s := new(Subject)
	// 创建观察者
	o := new(Observer)
	// 为主题添加观察者
	s.AddObservers(o)

	// 这里的被观察者要做各种更改...

	// 更改完毕，触发观察者
	s.NotifyObservers() // output: 已经触发了观察者
}
