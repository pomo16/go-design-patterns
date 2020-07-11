package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := &UnitedNationsSecurityCouncil{}
	usa := &USA{mediator}
	iraq := &Iraq{mediator}
	mediator.USA = usa
	mediator.Iraq = iraq
	usa.SendMessage("Hi Iraq, I love you~")      //美国收到对方消息: Hi USA, I love you too~~
	iraq.SendMessage("Hi USA, I love you too~~") //伊拉克收到对方消息: Hi Iraq, I love you~
}
