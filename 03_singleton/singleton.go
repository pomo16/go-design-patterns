package singleton

import "sync"

/*
	Singleton 单例模式：双重锁校验，线程安全
*/

//单例类
type Singleton struct {
	name string
}

var singleton *Singleton
var once sync.Once

//GetInstance 获取单例对象
func GetInstance() *Singleton {
	//sync.Once的Do方法可以实现在程序运行过程中只运行一次其中的回调
	once.Do(func() {
		singleton = new(Singleton)
		singleton.name = "第一次赋值单例"
	})
	return singleton
}

//传统做法，但是每次都要检查两遍
var mux sync.Mutex

func GetSingletonLow() *Singleton {
	if singleton == nil { // 单例没被实例化，才会加锁
		mux.Lock()
		defer mux.Unlock()
		if singleton == nil { // 单例没被实例化才会创建
			singleton = &Singleton{}
		}
	}
	return singleton
}
