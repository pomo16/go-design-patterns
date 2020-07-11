package proxy

/*
	Proxy 代理模式
*/

//Subject 代理人和被代理的抽象接口
type Subject interface {
	Do() string
}

//RealSubject 被代理的对象
type RealSubject struct{}

//RealSubject 被代理者的方法
func (RealSubject) Do() string {
	return "real"
}

//Proxy 代理者
type Proxy struct {
	real RealSubject
}

//Proxy 通过代理调用被代理者的方法
func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}
