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

package deque

import (
	"fmt"
	"testing"
)

const (
	refillCount = 3
	pushCount   = maxInternalSliceSize * 3 // Push to fill at least 3 internal slices
)

func TestNewShouldReturnInitiazedInstanceOfDeque(t *testing.T) {
	d := New()
	assertInvariants(t, d, nil)
}

func TestInvariantsWhenEmptyInMiddleOfSlice(t *testing.T) {
	d := new(Deque)
	d.PushBack(0)
	assertInvariants(t, d, nil)
	d.PushBack(1)
	assertInvariants(t, d, nil)
	d.PopFront()
	assertInvariants(t, d, nil)
	d.PopFront()
	// At this point, the queue is empty and hp will
	// not be pointing at the start of the slice.
	assertInvariants(t, d, nil)
}

func TestPushFrontPopBackShouldHaveAllInternalLinksInARing(t *testing.T) {
	d := New()
	pushValue, extraAddedItems, spareLinks := 0, 0, 0

	// Push maxFirstSliceSize items to fill the first array
	expectedHeadSliceSize := firstSliceSize
	for i := 1; i <= maxFirstSliceSize; i++ {
		pushValue++
		d.PushFront(pushValue)

		checkLinks(t, d, pushValue, expectedHeadSliceSize, expectedHeadSliceSize, spareLinks, d.head, d.head, d.head, d.head)
		if pushValue >= expectedHeadSliceSize {
			expectedHeadSliceSize *= sliceGrowthFactor
		}
	}

	// Push 1 extra item to force the creation of a new array
	pushValue++
	d.PushFront(pushValue)
	extraAddedItems++
	checkLinks(t, d, pushValue, maxInternalSliceSize, maxFirstSliceSize, spareLinks, d.tail, d.tail, d.head, d.head)

	// Push another maxInternalSliceSize-1 to fill the second array
	for i := 1; i <= maxInternalSliceSize-1; i++ {
		pushValue++
		d.PushFront(pushValue)
		checkLinks(t, d, pushValue, maxInternalSliceSize, maxFirstSliceSize, spareLinks, d.tail, d.tail, d.head, d.head)
	}

	// Push 1 extra item to force the creation of a new array (3 total)
	pushValue++
	d.PushFront(pushValue)
	checkLinks(t, d, pushValue, maxInternalSliceSize, maxFirstSliceSize, spareLinks, d.tail.p, d.tail, d.head, d.head.n)
	/// Check middle links
	if d.head.n.n != d.tail {
		t.Error("Expected: d.head.n.n == d.tail; Got: d.head.n.n != d.tail")
	}
	if d.head.n.p != d.head {
		t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
	}

	// Check final len after all pushes
	if d.Len() != maxFirstSliceSize+maxInternalSliceSize+extraAddedItems {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize+maxInternalSliceSize+extraAddedItems, d.Len())
	}

	// Pop maxFirstSliceSize-1 items to empty the tail slice
	popValue := 1
	for i := 1; i <= maxFirstSliceSize-1; i++ {
		if v, ok := d.PopBack(); !ok || v.(int) != popValue {
			t.Errorf("Expected: %d; Got: %d", popValue, v)
		}
		popValue++
		checkLinks(t, d, pushValue-popValue+extraAddedItems, maxInternalSliceSize, maxFirstSliceSize, spareLinks, d.tail.p, d.tail, d.head, d.head.n)
		/// Check middle links
		if d.head.n.n != d.tail {
			t.Error("Expected: d.head.n.n == d.tail; Got: d.head.n.n != d.tail")
		}
		if d.head.n.p != d.head {
			t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
		}
	}

	// Pop one extra item to force moving the tail to the middle slice. This also means the old tail
	// slice should have no items now, so spareLinks should be increased.
	if v, ok := d.PopBack(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	spareLinks++
	popValue++
	checkLinks(t, d, pushValue-popValue+extraAddedItems, maxInternalSliceSize, maxInternalSliceSize, spareLinks, d.tail, d.tail.n, d.head.p, d.head)
	//Check last slice links (not tail anymore; tail is the middle one)
	if d.head.p.n != d.head {
		t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
	}
	if d.head.p.p != d.tail {
		t.Error("Expected: d.head.p.p == d.tail; Got: d.head.p.p != d.tail")
	}

	// Pop maxInternalSliceSize-1 items to empty the tail (middle) slice
	for i := 1; i <= maxInternalSliceSize-1; i++ {
		if v, ok := d.PopBack(); !ok || v.(int) != popValue {
			t.Errorf("Expected: %d; Got: %d", popValue, v)
		}
		popValue++
		checkLinks(t, d, pushValue-popValue+extraAddedItems, maxInternalSliceSize, maxInternalSliceSize, spareLinks, d.tail, d.tail.n, d.head.p, d.head)
		//Check last slice links (not tail anymore; tail is the middle one)
		if d.head.p.n != d.head {
			t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
		}
		if d.head.p.p != d.tail {
			t.Error("Expected: d.head.p.p == d.tail; Got: d.head.p.p != d.tail")
		}
	}

	// Pop one extra item to force moving the tail to the head slice. This also means the old tail
	// slice should have no items now, so spareLinks should be increased.
	if v, ok := d.PopBack(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	spareLinks++
	checkLinks(t, d, pushValue-popValue, maxInternalSliceSize, maxInternalSliceSize, spareLinks, d.tail.p.p, d.tail.p, d.tail.p.p, d.tail.p)
	/// Check middle links
	if d.head.n.n != d.tail.p {
		t.Error("Expected: d.head.n.n == d.tail.p; Got: d.head.n.n != d.tail.p")
	}
	if d.head.n.p != d.head {
		t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
	}
	//Check last slice links (not tail anymore; tail is the first one)
	if d.head.p.n != d.head {
		t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
	}
	if d.head.p.p != d.tail.n {
		t.Error("Expected: d.head.p.p == d.tail.n; Got: d.head.p.p != d.tail.n")
	}

	// Pop one extra item emptying the deque
	popValue++
	if v, ok := d.PopBack(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	checkLinks(t, d, pushValue-popValue, maxInternalSliceSize, maxInternalSliceSize, spareLinks, d.tail.p.p, d.tail.p, d.tail.p.p, d.tail.p)
	/// Check middle links
	if d.head.n.n != d.tail.p {
		t.Error("Expected: d.head.n.n == d.tail.p; Got: d.head.n.n != d.tail.p")
	}
	if d.head.n.p != d.head {
		t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
	}
	//Check last slice links (not tail anymore; tail is the first one)
	if d.head.p.n != d.head {
		t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
	}
	if d.head.p.p != d.tail.n {
		t.Error("Expected: d.head.p.p == d.tail.n; Got: d.head.p.p != d.tail.n")
	}

	// The deque shoud be empty
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
	if _, ok := d.Back(); ok {
		t.Error("Expected: false; Got: true")
	}
	if _, ok := d.Front(); ok {
		t.Error("Expected: false; Got: true")
	}
	if len(d.head.v) != maxInternalSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxInternalSliceSize, len(d.head.v))
	}
	if len(d.tail.v) != maxInternalSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxInternalSliceSize, len(d.tail.v))
	}
	if d.head != d.tail {
		t.Error("Expected: d.head == d.tail; Got: d.head != d.tail")
	}
	if d.tp != maxInternalSliceSize-1 {
		t.Errorf("Expected: %d; Got: %d", 0, d.tp)
	}
	if d.hp != maxInternalSliceSize-1 {
		t.Errorf("Expected: %d; Got: %d", 0, d.hp)
	}
	if d.spareLinks != spareLinks {
		t.Errorf("Expected: %d; Got: %d", spareLinks, d.spareLinks)
	}
}

