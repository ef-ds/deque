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
//
// Below tests are used mostly for comparing the result of changes to deque and
// are not necessarily a replication of the comparison tests. For instance,
// the stack tests here use PushFront/PopFront instead of PushBack/PopBack.
//
// For comparing deque performance with other deques, refer to
// https://github.com/ef-ds/deque-bench-tests

package deque_test

import (
	"strconv"
	"testing"

	"github.com/ef-ds/deque"
)

// testData contains the number of items to add to the queues in each test.
type testData struct {
	count int
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

func BenchmarkFillQueue(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := deque.New()
				for i := 0; i < test.count; i++ {
					q.PushBack(nil)
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.PopFront()
				}
			}
		})
	}
}

func BenchmarkFillStack(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				q := deque.New()
				for i := 0; i < test.count; i++ {
					q.PushFront(nil)
				}
				for q.Len() > 0 {
					tmp, tmp2 = q.PopFront()
				}
			}
		})
	}
}

func BenchmarkRefillQueue(b *testing.B) {
	for i, test := range tests {
		// Doesn't run the first (0 items) and last (1mi) items tests
		// as 0 items makes no sense for this test and 1mi is too slow.
		if i == 0 || i > 6 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			q := deque.New()
			for n := 0; n < b.N; n++ {
				for n := 0; n < refillCount; n++ {
					for i := 0; i < test.count; i++ {
						q.PushBack(nil)
					}
					for q.Len() > 0 {
						tmp, tmp2 = q.PopFront()
					}
				}
			}
		})
	}
}

func BenchmarkRefillStack(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			q := deque.New()
			for n := 0; n < b.N; n++ {
				for n := 0; n < refillCount; n++ {
					for i := 0; i < test.count; i++ {
						q.PushFront(nil)
					}
					for q.Len() > 0 {
						tmp, tmp2 = q.PopFront()
					}
				}
			}
		})
	}
}

func BenchmarkRefillFullQueue(b *testing.B) {
	d := deque.New()
	for i := 0; i < fillCount; i++ {
		d.PushBack(nil)
	}

	for i, test := range tests {
		// Doesn't run the first (0 items) and last two (100k, 1mi) items tests
		// as 0 items makes no sense for this test and 1mi is too slow.
		if i == 0 || i > 5 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for k := 0; k < refillCount; k++ {
					for i := 0; i < test.count; i++ {
						d.PushBack(nil)
					}
					for i := 0; i < test.count; i++ {
						tmp, tmp2 = d.PopFront()
					}
				}
			}
		})
	}

	for d.Len() > 0 {
		tmp, tmp2 = d.PopFront()
	}
}

func BenchmarkRefillFullStack(b *testing.B) {
	d := deque.New()
	for i := 0; i < fillCount; i++ {
		d.PushBack(nil)
	}

	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for k := 0; k < refillCount; k++ {
					for i := 0; i < test.count; i++ {
						d.PushFront(nil)
					}
					for i := 0; i < test.count; i++ {
						tmp, tmp2 = d.PopFront()
					}
				}
			}
		})
	}

	for d.Len() > 0 {
		tmp, tmp2 = d.PopBack()
	}
}

func BenchmarkStableQueue(b *testing.B) {
	d := deque.New()
	for i := 0; i < fillCount; i++ {
		d.PushBack(nil)
	}

	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < test.count; i++ {
					d.PushBack(nil)
					tmp, tmp2 = d.PopFront()
				}
			}
		})
	}

	for d.Len() > 0 {
		tmp, tmp2 = d.PopFront()
	}
}

func BenchmarkStableStack(b *testing.B) {
	d := deque.New()
	for i := 0; i < fillCount; i++ {
		d.PushBack(nil)
	}

	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < test.count; i++ {
					d.PushFront(nil)
					tmp, tmp2 = d.PopFront()
				}
			}
		})
	}

	for d.Len() > 0 {
		tmp, tmp2 = d.PopFront()
	}
}

func BenchmarkSlowIncreaseQueue(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				d := deque.New()
				for i := 0; i < test.count; i++ {
					d.PushBack(nil)
					d.PushBack(nil)
					tmp, tmp2 = d.PopFront()
				}
				for d.Len() > 0 {
					tmp, tmp2 = d.PopFront()
				}
			}
		})
	}
}

