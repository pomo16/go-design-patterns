package memento

/*
	memento 备忘录模式
*/

//Memento 备忘录结构体
type Memento struct {
	state string // 这里就是保存的状态
}

func (m *Memento) SetState(s string) {
	m.state = s
}

func (m *Memento) GetState() string {
	return m.state
}

//Originator 发起人结构体
type Originator struct {
	state string // 这里就简单一点，要保存的状态就是一个字符串
}

func (o *Originator) SetState(s string) {
	o.state = s
}

func (o *Originator) GetState() string {
	return o.state
}

//CreateMemento 发起人创建备忘录，这里就是规定了要保存的状态范围
func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

//Caretaker 负责人
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}

func (c *Caretaker) SetMemento(m *Memento) {
	c.memento = m
}
