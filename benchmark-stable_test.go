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
	"container/list"
	"strconv"
	"testing"

	"github.com/christianrpetrin/queue-tests/queueimpl7"
	"github.com/ef-ds/deque"
	gammazero "github.com/gammazero/deque"
	juju "github.com/juju/utils/deque"
	phf "github.com/phf/go-queue/queue"
	cookiejar "gopkg.in/karalabe/cookiejar.v2/collections/deque"
)

func BenchmarkStableListQueue(b *testing.B) {
	var l *list.List
	benchmarkStable(
		b,
		func() {
			l = list.New()
		},
		func(v interface{}) {
			l.PushBack(v)
		},
		func() (interface{}, bool) {
			return l.Remove(l.Front()), true
		},
		func() bool {
			return l.Front() == nil
		},
	)
}

func BenchmarkStableListStack(b *testing.B) {
	var l *list.List
	benchmarkStable(
		b,
		func() {
			l = list.New()
		},
		func(v interface{}) {
			l.PushBack(v)
		},
		func() (interface{}, bool) {
			return l.Remove(l.Back()), true
		},
		func() bool {
			return l.Front() == nil
		},
	)
}

func BenchmarkStableSliceQueue(b *testing.B) {
	var q *CustomSliceQueue
	benchmarkStable(
		b,
		func() {
			q = NewCustomSliceQueue()
		},
		func(v interface{}) {
			q.PushBack(v.(*testValue))
		},
		func() (interface{}, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStableSliceStack(b *testing.B) {
	var q *CustomSliceQueue
	benchmarkStable(
		b,
		func() {
			q = NewCustomSliceQueue()
		},
		func(v interface{}) {
			q.PushBack(v.(*testValue))
		},
		func() (interface{}, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStableGammazeroQueue(b *testing.B) {
	var q *gammazero.Deque
	benchmarkStable(
		b,
		func() {
			q = new(gammazero.Deque)
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopFront(), true
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStableGammazeroStack(b *testing.B) {
	var q *gammazero.Deque
	benchmarkStable(
		b,
		func() {
			q = new(gammazero.Deque)
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopBack(), true
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStablePhfQueue(b *testing.B) {
	var q *phf.Queue
	benchmarkStable(
		b,
		func() {
			q = phf.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopFront(), true
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStablePhfStack(b *testing.B) {
	var q *phf.Queue
	benchmarkStable(
		b,
		func() {
			q = phf.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopBack(), true
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStableCookiejarQueue(b *testing.B) {
	var q *cookiejar.Deque
	benchmarkStable(
		b,
		func() {
			q = cookiejar.New()
		},
		func(v interface{}) {
			q.PushRight(v)
		},
		func() (interface{}, bool) {
			return q.PopLeft(), true
		},
		func() bool {
			return q.Size() == 0
		},
	)
}

func BenchmarkStableCookiejarStack(b *testing.B) {
	var q *cookiejar.Deque
	benchmarkStable(
		b,
		func() {
			q = cookiejar.New()
		},
		func(v interface{}) {
			q.PushRight(v)
		},
		func() (interface{}, bool) {
			return q.PopRight(), true
		},
		func() bool {
			return q.Size() == 0
		},
	)
}

func BenchmarkStableJujuQueue(b *testing.B) {
	var q *juju.Deque
	benchmarkStable(
		b,
		func() {
			q = juju.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStableJujuStack(b *testing.B) {
	var q *juju.Deque
	benchmarkStable(
		b,
		func() {
			q = juju.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStableImpl7Queue(b *testing.B) {
	var q *queueimpl7.Queueimpl7
	benchmarkStable(
		b,
		func() {
			q = queueimpl7.New()
		},
		func(v interface{}) {
			q.Push(v)
		},
		func() (interface{}, bool) {
			return q.Pop()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStableDequeQueue(b *testing.B) {
	var q *deque.Deque
	benchmarkStable(
		b,
		func() {
			q = deque.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkStableDequeStack(b *testing.B) {
	var q *deque.Deque
	benchmarkStable(
		b,
		func() {
			q = deque.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func benchmarkStable(b *testing.B, initInstance func(), push func(v interface{}), pop func() (interface{}, bool), empty func() bool) {
	initInstance()
	for i := 0; i < fillCount; i++ {
		push(getTestValue(i))
	}

	for i, test := range tests {
		// Doesn't run the first (0 items) test as 0 items makes no sense for this test.
		if i == 0 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < test.count; i++ {
					push(getTestValue(i))
					tmp, tmp2 = pop()
				}

			}
		})
	}

	for !empty() {
		tmp, tmp2 = pop()
	}
}
