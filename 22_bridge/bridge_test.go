package bridge

import "testing"

func TestBridge(t *testing.T) {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")

	m = NewCommonMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")

	m2 := NewUrgencyMessage(ViaSMS())
	m2.SendMessage("have a drink?", "bob")

	m2 = NewUrgencyMessage(ViaEmail())
	m2.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS
	// send have a drink? to bob via Email
	// send [Urgency] have a drink? to bob via SMS
	// send [Urgency] have a drink? to bob via Email
}