func TestPushFrontPopFrontShouldHaveAllInternalLinksInARing(t *testing.T) {
	d := New()
	pushValue, spareLinks := 0, 0

	// Push maxFirstSliceSize + maxInternalSliceSize + 1 items to fill the first, second
	// and have one item in the third slice
	for i := 1; i <= maxFirstSliceSize+maxInternalSliceSize+1; i++ {
		pushValue++
		d.PushFront(pushValue)
	}

	// Pop one item to force moving the head to the middle slice. This also means the old head
	// slice should have no items now, so spareLinks should be increased.
	popValue := pushValue
	if v, ok := d.PopFront(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	spareLinks++
	popValue--
	checkLinks(t, d, popValue, maxInternalSliceSize, maxFirstSliceSize, spareLinks, d.tail, d.tail.n, d.head.p, d.head)
	//Check first slice links (not head anymore; head is the middle one)
	if d.head.p.n != d.head {
		t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
	}
	if d.head.p.p != d.tail {
		t.Error("Expected: d.head.p.p == d.tail; Got: d.head.p.p != d.tail")
	}

	// Pop maxFirstSliceSize-1 items to empty the head, middle slice
	for i := 1; i <= maxInternalSliceSize-1; i++ {
		if v, ok := d.PopFront(); !ok || v.(int) != popValue {
			t.Errorf("Expected: %d; Got: %d", popValue, v)
		}
		popValue--
		checkLinks(t, d, popValue, maxInternalSliceSize, maxFirstSliceSize, spareLinks, d.tail, d.tail.n, d.head.p, d.head)
		/// Check first slice links
		if d.head.p.n != d.head {
			t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
		}
		if d.head.p.p != d.tail {
			t.Error("Expected: d.head.p.p == d.tail; Got: d.head.p.p != d.tail")
		}
	}

	// Pop one extra item to force moving the head to the last slice. This also means the old head
	// slice should have no items now, so spareLinks should be increased.
	if v, ok := d.PopFront(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	spareLinks++
	popValue--
	checkLinks(t, d, popValue, maxFirstSliceSize, maxFirstSliceSize, spareLinks, d.tail.p.p, d.tail.p, d.head.p.p, d.head.p)
	//Check first slice links
	if d.head.p.p.n != d.head.p {
		t.Error("Expected: d.head.p.p.n == d.head.p; Got: d.head.p.p.n != d.head.p")
	}
	if d.head.p.p.p != d.tail {
		t.Error("Expected: d.head.p.p.p == d.tail; Got: d.head.p.p.p != d.tail")
	}
	//Check middle  slice links
	if d.head.p.n != d.head {
		t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
	}
	if d.head.p.p != d.tail.n {
		t.Error("Expected: d.head.p.p == d.tail.n; Got: d.head.p.p != d.tail.n")
	}

	// Pop maxFirstSliceSize-1 items to empty the head/tail (last) slice
	for i := 1; i <= maxFirstSliceSize; i++ {
		if v, ok := d.PopFront(); !ok || v.(int) != popValue {
			t.Errorf("Expected: %d; Got: %d", popValue, v)
		}
		popValue--
		checkLinks(t, d, popValue, maxFirstSliceSize, maxFirstSliceSize, spareLinks, d.tail.p.p, d.tail.p, d.head.p.p, d.head.p)
		//Check first slice links
		if d.head.p.p.n != d.head.p {
			t.Error("Expected: d.head.p.p.n == d.head.p; Got: d.head.p.p.n != d.head.p")
		}
		if d.head.p.p.p != d.tail {
			t.Error("Expected: d.head.p.p.p == d.tail; Got: d.head.p.p.p != d.tail")
		}
		//Check middle  slice links
		if d.head.p.n != d.head {
			t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
		}
		if d.head.p.p != d.tail.n {
			t.Error("Expected: d.head.p.p == d.tail.n; Got: d.head.p.p != d.tail.n")
		}
	}

	// The deque shoud be empty
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
	if _, ok := d.Back(); ok {
		t.Error("Expected: false; Got: true")
	}
	if _, ok := d.Front(); ok {
		t.Error("Expected: false; Got: true")
	}
	if len(d.head.v) != maxFirstSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize, len(d.head.v))
	}
	if len(d.tail.v) != maxFirstSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize, len(d.tail.v))
	}
	if d.head != d.tail {
		t.Error("Expected: d.head == d.tail; Got: d.head != d.tail")
	}
	if d.tp != maxFirstSliceSize-1 {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize-1, d.tp)
	}
	if d.hp != maxFirstSliceSize-1 {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize-1, d.hp)
	}
	if d.spareLinks != spareLinks {
		t.Errorf("Expected: %d; Got: %d", spareLinks, d.spareLinks)
	}
}

