package adapter

/*
	Adapter 适配器模式
*/

//Target 适配的目标接口
type Target interface {
	Request() string
}

//Adaptee 被适配的接口
type Adaptee interface {
	SpecificRequest() string
}

//NewAdaptee 被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

//AdapteeImpl 被适配的目标结构体
type adapteeImpl struct{}

//SpecificRequest 目标结构体的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee SpecificRequest method"
}

//Adapter 转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

//NewAdapter 适配器的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

//Request 实现Target接口方法
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
