package iterator

import "testing"

func ExampleIterator() {
	var aggregate Aggregate
	aggregate = NewNumbers(1, 10)

	IteratorPrint(aggregate.Iterator())
}

func TestMain(m *testing.M) {
	ExampleIterator()
}
