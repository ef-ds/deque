package deque_test

import (
	"fmt"

	"github.com/ef-ds/deque"
)

func Example() {
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
