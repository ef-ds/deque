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
	"testing"
)

const (
	pushCount = maxInternalSliceSize * 3 // Push to fill at least 3 internal slices
)

// Unit tests ---------------------------------------------

func TestNewShouldReturnInitiazedInstanceOfDeque(t *testing.T) {
	d := New()

	if d == nil {
		t.Error("Expected: new instance of queue; Got: nil")
	}
	if d.Len() != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.Len())
	}
	if _, ok := d.Back(); ok {
		t.Error("Expected: false; Got: true")
	}
	if _, ok := d.Front(); ok {
		t.Error("Expected: false; Got: true")
	}
	if d.head != nil {
		t.Error("Expected: d.head == nil; Got: d.head != nil")
	}
	if d.tail != nil {
		t.Error("Expected: d.tail == nil; Got: d.tail != nil")
	}
	if d.tp != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.tp)
	}
	if d.hp != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.hp)
	}
	if d.spareLinks != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.spareLinks)
	}
}

func TestPushFrontPopBackShouldHaveAllInternalLinksInARing(t *testing.T) {
	d := New()
	pushValue, extraAddedItems, spareLinks := 0, 0, 0

	// Push maxFirstSliceSize items to fill the first array
	expectedHeadSliceSize := 1
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
	if d.tp != maxInternalSliceSize-2 {
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
	for i := 1; i <= maxFirstSliceSize-1; i++ {
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

	// Pop one extra item emptying the deque and making head move to next (first slice)
	if v, ok := d.PopFront(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	popValue--
	checkLinks(t, d, popValue, maxInternalSliceSize, maxFirstSliceSize, spareLinks, d.tail.p, d.tail, d.head, d.head.n)
	//Check middle  slice links
	if d.head.n.n != d.tail {
		t.Error("Expected: d.head.n.n == d.tail; Got: d.head.n.n != d.tail")
	}
	if d.head.n.p != d.tail.n {
		t.Error("Expected: d.head.n.p == d.tail.n; Got: d.head.n.p != d.tail.n")
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
	if len(d.tail.v) != maxFirstSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxFirstSliceSize, len(d.tail.v))
	}
	if d.head.p != d.tail {
		t.Error("Expected: d.head.p == d.tail; Got: d.head.p != d.tail")
	}
	if d.tp != -1 {
		t.Errorf("Expected: %d; Got: %d", -1, d.tp)
	}
	if d.hp != 0 {
		t.Errorf("Expected: %d; Got: %d", 0, d.hp)
	}
	if d.spareLinks != spareLinks {
		t.Errorf("Expected: %d; Got: %d", spareLinks, d.spareLinks)
	}
}

func TestPushBackPopBackShouldHaveAllInternalLinksInARing(t *testing.T) {
	d := New()
	pushValue, extraAddedItems, spareLinks := 0, 0, 0

	// Push maxFirstSliceSize items to fill the first array
	expectedHeadSliceSize := 1
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
	for i := 1; i <= maxFirstSliceSize-1; i++ {
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

	// Pop one extra item emptying the deque, making tail move to the previous, last slice
	if v, ok := d.PopBack(); !ok || v.(int) != popValue {
		t.Errorf("Expected: %d; Got: %d", popValue, v)
	}
	popValue--
	checkLinks(t, d, popValue, maxFirstSliceSize, maxInternalSliceSize, spareLinks, d.tail.p, d.tail, d.head, d.head.n)
	/// Check middle links
	if d.head.n.n != d.tail {
		t.Error("Expected: d.head.n.n == d.tail; Got: d.head.n.n != d.tail")
	}
	if d.head.n.p != d.head {
		t.Error("Expected: d.head.n.p == d.head; Got: d.head.n.p != d.head")
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
	if len(d.tail.v) != maxInternalSliceSize {
		t.Errorf("Expected: %d; Got: %d", maxInternalSliceSize, len(d.tail.v))
	}
	if d.head.n != d.tail.p {
		t.Error("Expected: d.head.n == d.tail.p; Got: d.head.n != d.tail.p")
	}
	if d.tp != maxInternalSliceSize-1 {
		t.Errorf("Expected: %d; Got: %d", maxInternalSliceSize-1, d.tp)
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
	if d.tp != 0 {
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
		d.PopFront()
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

// Integration tests ---------------------------------------------

func TestFillQueueShouldRetrieveAllElementsInOrder(t *testing.T) {
	d := New()

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
	const refillCount = 3
	d := New()

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
	const refillCount = 3
	d := New()
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
	const refillCount = 3
	d := New()

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
	d := New()
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
	d := New()

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

func TestFillStackShouldRetrieveAllElementsInOrder(t *testing.T) {
	d := New()

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
	const refillCount = 3
	d := New()

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
	const refillCount = 3
	d := New()
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
	const refillCount = 3
	d := New()

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
	d := New()
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
	d := New()

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

func TestPushPopFrontShouldRetrieveAllElementsInOrder(t *testing.T) {
	pushPopFrontBackShouldRetrieveAllElementsInOrder(
		t,
		func(d *Deque) (interface{}, bool) {
			return d.PopFront()
		},
		func(d *Deque) (interface{}, bool) {
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
		func(d *Deque) (interface{}, bool) {
			return d.PopBack()
		},
		func(d *Deque) (interface{}, bool) {
			return d.Back()
		},
		func(v, lastGet, lastPut interface{}) bool {
			return v == lastPut
		},
	)
}

func TestMixedPopFrontBackLenShouldReturnAllValuesInOrder(t *testing.T) {
	d := New()
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

// API tests -----------------------------------------------------

func TestPopFrontWithZeroValueShouldReturnReadyToUseDeque(t *testing.T) {
	var d Deque
	d.PushBack(1)
	d.PushBack(2)

	v, ok := d.PopFront()
	if !ok || v.(int) != 1 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	v, ok = d.PopFront()
	if !ok || v.(int) != 2 {
		t.Errorf("Expected: 2; Got: %d", v)
	}
	_, ok = d.PopFront()
	if ok {
		t.Error("Expected: empty slice (ok=false); Got: ok=true")
	}
}

func TestPopBackWithZeroValueShouldReturnReadyToUseDeque(t *testing.T) {
	var d Deque
	d.PushBack(1)
	d.PushBack(2)

	v, ok := d.PopBack()
	if !ok || v.(int) != 2 {
		t.Errorf("Expected: 2; Got: %d", v)
	}
	v, ok = d.PopBack()
	if !ok || v.(int) != 1 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	_, ok = d.PopBack()
	if ok {
		t.Error("Expected: empty slice (ok=false); Got: ok=true")
	}
}

func TestWithZeroValueAndEmptyShouldReturnAsEmpty(t *testing.T) {
	var d Deque
	if _, ok := d.Front(); ok {
		t.Error("Expected: false as the queue is empty; Got: true")
	}
	if _, ok := d.Back(); ok {
		t.Error("Expected: false as the queue is empty; Got: true")
	}
	if _, ok := d.PopFront(); ok {
		t.Error("Expected: false as the queue is empty; Got: true")
	}
	if _, ok := d.PopBack(); ok {
		t.Error("Expected: false as the queue is empty; Got: true")
	}
	if l := d.Len(); l != 0 {
		t.Errorf("Expected: 0 as the queue is empty; Got: %d", l)
	}
}

func TestInitShouldReturnEmptyDeque(t *testing.T) {
	var d Deque
	d.PushBack(1)

	d.Init()

	if _, ok := d.Front(); ok {
		t.Error("Expected: false as the queue is empty; Got: true")
	}
	if _, ok := d.Back(); ok {
		t.Error("Expected: false as the queue is empty; Got: true")
	}
	if _, ok := d.PopFront(); ok {
		t.Error("Expected: false as the queue is empty; Got: true")
	}
	if _, ok := d.PopBack(); ok {
		t.Error("Expected: false as the queue is empty; Got: true")
	}
	if l := d.Len(); l != 0 {
		t.Errorf("Expected: 0 as the queue is empty; Got: %d", l)
	}
}

func TestPopFrontWithNilValuesShouldReturnAllValuesInOrder(t *testing.T) {
	d := New()
	d.PushBack(1)
	d.PushBack(nil)
	d.PushBack(2)
	d.PushBack(nil)

	v, ok := d.PopFront()
	if !ok || v.(int) != 1 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	v, ok = d.PopFront()
	if !ok || v != nil {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	v, ok = d.PopFront()
	if !ok || v.(int) != 2 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	v, ok = d.PopFront()
	if !ok || v != nil {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	_, ok = d.PopFront()
	if ok {
		t.Error("Expected: empty slice (ok=false); Got: ok=true")
	}
}

func TestPopBackWithNilValuesShouldReturnAllValuesInOrder(t *testing.T) {
	d := New()
	d.PushBack(1)
	d.PushBack(nil)
	d.PushBack(2)
	d.PushBack(nil)

	v, ok := d.PopBack()
	if !ok || v != nil {
		t.Errorf("Expected: nil; Got: %d", v)
	}
	v, ok = d.PopBack()
	if !ok || v.(int) != 2 {
		t.Errorf("Expected: 2; Got: %d", v)
	}
	v, ok = d.PopBack()
	if !ok || v != nil {
		t.Errorf("Expected: nil; Got: %d", v)
	}
	v, ok = d.PopBack()
	if !ok || v.(int) != 1 {
		t.Errorf("Expected: 1; Got: %d", v)
	}
	_, ok = d.PopBack()
	if ok {
		t.Error("Expected: empty slice (ok=false); Got: ok=true")
	}
}

// Helper methods ------------------------------------------------

// Checks the internal slices and its links.
func checkLinks(t *testing.T, d *Deque, length, headSliceSize, tailSliceSize, spareLinks int, headNext, headPrevious, tailNext, tailPrevious *node) {
	if d.Len() != length {
		t.Errorf("Expected: %d; Got: %d", length, d.Len())
	}
	if len(d.head.v) != headSliceSize {
		t.Errorf("Expected: %d; Got: %d", headSliceSize, len(d.head.v))
	}
	if d.head.n != headNext {
		t.Error("Expected: d.head.n == headNext; Got: d.head.n != headNext")
	}
	if d.head.p != headPrevious {
		t.Error("Expected: d.head.p == headPrevious; Got: d.head.p != headPrevious")
	}
	if d.tail.n != tailNext {
		t.Error("Expected: d.tail.n == tailNext; Got: d.tail.n != tailNext")
	}
	if d.tail.p != tailPrevious {
		t.Error("Expected: d.tail.p == tailPrevious; Got: d.tail.p != tailPrevious")
	}
	if len(d.tail.v) != tailSliceSize {
		t.Errorf("Expected: %d; Got: %d", tailSliceSize, len(d.tail.v))
	}
	if d.spareLinks != spareLinks {
		t.Errorf("Expected: %d; Got: %d", spareLinks, d.spareLinks)
	}
}

func pushPopFrontBackShouldRetrieveAllElementsInOrder(t *testing.T, popFunc, frontbackFunc func(*Deque) (interface{}, bool), checkValueFunc func(v, lastGet, lastPut interface{}) bool) {
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
			d := New()
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
