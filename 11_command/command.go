package command

import "fmt"

/*
	Command 命令模式
*/

//Command 命令抽象接口
type Command interface {
	Execute()
}

//Box 请求结构体
type Box struct {
	button1 Command
	button2 Command
}

//NewBox 新建请求对象
func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

//PressButton1 封装命令执行函数
func (b *Box) PressButton1() {
	b.button1.Execute()
}

//PressButton2 封装命令执行函数
func (b *Box) PressButton2() {
	b.button2.Execute()
}

//StartCommand 具体命令
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

//RebootCommand 具体命令
type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

//MotherBoard 命令接收者、执行者
type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}
