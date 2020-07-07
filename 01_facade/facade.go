package facade

import "fmt"

/*
	Facade 外观模式：只暴露简单接口即可调用一系列接口的方法
*/

//API 对外暴露的简单接口
type API interface {
	DoSomeThing() string
}

//apiImpl 接口集合结构体
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

//NewAPI 新建对外简单接口
func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

//DoSomeThing 通过简单接口的方法调用一系列接口集合的方法
func (a *apiImpl) DoSomeThing() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

//AModuleAPI 具体接口及实现
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

//NewAModuleAPI 创建AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

//BModuleAPI 具体接口及实现
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

//NewBModuleAPI 创建BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}
