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

// import (
// 	"container/list"
// 	"strconv"
// 	"testing"

// 	"github.com/christianrpetrin/queue-tests/queueimpl7"
// 	phf "github.com/phf/go-queue/queue"
// 	cookiejar "gopkg.in/karalabe/cookiejar.v2/collections/queue"
// )

//egonelbre "github.com/egonelbre/exp/queue/blockqueue"

// Refill tests--------------------------------------------------------------------------------------

// func BenchmarkRefillList(b *testing.B) {
// 	var l *list.List
// 	benchmarkRefill(
// 		b,
// 		func() {
// 			l = list.New()
// 		},
// 		func(v interface{}) {
// 			l.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return l.Remove(l.Front()), true
// 		},
// 		func() bool {
// 			return l.Front() == nil
// 		},
// 	)
// }

// func BenchmarkRefillSlice(b *testing.B) {
// 	var q *CustomSliceQueue
// 	benchmarkRefill(
// 		b,
// 		func() {
// 			q = NewCustomSliceQueue()
// 		},
// 		func(v interface{}) {
// 			q.Push(v.(*testValue))
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkRefillPhf(b *testing.B) {
// 	var q *phf.Queue
// 	benchmarkRefill(
// 		b,
// 		func() {
// 			q = phf.New()
// 		},
// 		func(v interface{}) {
// 			q.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront(), true
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkRefillCookiejar(b *testing.B) {
// 	var q *cookiejar.Queue
// 	benchmarkRefill(
// 		b,
// 		func() {
// 			q = cookiejar.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop(), true
// 		},
// 		func() bool {
// 			return q.Size() == 0
// 		},
// 	)
// }

// // func BenchmarkRefillEgonelbre(b *testing.B) {
// // 	var q *egonelbre.Queue
// // 	benchmarkRefill(
// // 		b,
// // 		func() {
// // 			q = egonelbre.New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			return q.Pop()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func BenchmarkRefillImpl7(b *testing.B) {
// 	var q *queueimpl7.Queueimpl7
// 	benchmarkRefill(
// 		b,
// 		func() {
// 			q = queueimpl7.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkRefillDequePopFront(b *testing.B) {
// 	var q *Deque
// 	benchmarkRefill(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkRefillDequePopBack(b *testing.B) {
// 	var q *Deque
// 	benchmarkRefill(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopBack()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// // func BenchmarkRefillDequePopMixed(b *testing.B) {
// // 	i := 0
// // 	var q *Deque
// // 	benchmarkRefill(
// // 		b,
// // 		func() {
// // 			q = New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			i++
// // 			if i%2 == 0 {
// // 				return q.PopBack()
// // 			}
// // 			return q.PopFront()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func benchmarkRefill(b *testing.B, initInstance func(), push func(v interface{}), pop func() (interface{}, bool), empty func() bool) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			initInstance()
// 			for n := 0; n < b.N; n++ {
// 				for n := 0; n < refillCount; n++ {
// 					for i := 0; i < test.count; i++ {
// 						push(getTestValue(i))
// 					}
// 					for !empty() {
// 						tmp, tmp2 = pop()
// 					}
// 				}
// 			}
// 		})
// 	}
// }

// // Stable tests--------------------------------------------------------------------------------------

// func BenchmarkStableList(b *testing.B) {
// 	var l *list.List
// 	benchmarkStable(
// 		b,
// 		func() {
// 			l = list.New()
// 		},
// 		func(v interface{}) {
// 			l.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return l.Remove(l.Front()), true
// 		},
// 		func() bool {
// 			return l.Front() == nil
// 		},
// 	)
// }

// func BenchmarkStableSlice(b *testing.B) {
// 	var q *CustomSliceQueue
// 	benchmarkStable(
// 		b,
// 		func() {
// 			q = NewCustomSliceQueue()
// 		},
// 		func(v interface{}) {
// 			q.Push(v.(*testValue))
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkStablePhf(b *testing.B) {
// 	var q *phf.Queue
// 	benchmarkStable(
// 		b,
// 		func() {
// 			q = phf.New()
// 		},
// 		func(v interface{}) {
// 			q.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront(), true
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkStableCookiejar(b *testing.B) {
// 	var q *cookiejar.Queue
// 	benchmarkStable(
// 		b,
// 		func() {
// 			q = cookiejar.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop(), true
// 		},
// 		func() bool {
// 			return q.Size() == 0
// 		},
// 	)
// }

// // func BenchmarkStableEgonelbre(b *testing.B) {
// // 	var q *egonelbre.Queue
// // 	benchmarkStable(
// // 		b,
// // 		func() {
// // 			q = egonelbre.New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			return q.Pop()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func BenchmarkStableImpl7(b *testing.B) {
// 	var q *queueimpl7.Queueimpl7
// 	benchmarkStable(
// 		b,
// 		func() {
// 			q = queueimpl7.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkStableDequePopFront(b *testing.B) {
// 	var q *Deque
// 	benchmarkStable(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkStableDequePopBack(b *testing.B) {
// 	var q *Deque
// 	benchmarkStable(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopBack()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// // func BenchmarkStableDequePopMixed(b *testing.B) {
// // 	i := 0
// // 	var q *Deque
// // 	benchmarkStable(
// // 		b,
// // 		func() {
// // 			q = New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			i++
// // 			if i%2 == 0 {
// // 				return q.PopBack()
// // 			}
// // 			return q.PopFront()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func benchmarkStable(b *testing.B, initInstance func(), push func(v interface{}), pop func() (interface{}, bool), empty func() bool) {
// 	initInstance()
// 	for i := 0; i < fillCount; i++ {
// 		push(getTestValue(i))
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for i := 0; i < test.count; i++ {
// 					push(getTestValue(i))
// 					tmp, tmp2 = pop()
// 				}

// 			}
// 		})
// 	}

// 	for !empty() {
// 		tmp, tmp2 = pop()
// 	}
// }

// // Saw pattern tests--------------------------------------------------------------------------------------

// func BenchmarkRefillWhileFullList(b *testing.B) {
// 	var l *list.List
// 	benchmarkRefillWhileFull(
// 		b,
// 		func() {
// 			l = list.New()
// 		},
// 		func(v interface{}) {
// 			l.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return l.Remove(l.Front()), true
// 		},
// 		func() bool {
// 			return l.Front() == nil
// 		},
// 	)
// }

// func BenchmarkRefillWhileFullSlice(b *testing.B) {
// 	var q *CustomSliceQueue
// 	benchmarkRefillWhileFull(
// 		b,
// 		func() {
// 			q = NewCustomSliceQueue()
// 		},
// 		func(v interface{}) {
// 			q.Push(v.(*testValue))
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkRefillWhileFullPhf(b *testing.B) {
// 	var q *phf.Queue
// 	benchmarkRefillWhileFull(
// 		b,
// 		func() {
// 			q = phf.New()
// 		},
// 		func(v interface{}) {
// 			q.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront(), true
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkRefillWhileFullCookiejar(b *testing.B) {
// 	var q *cookiejar.Queue
// 	benchmarkRefillWhileFull(
// 		b,
// 		func() {
// 			q = cookiejar.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop(), true
// 		},
// 		func() bool {
// 			return q.Size() == 0
// 		},
// 	)
// }

// // func BenchmarkRefillWhileFullEgonelbre(b *testing.B) {
// // 	var q *egonelbre.Queue
// // 	benchmarkRefillWhileFull(
// // 		b,
// // 		func() {
// // 			q = egonelbre.New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			return q.Pop()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func BenchmarkRefillWhileFullImpl7(b *testing.B) {
// 	var q *queueimpl7.Queueimpl7
// 	benchmarkRefillWhileFull(
// 		b,
// 		func() {
// 			q = queueimpl7.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkRefillWhileFullDequePopFront(b *testing.B) {
// 	var q *Deque
// 	benchmarkRefillWhileFull(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkRefillWhileFullDequePopBack(b *testing.B) {
// 	var q *Deque
// 	benchmarkRefillWhileFull(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopBack()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// // func BenchmarkRefillWhileFullDequePopMixed(b *testing.B) {
// // 	i := 0
// // 	var q *Deque
// // 	benchmarkRefillWhileFull(
// // 		b,
// // 		func() {
// // 			q = New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			i++
// // 			if i%2 == 0 {
// // 				return q.PopBack()
// // 			}
// // 			return q.PopFront()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func benchmarkRefillWhileFull(b *testing.B, initInstance func(), push func(v interface{}), pop func() (interface{}, bool), empty func() bool) {
// 	initInstance()
// 	for i := 0; i < fillCount; i++ {
// 		push(getTestValue(i))
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for k := 0; k < refillCount; k++ {
// 					for i := 0; i < test.count; i++ {
// 						push(getTestValue(i))
// 					}
// 					for i := 0; i < test.count; i++ {
// 						tmp, tmp2 = pop()
// 					}
// 				}
// 			}
// 		})
// 	}
// }

// // Slow incereasing tests--------------------------------------------------------------------------------------

// func BenchmarkSlowIncreaseList(b *testing.B) {
// 	var l *list.List
// 	benchmarkSlowIncrease(
// 		b,
// 		func() {
// 			l = list.New()
// 		},
// 		func(v interface{}) {
// 			l.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return l.Remove(l.Front()), true
// 		},
// 		func() bool {
// 			return l.Front() == nil
// 		},
// 	)
// }

// func BenchmarkSlowIncreaseSlice(b *testing.B) {
// 	var q *CustomSliceQueue
// 	benchmarkSlowIncrease(
// 		b,
// 		func() {
// 			q = NewCustomSliceQueue()
// 		},
// 		func(v interface{}) {
// 			q.Push(v.(*testValue))
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkSlowIncreasePhf(b *testing.B) {
// 	var q *phf.Queue
// 	benchmarkSlowIncrease(
// 		b,
// 		func() {
// 			q = phf.New()
// 		},
// 		func(v interface{}) {
// 			q.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront(), true
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkSlowIncreaseCookiejar(b *testing.B) {
// 	var q *cookiejar.Queue
// 	benchmarkSlowIncrease(
// 		b,
// 		func() {
// 			q = cookiejar.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop(), true
// 		},
// 		func() bool {
// 			return q.Size() == 0
// 		},
// 	)
// }

// // func BenchmarkSlowIncreaseEgonelbre(b *testing.B) {
// // 	var q *egonelbre.Queue
// // 	benchmarkSlowIncrease(
// // 		b,
// // 		func() {
// // 			q = egonelbre.New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			return q.Pop()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func BenchmarkSlowIncreaseImpl7(b *testing.B) {
// 	var q *queueimpl7.Queueimpl7
// 	benchmarkSlowIncrease(
// 		b,
// 		func() {
// 			q = queueimpl7.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkSlowIncreaseDequePopFront(b *testing.B) {
// 	var q *Deque
// 	benchmarkSlowIncrease(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkSlowIncreaseDequePopBack(b *testing.B) {
// 	var q *Deque
// 	benchmarkSlowIncrease(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopBack()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// // func BenchmarkSlowIncreaseDequePopMixed(b *testing.B) {
// // 	i := 0
// // 	var q *Deque
// // 	benchmarkSlowIncrease(
// // 		b,
// // 		func() {
// // 			q = New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			i++
// // 			if i%2 == 0 {
// // 				return q.PopBack()
// // 			}
// // 			return q.PopFront()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func benchmarkSlowIncrease(b *testing.B, initInstance func(), push func(v interface{}), pop func() (interface{}, bool), empty func() bool) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				initInstance()
// 				for i := 0; i < test.count; i++ {
// 					push(getTestValue(i))
// 					push(getTestValue(i))
// 					tmp, tmp2 = pop()
// 				}
// 				for !empty() {
// 					tmp, tmp2 = pop()
// 				}
// 			}
// 		})
// 	}
// }

// // Slow decreasing tests--------------------------------------------------------------------------------------

// func BenchmarkSlowDecreaseList(b *testing.B) {
// 	var l *list.List
// 	benchmarkSlowDecrease(
// 		b,
// 		func() {
// 			l = list.New()
// 		},
// 		func(v interface{}) {
// 			l.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return l.Remove(l.Front()), true
// 		},
// 		func() bool {
// 			return l.Front() == nil
// 		},
// 	)
// }

// func BenchmarkSlowDecreaseSlice(b *testing.B) {
// 	var q *CustomSliceQueue
// 	benchmarkSlowDecrease(
// 		b,
// 		func() {
// 			q = NewCustomSliceQueue()
// 		},
// 		func(v interface{}) {
// 			q.Push(v.(*testValue))
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkSlowDecreasePhf(b *testing.B) {
// 	var q *phf.Queue
// 	benchmarkSlowDecrease(
// 		b,
// 		func() {
// 			q = phf.New()
// 		},
// 		func(v interface{}) {
// 			q.PushBack(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront(), true
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkSlowDecreaseCookiejar(b *testing.B) {
// 	var q *cookiejar.Queue
// 	benchmarkSlowDecrease(
// 		b,
// 		func() {
// 			q = cookiejar.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop(), true
// 		},
// 		func() bool {
// 			return q.Size() == 0
// 		},
// 	)
// }

// // func BenchmarkSlowDecreaseEgonelbre(b *testing.B) {
// // 	var q *egonelbre.Queue
// // 	benchmarkSlowDecrease(
// // 		b,
// // 		func() {
// // 			q = egonelbre.New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			return q.Pop()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func BenchmarkSlowDecreaseImpl7(b *testing.B) {
// 	var q *queueimpl7.Queueimpl7
// 	benchmarkSlowDecrease(
// 		b,
// 		func() {
// 			q = queueimpl7.New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.Pop()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkSlowDecreaseDequePopFront(b *testing.B) {
// 	var q *Deque
// 	benchmarkSlowDecrease(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopFront()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// func BenchmarkSlowDecreaseDequePopBack(b *testing.B) {
// 	var q *Deque
// 	benchmarkSlowDecrease(
// 		b,
// 		func() {
// 			q = New()
// 		},
// 		func(v interface{}) {
// 			q.Push(v)
// 		},
// 		func() (interface{}, bool) {
// 			return q.PopBack()
// 		},
// 		func() bool {
// 			return q.Len() == 0
// 		},
// 	)
// }

// // func BenchmarkSlowDecreaseDequePopMixed(b *testing.B) {
// // 	i := 0
// // 	var q *Deque
// // 	benchmarkSlowDecrease(
// // 		b,
// // 		func() {
// // 			q = New()
// // 		},
// // 		func(v interface{}) {
// // 			q.Push(v)
// // 		},
// // 		func() (interface{}, bool) {
// // 			i++
// // 			if i%2 == 0 {
// // 				return q.PopBack()
// // 			}
// // 			return q.PopFront()
// // 		},
// // 		func() bool {
// // 			return q.Len() == 0
// // 		},
// // 	)
// // }

// func benchmarkSlowDecrease(b *testing.B, initInstance func(), push func(v interface{}), pop func() (interface{}, bool), empty func() bool) {
// 	initInstance()
// 	for i := 0; i <= tests[len(tests)-1].count; i++ {
// 		push(getTestValue(i))
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for i := 0; i < test.count; i++ {
// 					push(getTestValue(i))
// 					tmp, tmp2 = pop()
// 					if !empty() {
// 						tmp, tmp2 = pop()
// 					}
// 				}
// 			}
// 		})
// 	}

// 	for !empty() {
// 		tmp, tmp2 = pop()
// 	}
// }

// Direct Deque tests--------------------------------------------------------------------------------------
// The above generic tests use method expressions to call the different queue/deque implementations.
// This adds some overhead in terns of performance and memory. For comparison purposes, however, using
// method expressions should be fine as all queues are subject to the same overhead.
// However, above tests should not be very representative of real world usage, so the performance
// and memory footprint is likely worse than seeing in real world applications. For this reason,
// all the same tests were replicated using the deque directly.
// Keeping below tests commmented out to not create any confusion with the other tests.
// If you are interested in checking the deque performance in a "more" real world setup,
// uncommented and run below tests.

// func BenchmarkFillDequePopFrontDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				q := New()
// 				for i := 0; i < test.count; i++ {
// 					q.Push(getTestValue(i))
// 				}
// 				for q.Len() > 0 {
// 					tmp, tmp2 = q.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkFillDequePopBackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				q := New()
// 				for i := 0; i < test.count; i++ {
// 					q.Push(getTestValue(i))
// 				}
// 				for q.Len() > 0 {
// 					tmp, tmp2 = q.PopBack()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkFillDequePopMixedDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				i := 0
// 				q := New()
// 				for i := 0; i < test.count; i++ {
// 					q.Push(getTestValue(i))
// 				}
// 				for q.Len() > 0 {
// 					i++
// 					if i%2 == 0 {
// 						tmp, tmp2 = q.PopBack()
// 					}
// 					tmp, tmp2 = q.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillDequePopBackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			q := New()
// 			for n := 0; n < b.N; n++ {
// 				for n := 0; n < refillCount; n++ {
// 					for i := 0; i < test.count; i++ {
// 						q.Push(getTestValue(i))
// 					}
// 					for q.Len() > 0 {
// 						tmp, tmp2 = q.PopFront()
// 					}
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillDequePopBackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			q := New()
// 			for n := 0; n < b.N; n++ {
// 				for n := 0; n < refillCount; n++ {
// 					for i := 0; i < test.count; i++ {
// 						q.Push(getTestValue(i))
// 					}
// 					for q.Len() > 0 {
// 						tmp, tmp2 = q.PopBack()
// 					}
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillDequePopMixedDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			i := 0
// 			q := New()
// 			for n := 0; n < b.N; n++ {
// 				for n := 0; n < refillCount; n++ {
// 					for i := 0; i < test.count; i++ {
// 						q.Push(getTestValue(i))
// 					}
// 					for q.Len() > 0 {
// 						i++
// 						if i%2 == 0 {
// 							tmp, tmp2 = q.PopBack()
// 						}
// 						tmp, tmp2 = q.PopFront()
// 					}
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkStableDequePopFrontDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for i := 0; i < test.count; i++ {
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopFront()
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkStableDequePopBackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for i := 0; i < test.count; i++ {
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopBack()
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopBack()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkStableDequePopMixedDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				i := 0
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for i := 0; i < test.count; i++ {
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopBack()
// 				}
// 				for d.Len() > 0 {
// 					i++
// 					if i%2 == 0 {
// 						tmp, tmp2 = d.PopBack()
// 					}
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillWhileFullDequePopFrontDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for k := 0; k < 1000; k++ {
// 					for i := 0; i < test.count; i++ {
// 						d.Push(getTestValue(i))
// 					}
// 					for i := 0; i < test.count; i++ {
// 						tmp, tmp2 = d.PopFront()
// 					}
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillWhileFullDequePopBackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for k := 0; k < 1000; k++ {
// 					for i := 0; i < test.count; i++ {
// 						d.Push(getTestValue(i))
// 					}
// 					for i := 0; i < test.count; i++ {
// 						tmp, tmp2 = d.PopBack()
// 					}
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopBack()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillWhileFullDequePopMixedDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				i := 0
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for k := 0; k < 1000; k++ {
// 					for i := 0; i < test.count; i++ {
// 						d.Push(getTestValue(i))
// 					}
// 					for i := 0; i < test.count; i++ {
// 						tmp, tmp2 = d.PopBack()
// 					}
// 				}
// 				for d.Len() > 0 {
// 					i++
// 					if i%2 == 0 {
// 						tmp, tmp2 = d.PopBack()
// 					}
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkSlowIncreaseDequePopFrontDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := New()
// 				for i := 0; i < test.count && d.Len() > 0; i++ {
// 					d.Push(getTestValue(i))
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopFront()
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkSlowIncreaseDequePopBackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := New()
// 				for i := 0; i < test.count && d.Len() > 0; i++ {
// 					d.Push(getTestValue(i))
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopBack()
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopBack()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkSlowIncreaseDequePopMixedDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				i := 0
// 				d := New()
// 				for i := 0; i < test.count && d.Len() > 0; i++ {
// 					d.Push(getTestValue(i))
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopBack()
// 				}
// 				for d.Len() > 0 {
// 					i++
// 					if i%2 == 0 {
// 						tmp, tmp2 = d.PopBack()
// 					}
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkSlowDecreaseDequePopFrontDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for i := 0; i < test.count && d.Len() > 0; i++ {
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopFront()
// 					tmp, tmp2 = d.PopFront()
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkSlowDecreaseDequePopBackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for i := 0; i < test.count && d.Len() > 0; i++ {
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopBack()
// 					tmp, tmp2 = d.PopBack()
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopBack()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkSlowDecreaseDequePopMixedDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				i := 0
// 				d := New()
// 				for i := 0; i < fillCount; i++ {
// 					d.Push(getTestValue(i))
// 				}
// 				for i := 0; i < test.count && d.Len() > 0; i++ {
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.PopBack()
// 					tmp, tmp2 = d.PopBack()
// 				}
// 				for d.Len() > 0 {
// 					i++
// 					if i%2 == 0 {
// 						tmp, tmp2 = d.PopBack()
// 					}
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkStableEgonelbre2(b *testing.B) {
// 	d := egonelbre.New()
// 	for i := 0; i < fillCount; i++ {
// 		d.Push(getTestValue(i))
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for i := 0; i < test.count; i++ {
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.Pop()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkStableImpl92(b *testing.B) {
// 	d := NewImpl9()
// 	for i := 0; i < fillCount; i++ {
// 		d.Push(getTestValue(i))
// 	}
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for i := 0; i < test.count; i++ {
// 					d.Push(getTestValue(i))
// 					tmp, tmp2 = d.Pop()
// 				}
// 			}
// 		})
// 	}
// }