func TestPushBackPopBackShouldHaveAllInternalLinksInARing(t *testing.T) {
	d := New()
	pushValue, extraAddedItems, spareLinks := 0, 0, 0

	// Push maxFirstSliceSize items to fill the first array
	expectedHeadSliceSize := firstSliceSize
	for i := 1; i <= maxFirstSliceSize; i++ {
		pushValue++
		d.PushBack(pushValue)

		checkLinks(t, d, pushValue, expectedHeadSliceSize, expectedHeadSliceSize, spareLinks, d.head, d.head, d.head, d.head)
		if pushValue >= expectedHeadSliceSize {
			expectedHeadSliceSize *= sliceGrowthFactor
		}
	}

	// Push 1 extra item to force the creation of a new array
	pushValue++
	d.PushBack(pushValue)
	extraAddedItems++
	checkLinks(t, d, pushValue, maxFirstSliceSize, maxInternalSliceSize, spareLinks, d.tail, d.tail, d.head, d.head)

	// Push another maxInternalSliceSize-1 to fill the second array
	for i := 1; i <= maxInternalSliceSize-1; i++ {
		pushValue++
		d.PushBack(pushValue)
		checkLinks(t, d, pushValue, maxFirstSliceSize, maxInternalSliceSize, spareLinks, d.tail, d.tail, d.head, d.head)
	}

	// Push 1 extra item to force the creation of a new array (3 total)
	pushValue++
	d.PushBack(pushValue)
	checkLinks(t, d, pushValue, maxFirstSliceSize, maxInternalSliceSize, spareLinks, d.tail.p, d.tail, d.head, d.head.n)
	/// Check middle links
	if d.head.n.n != d.tail {
		t.Error("Expected: d.head.n.n == d.tail; Got: d.head.n.n != d.tail")
	}
	if d.head.n.p != d.head {
		t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
	}

	// Check final len after all pushes
	if d.Len() != maxFirstSliceSize+maxInternalSliceSize+extraAddedItems {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize+maxInternalSliceSize+extraAddedItems, d.Len())
	}

	// Pop one item to force moving the tail to the middle slice. This also means the old tail
	// slice should have no items now, so spareLinks should be increased.
	popValue := d.Len()
	if v, ok := d.PopBack(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	spareLinks++
	popValue--
	checkLinks(t, d, popValue, maxFirstSliceSize, maxInternalSliceSize, spareLinks, d.tail, d.tail.n, d.head.p, d.head)
	//Check last slice links (not tail anymore; tail is the middle one)
	if d.head.p.n != d.head {
		t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
	}
	if d.head.p.p != d.tail {
		t.Error("Expected: d.head.p.p == d.tail; Got: d.head.p.p != d.tail")
	}

	// Pop maxInternalSliceSize-1 items to empty the tail (middle) slice
	for i := 1; i <= maxInternalSliceSize-1; i++ {
		if v, ok := d.PopBack(); !ok || v.(int) != popValue {
			t.Errorf("Expected: %d; Got: %d", popValue, v)
		}
		popValue--
		checkLinks(t, d, popValue, maxFirstSliceSize, maxInternalSliceSize, spareLinks, d.tail, d.tail.n, d.head.p, d.head)
		/// Check last slice links
		if d.head.p.n != d.head {
			t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
		}
		if d.head.p.p != d.tail {
			t.Error("Expected: d.head.p.p == d.tail; Got: d.head.p.p != d.tail")
		}
	}

	// Pop one extra item to force moving the tail to the head (first) slice. This also means the old tail
	// slice should have no items now, so spareLinks should be increased.
	if v, ok := d.PopBack(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	popValue--
	spareLinks++
	checkLinks(t, d, popValue, maxFirstSliceSize, maxFirstSliceSize, spareLinks, d.tail.p.p, d.tail.p, d.tail.p.p, d.tail.p)
	/// Check middle links
	if d.head.n.n != d.tail.p {
		t.Error("Expected: d.head.n.n == d.tail.p; Got: d.head.n.n != d.tail.p")
	}
	if d.head.n.p != d.head {
		t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
	}
	//Check last slice links (not tail anymore; tail is the first one)
	if d.head.p.n != d.head {
		t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
	}
	if d.head.p.p != d.tail.n {
		t.Error("Expected: d.head.p.p == d.tail.n; Got: d.head.p.p != d.tail.n")
	}

	// Pop maxFirstSliceSize-1 items to empty the head (first) slice
	for i := 1; i <= maxFirstSliceSize; i++ {
		if v, ok := d.PopBack(); !ok || v.(int) != popValue {
			t.Errorf("Expected: %d; Got: %d", popValue, v)
		}
		popValue--
		checkLinks(t, d, popValue, maxFirstSliceSize, maxFirstSliceSize, spareLinks, d.tail.p.p, d.tail.p, d.tail.p.p, d.tail.p)
		/// Check middle links
		if d.head.n.n != d.tail.p {
			t.Error("Expected: d.head.n.n == d.tail.p; Got: d.head.n.n != d.tail.p")
		}
		if d.head.n.p != d.head {
			t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
		}
		//Check last slice links (not tail anymore; tail is the first one)
		if d.head.p.n != d.head {
			t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
		}
		if d.head.p.p != d.tail.n {
			t.Error("Expected: d.head.p.p == d.tail.n; Got: d.head.p.p != d.tail.n")
		}
	}

	// The deque shoud be empty
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
	if _, ok := d.Back(); ok {
		t.Error("Expected: false; Got: true")
	}
	if _, ok := d.Front(); ok {
		t.Error("Expected: false; Got: true")
	}
	if len(d.head.v) != maxFirstSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize, len(d.head.v))
	}
	if len(d.tail.v) != maxFirstSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize, len(d.tail.v))
	}
	if d.head.n == d.tail.p {
		t.Error("Expected: d.head.n != d.tail.p; Got: d.head.n == d.tail.p")
	}
	if d.tp != 0 {
		t.Errorf("Expected: %d; Got: %d", -1, d.tp)
	}
	if d.hp != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.hp)
	}
	if d.spareLinks != spareLinks {
		t.Errorf("Expected: %d; Got: %d", spareLinks, d.spareLinks)
	}
}

