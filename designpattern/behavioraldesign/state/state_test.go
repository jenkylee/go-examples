package state

import "testing"

func ExampleWeek() {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := 0; i < 8; i++ {
		todayAndNext()
	}
}

func TestMain(m *testing.M) {
	ExampleWeek()
}
