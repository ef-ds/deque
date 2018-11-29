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

// All benchmark tests use method expressions to call the different queue implementations,
// which adds some overhead in terns of performance and memory. For comparison purposes, however, using
// method expressions should be fine as all queues are subject to the same overhead.
// Below are all the benchmark tests replicated using deque directly, so below tests
// are not subject to any additional overhead. Keeping below tests commented out
// to avoid any confusion with the other tests.
// If you are interested in checking the deque performance without the overhead,
// uncommented and run below tests.

// func BenchmarkFillDequeQueueDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				q := new(deque.Deque)
// 				for i := 0; i < test.count; i++ {
// 					q.PushBack(getTestValue(i))
// 				}
// 				for q.Len() > 0 {
// 					tmp, tmp2 = q.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkFillDequeStackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				q := new(deque.Deque)
// 				for i := 0; i < test.count; i++ {
// 					q.PushBack(getTestValue(i))
// 				}
// 				for q.Len() > 0 {
// 					tmp, tmp2 = q.PopBack()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillDequeQueueDirect(b *testing.B) {
// 	for i, test := range tests {
// 		// Doesn't run the first (0 items) and last (1mi) items tests
// 		// as 0 items makes no sense for this test and 1mi is too slow.
// 		if i == 0 || i > 6 {
// 			continue
// 		}

// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			q := new(deque.Deque)
// 			for n := 0; n < b.N; n++ {
// 				for n := 0; n < refillCount; n++ {
// 					for i := 0; i < test.count; i++ {
// 						q.PushBack(getTestValue(i))
// 					}
// 					for q.Len() > 0 {
// 						tmp, tmp2 = q.PopFront()
// 					}
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillDequeStackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			q := new(deque.Deque)
// 			for n := 0; n < b.N; n++ {
// 				for n := 0; n < refillCount; n++ {
// 					for i := 0; i < test.count; i++ {
// 						q.PushBack(getTestValue(i))
// 					}
// 					for q.Len() > 0 {
// 						tmp, tmp2 = q.PopBack()
// 					}
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkRefillFullDequeQueueDirect(b *testing.B) {
// 	d := new(deque.Deque)
// 	for i := 0; i < fillCount; i++ {
// 		d.PushBack(getTestValue(i))
// 	}

// 	for i, test := range tests {
// 		// Doesn't run the first (0 items) and last (1mi) items tests
// 		// as 0 items makes no sense for this test and 1mi is too slow.
// 		if i == 0 || i > 6 {
// 			continue
// 		}

// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for k := 0; k < refillCount; k++ {
// 					for i := 0; i < test.count; i++ {
// 						d.PushBack(getTestValue(i))
// 					}
// 					for i := 0; i < test.count; i++ {
// 						tmp, tmp2 = d.PopFront()
// 					}
// 				}
// 			}
// 		})
// 	}

// 	for d.Len() > 0 {
// 		tmp, tmp2 = d.PopFront()
// 	}
// }

// func BenchmarkRefillFullDequeStackDirect(b *testing.B) {
// 	d := new(deque.Deque)
// 	for i := 0; i < fillCount; i++ {
// 		d.PushBack(getTestValue(i))
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for k := 0; k < refillCount; k++ {
// 					for i := 0; i < test.count; i++ {
// 						d.PushBack(getTestValue(i))
// 					}
// 					for i := 0; i < test.count; i++ {
// 						tmp, tmp2 = d.PopBack()
// 					}
// 				}
// 			}
// 		})
// 	}

// 	for d.Len() > 0 {
// 		tmp, tmp2 = d.PopBack()
// 	}
// }

// func BenchmarkStableDequeQueueDirect(b *testing.B) {
// 	d := new(deque.Deque)
// 	for i := 0; i < fillCount; i++ {
// 		d.PushBack(getTestValue(i))
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for i := 0; i < test.count; i++ {
// 					d.PushBack(getTestValue(i))
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}

// 	for d.Len() > 0 {
// 		tmp, tmp2 = d.PopFront()
// 	}
// }

// func BenchmarkStableDequeStackDirect(b *testing.B) {
// 	d := new(deque.Deque)
// 	for i := 0; i < fillCount; i++ {
// 		d.PushBack(getTestValue(i))
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for i := 0; i < test.count; i++ {
// 					d.PushBack(getTestValue(i))
// 					tmp, tmp2 = d.PopBack()
// 				}
// 			}
// 		})
// 	}

// 	for d.Len() > 0 {
// 		tmp, tmp2 = d.PopBack()
// 	}
// }

// func BenchmarkSlowIncreaseDequeQueueDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := new(deque.Deque)
// 				for i := 0; i < test.count; i++ {
// 					d.PushBack(getTestValue(i))
// 					d.PushBack(getTestValue(i))
// 					tmp, tmp2 = d.PopFront()
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopFront()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkSlowIncreaseDequeStackDirect(b *testing.B) {
// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				d := new(deque.Deque)
// 				for i := 0; i < test.count; i++ {
// 					d.PushBack(getTestValue(i))
// 					d.PushBack(getTestValue(i))
// 					tmp, tmp2 = d.PopBack()
// 				}
// 				for d.Len() > 0 {
// 					tmp, tmp2 = d.PopBack()
// 				}
// 			}
// 		})
// 	}
// }

// func BenchmarkSlowDecreaseDequeQueueDirect(b *testing.B) {
// 	d := new(deque.Deque)
// 	for _, test := range tests {
// 		items := test.count / 2
// 		for i := 0; i <= items; i++ {
// 			d.PushBack(getTestValue(i))
// 		}
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for i := 0; i < test.count; i++ {
// 					d.PushBack(getTestValue(i))
// 					tmp, tmp2 = d.PopFront()
// 					if d.Len() > 0 {
// 						tmp, tmp2 = d.PopFront()
// 					}
// 				}
// 			}
// 		})
// 	}

// 	for d.Len() > 0 {
// 		tmp, tmp2 = d.PopFront()
// 	}
// }

// func BenchmarkSlowDecreaseDequeStackDirect(b *testing.B) {
// 	d := new(deque.Deque)
// 	for _, test := range tests {
// 		items := test.count / 2
// 		for i := 0; i <= items; i++ {
// 			d.PushBack(getTestValue(i))
// 		}
// 	}

// 	for _, test := range tests {
// 		b.Run(strconv.Itoa(test.count), func(b *testing.B) {
// 			for n := 0; n < b.N; n++ {
// 				for i := 0; i < test.count; i++ {
// 					d.PushBack(getTestValue(i))
// 					tmp, tmp2 = d.PopFront()
// 					if d.Len() > 0 {
// 						tmp, tmp2 = d.PopBack()
// 					}
// 				}
// 			}
// 		})
// 	}

// 	for d.Len() > 0 {
// 		tmp, tmp2 = d.PopBack()
// 	}
// }
