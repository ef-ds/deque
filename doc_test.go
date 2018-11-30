package deque_test

import (
	"fmt"

	"github.com/ef-ds/deque"
)

func Example_fIFOQueue() {
	d := deque.New()

	for i := 0; i < 5; i++ {
		d.PushBack(i)
	}
	for d.Len() > 0 {
		v, _ := d.PopFront()
		fmt.Print(v)
	}
	// Output: 01234
}

func Example_stack() {
	d := deque.New()

	for i := 0; i < 5; i++ {
		d.PushBack(i)
	}
	for d.Len() > 0 {
		v, _ := d.PopBack()
		fmt.Print(v)
	}
	// Output: 43210
}