func TestPushBackPopFrontShouldHaveAllInternalLinksInARing(t *testing.T) {
	d := New()
	pushValue, spareLinks := 0, 0

	// Push maxFirstSliceSize + maxInternalSliceSize + 1 items to fill the first, second
	// and have one item in the third slice
	for i := 1; i <= maxFirstSliceSize+maxInternalSliceSize+1; i++ {
		pushValue++
		d.PushBack(pushValue)
	}

	// Pop maxFirstSliceSize-1 items to empty the head slice
	popValue := 0
	for i := 1; i <= maxFirstSliceSize-1; i++ {
		popValue++
		if v, ok := d.PopFront(); !ok || v.(int) != popValue {
			t.Errorf("Expected: %d; Got: %d", popValue, v)
		}
		checkLinks(t, d, pushValue-popValue, maxFirstSliceSize, maxInternalSliceSize, spareLinks, d.tail.p, d.tail, d.head, d.head.n)
		/// Check middle slice links
		if d.head.n.n != d.tail {
			t.Error("Expected: d.head.n.n == d.tail; Got: d.head.n.n != d.tail")
		}
		if d.head.n.p != d.head {
			t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
		}
	}

	// Pop one item to force moving the head to the middle slice. This also means the old head
	// slice should have no items now, so spareLinks should be increased.
	popValue++
	if v, ok := d.PopFront(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	spareLinks++
	checkLinks(t, d, pushValue-popValue, maxInternalSliceSize, maxInternalSliceSize, spareLinks, d.tail, d.tail.n, d.head.p, d.head)
	//Check first slice links (not head anymore; head is the middle one)
	if d.head.p.n != d.head {
		t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
	}
	if d.head.p.p != d.tail {
		t.Error("Expected: d.head.p.p == d.tail; Got: d.head.p.p != d.tail")
	}

	// Pop maxInternalSliceSize-1 items to empty the head (middle) slice
	for i := 1; i <= maxInternalSliceSize-1; i++ {
		popValue++
		if v, ok := d.PopFront(); !ok || v.(int) != popValue {
			t.Errorf("Expected: %d; Got: %d", popValue, v)
		}
		checkLinks(t, d, pushValue-popValue, maxInternalSliceSize, maxInternalSliceSize, spareLinks, d.tail, d.tail.n, d.head.p, d.head)
		//Check first slice links
		if d.head.p.n != d.head {
			t.Error("Expected: d.head.p.n == d.head; Got: d.head.p.n != d.head")
		}
		if d.head.p.p != d.tail {
			t.Error("Expected: d.head.p.p == d.tail; Got: d.head.p.p != d.tail")
		}
	}

	// Pop one extra item making head move to next (last slice). This also means the old head
	// slice should have no items now, so spareLinks should be increased.
	popValue++
	spareLinks++
	if v, ok := d.PopFront(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	checkLinks(t, d, pushValue-popValue, maxInternalSliceSize, maxInternalSliceSize, spareLinks, d.tail.p.p, d.tail.p, d.head.p.p, d.head.p)
	//Check first slice links
	if d.head.n.n != d.head.p {
		t.Error("Expected: d.head.n.n == d.head.p; Got: d.head.n.n != d.head.p")
	}
	if d.head.n.p != d.tail {
		t.Error("Expected: d.head.n.p == d.tail; Got: d.head.n.p != d.tail")
	}
	//Check middle slice links
	if d.head.p.n != d.tail {
		t.Error("Expected: d.head.p.n == d.tail; Got: d.head.p.n != d.tail")
	}
	if d.head.p.p != d.tail.n {
		t.Error("Expected: d.head.p.p == d.tail.n; Got: d.head.p.p != d.tail.n")
	}

	// Pop one extra item emptying the deque.
	popValue++
	if v, ok := d.PopFront(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	checkLinks(t, d, pushValue-popValue, maxInternalSliceSize, maxInternalSliceSize, spareLinks, d.tail.p.p, d.tail.p, d.head.p.p, d.head.p)
	//Check first slice links
	if d.head.n.n != d.head.p {
		t.Error("Expected: d.head.n.n == d.head.p; Got: d.head.n.n != d.head.p")
	}
	if d.head.n.p != d.tail {
		t.Error("Expected: d.head.n.p == d.tail; Got: d.head.n.p != d.tail")
	}
	//Check middle slice links
	if d.head.p.n != d.tail {
		t.Error("Expected: d.head.p.n == d.tail; Got: d.head.p.n != d.tail")
	}
	if d.head.p.p != d.tail.n {
		t.Error("Expected: d.head.p.p == d.tail.n; Got: d.head.p.p != d.tail.n")
	}

	// The deque shoud be empty
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
	if _, ok := d.Back(); ok {
		t.Error("Expected: false; Got: true")
	}
	if _, ok := d.Front(); ok {
		t.Error("Expected: false; Got: true")
	}
	if len(d.head.v) != maxInternalSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxInternalSliceSize, len(d.head.v))
	}
	if len(d.tail.v) != maxInternalSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxInternalSliceSize, len(d.tail.v))
	}
	if d.head.p != d.tail.p {
		t.Error("Expected: d.head.p == d.tail.p; Got: d.head.p != d.tail.p")
	}
	if d.tp != 1 {
		t.Errorf("Expected: %d; Got: %d", 0, d.tp)
	}
	if d.hp != 1 {
		t.Errorf("Expected: %d; Got: %d", 1, d.hp)
	}
	if d.spareLinks != spareLinks {
		t.Errorf("Expected: %d; Got: %d", spareLinks, d.spareLinks)
	}
}

