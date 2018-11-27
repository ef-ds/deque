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

func TestPopFrontWithZeroValueShouldReturnReadyToUseDeque(t *testing.T) {
	var d deque.Deque
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
	var d deque.Deque
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
	var d deque.Deque
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
	var d deque.Deque
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
	d := deque.New()
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
	d := deque.New()
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
