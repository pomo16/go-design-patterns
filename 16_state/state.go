package state

import "fmt"

//Context 环境类
type Context struct {
	State       ActionState
	HealthValue int
}

//ActionState 状态接口，包含所有行为
type ActionState interface {
	View()
	Comment()
	Post()
}

//NewContext 创建环境
func NewContext(health int) *Context {
	a := &Context{
		HealthValue: health,
	}
	a.changeState()
	return a
}

//SetHealth 上下文环境变化
func (a *Context) SetHealth(value int) {
	a.HealthValue = value
	a.changeState()
}

//changeState 上下文环境变更
func (a *Context) changeState() {
	if a.HealthValue <= -10 {
		a.State = &CloseState{}
	} else if a.HealthValue > -10 && a.HealthValue <= 0 {
		a.State = &Retricted{}
	} else if a.HealthValue > 0 {
		a.State = &NormalState{}
	}
}

//State 三种状态
type NormalState struct{}
type Retricted struct{}
type CloseState struct{}

//封装三种状态的不同行为
func (c *CloseState) View() {
	fmt.Println("无法查看")
}

func (c *CloseState) Comment() {
	fmt.Println("不能评论")
}

func (c *CloseState) Post() {
	fmt.Println("不能发布")
}

func (r *Retricted) View() {
	fmt.Println("可以查看")
}

func (r *Retricted) Comment() {
	fmt.Println("不能评论")
}

func (r *Retricted) Post() {
	fmt.Println("不能发布")
}

func (r *NormalState) View() {
	fmt.Println("可以查看")
}

func (r *NormalState) Comment() {
	fmt.Println("可以评论")
}

func (r *NormalState) Post() {
	fmt.Println("可以发布")
}