func TestPushFrontShouldReuseSpareLinks(t *testing.T) {
	d := New()
	count := maxInternalSliceSize * 3
	// Fills the deque
	for i := 0; i < count; i++ {
		d.PushFront(i)
	}
	// Pop the items to generate spare links
	for i := 0; i < maxInternalSliceSize; i++ {
		d.PopFront()
	}
	if d.spareLinks != 1 {
		t.Errorf("Expected: %d; Got: %d", 1, d.spareLinks)
	}

	// Push the items back using PushFront
	for i := 0; i < count; i++ {
		d.PushFront(i)
	}

	// The spare links should've been used up
	if d.spareLinks != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.spareLinks)
	}
}

func TestPushBackShouldReuseSpareLinks(t *testing.T) {
	d := New()
	count := maxInternalSliceSize * 3
	// Fills the deque
	for i := 0; i < count; i++ {
		d.PushBack(i)
	}
	// Pop the items to generate spare links
	for i := 0; i < maxInternalSliceSize; i++ {
		d.PopBack()
	}
	if d.spareLinks != 1 {
		t.Errorf("Expected: %d; Got: %d", 1, d.spareLinks)
	}

	// Push the items back using PushFront
	for i := 0; i < count; i++ {
		d.PushBack(i)
	}

	// The spare links should've been used up
	if d.spareLinks != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.spareLinks)
	}
}

