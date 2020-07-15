package bridge

import "fmt"

/*
	bridge 桥接模式
*/

//AbstractMessage 抽象信息类
type AbstractMessage interface {
	SendMessage(text, to string)
}

//MessageImplementer 抽象发送器类
type MessageImplementer interface {
	Send(text, to string)
}

//MessageSMS SMS发送器
type MessageSMS struct{}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS\n", text, to)
}

//MessageEmail Email发送器
type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email\n", text, to)
}

//CommonMessage 通用信息，实现抽象信息类
type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

//UrgencyMessage 紧急信息，实现抽象信息类
type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
