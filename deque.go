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

// Package deque implements a high performance, general purpose, dynamically growing,
// ring shaped, linked arrays double-ended queue.
package deque

const (
	// firstSliceSize holds the size of the first slice.
	firstSliceSize = 1

	// sliceGrowthFactor ...
	sliceGrowthFactor = 4

	// firstSliceLastPosition ...
	firstSliceLastPosition = 15

	// maxFirstSliceSize holds the maximum size of the first slice.
	maxFirstSliceSize = 16

	// maxInternalSliceSize holds the maximum size of each internal slice.
	internalSliceLastPosition = 255

	// maxInternalSliceSize holds the maximum size of each internal slice.
	maxInternalSliceSize = 256

	// maxSpareLinks ...
	maxSpareLinks = 4
)

// Deque implements an unbounded, dynamically growing double-ended-queue (deque).
// The zero value for deque is an empty deque ready to use.
type Deque struct {
	// Head points to the first node of the linked list.
	head *node

	// Tail points to the last node of the linked list.
	// In an empty deque, head and tail points to the same node.
	tail *node

	// Hp is the index pointing to the current first element in the deque
	// (i.e. first element added in the current deque values).
	hp int

	// tp is the index pointing to the current last element in the deque
	// (i.e. last element added in the current deque values).
	tp int

	// Len holds the current deque values length.
	len int

	// lastTailPosition holds the index pointing to the last tail position.
	lastTailPosition int

	// spareLinks holds the number of already used, but now empty, ready-to-be-reused, slices.
	spareLinks int
}

// Node represents a deque node.
// Each node holds a slice of user managed values.
type node struct {
	// v holds the list of user added values in this node.
	v []interface{}

	// n points to the next node in the linked list.
	n *node

	// p points to the previous node in the linked list.
	p *node
}

// New returns an initialized deque.
func New() *Deque {
	return new(Deque).Init()
}

// Init initializes or clears deque d.
func (d *Deque) Init() *Deque {
	d.head = nil
	d.tail = nil
	d.hp = 0
	d.tp = 0
	d.len = 0
	d.spareLinks = 0
	return d
}

// Len returns the number of elements of deque d.
// The complexity is O(1).
func (d *Deque) Len() int { return d.len }

// Front returns the first element of deque q or nil if the deque is empty.
// The second, bool result indicates whether a valid value was returned;
//   if the deque is empty, false will be returned.
// The complexity is O(1).
func (d *Deque) Front() (interface{}, bool) {
	if d.len == 0 {
		return nil, false
	}
	return d.head.v[d.hp], true
}

// Back returns the last element of deque q or nil if the deque is empty.
// The second, bool result indicates whether a valid value was returned;
//   if the deque is empty, false will be returned.
// The complexity is O(1).
func (d *Deque) Back() (interface{}, bool) {
	if d.len == 0 {
		return nil, false
	}
	return d.tail.v[d.tp], true
}

// PushFront adds value v to the the front of the deque.
// The complexity is O(1).
func (d *Deque) PushFront(v interface{}) {
	if d.head == nil {
		h := &node{v: make([]interface{}, firstSliceSize)}
		h.n = h
		h.p = h
		d.head = h
		d.tail = h
		d.lastTailPosition = firstSliceLastPosition
	} else {
		if d.hp > 0 {
			d.hp--
		} else if d.head.p != d.tail {
			d.head = d.head.p
			d.hp = len(d.head.v) - 1
			d.spareLinks--
		} else {
			if d.head == d.tail {
				d.tp++
				if d.tp >= len(d.head.v) && len(d.head.v) < maxFirstSliceSize {
					d.head.v = resize(d.head.v)
					copy(d.head.v[d.hp+1:], d.head.v[d.hp:])
				} else if d.tp > d.lastTailPosition {
					n := &node{v: make([]interface{}, maxInternalSliceSize)}
					n.n = d.head
					n.p = d.tail
					d.head.p = n
					d.tail.n = n
					d.head = n
					d.hp = internalSliceLastPosition
					d.tp = len(d.tail.v) - 1
					d.lastTailPosition = internalSliceLastPosition
				} else {
					copy(d.head.v[d.hp+1:], d.head.v[d.hp:])
				}
			} else {
				n := &node{v: make([]interface{}, maxInternalSliceSize)}
				n.n = d.head
				n.p = d.tail
				d.head.p = n
				d.tail.n = n
				d.head = n
				d.hp = internalSliceLastPosition
			}
		}
	}

	d.len++
	d.head.v[d.hp] = v
}

// PushBack adds value v to the the back of the deque.
// The complexity is O(1).
func (d *Deque) PushBack(v interface{}) {
	if d.head == nil {
		h := &node{v: make([]interface{}, firstSliceSize)}
		h.n = h
		h.p = h
		d.head = h
		d.tail = h
		d.lastTailPosition = firstSliceLastPosition
	} else if d.tp >= d.lastTailPosition {
		var n *node
		if d.tail.n != d.head {
			d.spareLinks--
			n = d.tail.n
		} else {
			n = &node{v: make([]interface{}, maxInternalSliceSize)}
			n.n = d.head
			n.p = d.tail
			d.tail.n = n
			d.head.p = n
			d.lastTailPosition = internalSliceLastPosition
		}
		d.tp = 0
		d.tail = n
	} else {
		d.tp++
		if d.tp >= len(d.tail.v) {
			d.tail.v = resize(d.tail.v)
		}
	}

	d.len++
	d.tail.v[d.tp] = v
}

// PopFront retrieves and removes the current element from the front of the deque.
// The second, bool result indicates whether a valid value was returned;
// 	if the deque is empty, false will be returned.
// The complexity is O(1).
func (d *Deque) PopFront() (interface{}, bool) {
	if d.len == 0 {
		return nil, false
	}

	v := d.head.v[d.hp]
	d.head.v[d.hp] = nil // Avoid memory leaks
	d.len--
	d.hp++
	if d.hp >= maxFirstSliceSize && d.hp >= len(d.head.v) {
		d.hp = 0
		if d.head == d.tail {
			d.tp = -1
		} else {
			d.spareLinks++
			if d.spareLinks >= maxSpareLinks {
				d.head.p.n = d.head.n // Eliminate this link
			}
		}
		d.head = d.head.n
	}
	return v, true
}

// PopBack retrieves and removes the current element from the back of the deque.
// The second, bool result indicates whether a valid value was returned;
// 	if the deque is empty, false will be returned.
// The complexity is O(1).
func (d *Deque) PopBack() (interface{}, bool) {
	if d.len == 0 {
		return nil, false
	}

	d.len--
	v := d.tail.v[d.tp]
	d.tail.v[d.tp] = nil // Avoid memory leaks
	d.tp--
	if d.tp < 0 {
		if d.head != d.tail {
			d.spareLinks++
			if d.spareLinks >= maxSpareLinks {
				d.head.p.n = d.head.n // Eliminate this link
			}
		}
		d.tail = d.tail.p
		d.tp = len(d.tail.v) - 1
	}
	return v, true
}

func resize(s []interface{}) []interface{} {
	n := make([]interface{}, len(s)*sliceGrowthFactor)
	copy(n, s)
	return n
}
