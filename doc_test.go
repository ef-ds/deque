package deque_test

import (
	"fmt"

	deque "github.com/ef-ds/deque/v2"
)

func Example_fIFOQueue() {
	var d deque.Deque[int]
	for i := 1; i <= 5; i++ {
		d.PushBack(i)
	}
	for d.Len() > 0 {
		v, _ := d.PopFront()
		fmt.Print(v)
	}
	// Output: 12345
}

func Example_stack() {
	var d deque.Deque[int]
	for i := 1; i <= 5; i++ {
		d.PushBack(i)
	}
	for d.Len() > 0 {
		v, _ := d.PopBack()
		fmt.Print(v)
	}
	// Output: 54321
}
