package deque_test

import (
	"fmt"

	"github.com/ef-ds/deque"
)

func Example() {
	q := deque.New()

	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}

	for q.Len() > 0 {
		v, _ := q.PopFront()
		fmt.Print(v)
	}

	// Output: 01234
}
