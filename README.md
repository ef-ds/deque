# deque [![Build Status](https://travis-ci.com/ef-ds/deque.svg?branch=master)](https://travis-ci.com/ef-ds/deque) [![codecov](https://codecov.io/gh/ef-ds/deque/branch/master/graph/badge.svg)](https://codecov.io/gh/ef-ds/deque) [![Go Report Card](https://goreportcard.com/badge/github.com/ef-ds/deque)](https://goreportcard.com/report/github.com/ef-ds/deque)  [![GoDoc](https://godoc.org/github.com/ef-ds/deque?status.svg)](https://godoc.org/github.com/ef-ds/deque)

Package deque implements a dynamically growing, ring shaped, linked arrays double-ended queue. The deque is a general purpose deque data structure that performs under virtually all scenarios and is specifically optimized to perform when running in production-like traffic Microservice systems.

## Install
From a configured [Go environment](https://golang.org/doc/install#testing):
```sh
go get -u github.com/ef-ds/deque
```

If you are using dep:
```sh
dep ensure -add github.com/ef-ds/deque
```

## How to Use
```go
package main

import (
	"fmt"

	"github.com/ef-ds/deque"
)

func main() {
	d := deque.New()

	for i := 1; i <= 5; i++ {
		d.PushBack(i)
	}
	for v, ok := d.PopFront(); ok; v, ok = d.PopFront() {
		fmt.Println(v)
	}
}
```

Output:
```
1
2
3
4
5
```

Also refer to the tests at [deque_test.go](deque_test.go).


## Supported Data Types
Similarly to Go's standard library list, [list](https://github.com/golang/go/tree/master/src/container/list), 
[ring](https://github.com/golang/go/tree/master/src/container/ring) and [heap](https://github.com/golang/go/blob/master/src/container/heap/heap.go) packages, deque supports "interface{}" as its data type. This means it can be used with any Go data types, including int, float, string and any user defined structs and pointers to interfaces.

The data types pushed into the deque can even be mixed, meaning, it's possible to push ints, floats and struct instances into the same deque.


## Design
The efficient data structures (ef-ds) deque employs a new, mordern queue design: a ring shaped, auto shrinking, linked slices design.

That means the queue is a doubly-linked list where each node value is a fixed size slice. It is ring in shape because the linked list is a circular one, where the tail always points to the head in the ring. It also auto shrinks as the values
are popped off the deque, keeping itself as a lean and low memory footprint data structure even after heavy use.


### Design Considerations

Deque uses linked slices as its underlying data structure. The reason for the choice comes from two main observations of slice based deques/queues/stacks:

1. When the deque/queue/stack needs to expand to accommodate new values, [a new, larger slice needs to be allocated](https://en.wikipedia.org/wiki/Dynamic_array#Geometric_expansion_and_amortized_cost) and used
2. Allocating and managing large slices is expensive, especially in an overloaded system with little available physical memory

To help clarify the scenario, below is what happens when a slice based queue that already holds, say 1bi items, needs to expand to accommodate a new item

Slice based implementation

- Allocate a new, twice the size of the previous allocated one, say 2 billion positions slice
- Copy over all 1 billion items from the previous slice into the new one
- Add the new value into the first unused position in the new slice, position 1000000001.
- The same scenario for impl7 plays out like below.

Deque

- Allocate a new 256 size slice
- Set the next pointer
- Add the value into the first position of the new slice, position 0

Deque never copies data around, but slice based ones do, and if the data set is large, it doesn't matter how fast the copying algorithm is. The copying has to be done and will take some time.

The decision to use linked slices was also the result of the observation that slices goes to great length to provide predictive, indexed positions. A hash table, for instance, absolutely need this property, but not a queue. So impl7 completely gives up this property and focus on what really matters: add to end, retrieve from head. No copying around and repositioning of elements is needed for that. So when a slice goes to great length to provide that functionality, the whole work of allocating new arrays, copying data around is all wasted work. None of that is necessary. And this work costs dearly for large data sets as observed in the tests.

While linked slices design solve the slice expansion problem very effectively, it doesn't help with many real world
usage scenarios such as in a stable processing environment where small amount of items are pushed and popped from
the deque in a sequential way. This is a very common scenario for TCP/HTTP services, for instance, where the service
is able to handle the current traffic without stress (i.e., on each request, an item is added to the queue; when the request is fulfilled, the item is removed from the queue).

To address the stable scenario in an effective way, deque keeps its internal linked arrays in a circular, ring shape.
This way when items are pushed to the deque after some of them have been removed, the deque will automatically move
over its tail slice back to the old head of the deque, effectively reusing the same already allocated slice. The result
is a deque that will run through its ring reusing the ring to store the new values, instead of allocating new
slices for the new values.

Another important real world scenario is when a service is subject to high and low traffic patterns over time.
It is desired the service to be able to scale out quickly and effectively to handle the extra traffic when needed,
but also to scale in, decreasing resource usage, such as CPU and memory, when the traffic subsides, so
the cost to run the system over time is optimized.

To solve this problem, deque employs an automatically shrinking mechanism that will shrink, releasing the extra
resources from memory, as the number of items in the deque decreases. This mechanism allows
the queue to shrink when needed, keeping itself lean, but also solves a major problem of most ring based
implementations: once inflated, the data structure either never shrinks or require manual "shrink" calls,
which can be tricky to use as the ideal, most optimized moment to shrink is not always clear and well defined. However, a data structure that completely shrinks after use, when it is used again, it means it has to expand again to accomodate the new values, hindering performance on refill scenarios (where a number of items is added and removed from the queue successively). To address this scenario, deque keeps a configurable number of internal, empty, slices in its ring. This way in refill, scenarios, deque is able to scale out very quickly, still managing to keep the memory footprint very low.


## Tests

Besides having 100% code coverage, deque has a extensive set of unit, integration and API tests covering all happy, sad and edge cases. The tests can be found [here](deque_test.go).

Performance and efficiency is a major concern, so deque has a extensive set of benchmark tests as well comparing the deque performance
with a variety of high quality open source deque implementations. See the [benchmark tests](BENCHMARK_TESTS.md) for details.


## Performance

Deque has constant time (O(1)) on all its operations (PushFront/PushBack/PopFront/PopBack/Len). It's not amortized constant as it is with most [slice based deques](https://en.wikipedia.org/wiki/Dynamic_array#Geometric_expansion_and_amortized_cost) because it never copies more than 8 (maxFirstSliceSize/2) items and when it expands or grow, it never does so by more than 256 (maxInternalArraySize) items in a single operation.

Deque, either used as a FIFO queue or LIFO stack, offers either the best or very competetive performance across all test sets, batteries and ranges.

As a general purpose FIFO deque or LIFO stack, deque offers, by far, the most balanced and consistent performance of all tested data structures.

See [performance](PERFORMANCE.md) for details.



## Support

As the [CloudLogger](https://github.com/cloud-logger/docs) project needed a high performance unbounded queue and, given the fact that
Go doesn't provide such queue in its standard library, I built a new queue and proposed it to be added to the standard library.

The initial [proposal](https://github.com/golang/go/issues/27935) was to add [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) to the standard library.

Given the suggestions to build a deque instead of a FIFO queue as a deque is a muh more flexible data structure, coupled with suggestions to build a proper external package,
this deque package was built.

I truly believe in the deque and I believe it should have a place in the Go's standard library, so all Go users can benefit from this
data structure, and not only the Go insider ones or the lucky ones who found out about the deque by chance or were lucky enough to find it through the seach engines.

If you like deque, please help us support it by thumbing up the [proposal](https://github.com/golang/go/issues/27935) and leaving comments.

We need to gather real world usage data to help validate the queue and make it into the standard library, so if you are using it, please make sure to share you experience with us.

I'm also extremely interested in improving the deque. Please let me know your suggestions for possible improvements.



## Competition

Do you know a better, faster deque implementation? Please let me know and I'm very glad to analyse it.



## Releases
We're committed to a CI/CD lifecycle releasing frequent, but only stable, production ready versions with all proper tests in place.

We strive as much as possible to keep backwards compatibility with previous versions, so breaking changes are a no-go.

For a list of changes in each released version, see [CHANGELOG.md](CHANGELOG.md).


## Supported Go Versions
See [supported_go_versions.md](https://github.com/ef-ds/docs/blob/master/supported_go_versions.md).


## License
MIT, see [LICENSE](LICENSE).

"Use, abuse, have fun and contribute back!"


## Contributions
See [contributing.md](https://github.com/ef-ds/docs/blob/master/contributing.md).

