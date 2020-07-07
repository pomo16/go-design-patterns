package mediator

import "fmt"

/*
	Mediator 中介者模式
*/

/*
	Mediator: 抽象中介者
	ConcreteMediator: 具体中介者
	Colleague: 抽象同事类
	ConcreteColleague: 具体同事类
*/

//UnitedNations 抽象中介者
type UnitedNations interface {
	SendMessage(message string, country Country)
}

//UnitedNationsSecurityCouncil 具体中介者
type UnitedNationsSecurityCouncil struct {
	*USA
	*Iraq
}

func (u UnitedNationsSecurityCouncil) SendMessage(message string, country Country) {
	switch country.(type) {
	case USA:
		u.Iraq.GetMessage(message)
	case Iraq:
		u.USA.GetMessage(message)
	default:
		fmt.Printf("The country is not a member of UNSC")
	}
}

//Country 抽象同事类
type Country interface {
	SendMessage(message string)
	GetMessage(message string)
}

//具体同事类
type USA struct {
	UnitedNations
}

func (usa USA) SendMessage(message string) {
	usa.UnitedNations.SendMessage(message, usa)
}

func (usa USA) GetMessage(message string) {
	fmt.Printf("美国收到对方消息: %s\n", message)
}

type Iraq struct {
	UnitedNations
}

func (iraq Iraq) SendMessage(message string) {
	iraq.UnitedNations.SendMessage(message, iraq)
}

func (iraq Iraq) GetMessage(message string) {
	fmt.Printf("伊拉克收到对方消息: %s\n", message)
}
