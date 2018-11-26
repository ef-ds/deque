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

// testData contains the number of items to add to the queues in each test.
type testData struct {
	count int
}

// testValue is used as the value added in each push call to the queues.
// A struct is being used as structs should be more representative of real
// world uses of a queue. A second f2 field was added as the users structs
// are likely to contain more than one field.
type testValue struct {
	count int
	f2    int
}

var (
	tests = []testData{
		{count: 0},
		{count: 1},
		{count: 10},
		{count: 100},
		{count: 1000},    // 1k
		{count: 10000},   //10k
		{count: 100000},  // 100k
		{count: 1000000}, // 1mi
	}

	// Used to store temp values, avoiding any compiler optimizations.
	tmp  interface{}
	tmp2 bool

	fillCount   = 10000
	refillCount = 100
)

// Pure slice based test queue implementation-------------------------------------------------------

// CustomSliceQueue represents an unbounded, dynamically growing deque customized
// to operate on testVale struct.
type CustomSliceQueue struct {
	// The queue values.
	v []*testValue
}

func NewCustomSliceQueue() *CustomSliceQueue {
	return new(CustomSliceQueue).Init()
}

func (q *CustomSliceQueue) Init() *CustomSliceQueue {
	q.v = make([]*testValue, 0)
	return q
}

func (q *CustomSliceQueue) Len() int { return len(q.v) }

func (q *CustomSliceQueue) Front() (*testValue, bool) {
	if len(q.v) == 0 {
		return nil, false
	}
	return q.v[0], true
}

func (q *CustomSliceQueue) Back() (*testValue, bool) {
	if len(q.v) == 0 {
		return nil, false
	}
	return q.v[len(q.v)-1], true
}

func (q *CustomSliceQueue) PushFront(v *testValue) {
	q.v = append(q.v, v)
	copy(q.v[1:], q.v[0:])
	q.v[0] = v
}

func (q *CustomSliceQueue) PushBack(v *testValue) {
	q.v = append(q.v, v)
}

func (q *CustomSliceQueue) PopFront() (*testValue, bool) {
	if len(q.v) == 0 {
		return nil, false
	}

	v := q.v[0]
	q.v[0] = nil // Avoid memory leaks
	q.v = q.v[1:]
	return v, true
}

func (q *CustomSliceQueue) PopBack() (*testValue, bool) {
	if len(q.v) == 0 {
		return nil, false
	}

	tp := len(q.v) - 1
	v := q.v[tp]
	q.v[tp] = nil // Avoid memory leaks
	q.v = q.v[:tp]
	return v, true
}

// Helper methods-----------------------------------------------------------------------------------

func getTestValue(i int) *testValue {
	return &testValue{
		count: i,
		f2:    1, // Initializes f2 to some random value (1).
	}
}