func TestPopFrontWithRefillShouldKeepMaxSpareLinks(t *testing.T) {
	d := New()
	count := maxInternalSliceSize * (maxSpareLinks + 2)
	for i := 0; i < refillCount; i++ {
		for j := 0; j < count; j++ {
			d.PushBack(j)
		}
		for j := 0; j < count; j++ {
			if v, ok := d.PopFront(); !ok || v.(int) != j {
				t.Errorf("Expected: %d; Got: %d", j, v)
			}
			if d.spareLinks > maxSpareLinks {
				t.Fatalf("Too many spare links : got %d; want <= %d", d.spareLinks, maxSpareLinks)
			}
		}

		// Count the actual number of spare links
		actualSpareLinks := 0
		tmp := d.head.n // Initially head and tail points to the same slice.
		for tmp != nil {
			if tmp != d.tail {
				actualSpareLinks++
			} else {
				break
			}
			tmp = tmp.n
		}
		if actualSpareLinks != maxSpareLinks {
			t.Errorf("Expected: %d; Got: %d", d.spareLinks, actualSpareLinks)
		}
	}
}

func TestPopBackWithRefillShouldKeepMaxSpareLinks(t *testing.T) {
	d := New()
	count := maxInternalSliceSize * (maxSpareLinks + 2)
	for i := 0; i < refillCount; i++ {
		for j := 0; j < count; j++ {
			d.PushFront(j)
		}
		for j := 0; j < count; j++ {
			if v, ok := d.PopBack(); !ok || v.(int) != j {
				t.Errorf("Expected: %d; Got: %d", j, v)
			}
			if d.spareLinks > maxSpareLinks {
				t.Fatalf("Too many spare links : got %d; want <= %d", d.spareLinks, maxSpareLinks)
			}
		}

		// Count the actual number of spare links
		actualSpareLinks := 0
		tmp := d.head.n // Initially head and tail points to the same slice.
		for tmp != nil {
			if tmp != d.tail {
				actualSpareLinks++
			} else {
				break
			}
			tmp = tmp.n
		}
		if actualSpareLinks != maxSpareLinks {
			t.Errorf("Expected: %d; Got: %d", maxSpareLinks, actualSpareLinks)
		}
	}
}

