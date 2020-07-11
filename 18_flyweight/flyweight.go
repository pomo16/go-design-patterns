package flyweight

import "fmt"

/*
	flyweight 享元模式
*/

//FlyweightFactory 享元工厂结构体
type FlyweightFactory struct {
	maps map[string]*Flyweight
}

var flyweightFactory *FlyweightFactory

//GetFlyweightFactory 创建享元工厂
func GetFlyweightFactory() *FlyweightFactory {
	if flyweightFactory == nil {
		flyweightFactory = &FlyweightFactory{
			maps: make(map[string]*Flyweight),
		}
	}
	return flyweightFactory
}

//GetFlyweight 获取享元
func (f *FlyweightFactory) GetFlyweight(state string) *Flyweight {
	flyweight := f.maps[state]
	if flyweight == nil {
		flyweight = NewFlyweight(state)
		f.maps[state] = flyweight
	}
	return flyweight
}

//Flyweight 享元
type Flyweight struct {
	state string
}

//NewFlyweight 创建享元
func NewFlyweight(state string) *Flyweight {
	state = fmt.Sprintf("flyweight state %s", state)
	return &Flyweight{
		state: state,
	}
}

//GetState 获取享元状态
func (i *Flyweight) GetState() string {
	return i.state
}

type Viewer struct {
	*Flyweight
}

func NewViewer(state string) *Viewer {
	flyweight := GetFlyweightFactory().GetFlyweight(state)
	return &Viewer{
		Flyweight: flyweight,
	}
}

func (i *Viewer) Display() {
	fmt.Printf("Display: %s\n", i.GetState())
}
