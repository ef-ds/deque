// Copyright (c) 2018 ef-ds
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package deque_test

import (
	"testing"

	"github.com/ef-ds/deque"
)

const (
	pushCount = 256 * 3 // Push to fill at least 3 internal slices
)

func TestFillQueueShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
	}
	for i := 0; i < pushCount; i++ {
		if v, ok := d.PopFront(); !ok || v.(int) != i {
			t.Errorf("Expected: %d; Got: %d", i, v)
		}
	}
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
}

func TestRefillQueueShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < refillCount; i++ {
		for j := 0; j < pushCount; j++ {
			d.PushBack(j)
		}
		for j := 0; j < pushCount; j++ {
			if v, ok := d.PopFront(); !ok || v.(int) != j {
				t.Errorf("Expected: %d; Got: %d", i, v)
			}
		}
		if d.Len() != 0 {
			t.Errorf("Expected: %d; Got: %d", 0, d.Len())
		}
	}
}

func TestRefillFullQueueShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque
	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
	}

	for i := 0; i < refillCount; i++ {
		for j := 0; j < pushCount; j++ {
			d.PushBack(j)
		}
		for j := 0; j < pushCount; j++ {
			if v, ok := d.PopFront(); !ok || v.(int) != j {
				t.Errorf("Expected: %d; Got: %d", j, v)
			}
		}
		if d.Len() != pushCount {
			t.Errorf("Expected: %d; Got: %d", pushCount, d.Len())
		}
	}
}

func TestSlowIncreaseQueueShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque

	count := 0
	for i := 0; i < pushCount; i++ {
		count++
		d.PushBack(count)
		count++
		d.PushBack(count)
		if v, ok := d.PopFront(); !ok || v.(int) != i+1 {
			t.Errorf("Expected: %d; Got: %d", i+1, v)
		}
	}
	if d.Len() != pushCount {
		t.Errorf("Expected: %d; Got: %d", pushCount, d.Len())
	}
}

func TestSlowDecreaseQueueShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque
	push := 0
	for i := 0; i < pushCount; i++ {
		d.PushBack(push)
		push++
	}

	count := -1
	for i := 0; i < pushCount-1; i++ {
		count++
		if v, ok := d.PopFront(); !ok || v.(int) != count {
			t.Errorf("Expected: %d; Got: %d", count, v)
		}
		count++
		if v, ok := d.PopFront(); !ok || v.(int) != count {
			t.Errorf("Expected: %d; Got: %d", count, v)
		}

		d.PushBack(push)
		push++
	}
	count++
	if v, ok := d.PopFront(); !ok || v.(int) != count {
		t.Errorf("Expected: %d; Got: %d", count, v)
	}
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
}

func TestStableQueueShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
		if v, ok := d.PopFront(); !ok || v.(int) != i {
			t.Errorf("Expected: %d; Got: %d", i, v)
		}
	}
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
}

func TestStableQueueFullShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque
	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
	}

	count := 0
	for i := 0; i < pushCount-1; i++ {
		d.PushBack(i)
		if v, ok := d.PopFront(); !ok || v.(int) != count {
			t.Errorf("Expected: %d; Got: %d", count, v)
		}
		count++
	}
	if d.Len() != pushCount {
		t.Errorf("Expected: %d; Got: %d", pushCount, d.Len())
	}
}

func TestFillStackShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
	}
	for i := pushCount - 1; i >= 0; i-- {
		if v, ok := d.PopBack(); !ok || v.(int) != i {
			t.Errorf("Expected: %d; Got: %d", i, v)
		}
	}
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
}

func TestRefillStackShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < refillCount; i++ {
		for j := 0; j < pushCount; j++ {
			d.PushBack(j)
		}
		for j := pushCount - 1; j >= 0; j-- {
			if v, ok := d.PopBack(); !ok || v.(int) != j {
				t.Errorf("Expected: %d; Got: %d", i, v)
			}
		}
		if d.Len() != 0 {
			t.Errorf("Expected: %d; Got: %d", 0, d.Len())
		}
	}
}

func TestRefillFullStackShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque
	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
	}

	for i := 0; i < refillCount; i++ {
		for j := 0; j < pushCount; j++ {
			d.PushBack(j)
		}
		for j := pushCount - 1; j >= 0; j-- {
			if v, ok := d.PopBack(); !ok || v.(int) != j {
				t.Errorf("Expected: %d; Got: %d", j, v)
			}
		}
		if d.Len() != pushCount {
			t.Errorf("Expected: %d; Got: %d", pushCount, d.Len())
		}
	}
}

func TestSlowIncreaseStackShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque

	count := 0
	for i := 0; i < pushCount; i++ {
		count++
		d.PushBack(count)
		count++
		d.PushBack(count)
		if v, ok := d.PopBack(); !ok || v.(int) != count {
			t.Errorf("Expected: %d; Got: %d", count, v)
		}
	}
	if d.Len() != pushCount {
		t.Errorf("Expected: %d; Got: %d", pushCount, d.Len())
	}
}

func TestSlowDecreaseStackShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque
	push := 0
	for i := 0; i < pushCount; i++ {
		d.PushBack(push)
		push++
	}

	count := push
	for i := 0; i < pushCount-1; i++ {
		count--
		if v, ok := d.PopBack(); !ok || v.(int) != count {
			t.Errorf("Expected: %d; Got: %d", count, v)
		}
		count--
		if v, ok := d.PopBack(); !ok || v.(int) != count {
			t.Errorf("Expected: %d; Got: %d", count, v)
		}

		d.PushBack(count)
		count++
	}
	count--
	if v, ok := d.PopBack(); !ok || v.(int) != count {
		t.Errorf("Expected: %d; Got: %d", count, v)
	}
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
}

func TestStableStackShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
		if v, ok := d.PopBack(); !ok || v.(int) != i {
			t.Errorf("Expected: %d; Got: %d", i, v)
		}
	}
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
}

func TestStableFullStackShouldRetrieveAllElementsInOrder(t *testing.T) {
	var d deque.Deque
	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
	}

	count := 0
	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
		if v, ok := d.PopFront(); !ok || v.(int) != count {
			t.Errorf("Expected: %d; Got: %d", count, v)
		}
		count++
	}
	if d.Len() != pushCount {
		t.Errorf("Expected: %d; Got: %d", pushCount, d.Len())
	}
}

func TestPushPopFrontShouldRetrieveAllElementsInOrder(t *testing.T) {
	pushPopFrontBackShouldRetrieveAllElementsInOrder(
		t,
		func(d *deque.Deque) (interface{}, bool) {
			return d.PopFront()
		},
		func(d *deque.Deque) (interface{}, bool) {
			return d.Front()
		},
		func(v, lastGet, lastPut interface{}) bool {
			return v == lastGet
		},
	)
}

func TestPushPopBackShouldRetrieveAllElementsInOrder(t *testing.T) {
	pushPopFrontBackShouldRetrieveAllElementsInOrder(
		t,
		func(d *deque.Deque) (interface{}, bool) {
			return d.PopBack()
		},
		func(d *deque.Deque) (interface{}, bool) {
			return d.Back()
		},
		func(v, lastGet, lastPut interface{}) bool {
			return v == lastPut
		},
	)
}

func TestMixedPushBackPopFrontBackShouldReturnAllValuesInOrder(t *testing.T) {
	var d deque.Deque
	for i := 0; i < pushCount; i++ {
		d.PushBack(i)
	}
	count, popFrontCount, popBackCount := 0, 0, pushCount-1
	for d.Len() > 0 {
		if count%2 == 0 {
			if v, ok := d.PopFront(); !ok || v != popFrontCount {
				t.Errorf("Expected: %d; Got: %d", popFrontCount, v)
			}
			popFrontCount++
		} else {
			if v, ok := d.PopBack(); !ok || v != popBackCount {
				t.Errorf("Expected: %d; Got: %d", popBackCount, v)
			}
			popBackCount--
		}
		if d.Len() != pushCount-count-1 {
			t.Errorf("Expected: %d; Got: %d", pushCount-count-1, d.Len())
		}
		count++
	}
	if v, ok := d.PopBack(); ok || v != nil {
		t.Errorf("Expected: nil; Got: %d", v)
	}
	if d.Len() != 0 {
		t.Errorf("Expected: 0; Got: %d", d.Len())
	}
}