// Helper methods-----------------------------------------------------------------------------------

// Checks the internal slices and its links.
func checkLinks(t *testing.T, d *Deque, length, headSliceSize, tailSliceSize, spareLinks int, headNext, headPrevious, tailNext, tailPrevious *node) {
	t.Helper()
	if d.Len() != length {
		t.Errorf("unexpected length; Expected: %d; Got: %d", length, d.Len())
	}
	if len(d.head.v) != headSliceSize {
		t.Errorf("unexpected head size; Expected: %d; Got: %d", headSliceSize, len(d.head.v))
	}
	if d.head.n != headNext {
		t.Error("unexpected head node; Expected: d.head.n == headNext; Got: d.head.n != headNext")
	}
	if d.head.p != headPrevious {
		t.Error("unexpected head; Expected: d.head.p == headPrevious; Got: d.head.p != headPrevious")
	}
	if d.tail.n != tailNext {
		t.Error("unexpected tailNext; Expected: d.tail.n == tailNext; Got: d.tail.n != tailNext")
	}
	if d.tail.p != tailPrevious {
		t.Error("unexpected tailPrevious; Expected: d.tail.p == tailPrevious; Got: d.tail.p != tailPrevious")
	}
	if len(d.tail.v) != tailSliceSize {
		t.Errorf("unexpected tail size; Expected: %d; Got: %d", tailSliceSize, len(d.tail.v))
	}
	if d.spareLinks != spareLinks {
		t.Errorf("unexpected spare link count; Expected: %d; Got: %d", spareLinks, d.spareLinks)
	}
	if t.Failed() {
		t.FailNow()
	}
}

