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

func BenchmarkMicroserviceListQueue(b *testing.B) {
	var l *list.List
	benchmarkMicroservice(
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

func BenchmarkMicroserviceListStack(b *testing.B) {
	var l *list.List
	benchmarkMicroservice(
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

func BenchmarkMicroserviceSliceQueue(b *testing.B) {
	var q *CustomSliceQueue
	benchmarkMicroservice(
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

func BenchmarkMicroserviceSliceStack(b *testing.B) {
	var q *CustomSliceQueue
	benchmarkMicroservice(
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

func BenchmarkMicroserviceGammazeroQueue(b *testing.B) {
	var q *gammazero.Deque
	benchmarkMicroservice(
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

func BenchmarkMicroserviceGammazeroStack(b *testing.B) {
	var q *gammazero.Deque
	benchmarkMicroservice(
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

func BenchmarkMicroservicePhfQueue(b *testing.B) {
	var q *phf.Queue
	benchmarkMicroservice(
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

func BenchmarkMicroservicePhfStack(b *testing.B) {
	var q *phf.Queue
	benchmarkMicroservice(
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

func BenchmarkMicroserviceCookiejarQueue(b *testing.B) {
	var q *cookiejar.Deque
	benchmarkMicroservice(
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

func BenchmarkMicroserviceCookiejarStack(b *testing.B) {
	var q *cookiejar.Deque
	benchmarkMicroservice(
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

func BenchmarkMicroserviceJujuQueue(b *testing.B) {
	var q *juju.Deque
	benchmarkMicroservice(
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

func BenchmarkMicroserviceJujuStack(b *testing.B) {
	var q *juju.Deque
	benchmarkMicroservice(
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

func BenchmarkMicroserviceImpl7Queue(b *testing.B) {
	var q *queueimpl7.Queueimpl7
	benchmarkMicroservice(
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

func BenchmarkMicroserviceDequeQueue(b *testing.B) {
	var q *deque.Deque
	benchmarkMicroservice(
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

func BenchmarkMicroserviceDequeStack(b *testing.B) {
	var q *deque.Deque
	benchmarkMicroservice(
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

func benchmarkMicroservice(b *testing.B, initInstance func(), push func(v interface{}), pop func() (interface{}, bool), empty func() bool) {
	for i, test := range tests {
		// Doesn't run the first (0 items) as 0 items makes no sense for this test.
		if i == 0 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				initInstance()

				// Simulate stable traffic
				for i := 0; i < test.count; i++ {
					push(getTestValue(i))
					pop()
				}

				// Simulate slowly increasing traffic
				for i := 0; i < test.count; i++ {
					push(getTestValue(i))
					push(getTestValue(i))
					pop()
				}

				// Simulate slowly decreasing traffic, bringing traffic back to normal
				for i := 0; i < test.count; i++ {
					pop()
					if !empty() {
						pop()
					}
					push(getTestValue(i))
				}

				// Simulate quick traffic spike (DDOS attack, etc)
				for i := 0; i < test.count; i++ {
					push(getTestValue(i))
				}

				// Simulate stable traffic while at high traffic
				for i := 0; i < test.count; i++ {
					push(getTestValue(i))
					pop()
				}

				// Simulate going back to normal (DDOS attack fended off)
				for i := 0; i < test.count; i++ {
					pop()
				}

				// Simulate stable traffic (now that is back to normal)
				for i := 0; i < test.count; i++ {
					push(getTestValue(i))
					pop()
				}
			}
		})
	}
}
