package bridge

import "testing"

func ExampleCommonSMS() {
	m := NewCommonMessge(ViaSMS())

	m.SendMessage("have a drink?", "bob")
}

func ExampleCommonEmail() {
	m := NewCommonMessge(ViaEmail())
	m.SendMessage("have ad drink?", "bob")
}

func ExampleUrgencySMS() {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("have ad drink?", "bob")
}

func ExampleUrgencyEmail() {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have ad drink?", "bob")
}

func TestMain(m *testing.M) {
	ExampleCommonSMS()
	ExampleCommonEmail()
	ExampleUrgencySMS()
	ExampleUrgencyEmail()
}
