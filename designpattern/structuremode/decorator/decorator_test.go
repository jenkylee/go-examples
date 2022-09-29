package decoractor

import (
	"fmt"
	"testing"
)

func ExampleDecorator() {
	var c Component = &ConcreteComponent{}

	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)

	res := c.Calc()

	fmt.Printf("res %d\n", res)
}

func TestMain(m *testing.M) {
	ExampleDecorator()
}