// assertInvariants checks all the invariant conditions in d that we can think of.
// If val is non-nil it is used to find the expected value for an item at index
// i measured from the head of the queue.
func assertInvariants(t *testing.T, d *Deque, val func(i int) interface{}) {
	t.Helper()
	fail := func(what string, got, want interface{}) {
		t.Errorf("invariant fail: %s; got %v want %v", what, got, want)
	}
	if d == nil {
		fail("non-nil Deque", d, "non-nil")
	}
	if d.head == nil {
		// Zero value.
		if d.tail != nil {
			fail("nil tail when zero", d.tail, nil)
		}
		if d.len != 0 {
			fail("zero length when zero", d.len, 0)
		}
		if d.hp != 0 {
			fail("zero hp when zero", d.hp, 0)
		}
		if d.tp != 0 {
			fail("zero tp when zero", d.tp, 0)
		}
		if d.spareLinks != 0 {
			fail("no spare links when zero", d.spareLinks, 0)
		}
		return
	}
	if d.hp < 0 {
		fail("head index non-negative", d.hp, 0)
	}
	if d.hp > len(d.head.v) {
		fail("head index out of range", d.hp, len(d.head.v))
	}
	if d.tp < 0 {
		fail("tail index non-negative", d.tp, 0)
	}
	if d.tp > len(d.tail.v) {
		fail("head index out of range", d.tp, len(d.tail.v))
	}

	if d.head == d.tail {
		if d.tp-d.hp != d.len {
			fail("tail index - head index == len", d.tp-d.hp, d.len)
		}
	}
	spareLinkCount := 0
	inQueue := true
	elemCount := 0
	smallNodeCount := 0
	index := 0
	walkLinks(t, d, func(n *node) {
		if len(n.v) < maxInternalSliceSize {
			smallNodeCount++
			if len(n.v) > maxFirstSliceSize {
				fail("first node within bounds", len(n.v), maxFirstSliceSize)
			}
		}
		if len(n.v) > maxInternalSliceSize {
			fail("slice too big", len(n.v), maxInternalSliceSize)
		}
		for i, v := range n.v {
			failElem := func(what string, got, want interface{}) {
				fail(fmt.Sprintf("at elem %d, node %p, %s", i, n, what), got, want)
				t.FailNow()
			}
			if !inQueue {
				if v != nil {
					failElem("all values outside queue nil", v, nil)
				}
				continue
			}
			elemInQueue := inQueue
			switch {
			case n == d.head && n == d.tail:
				elemInQueue = i >= d.hp && i < d.tp
			case n == d.head:
				elemInQueue = i >= d.hp
			case n == d.tail:
				elemInQueue = i < d.tp
			}
			if elemInQueue {
				if v == nil {
					failElem("all values inside queue non-nil", v, "non-nil")
				}
			} else {
				if v != nil {
					failElem("all values outside queue nil", v, nil)
				}
			}
			if v != nil {
				if val != nil {
					want := val(index)
					if want != v {
						failElem(fmt.Sprintf("element %d has expected value", index), v, want)
					}
				}
				elemCount++
				index++
			}
		}
		if !inQueue {
			spareLinkCount++
		}
		if n == d.tail {
			inQueue = false
		}
	})
	if inQueue {
		// We never encountered the tail pointer.
		t.Errorf("tail does not point to element in list")
	}
	if spareLinkCount > maxSpareLinks {
		fail("spare link count <= maxSpareLinks", spareLinkCount, maxSpareLinks)
	}
	if elemCount != d.len {
		fail("element count == d.len", elemCount, d.len)
	}
	if smallNodeCount > 1 {
		fail("only one first node", smallNodeCount, 1)
	}
	if t.Failed() {
		t.FailNow()
	}
}

// walkLinks calls f for each node in the linked list.
// It also checks link invariants:
func walkLinks(t *testing.T, d *Deque, f func(n *node)) {
	t.Helper()
	fail := func(what string, got, want interface{}) {
		t.Errorf("link invariant %s fail; got %v want %v", what, got, want)
	}
	n := d.head
	for {
		if n.n.p != n {
			fail("node.n.p == node", n.n.p, n)
		}
		if n.p.n != n {
			fail("node.p.n == node", n.p.n, n)
		}
		f(n)
		n = n.n
		if n == d.head {
			break
		}
	}
}