func TestMixedPushFrontPopFrontBackShouldReturnAllValuesInOrder(t *testing.T) {
	var d deque.Deque
	for i := 0; i < pushCount; i++ {
		d.PushFront(i)
	}
	count, popFrontCount, popBackCount := 0, pushCount-1, 0
	for d.Len() > 0 {
		if count%2 == 0 {
			if v, ok := d.PopFront(); !ok || v != popFrontCount {
				t.Errorf("Expected: %d; Got: %d", popFrontCount, v)
			}
			popFrontCount--
		} else {
			if v, ok := d.PopBack(); !ok || v != popBackCount {
				t.Errorf("Expected: %d; Got: %d", popBackCount, v)
			}
			popBackCount++
		}
		if d.Len() != pushCount-count-1 {
			t.Errorf("Expected: %d; Got: %d", pushCount-count-1, d.Len())
		}
		count++
	}
	if v, ok := d.PopBack(); ok || v != nil {
		t.Errorf("Expected: nil; Got: %d", v)
	}
	if d.Len() != 0 {
		t.Errorf("Expected: 0; Got: %d", d.Len())
	}
}

func TestMixedPushFrontBackPopFrontShouldReturnAllValuesInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < pushCount; i++ {
		if i%2 == 0 {
			d.PushFront(i)
		} else {
			d.PushBack(i)
		}
	}
	decValue := true
	expectedValue := pushCount - 2
	for d.Len() > 0 {
		if v, ok := d.PopFront(); !ok || v != expectedValue {
			t.Errorf("Expected: %d; Got: %d", expectedValue, v)
		}
		if decValue {
			expectedValue -= 2
		} else {
			expectedValue += 2
		}
		if expectedValue == 0 {
			decValue = false
		}
		if !decValue && expectedValue == 2 {
			expectedValue = 1
		}
	}
	if v, ok := d.PopFront(); ok || v != nil {
		t.Errorf("Expected: nil; Got: %d", v)
	}
	if d.Len() != 0 {
		t.Errorf("Expected: 0; Got: %d", d.Len())
	}
}

func TestMixedPushFrontBackPopBackShouldReturnAllValuesInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < pushCount; i++ {
		if i%2 == 0 {
			d.PushFront(i)
		} else {
			d.PushBack(i)
		}
	}
	decValue := true
	expectedValue := pushCount - 1
	for d.Len() > 0 {
		if v, ok := d.PopBack(); !ok || v != expectedValue {
			t.Errorf("Expected: %d; Got: %d", expectedValue, v)
		}
		if decValue {
			expectedValue -= 2
		} else {
			expectedValue += 2
		}
		if expectedValue == -1 {
			decValue = false
			expectedValue = 0
		}
	}
	if v, ok := d.PopBack(); ok || v != nil {
		t.Errorf("Expected: nil; Got: %d", v)
	}
	if d.Len() != 0 {
		t.Errorf("Expected: 0; Got: %d", d.Len())
	}
}

func TestPushFrontPopFrontRefillWith0ToPushCountItemsShouldReturnAllValuesInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < refillCount; i++ {
		for k := 0; k < pushCount; k++ {
			for j := 0; j < k; j++ {
				d.PushFront(j)
			}
			for j := k; j > 0; j-- {
				v, ok := d.PopFront()
				if !ok || v == nil || v.(int) != j-1 {
					t.Errorf("Expected: %d; Got: %d", j-1, v)
				}
			}
			if d.Len() != 0 {
				t.Errorf("Expected: %d; Got: %d", 0, d.Len())
			}
		}
	}
}