func BenchmarkSlowIncreaseStack(b *testing.B) {
	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				d := deque.New()
				for i := 0; i < test.count; i++ {
					d.PushFront(nil)
					d.PushFront(nil)
					tmp, tmp2 = d.PopFront()
				}
				for d.Len() > 0 {
					tmp, tmp2 = d.PopFront()
				}
			}
		})
	}
}

func BenchmarkSlowDecreaseQueue(b *testing.B) {
	d := deque.New()
	for _, test := range tests {
		items := test.count / 2
		for i := 0; i <= items; i++ {
			d.PushBack(nil)
		}
	}

	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < test.count; i++ {
					d.PushBack(nil)
					tmp, tmp2 = d.PopFront()
					if d.Len() > 0 {
						tmp, tmp2 = d.PopFront()
					}
				}
			}
		})
	}

	for d.Len() > 0 {
		tmp, tmp2 = d.PopFront()
	}
}

func BenchmarkSlowDecreaseStack(b *testing.B) {
	d := deque.New()
	for _, test := range tests {
		items := test.count / 2
		for i := 0; i <= items; i++ {
			d.PushFront(nil)
		}
	}

	for _, test := range tests {
		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < test.count; i++ {
					d.PushFront(nil)
					tmp, tmp2 = d.PopFront()
					if d.Len() > 0 {
						tmp, tmp2 = d.PopFront()
					}
				}
			}
		})
	}

	for d.Len() > 0 {
		tmp, tmp2 = d.PopFront()
	}
}

func BenchmarkMicroserviceQueue(b *testing.B) {
	for i, test := range tests {
		// Doesn't run the first (0 items) as 0 items makes no sense for this test.
		if i == 0 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				d := deque.New()

				// Simulate stable traffic
				for i := 0; i < test.count; i++ {
					d.PushBack(nil)
					d.PopFront()
				}

				// Simulate slowly increasing traffic
				for i := 0; i < test.count; i++ {
					d.PushBack(nil)
					d.PushBack(nil)
					d.PopFront()
				}

				// Simulate slowly decreasing traffic, bringing traffic back to normal
				for i := 0; i < test.count; i++ {
					d.PopFront()
					if d.Len() > 0 {
						d.PopFront()
					}
					d.PushBack(nil)
				}

				// Simulate quick traffic spike (DDOS attack, etc)
				for i := 0; i < test.count; i++ {
					d.PushBack(nil)
				}

				// Simulate stable traffic while at high traffic
				for i := 0; i < test.count; i++ {
					d.PushBack(nil)
					d.PopFront()
				}

				// Simulate going back to normal (DDOS attack fended off)
				for i := 0; i < test.count; i++ {
					d.PopFront()
				}

				// Simulate stable traffic (now that is back to normal)
				for i := 0; i < test.count; i++ {
					d.PushBack(nil)
					d.PopFront()
				}
			}
		})
	}
}

func BenchmarkMicroserviceStack(b *testing.B) {
	for i, test := range tests {
		// Doesn't run the first (0 items) as 0 items makes no sense for this test.
		if i == 0 {
			continue
		}

		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				d := deque.New()

				// Simulate stable traffic
				for i := 0; i < test.count; i++ {
					d.PushFront(nil)
					d.PopFront()
				}

				// Simulate slowly increasing traffic
				for i := 0; i < test.count; i++ {
					d.PushFront(nil)
					d.PushFront(nil)
					d.PopFront()
				}

				// Simulate slowly decreasing traffic, bringing traffic Front to normal
				for i := 0; i < test.count; i++ {
					d.PopFront()
					if d.Len() > 0 {
						d.PopFront()
					}
					d.PushFront(nil)
				}

				// Simulate quick traffic spike (DDOS attack, etc)
				for i := 0; i < test.count; i++ {
					d.PushFront(nil)
				}

				// Simulate stable traffic while at high traffic
				for i := 0; i < test.count; i++ {
					d.PushFront(nil)
					d.PopFront()
				}

				// Simulate going Front to normal (DDOS attack fended off)
				for i := 0; i < test.count; i++ {
					d.PopFront()
				}

				// Simulate stable traffic (now that is Front to normal)
				for i := 0; i < test.count; i++ {
					d.PushFront(nil)
					d.PopFront()
				}
			}
		})
	}
}