func TestPushFrontPopBackRefillWith0ToPushCountItemsShouldReturnAllValuesInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < refillCount; i++ {
		for k := 0; k < pushCount; k++ {
			for j := 0; j < k; j++ {
				d.PushFront(j)
			}
			for j := 0; j < k; j++ {
				v, ok := d.PopBack()
				if !ok || v == nil || v.(int) != j {
					t.Errorf("Expected: %d; Got: %d", j, v)
				}
			}
			if d.Len() != 0 {
				t.Errorf("Expected: %d; Got: %d", 0, d.Len())
			}
		}
	}
}

func TestPushBackPopFrontRefillWith0ToPushCountItemsShouldReturnAllValuesInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < refillCount; i++ {
		for k := 1; k < pushCount; k++ {
			for j := 0; j < k; j++ {
				d.PushBack(j)
			}
			for j := 0; j < k; j++ {
				v, ok := d.PopFront()
				if !ok || v == nil || v.(int) != j {
					t.Errorf("Expected: %d; Got: %d", j, v)
				}
			}
			if d.Len() != 0 {
				t.Errorf("Expected: %d; Got: %d", 0, d.Len())
			}
		}
	}
}

func TestPushBackPopBackRefillWith0ToPushCountItemsShouldReturnAllValuesInOrder(t *testing.T) {
	var d deque.Deque

	for i := 0; i < refillCount; i++ {
		for k := 1; k < pushCount; k++ {
			for j := 0; j < k; j++ {
				d.PushBack(j)
			}
			for j := k; j > 0; j-- {
				v, ok := d.PopBack()
				if !ok || v == nil || v.(int) != j-1 {
					t.Errorf("Expected: %d; Got: %d", j-1, v)
				}
			}
			if d.Len() != 0 {
				t.Errorf("Expected: %d; Got: %d", 0, d.Len())
			}
		}
	}
}

// Helper methods ------------------------------------------------

func pushPopFrontBackShouldRetrieveAllElementsInOrder(t *testing.T, popFunc, frontbackFunc func(*deque.Deque) (interface{}, bool), checkValueFunc func(v, lastGet, lastPut interface{}) bool) {
	tests := map[string]struct {
		putCount []int
		getCount []int
	}{
		"Test 1 item": {
			putCount: []int{1},
			getCount: []int{1},
		},
		"Test 10 item": {
			putCount: []int{10},
			getCount: []int{10},
		},
		"Test 100 items": {
			putCount: []int{100},
			getCount: []int{100},
		},
		"Test 1000 items": {
			putCount: []int{1000},
			getCount: []int{1000},
		},
		"Test 10000 items": {
			putCount: []int{10000},
			getCount: []int{10000},
		},
		"Test 100000 items": {
			putCount: []int{100000},
			getCount: []int{100000},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := deque.New()
			lastPut := 0
			lastGet := 0
			var ok bool
			var v interface{}
			for count := 0; count < len(test.getCount); count++ {
				for i := 1; i <= test.putCount[count]; i++ {
					lastPut++
					d.PushBack(lastPut)
					if v, ok = frontbackFunc(d); !ok || !checkValueFunc(v, lastGet+1, lastPut) {
						t.Errorf("Expected: %d; Got: %d", lastGet, v)
					}
				}

				for i := 1; i <= test.getCount[count]; i++ {
					lastGet++
					v, ok = frontbackFunc(d)
					if !ok || !checkValueFunc(v.(int), lastGet, lastPut-lastGet+1) {
						t.Errorf("Expected: %d; Got: %d or %d", lastGet, lastPut-lastGet+1, v)
					}
					v, ok = popFunc(d)
					if !ok || !checkValueFunc(v.(int), lastGet, lastPut-lastGet+1) {
						t.Errorf("Expected: %d; Got: %d or %d", lastGet, lastPut-lastGet+1, v)
					}
				}
			}

			if d.Len() != 0 {
				t.Errorf("Expected: %d; Got: %d", 0, d.Len())
			}
			if v, ok = frontbackFunc(d); ok || v != nil {
				t.Errorf("Expected: nil as the queue should be empty; Got: %d", v)
			}
			if v, ok = popFunc(d); ok || v != nil {
				t.Errorf("Expected: nil as the queue should be empty; Got: %d", v)
			}
		})
	}
}
