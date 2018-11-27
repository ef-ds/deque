# Performance

Below compares the deque [benchmark tests](BENCHMARK_TESTS.md) results with the other tested queues.

## Running the Tests
In the "testdata" directory, we have included the result of local test runs for all queues. Below uses this run to compare the queues, but it's possible and we higly encourage you to run the tests youself to help validate the results.

To run the tests locally, close the deque repo, cd to the deque main directory and run below command.

```
go test -benchmem -timeout 60m -bench=. -run=^$
```

This command will run all tests for all queues locally once. This should be good enouh to give you a sense of the queues performance, but to
do a proper comparison, elimating test variations, we recommend you to run the tests as detailed [here](BENCHMARK_TESTS.md) by running the tests with multiple counts, splitting the files with [test-splitter](https://github.com/ef-ds/tools/tree/master/testsplitter) and using the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool to aggeegate the results.


## Bottom Line
As a general purpose double-ended queue, deque is the queue that displays the most balanced performance, performing either very competitively or besting all other queues in all the different test scenarios.

Having said that, the tests show there's some room for improvements. We're actively working on the deque. We expect to release better performant deque versions in the near future.

## Recommendations
Using a double-ended queue as a stack is possible and works very well. However, given the stack inverted properties (LIFO) when comparing to a FIFO queue, using a deque as a stack is not the most efficient solution.

For a stack solution, we recommend building a stack using a simple slice, such as the [CustomSliceQueue](testdata_test.go).

For all other uses, we recommend using the deque as it performs very well on both low and high load scenarios on all tests.


## Results

Given the enormous amount of test data, I'll focus the analysis here on the results
of the, arguably, the most important test: the Microservice test. All the other test results are posted here without comments.

### Microservice Test Results

deque vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4           525ns ± 4%     646ns ±10%  +23.18%  (p=0.000 n=10+9)
/10-4         3.66µs ± 5%    4.81µs ± 5%  +31.38%  (p=0.000 n=10+10)
/100-4        25.9µs ± 2%    32.3µs ± 3%  +24.42%  (p=0.000 n=10+10)
/1000-4        241µs ± 3%     313µs ± 5%  +30.03%  (p=0.000 n=10+10)
/10000-4      2.44ms ± 2%    3.17ms ±11%  +30.12%  (p=0.000 n=10+10)
/100000-4     27.1ms ± 3%    33.1ms ± 2%  +22.41%  (p=0.000 n=10+8)
/1000000-4     282ms ± 6%     348ms ± 8%  +23.13%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            560B ± 0%      432B ± 0%  -22.86%  (p=0.000 n=10+10)
/10-4         5.71kB ± 0%    6.91kB ± 0%  +21.01%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    29.6kB ± 0%  +41.70%  (p=0.000 n=10+10)
/1000-4        138kB ± 0%     261kB ± 0%  +88.73%  (p=0.000 n=10+10)
/10000-4      1.54MB ± 0%    2.58MB ± 0%  +67.91%  (p=0.000 n=10+10)
/100000-4     15.3MB ± 0%    25.8MB ± 0%  +68.77%  (p=0.000 n=9+9)
/1000000-4     152MB ± 0%     258MB ± 0%  +68.88%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      17.0 ± 0%  +41.67%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%     109.0 ± 0%  +41.56%  (p=0.000 n=10+10)
/100-4           709 ± 0%       927 ± 0%  +30.75%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%     9.11k ± 0%  +29.84%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     91.0k ± 0%  +29.56%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      909k ± 0%  +29.55%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     9.09M ± 0%  +29.55%  (p=0.000 n=10+10)
```

Deque was based on impl7 and due to its optimizations, it performs
considerably better than impl7 on all test loads. The biggest gain is,
however, in the memory footprint.

deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceListQueue.txt
name        old time/op    new time/op    delta
/1-4           525ns ± 4%     511ns ± 6%      ~     (p=0.082 n=10+9)
/10-4         3.66µs ± 5%    4.81µs ± 2%   +31.23%  (p=0.000 n=10+10)
/100-4        25.9µs ± 2%    46.8µs ± 1%   +80.26%  (p=0.000 n=10+8)
/1000-4        241µs ± 3%     500µs ± 5%  +107.74%  (p=0.000 n=10+8)
/10000-4      2.44ms ± 2%    5.39ms ± 4%  +121.14%  (p=0.000 n=10+9)
/100000-4     27.1ms ± 3%    75.0ms ± 8%  +177.19%  (p=0.000 n=10+10)
/1000000-4     282ms ± 6%     815ms ±13%  +188.65%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            560B ± 0%      496B ± 0%   -11.43%  (p=0.000 n=10+10)
/10-4         5.71kB ± 0%    4.53kB ± 0%   -20.73%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    44.8kB ± 0%  +114.46%  (p=0.000 n=10+10)
/1000-4        138kB ± 0%     448kB ± 0%  +224.00%  (p=0.000 n=10+10)
/10000-4      1.54MB ± 0%    4.48MB ± 0%  +191.70%  (p=0.000 n=10+10)
/100000-4     15.3MB ± 0%    44.8MB ± 0%  +193.61%  (p=0.000 n=9+9)
/1000000-4     152MB ± 0%     448MB ± 0%  +193.81%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      15.0 ± 0%   +25.00%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%     141.0 ± 0%   +83.12%  (p=0.000 n=10+10)
/100-4           709 ± 0%      1401 ± 0%   +97.60%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%    14.00k ± 0%   +99.53%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%    140.0k ± 0%   +99.42%  (p=0.000 n=10+10)
/100000-4       702k ± 0%     1400k ± 0%   +99.44%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%    14.00M ± 0%   +99.44%  (p=0.000 n=10+10)
```

The standlist list package can be used as a deque. However, Deque is much faster and displays a much lower memory footprint.


deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceListStack.txt
name        old time/op    new time/op    delta
/1-4           419ns ± 8%     510ns ± 4%   +21.89%  (p=0.000 n=9+9)
/10-4         2.61µs ± 6%    4.87µs ± 6%   +86.97%  (p=0.000 n=10+9)
/100-4        25.2µs ± 5%    47.1µs ± 2%   +86.49%  (p=0.000 n=10+8)
/1000-4        234µs ± 6%     488µs ±14%  +108.95%  (p=0.000 n=10+10)
/10000-4      2.36ms ± 7%    5.28ms ± 7%  +123.49%  (p=0.000 n=10+9)
/100000-4     24.2ms ± 2%    75.2ms ± 3%  +210.64%  (p=0.000 n=9+9)
/1000000-4     249ms ± 2%     941ms ±12%  +278.47%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            304B ± 0%      496B ± 0%   +63.16%  (p=0.000 n=10+10)
/10-4         1.57kB ± 0%    4.53kB ± 0%  +188.78%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    44.8kB ± 0%  +114.46%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     448kB ± 0%  +244.65%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    4.48MB ± 0%  +248.05%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    44.8MB ± 0%  +249.41%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     448MB ± 0%  +249.48%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      15.0 ± 0%   +36.36%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     141.0 ± 0%   +88.00%  (p=0.000 n=10+10)
/100-4           709 ± 0%      1401 ± 0%   +97.60%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.64%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%    140.0k ± 0%   +99.76%  (p=0.000 n=10+10)
/100000-4       701k ± 0%     1400k ± 0%   +99.78%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%    14.00M ± 0%   +99.78%  (p=0.000 n=10+10)
```

The standlist list package when used as a stack, displays similar levels of performance when compared to using it as a FIFO queue. However, Deque is much faster and displays a much lower memory footprint when used as a stack as well.


deque vs [CustomSliceQueue](testdata_test.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceSliceQueue.txt
name        old time/op    new time/op    delta
/1-4           525ns ± 4%     475ns ± 2%    -9.45%  (p=0.000 n=10+9)
/10-4         3.66µs ± 5%    3.54µs ± 2%    -3.36%  (p=0.003 n=10+10)
/100-4        25.9µs ± 2%    27.0µs ± 5%    +3.90%  (p=0.000 n=10+10)
/1000-4        241µs ± 3%     278µs ± 2%   +15.46%  (p=0.000 n=10+10)
/10000-4      2.44ms ± 2%    3.08ms ± 2%   +26.28%  (p=0.000 n=10+10)
/100000-4     27.1ms ± 3%    47.4ms ± 3%   +75.04%  (p=0.000 n=10+10)
/1000000-4     282ms ± 6%     596ms ± 7%  +111.04%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            560B ± 0%      232B ± 0%   -58.57%  (p=0.000 n=10+10)
/10-4         5.71kB ± 0%    2.17kB ± 0%   -62.04%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    21.3kB ± 0%    +2.03%  (p=0.000 n=10+10)
/1000-4        138kB ± 0%     214kB ± 0%   +55.09%  (p=0.000 n=10+10)
/10000-4      1.54MB ± 0%    2.95MB ± 0%   +91.93%  (p=0.000 n=10+10)
/100000-4     15.3MB ± 0%    33.1MB ± 0%  +117.19%  (p=0.000 n=9+10)
/1000000-4     152MB ± 0%     338MB ± 0%  +121.58%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      14.0 ± 0%   +16.67%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%     101.0 ± 0%   +31.17%  (p=0.000 n=10+10)
/100-4           709 ± 0%       822 ± 0%   +15.94%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%     8.77k ± 0%   +24.93%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     87.8k ± 0%   +25.11%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      875k ± 0%   +24.72%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     8.86M ± 0%   +26.25%  (p=0.000 n=10+10)
```

Using a simple slice as a FIFO queue, performs marginally better than deque with very small data sets, but fall behind quite dramatically with large data sets, performance and memory wise. This happens because [CustomSliceQueue](testdata_test.go) is a naive approach to building queues, one that doesn't consider that the slice is ever expanding, consuming considerably more memory as the number of items is pushed to the slice. Building a non-naive queue using simple slices requires much more effort.


deque vs [CustomSliceQueue](testdata_test.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceSliceStack.txt
name        old time/op    new time/op    delta
/1-4           419ns ± 8%     366ns ± 2%  -12.69%  (p=0.000 n=9+8)
/10-4         2.61µs ± 6%    2.15µs ± 2%  -17.41%  (p=0.000 n=10+9)
/100-4        25.2µs ± 5%    18.4µs ± 1%  -27.17%  (p=0.000 n=10+10)
/1000-4        234µs ± 6%     177µs ± 1%  -24.46%  (p=0.000 n=10+10)
/10000-4      2.36ms ± 7%    1.88ms ± 2%  -20.17%  (p=0.000 n=10+10)
/100000-4     24.2ms ± 2%    25.7ms ± 4%   +6.23%  (p=0.000 n=9+10)
/1000000-4     249ms ± 2%     283ms ± 5%  +13.78%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            304B ± 0%      200B ± 0%  -34.21%  (p=0.000 n=10+10)
/10-4         1.57kB ± 0%    1.40kB ± 0%  -10.71%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    13.3kB ± 0%  -36.53%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     128kB ± 0%   -1.22%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    1.51MB ± 0%  +17.03%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    15.9MB ± 0%  +23.65%  (p=0.000 n=10+9)
/1000000-4     128MB ± 0%     157MB ± 0%  +22.62%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      11.0 ± 0%     ~     (all equal)
/10-4           75.0 ± 0%      76.0 ± 0%   +1.33%  (p=0.000 n=10+10)
/100-4           709 ± 0%       709 ± 0%     ~     (all equal)
/1000-4        7.01k ± 0%     7.01k ± 0%   -0.01%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%     70.0k ± 0%   -0.09%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      700k ± 0%   -0.11%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     7.00M ± 0%   -0.11%  (p=0.000 n=10+10)
```

This is the only scenario where deque doesn't perform particularly well against
a specific data structure. A simple slice based stack manages to perform
considerably better than deque for data sets up to 10.000 items. Past that,
deque performs better. This is due to the fact that using a simple slice as a stack is a much better choice when compared to a queue due to the stack inverted properties (remove from tail instead of head).



deque vs [gammazero](https://github.com/gammazero/deque) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4           525ns ± 4%     379ns ± 8%  -27.72%  (p=0.000 n=10+10)
/10-4         3.66µs ± 5%    2.57µs ±11%  -29.80%  (p=0.000 n=10+9)
/100-4        25.9µs ± 2%    27.0µs ± 3%   +4.00%  (p=0.000 n=10+10)
/1000-4        241µs ± 3%     251µs ± 2%   +4.26%  (p=0.000 n=10+9)
/10000-4      2.44ms ± 2%    2.75ms ±18%  +12.75%  (p=0.000 n=10+10)
/100000-4     27.1ms ± 3%    33.3ms ±10%  +22.87%  (p=0.000 n=10+9)
/1000000-4     282ms ± 6%     341ms ± 5%  +20.67%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            560B ± 0%      416B ± 0%  -25.71%  (p=0.000 n=10+10)
/10-4         5.71kB ± 0%    1.42kB ± 0%  -75.07%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    22.3kB ± 0%   +6.43%  (p=0.000 n=10+10)
/1000-4        138kB ± 0%     209kB ± 0%  +51.19%  (p=0.000 n=10+10)
/10000-4      1.54MB ± 0%    2.69MB ± 0%  +75.26%  (p=0.000 n=10+10)
/100000-4     15.3MB ± 0%    23.8MB ± 0%  +55.86%  (p=0.000 n=9+9)
/1000000-4     152MB ± 0%     213MB ± 0%  +39.47%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%       9.0 ± 0%  -25.00%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%      72.0 ± 0%   -6.49%  (p=0.000 n=10+10)
/100-4           709 ± 0%       714 ± 0%   +0.71%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%     7.03k ± 0%   +0.13%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.23%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.27%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.28%  (p=0.000 n=10+10)
```

The gammazero deque performs really well with very small data sets (<= 10), beating deque by ~30%. However, it is considerably slower with larger data sets and uses considerably more memory than deque.


deque vs gammazero - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4           419ns ± 8%     375ns ± 4%   -10.56%  (p=0.000 n=9+9)
/10-4         2.61µs ± 6%    2.56µs ± 5%      ~     (p=0.247 n=10+10)
/100-4        25.2µs ± 5%    27.6µs ± 7%    +9.28%  (p=0.000 n=10+10)
/1000-4        234µs ± 6%     270µs ± 4%   +15.47%  (p=0.000 n=10+9)
/10000-4      2.36ms ± 7%    2.67ms ± 2%   +13.07%  (p=0.000 n=10+9)
/100000-4     24.2ms ± 2%    33.5ms ±19%   +38.31%  (p=0.000 n=9+10)
/1000000-4     249ms ± 2%     342ms ± 6%   +37.49%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            304B ± 0%      416B ± 0%   +36.84%  (p=0.000 n=10+10)
/10-4         1.57kB ± 0%    1.42kB ± 0%    -9.18%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    22.3kB ± 0%    +6.43%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     209kB ± 0%   +60.82%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    2.69MB ± 0%  +109.11%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    23.8MB ± 0%   +85.48%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     213MB ± 0%   +65.89%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%       9.0 ± 0%   -18.18%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      72.0 ± 0%    -4.00%  (p=0.000 n=10+10)
/100-4           709 ± 0%       714 ± 0%    +0.71%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%    +0.19%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%     70.0k ± 0%    -0.06%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      700k ± 0%    -0.10%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     7.00M ± 0%    -0.11%  (p=0.000 n=10+10)
```

The gammazero deque, when used as a stack, doesn't seem to hold on to its better
performance, when used as a queue, with very small data sets.
The biggest difference, again here, is in the deque's considerably lower memory footprint.


deque vs phf - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroservicePhfQueue.txt
name        old time/op    new time/op    delta
/1-4           525ns ± 4%     444ns ± 2%  -15.32%  (p=0.000 n=10+10)
/10-4         3.66µs ± 5%    3.39µs ± 7%   -7.57%  (p=0.000 n=10+10)
/100-4        25.9µs ± 2%    29.5µs ± 6%  +13.77%  (p=0.000 n=10+10)
/1000-4        241µs ± 3%     260µs ± 2%   +8.05%  (p=0.000 n=10+9)
/10000-4      2.44ms ± 2%    2.79ms ± 3%  +14.37%  (p=0.000 n=10+10)
/100000-4     27.1ms ± 3%    33.6ms ± 4%  +24.26%  (p=0.000 n=10+10)
/1000000-4     282ms ± 6%     361ms ± 9%  +27.75%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            560B ± 0%      272B ± 0%  -51.43%  (p=0.000 n=10+10)
/10-4         5.71kB ± 0%    2.18kB ± 0%  -61.90%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    23.0kB ± 0%  +10.02%  (p=0.000 n=10+10)
/1000-4        138kB ± 0%     210kB ± 0%  +51.73%  (p=0.000 n=10+10)
/10000-4      1.54MB ± 0%    2.69MB ± 0%  +75.31%  (p=0.000 n=10+10)
/100000-4     15.3MB ± 0%    23.8MB ± 0%  +55.86%  (p=0.000 n=9+9)
/1000000-4     152MB ± 0%     213MB ± 0%  +39.47%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      11.0 ± 0%   -8.33%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%      79.0 ± 0%   +2.60%  (p=0.000 n=10+10)
/100-4           709 ± 0%       721 ± 0%   +1.69%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%     7.03k ± 0%   +0.23%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.22%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.27%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.28%  (p=0.000 n=10+10)
```

phf deque seems to perform very similarly to gammazero's deque. This is not a surprise as both deques use a similar design: ring around a slice.
The biggest difference, again, is in the deque's considerably lower memory footprint.


deque vs phf - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroservicePhfStack.txt
name        old time/op    new time/op    delta
/1-4           419ns ± 8%     458ns ± 9%    +9.44%  (p=0.001 n=9+10)
/10-4         2.61µs ± 6%    3.22µs ± 2%   +23.56%  (p=0.000 n=10+9)
/100-4        25.2µs ± 5%    29.0µs ± 4%   +14.94%  (p=0.000 n=10+10)
/1000-4        234µs ± 6%     265µs ± 5%   +13.20%  (p=0.000 n=10+10)
/10000-4      2.36ms ± 7%    2.90ms ± 3%   +22.88%  (p=0.000 n=10+9)
/100000-4     24.2ms ± 2%    36.0ms ± 3%   +48.56%  (p=0.000 n=9+7)
/1000000-4     249ms ± 2%     359ms ± 7%   +44.49%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            304B ± 0%      272B ± 0%   -10.53%  (p=0.000 n=10+10)
/10-4         1.57kB ± 0%    2.18kB ± 0%   +38.78%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    23.0kB ± 0%   +10.02%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     210kB ± 0%   +61.40%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    2.69MB ± 0%  +109.17%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    23.8MB ± 0%   +85.49%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     213MB ± 0%   +65.89%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      11.0 ± 0%      ~     (all equal)
/10-4           75.0 ± 0%      79.0 ± 0%    +5.33%  (p=0.000 n=10+10)
/100-4           709 ± 0%       721 ± 0%    +1.69%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%    +0.29%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%     70.0k ± 0%    -0.05%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      700k ± 0%    -0.10%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     7.00M ± 0%    -0.11%  (p=0.000 n=10+10)
```

phf's deque, when used as a stack, offers considerably worse performance 
and memory footprint when compared to deque.


deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4           525ns ± 4%   10506ns ± 2%   +1901.92%  (p=0.000 n=10+10)
/10-4         3.66µs ± 5%   12.85µs ± 7%    +250.61%  (p=0.000 n=10+10)
/100-4        25.9µs ± 2%    31.3µs ± 3%     +20.48%  (p=0.000 n=10+10)
/1000-4        241µs ± 3%     226µs ± 7%      -5.92%  (p=0.000 n=10+10)
/10000-4      2.44ms ± 2%    2.08ms ± 4%     -14.85%  (p=0.000 n=10+10)
/100000-4     27.1ms ± 3%    24.3ms ± 4%     -10.11%  (p=0.000 n=10+10)
/1000000-4     282ms ± 6%     242ms ± 4%     -14.23%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            560B ± 0%    65792B ± 0%  +11648.57%  (p=0.000 n=10+10)
/10-4         5.71kB ± 0%   66.80kB ± 0%   +1069.47%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    76.9kB ± 0%    +267.64%  (p=0.000 n=10+10)
/1000-4        138kB ± 0%     243kB ± 0%     +75.91%  (p=0.000 n=10+10)
/10000-4      1.54MB ± 0%    1.38MB ± 0%      -9.98%  (p=0.000 n=10+10)
/100000-4     15.3MB ± 0%    12.9MB ± 0%     -15.37%  (p=0.000 n=9+10)
/1000000-4     152MB ± 0%     129MB ± 0%     -15.46%  (p=0.002 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      10.0 ± 0%     -16.67%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%      73.0 ± 0%      -5.19%  (p=0.000 n=10+10)
/100-4           709 ± 0%       703 ± 0%      -0.85%  (p=0.000 n=10+10)
/1000-4        7.02k ± 0%     7.00k ± 0%      -0.17%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%      -0.28%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%      -0.27%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%      -0.27%  (p=0.000 n=10+10)
```

For very high load scenarios, the cookiejar deque actually performs a bit better than deque. This is expected as the coookiejar queue was clearly optimized for dealing with very large data sets, as its internal block size (4096) indicates. It also implements a very interesting circular slice of blocks and its implementation is higly optimized for performance and efficiency.

In the era of [Microservices](https://en.wikipedia.org/wiki/Microservices) and [serverless computing](https://en.wikipedia.org/wiki/Serverless_computing), 
it is very important for the applications to remain lean and be able to start up (and shutdown) quickly as they are frequently deployed and moved around
to different pods, nodes to answer to different demands. The cookiejar deque, being orders of magnitude slower with small data sets (<= 100), would incur a considerable penalty on the application startup time, making it not a particularly good option for microservice and serverless applications.

A word of caution if you decide to use cookiejar's deque: the implementation is highly optimized for performance and so it doesn't perform even basic bound checks on its pop methods. So some care has to be taken when using the cookiejar deque. Deque, on the other hand, makes no concessions on security for performance gains.


deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4           419ns ± 8%   10783ns ± 8%   +2474.85%  (p=0.000 n=9+10)
/10-4         2.61µs ± 6%   13.09µs ± 8%    +402.48%  (p=0.000 n=10+10)
/100-4        25.2µs ± 5%    30.6µs ± 4%     +21.43%  (p=0.000 n=10+8)
/1000-4        234µs ± 6%     220µs ± 6%      -6.02%  (p=0.001 n=10+10)
/10000-4      2.36ms ± 7%    2.07ms ± 5%     -12.49%  (p=0.000 n=10+10)
/100000-4     24.2ms ± 2%    24.0ms ± 5%        ~     (p=0.315 n=9+10)
/1000000-4     249ms ± 2%     250ms ± 5%        ~     (p=0.631 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            304B ± 0%    65792B ± 0%  +21542.11%  (p=0.000 n=10+10)
/10-4         1.57kB ± 0%   66.80kB ± 0%   +4160.20%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    76.9kB ± 0%    +267.64%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     178kB ± 0%     +36.68%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    1.32MB ± 0%      +2.31%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    12.8MB ± 0%      +0.19%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     129MB ± 0%      +0.50%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      10.0 ± 0%      -9.09%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      73.0 ± 0%      -2.67%  (p=0.000 n=10+10)
/100-4           709 ± 0%       703 ± 0%      -0.85%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.00k ± 0%      -0.14%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%     70.0k ± 0%      -0.11%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      700k ± 0%      -0.11%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     7.00M ± 0%      -0.10%  (p=0.000 n=10+10)
```

As a stack, deque outperforms cookiejar deque for small data sets quite dramatically and displays a pretty matched performance for large ones, even though cookiejar deque was specifically optimized for large data sets.



### Other Test Results

#### deque vs impl7 - FIFO queue

deque vs impl7 - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillImpl7Queue.txt
name        old time/op    new time/op    delta
/0-4          39.9ns ± 9%    35.9ns ± 3%   -9.88%  (p=0.000 n=10+9)
/1-4           142ns ± 1%     133ns ± 1%   -6.61%  (p=0.000 n=10+10)
/10-4          636ns ± 1%     764ns ± 7%  +20.26%  (p=0.000 n=9+9)
/100-4        4.74µs ± 3%    4.28µs ± 3%   -9.67%  (p=0.000 n=9+9)
/1000-4       43.0µs ±23%    38.8µs ± 7%   -9.84%  (p=0.004 n=10+10)
/10000-4       450µs ±19%     388µs ± 5%  -13.70%  (p=0.001 n=10+10)
/100000-4     4.24ms ± 4%    3.95ms ± 2%   -6.84%  (p=0.000 n=10+8)
/1000000-4    46.8ms ± 1%    45.9ms ± 4%   -1.87%  (p=0.008 n=8+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%  -25.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%      112B ± 0%  -22.22%  (p=0.000 n=10+10)
/10-4           608B ± 0%      736B ± 0%  +21.05%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    4.26kB ± 0%  -31.27%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    33.2kB ± 0%   +0.58%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     323kB ± 0%   +0.23%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.23MB ± 0%   +0.20%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    32.3MB ± 0%   +0.19%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           15.0 ± 0%      17.0 ± 0%  +13.33%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%   +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.02k ± 0%   +0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.2k ± 0%   +0.79%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      102k ± 0%   +0.78%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.02M ± 0%   +0.78%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillImpl7Queue.txt
name       old time/op    new time/op    delta
/1-4         3.79µs ± 1%   10.03µs ± 6%  +165.06%  (p=0.000 n=9+10)
/10-4        37.8µs ± 7%    74.7µs ± 5%   +97.58%  (p=0.000 n=10+10)
/100-4        361µs ± 7%     442µs ± 4%   +22.25%  (p=0.000 n=9+10)
/1000-4      3.75ms ± 4%    3.97ms ± 3%    +6.10%  (p=0.000 n=10+9)
/10000-4     36.5ms ± 3%    39.1ms ± 6%    +7.17%  (p=0.000 n=10+10)
/100000-4     380ms ± 1%     421ms ±24%   +10.67%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    68.8kB ± 0%  +330.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     421kB ± 0%  +163.00%  (p=0.000 n=8+10)
/1000-4      2.42MB ± 0%    3.32MB ± 0%   +37.30%  (p=0.000 n=10+10)
/10000-4     17.0MB ± 0%    32.3MB ± 0%   +89.36%  (p=0.000 n=8+10)
/100000-4     162MB ± 0%     323MB ± 0%   +99.52%  (p=0.000 n=8+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.60k ± 0%   +60.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +1.80%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.02M ± 0%    +1.57%  (p=0.000 n=8+10)
/100000-4     10.0M ± 0%     10.2M ± 0%    +1.56%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullImpl7Queue.txt
name       old time/op    new time/op    delta
/1-4         3.75µs ± 4%    4.24µs ± 4%   +13.02%  (p=0.000 n=10+10)
/10-4        37.8µs ± 4%    43.3µs ±24%   +14.50%  (p=0.000 n=10+10)
/100-4        371µs ± 3%     398µs ± 5%    +7.21%  (p=0.000 n=10+10)
/1000-4      4.02ms ± 5%    4.01ms ± 7%      ~     (p=0.739 n=10+10)
/10000-4     39.4ms ± 5%    39.9ms ± 7%      ~     (p=0.356 n=10+9)
/100000-4     392ms ± 3%     412ms ±15%    +5.05%  (p=0.028 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     322kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4      2.40MB ± 0%    3.23MB ± 0%   +34.21%  (p=0.000 n=10+10)
/10000-4     16.6MB ± 0%    32.3MB ± 0%   +93.72%  (p=0.000 n=10+10)
/100000-4     161MB ± 0%     322MB ± 0%  +100.33%  (p=0.000 n=8+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       101 ± 0%    +1.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.01k ± 0%    +1.50%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.2k ± 0%    +1.56%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +1.17%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.02M ± 0%    +1.53%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.2M ± 0%    +1.56%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4           244ns ± 1%     230ns ± 3%   -5.98%  (p=0.000 n=10+10)
/10-4         1.86µs ± 1%    1.61µs ± 6%  -13.41%  (p=0.000 n=10+10)
/100-4        8.02µs ± 1%    7.87µs ± 2%   -1.85%  (p=0.007 n=10+9)
/1000-4       73.5µs ± 1%    71.5µs ± 1%   -2.70%  (p=0.000 n=10+9)
/10000-4       725µs ± 1%     723µs ± 3%     ~     (p=0.400 n=10+9)
/100000-4     8.20ms ± 0%    8.50ms ± 9%     ~     (p=0.156 n=9+10)
/1000000-4    86.4ms ± 1%    93.4ms ± 4%   +8.11%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      160B ± 0%  -28.57%  (p=0.000 n=10+10)
/10-4         4.91kB ± 0%    2.98kB ± 0%  -39.41%  (p=0.000 n=10+10)
/100-4        7.79kB ± 0%    7.94kB ± 0%   +1.85%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    65.9kB ± 0%  +21.63%  (p=0.000 n=10+10)
/10000-4       491kB ± 0%     647kB ± 0%  +31.71%  (p=0.000 n=10+10)
/100000-4     4.83MB ± 0%    6.45MB ± 0%  +33.69%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    64.5MB ± 0%  +33.83%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%     ~     (all equal)
/10-4           27.0 ± 0%      29.0 ± 0%   +7.41%  (p=0.000 n=10+10)
/100-4           207 ± 0%       211 ± 0%   +1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.04k ± 0%   +1.19%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.3k ± 0%   +1.16%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      203k ± 0%   +1.17%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.03M ± 0%   +1.17%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4          40.8ns ± 3%    99.0ns ±12%  +142.53%  (p=0.000 n=9+10)
/10-4          413ns ± 1%    1013ns ± 8%  +145.03%  (p=0.000 n=8+10)
/100-4        4.05µs ± 5%    9.67µs ± 5%  +139.11%  (p=0.000 n=10+9)
/1000-4       42.7µs ± 9%    93.4µs ± 3%  +118.47%  (p=0.000 n=10+8)
/10000-4       436µs ± 6%     983µs ±12%  +125.40%  (p=0.000 n=9+10)
/100000-4     3.92ms ± 2%   10.70ms ±22%  +172.84%  (p=0.000 n=9+10)
/1000000-4    39.2ms ± 1%    92.9ms ± 6%  +136.91%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      3.00 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      30.0 ± 0%  +200.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     3.00k ± 0%  +200.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     30.0k ± 0%  +200.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      300k ± 0%  +200.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     3.00M ± 0%  +200.00%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    38.1ns ± 5%    +9.48%  (p=0.000 n=10+9)
/10-4          353ns ± 2%     391ns ± 7%   +10.92%  (p=0.000 n=10+10)
/100-4        3.45µs ± 1%    4.13µs ±19%   +19.82%  (p=0.000 n=10+10)
/1000-4       34.5µs ± 1%    39.6µs ±12%   +14.95%  (p=0.000 n=10+10)
/10000-4       346µs ± 2%     385µs ± 4%   +11.40%  (p=0.000 n=10+10)
/100000-4     3.43ms ± 1%    3.85ms ±11%   +12.11%  (p=0.000 n=10+10)
/1000000-4    34.4ms ± 1%    38.3ms ± 6%   +11.13%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     32.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      322B ± 0%  +101.25%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     322kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    3.23MB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    32.3MB ± 0%  +101.56%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       101 ± 0%    +1.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     1.01k ± 0%    +1.50%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.2k ± 0%    +1.56%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      102k ± 0%    +1.56%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.02M ± 0%    +1.56%  (p=0.000 n=10+10)
```


#### deque vs list - FIFO queue
deque vs list - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillListQueue.txt
name        old time/op    new time/op    delta
/0-4          39.9ns ± 9%    39.1ns ± 2%      ~     (p=0.233 n=10+9)
/1-4           142ns ± 1%     107ns ± 1%   -24.99%  (p=0.000 n=10+9)
/10-4          636ns ± 1%     726ns ± 1%   +14.26%  (p=0.000 n=9+10)
/100-4        4.74µs ± 3%    6.74µs ± 1%   +42.11%  (p=0.000 n=9+9)
/1000-4       43.0µs ±23%    69.1µs ± 1%   +60.60%  (p=0.000 n=10+8)
/10000-4       450µs ±19%     712µs ± 1%   +58.35%  (p=0.000 n=10+10)
/100000-4     4.24ms ± 4%   20.61ms ±10%  +386.38%  (p=0.000 n=10+10)
/1000000-4    46.8ms ± 1%   148.7ms ± 2%  +217.80%  (p=0.000 n=8+7)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%      112B ± 0%   -22.22%  (p=0.000 n=10+10)
/10-4           608B ± 0%      688B ± 0%   +13.16%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    6.45kB ± 0%    +4.13%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    64.0kB ± 0%   +93.94%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     640kB ± 0%   +98.73%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +98.71%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    64.0MB ± 0%   +98.81%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      21.0 ± 0%   +40.00%  (p=0.000 n=10+10)
/100-4           107 ± 0%       201 ± 0%   +87.85%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     2.00k ± 0%   +97.53%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     20.0k ± 0%   +98.36%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      200k ± 0%   +98.44%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     2.00M ± 0%   +98.45%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillListQueue.txt
name       old time/op    new time/op    delta
/1-4         3.79µs ± 1%    7.36µs ± 4%   +94.38%  (p=0.000 n=9+9)
/10-4        37.8µs ± 7%    74.2µs ±10%   +96.14%  (p=0.000 n=10+10)
/100-4        361µs ± 7%     716µs ± 3%   +98.26%  (p=0.000 n=9+9)
/1000-4      3.75ms ± 4%    7.13ms ± 4%   +90.34%  (p=0.000 n=10+8)
/10000-4     36.5ms ± 3%    75.9ms ± 9%  +108.18%  (p=0.000 n=10+10)
/100000-4     380ms ± 1%    2066ms ±11%  +443.12%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=8+10)
/1000-4      2.42MB ± 0%    6.40MB ± 0%  +164.92%  (p=0.000 n=10+9)
/10000-4     17.0MB ± 0%    64.0MB ± 0%  +275.47%  (p=0.000 n=8+9)
/100000-4     162MB ± 0%     640MB ± 0%  +295.68%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%   +99.22%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%   +99.90%  (p=0.000 n=8+10)
/100000-4     10.0M ± 0%     20.0M ± 0%   +99.98%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullListQueue.txt
name       old time/op    new time/op    delta
/1-4         3.75µs ± 4%    8.21µs ± 5%  +118.90%  (p=0.000 n=10+10)
/10-4        37.8µs ± 4%    81.5µs ± 4%  +115.45%  (p=0.000 n=10+10)
/100-4        371µs ± 3%     803µs ± 3%  +116.29%  (p=0.000 n=10+10)
/1000-4      4.02ms ± 5%    8.38ms ± 9%  +108.42%  (p=0.000 n=10+10)
/10000-4     39.4ms ± 5%    86.6ms ± 7%  +119.52%  (p=0.000 n=10+10)
/100000-4     392ms ± 3%    1608ms ± 6%  +310.40%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      2.40MB ± 0%    6.40MB ± 0%  +166.35%  (p=0.000 n=10+10)
/10000-4     16.6MB ± 0%    64.0MB ± 0%  +284.44%  (p=0.000 n=10+9)
/100000-4     161MB ± 0%     640MB ± 0%  +297.55%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%   +99.23%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%   +99.94%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%   +99.99%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseListQueue.txt
name        old time/op    new time/op    delta
/1-4           244ns ± 1%     179ns ± 5%   -26.62%  (p=0.000 n=10+10)
/10-4         1.86µs ± 1%    1.44µs ± 4%   -22.59%  (p=0.000 n=10+10)
/100-4        8.02µs ± 1%   14.12µs ± 7%   +76.19%  (p=0.000 n=10+9)
/1000-4       73.5µs ± 1%   139.3µs ± 5%   +89.53%  (p=0.000 n=10+9)
/10000-4       725µs ± 1%    1473µs ± 5%  +103.11%  (p=0.000 n=10+8)
/100000-4     8.20ms ± 0%   22.99ms ± 7%  +180.28%  (p=0.000 n=9+10)
/1000000-4    86.4ms ± 1%   268.3ms ± 8%  +210.70%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      176B ± 0%   -21.43%  (p=0.000 n=10+10)
/10-4         4.91kB ± 0%    1.33kB ± 0%   -72.96%  (p=0.000 n=10+10)
/100-4        7.79kB ± 0%   12.85kB ± 0%   +64.89%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%   128.0kB ± 0%  +136.50%  (p=0.000 n=10+10)
/10000-4       491kB ± 0%    1280kB ± 0%  +160.53%  (p=0.000 n=10+10)
/100000-4     4.83MB ± 0%   12.80MB ± 0%  +165.24%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   128.0MB ± 0%  +165.58%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      41.0 ± 0%   +51.85%  (p=0.000 n=10+10)
/100-4           207 ± 0%       401 ± 0%   +93.72%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     4.00k ± 0%   +98.56%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     40.0k ± 0%   +99.14%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      400k ± 0%   +99.21%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     4.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseListQueue.txt
name        old time/op    new time/op    delta
/1-4          40.8ns ± 3%    70.4ns ±12%   +72.58%  (p=0.000 n=9+9)
/10-4          413ns ± 1%     703ns ± 4%   +70.01%  (p=0.000 n=8+9)
/100-4        4.05µs ± 5%    6.73µs ± 1%   +66.26%  (p=0.000 n=10+8)
/1000-4       42.7µs ± 9%    70.8µs ± 6%   +65.59%  (p=0.000 n=10+9)
/10000-4       436µs ± 6%     696µs ± 3%   +59.59%  (p=0.000 n=9+8)
/100000-4     3.92ms ± 2%    7.25ms ± 6%   +84.75%  (p=0.000 n=9+8)
/1000000-4    39.2ms ± 1%    69.4ms ± 8%   +76.96%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableListQueue.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    86.9ns ± 6%  +149.90%  (p=0.000 n=10+10)
/10-4          353ns ± 2%     821ns ± 4%  +132.68%  (p=0.000 n=10+9)
/100-4        3.45µs ± 1%    7.76µs ± 1%  +124.84%  (p=0.000 n=10+8)
/1000-4       34.5µs ± 1%    90.0µs ±37%  +160.87%  (p=0.000 n=10+10)
/10000-4       346µs ± 2%     791µs ± 7%  +128.74%  (p=0.000 n=10+9)
/100000-4     3.43ms ± 1%    9.03ms ±26%  +163.07%  (p=0.000 n=10+9)
/1000000-4    34.4ms ± 1%    83.4ms ± 9%  +142.21%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+9)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```


#### deque vs list - LIFO stack

deque vs list - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillListStack.txt
name        old time/op    new time/op    delta
/0-4          39.0ns ± 6%    39.0ns ± 1%      ~     (p=0.244 n=9+9)
/1-4           145ns ± 2%     107ns ± 1%   -25.74%  (p=0.000 n=8+10)
/10-4          615ns ± 0%     725ns ± 1%   +17.86%  (p=0.000 n=8+10)
/100-4        4.42µs ± 1%    6.75µs ± 1%   +52.80%  (p=0.000 n=10+10)
/1000-4       36.2µs ± 1%    68.4µs ± 0%   +88.95%  (p=0.000 n=10+10)
/10000-4       383µs ± 7%     700µs ± 0%   +82.86%  (p=0.000 n=10+9)
/100000-4     4.10ms ± 9%   21.25ms ±10%  +418.15%  (p=0.000 n=10+10)
/1000000-4    44.5ms ± 4%   157.6ms ±10%  +253.86%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%      112B ± 0%   -22.22%  (p=0.000 n=10+10)
/10-4           608B ± 0%      688B ± 0%   +13.16%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    6.45kB ± 0%    +4.13%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    64.0kB ± 0%   +93.94%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     640kB ± 0%   +98.73%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +98.71%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    64.0MB ± 0%   +98.81%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      21.0 ± 0%   +40.00%  (p=0.000 n=10+10)
/100-4           107 ± 0%       201 ± 0%   +87.85%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     2.00k ± 0%   +97.53%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     20.0k ± 0%   +98.36%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      200k ± 0%   +98.44%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     2.00M ± 0%   +98.45%  (p=0.000 n=10+10)
```

deque vs list - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillListStack.txt
name       old time/op    new time/op    delta
/1-4         3.62µs ± 1%    7.28µs ± 5%  +101.09%  (p=0.000 n=10+10)
/10-4        34.1µs ± 5%    71.6µs ± 4%  +110.08%  (p=0.000 n=10+8)
/100-4        332µs ± 2%     718µs ± 6%  +116.53%  (p=0.000 n=10+10)
/1000-4      3.28ms ± 2%    7.04ms ± 5%  +114.73%  (p=0.000 n=10+10)
/10000-4     34.0ms ± 3%    77.3ms ±16%  +127.12%  (p=0.000 n=9+9)
/100000-4     370ms ± 5%    2156ms ±11%  +483.37%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=10+9)
/10000-4     16.0MB ± 0%    64.0MB ± 0%  +299.92%  (p=0.000 n=10+9)
/100000-4     161MB ± 0%     640MB ± 0%  +298.65%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%   +99.99%  (p=0.000 n=10+9)
```

deque vs list - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullListStack.txt
name       old time/op    new time/op    delta
/1-4         3.90µs ± 6%    8.28µs ± 4%  +112.17%  (p=0.000 n=10+9)
/10-4        34.6µs ± 2%    83.8µs ± 2%  +142.14%  (p=0.000 n=9+9)
/100-4        337µs ± 4%     831µs ± 2%  +146.18%  (p=0.000 n=9+10)
/1000-4      3.58ms ± 8%    8.55ms ±11%  +138.62%  (p=0.000 n=10+9)
/10000-4     34.9ms ± 7%    96.3ms ±10%  +175.77%  (p=0.000 n=10+9)
/100000-4     339ms ± 2%    1619ms ± 9%  +378.07%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+9)
/100000-4     160MB ± 0%     640MB ± 0%  +300.00%  (p=0.000 n=10+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%  +100.00%  (p=0.000 n=10+8)
```

deque vs list - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseListStack.txt
name        old time/op    new time/op    delta
/1-4           240ns ± 0%     180ns ± 6%   -25.01%  (p=0.000 n=9+10)
/10-4          929ns ± 1%    1465ns ± 9%   +57.71%  (p=0.000 n=10+9)
/100-4        8.65µs ± 1%   13.99µs ± 9%   +61.78%  (p=0.000 n=10+9)
/1000-4       66.7µs ± 1%   139.6µs ± 2%  +109.38%  (p=0.000 n=10+8)
/10000-4       666µs ± 1%    1520µs ±11%  +128.26%  (p=0.000 n=10+10)
/100000-4     7.78ms ± 1%   27.93ms ± 9%  +259.18%  (p=0.000 n=9+10)
/1000000-4    81.9ms ± 1%   357.6ms ±29%  +336.45%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      176B ± 0%   -21.43%  (p=0.000 n=10+10)
/10-4           768B ± 0%     1328B ± 0%   +72.92%  (p=0.000 n=10+10)
/100-4        12.9kB ± 0%    12.8kB ± 0%    -0.50%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%   128.0kB ± 0%  +156.10%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%    1280kB ± 0%  +162.74%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   12.80MB ± 0%  +165.47%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   128.0MB ± 0%  +165.60%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      41.0 ± 0%   +64.00%  (p=0.000 n=10+10)
/100-4           209 ± 0%       401 ± 0%   +91.87%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     4.00k ± 0%   +98.76%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     40.0k ± 0%   +99.16%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      400k ± 0%   +99.22%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     4.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

deque vs list - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseListStack.txt
name        old time/op    new time/op    delta
/1-4          35.0ns ± 1%    72.6ns ± 3%  +107.39%  (p=0.000 n=8+8)
/10-4          352ns ± 2%     717ns ± 3%  +103.46%  (p=0.000 n=10+10)
/100-4        3.45µs ± 2%    7.01µs ± 4%  +103.48%  (p=0.000 n=10+9)
/1000-4       34.4µs ± 1%    81.4µs ±36%  +136.61%  (p=0.000 n=9+10)
/10000-4       343µs ± 1%     699µs ± 5%  +104.10%  (p=0.000 n=9+8)
/100000-4     3.43ms ± 1%    7.22ms ± 6%  +110.60%  (p=0.000 n=10+8)
/1000000-4    34.3ms ± 1%    71.9ms ± 7%  +109.82%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+8)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

deque vs list - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableListStack.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 2%    87.1ns ±12%  +146.75%  (p=0.000 n=10+9)
/10-4          355ns ± 0%     920ns ±31%  +159.55%  (p=0.000 n=7+10)
/100-4        3.46µs ± 1%    8.71µs ±17%  +151.36%  (p=0.000 n=10+10)
/1000-4       34.6µs ± 1%    87.1µs ±31%  +151.75%  (p=0.000 n=8+9)
/10000-4       406µs ±15%     801µs ± 6%   +97.45%  (p=0.000 n=10+9)
/100000-4     3.90ms ±12%    8.66ms ±23%  +121.94%  (p=0.000 n=10+9)
/1000000-4    38.3ms ±18%    85.9ms ±13%  +124.24%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+9)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=9+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

#### deque vs slice - FIFO queue

deque vs slice - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillSliceQueue.txt
name        old time/op    new time/op    delta
/0-4          39.9ns ± 9%    41.1ns ± 4%    +3.19%  (p=0.029 n=10+10)
/1-4           142ns ± 1%     101ns ± 7%   -28.94%  (p=0.000 n=10+8)
/10-4          636ns ± 1%     662ns ±14%      ~     (p=0.713 n=9+9)
/100-4        4.74µs ± 3%    4.64µs ±30%      ~     (p=0.780 n=9+10)
/1000-4       43.0µs ±23%    33.1µs ± 3%   -23.01%  (p=0.000 n=10+9)
/10000-4       450µs ±19%     410µs ± 9%      ~     (p=0.089 n=10+10)
/100000-4     4.24ms ± 4%    9.28ms ±10%  +119.07%  (p=0.000 n=10+9)
/1000000-4    46.8ms ± 1%   115.2ms ±33%  +146.21%  (p=0.000 n=8+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     32.0B ± 0%   -50.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%       56B ± 0%   -61.11%  (p=0.000 n=10+10)
/10-4           608B ± 0%      440B ± 0%   -27.63%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    3.67kB ± 0%   -40.70%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    32.4kB ± 0%    -1.87%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     546kB ± 0%   +69.63%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.19%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    61.2MB ± 0%   +90.08%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      16.0 ± 0%    +6.67%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    -0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.61%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs slice - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillSliceQueue.txt
name       old time/op    new time/op    delta
/1-4         3.79µs ± 1%    6.00µs ± 9%   +58.50%  (p=0.000 n=9+10)
/10-4        37.8µs ± 7%    45.8µs ±11%   +21.09%  (p=0.000 n=10+10)
/100-4        361µs ± 7%     330µs ± 4%    -8.63%  (p=0.000 n=9+10)
/1000-4      3.75ms ± 4%    3.77ms ±26%      ~     (p=0.481 n=10+10)
/10000-4     36.5ms ± 3%    38.6ms ±13%      ~     (p=0.123 n=10+10)
/100000-4     380ms ± 1%     813ms ±10%  +113.65%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.0kB ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     314kB ± 0%   +96.00%  (p=0.000 n=8+10)
/1000-4      2.42MB ± 0%    3.15MB ± 0%   +30.34%  (p=0.000 n=10+10)
/10000-4     17.0MB ± 0%    48.8MB ± 0%  +186.12%  (p=0.000 n=8+10)
/100000-4     162MB ± 0%     548MB ± 0%  +238.73%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.1k ± 0%    +1.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    -0.19%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    +0.01%  (p=0.000 n=8+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.00%  (p=0.000 n=10+10)
```

deque vs slice - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullSliceQueue.txt
name       old time/op    new time/op    delta
/1-4         3.75µs ± 4%    4.16µs ±10%   +10.90%  (p=0.000 n=10+9)
/10-4        37.8µs ± 4%    40.4µs ± 7%    +6.80%  (p=0.001 n=10+10)
/100-4        371µs ± 3%     406µs ±27%    +9.52%  (p=0.000 n=10+9)
/1000-4      4.02ms ± 5%    4.11ms ± 5%      ~     (p=0.173 n=10+8)
/10000-4     39.4ms ± 5%    40.4ms ± 6%      ~     (p=0.113 n=10+9)
/100000-4     392ms ± 3%     929ms ±18%  +137.03%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    48.2kB ± 0%  +201.08%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     483kB ± 0%  +201.70%  (p=0.000 n=10+10)
/1000-4      2.40MB ± 0%    5.15MB ± 0%  +114.32%  (p=0.000 n=10+8)
/10000-4     16.6MB ± 0%    52.0MB ± 0%  +212.63%  (p=0.000 n=10+8)
/100000-4     161MB ± 0%     551MB ± 0%  +242.32%  (p=0.000 n=8+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    -0.35%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    -0.01%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.00%  (p=0.000 n=10+9)
```

deque vs slice - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseSliceQueue.txt
name        old time/op    new time/op    delta
/1-4           244ns ± 1%     160ns ± 7%   -34.38%  (p=0.000 n=10+8)
/10-4         1.86µs ± 1%    1.10µs ± 4%   -40.61%  (p=0.000 n=10+9)
/100-4        8.02µs ± 1%    7.23µs ± 5%    -9.78%  (p=0.000 n=10+9)
/1000-4       73.5µs ± 1%    66.7µs ± 4%    -9.28%  (p=0.000 n=10+9)
/10000-4       725µs ± 1%     729µs ± 3%      ~     (p=0.796 n=10+10)
/100000-4     8.20ms ± 0%   15.10ms ±23%   +84.08%  (p=0.000 n=9+9)
/1000000-4    86.4ms ± 1%   169.3ms ± 9%   +96.08%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%       88B ± 0%   -60.71%  (p=0.000 n=10+10)
/10-4         4.91kB ± 0%    0.78kB ± 0%   -84.20%  (p=0.000 n=10+10)
/100-4        7.79kB ± 0%    6.66kB ± 0%   -14.48%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    74.0kB ± 0%   +36.66%  (p=0.000 n=10+10)
/10000-4       491kB ± 0%     991kB ± 0%  +101.62%  (p=0.000 n=10+10)
/100000-4     4.83MB ± 0%   11.42MB ± 0%  +136.55%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   114.6MB ± 0%  +137.72%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      29.0 ± 0%    +7.41%  (p=0.000 n=10+10)
/100-4           207 ± 0%       214 ± 0%    +3.38%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.02k ± 0%    +0.25%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.26%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs slice - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseSliceQueue.txt
name        old time/op    new time/op    delta
/1-4          40.8ns ± 3%    56.3ns ± 3%   +37.89%  (p=0.000 n=9+8)
/10-4          413ns ± 1%     574ns ± 4%   +38.87%  (p=0.000 n=8+8)
/100-4        4.05µs ± 5%    5.50µs ± 3%   +35.91%  (p=0.000 n=10+9)
/1000-4       42.7µs ± 9%    55.6µs ± 8%   +30.01%  (p=0.000 n=10+9)
/10000-4       436µs ± 6%     547µs ± 2%   +25.39%  (p=0.000 n=9+8)
/100000-4     3.92ms ± 2%    6.84ms ±61%   +74.36%  (p=0.000 n=9+10)
/1000000-4    39.2ms ± 1%    56.0ms ± 5%   +42.97%  (p=0.000 n=9+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     24.0B ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      240B ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    24.0kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     240kB ± 0%   +50.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    2.40MB ± 0%   +50.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    24.0MB ± 0%   +50.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

deque vs slice - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableSliceQueue.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    39.2ns ± 5%   +12.71%  (p=0.000 n=10+8)
/10-4          353ns ± 2%     452ns ±37%   +28.07%  (p=0.000 n=10+10)
/100-4        3.45µs ± 1%    3.77µs ± 8%    +9.29%  (p=0.000 n=10+9)
/1000-4       34.5µs ± 1%    39.7µs ±25%   +15.02%  (p=0.000 n=10+9)
/10000-4       346µs ± 2%     400µs ± 4%   +15.55%  (p=0.000 n=10+9)
/100000-4     3.43ms ± 1%    3.91ms ± 7%   +14.01%  (p=0.000 n=10+9)
/1000000-4    34.4ms ± 1%    39.4ms ± 5%   +14.38%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     48.0B ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      481B ± 0%  +200.62%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    48.2kB ± 0%  +200.96%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     482kB ± 0%  +200.97%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    4.82MB ± 0%  +200.97%  (p=0.000 n=10+8)
/1000000-4    16.0MB ± 0%    48.2MB ± 0%  +200.97%  (p=0.000 n=10+7)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       100 ± 0%      ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      100k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.00M ± 0%    +0.03%  (p=0.000 n=10+10)
```

#### deque vs slice - LIFO stack

deque vs slice - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillSliceStack.txt
name        old time/op    new time/op    delta
/0-4          39.0ns ± 6%    43.4ns ± 3%   +11.35%  (p=0.000 n=9+9)
/1-4           145ns ± 2%     100ns ± 4%   -30.96%  (p=0.000 n=8+9)
/10-4          615ns ± 0%     623ns ± 5%      ~     (p=0.938 n=8+8)
/100-4        4.42µs ± 1%    4.79µs ±37%      ~     (p=0.631 n=10+10)
/1000-4       36.2µs ± 1%    36.0µs ± 7%      ~     (p=0.529 n=10+10)
/10000-4       383µs ± 7%     426µs ± 6%   +11.32%  (p=0.000 n=10+9)
/100000-4     4.10ms ± 9%   10.34ms ±14%  +152.06%  (p=0.000 n=10+10)
/1000000-4    44.5ms ± 4%   116.0ms ± 9%  +160.57%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     32.0B ± 0%   -50.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%       56B ± 0%   -61.11%  (p=0.000 n=10+10)
/10-4           608B ± 0%      440B ± 0%   -27.63%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    3.67kB ± 0%   -40.70%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    32.4kB ± 0%    -1.87%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     546kB ± 0%   +69.63%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.19%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    61.2MB ± 0%   +90.08%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      16.0 ± 0%    +6.67%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    -0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.61%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs slice - LIFO stack - [refill tests](benchmark-refill_test.go)
```
 benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillSliceStack.txt
name       old time/op    new time/op    delta
/1-4         3.62µs ± 1%    3.01µs ± 5%  -16.83%  (p=0.000 n=10+10)
/10-4        34.1µs ± 5%    31.7µs ±24%     ~     (p=0.138 n=10+10)
/100-4        332µs ± 2%     286µs ±17%  -13.64%  (p=0.001 n=10+9)
/1000-4      3.28ms ± 2%    2.66ms ± 2%  -18.90%  (p=0.000 n=10+9)
/10000-4     34.0ms ± 3%    29.4ms ±16%  -13.61%  (p=0.001 n=9+9)
/100000-4     370ms ± 5%     343ms ± 3%   -7.12%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   -0.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=10+9)
/10000-4     16.0MB ± 0%    16.0MB ± 0%   +0.03%  (p=0.000 n=10+9)
/100000-4     161MB ± 0%     162MB ± 0%   +0.63%  (p=0.000 n=9+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%   -0.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%   -0.00%  (p=0.000 n=10+8)
```

deque vs slice - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullSliceStack.txt
name       old time/op    new time/op    delta
/1-4         3.90µs ± 6%    3.11µs ± 9%  -20.28%  (p=0.000 n=10+9)
/10-4        34.6µs ± 2%    29.6µs ± 9%  -14.55%  (p=0.000 n=9+9)
/100-4        337µs ± 4%     288µs ± 6%  -14.56%  (p=0.000 n=9+9)
/1000-4      3.58ms ± 8%    2.96ms ±12%  -17.24%  (p=0.000 n=10+9)
/10000-4     34.9ms ± 7%    29.3ms ± 7%  -16.17%  (p=0.000 n=10+10)
/100000-4     339ms ± 2%     336ms ± 2%     ~     (p=0.277 n=9+8)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%   -0.00%  (p=0.022 n=10+9)
/100000-4     160MB ± 0%     160MB ± 0%     ~     (all equal)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%     ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%     ~     (all equal)
```

deque vs slice - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseSliceStack.txt
name        old time/op    new time/op    delta
/1-4           240ns ± 0%     158ns ± 5%  -34.41%  (p=0.000 n=9+10)
/10-4          929ns ± 1%     891ns ± 6%   -4.12%  (p=0.001 n=10+10)
/100-4        8.65µs ± 1%    6.52µs ± 2%  -24.63%  (p=0.000 n=10+8)
/1000-4       66.7µs ± 1%    65.0µs ±25%     ~     (p=0.143 n=10+10)
/10000-4       666µs ± 1%     639µs ± 3%   -4.05%  (p=0.000 n=10+9)
/100000-4     7.78ms ± 1%   11.38ms ± 5%  +46.29%  (p=0.000 n=9+9)
/1000000-4    81.9ms ± 1%   138.6ms ± 8%  +69.13%  (p=0.000 n=9+8)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%       88B ± 0%  -60.71%  (p=0.000 n=10+10)
/10-4           768B ± 0%      600B ± 0%  -21.88%  (p=0.000 n=10+10)
/100-4        12.9kB ± 0%     5.3kB ± 0%  -59.17%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    48.4kB ± 0%   -3.18%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     706kB ± 0%  +44.98%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    7.85MB ± 0%  +62.90%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    77.2MB ± 0%  +60.17%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%  -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      26.0 ± 0%   +4.00%  (p=0.000 n=10+10)
/100-4           209 ± 0%       209 ± 0%     ~     (all equal)
/1000-4        2.01k ± 0%     2.01k ± 0%   -0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%   -0.32%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%   -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%   -0.39%  (p=0.000 n=10+10)
```

deque vs slice - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseSliceStack.txt
name        old time/op    new time/op    delta
/1-4          35.0ns ± 1%    32.3ns ± 6%   -7.57%  (p=0.000 n=8+8)
/10-4          352ns ± 2%     307ns ± 2%  -12.76%  (p=0.000 n=10+9)
/100-4        3.45µs ± 2%    3.04µs ± 4%  -11.83%  (p=0.000 n=10+10)
/1000-4       34.4µs ± 1%    31.8µs ±11%   -7.71%  (p=0.004 n=9+9)
/10000-4       343µs ± 1%     317µs ± 6%   -7.40%  (p=0.000 n=9+10)
/100000-4     3.43ms ± 1%    3.08ms ± 5%  -10.01%  (p=0.000 n=10+10)
/1000000-4    34.3ms ± 1%    30.8ms ± 5%  -10.03%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs slice - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableSliceStack.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 2%    27.4ns ± 5%  -22.48%  (p=0.000 n=10+9)
/10-4          355ns ± 0%     282ns ± 6%  -20.47%  (p=0.000 n=7+8)
/100-4        3.46µs ± 1%    2.70µs ± 3%  -22.19%  (p=0.000 n=10+8)
/1000-4       34.6µs ± 1%    26.8µs ± 6%  -22.47%  (p=0.000 n=8+8)
/10000-4       406µs ±15%     265µs ± 6%  -34.67%  (p=0.000 n=10+9)
/100000-4     3.90ms ±12%    2.64ms ± 4%  -32.31%  (p=0.000 n=10+8)
/1000000-4    38.3ms ±18%    26.9ms ± 7%  -29.71%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```


#### deque vs gammazero - FIFO queue

deque vs gammazero - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillGammazeroQueue.txt
name        old time/op    new time/op    delta
/0-4          39.9ns ± 9%    38.7ns ± 6%      ~     (p=0.183 n=10+10)
/1-4           142ns ± 1%     176ns ± 5%   +23.77%  (p=0.000 n=10+9)
/10-4          636ns ± 1%     510ns ± 7%   -19.82%  (p=0.000 n=9+9)
/100-4        4.74µs ± 3%    5.47µs ±11%   +15.39%  (p=0.000 n=9+10)
/1000-4       43.0µs ±23%    44.7µs ± 3%      ~     (p=0.156 n=10+9)
/10000-4       450µs ±19%     509µs ± 2%      ~     (p=0.173 n=10+8)
/100000-4     4.24ms ± 4%    7.22ms ±20%   +70.34%  (p=0.000 n=10+9)
/1000000-4    46.8ms ± 1%    81.8ms ±15%   +74.75%  (p=0.000 n=8+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%      320B ± 0%  +122.22%  (p=0.000 n=10+10)
/10-4           608B ± 0%      464B ± 0%   -23.68%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    7.28kB ± 0%   +17.57%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    64.7kB ± 0%   +95.88%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.72%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.00%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.05%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      12.0 ± 0%   -20.00%  (p=0.000 n=10+10)
/100-4           107 ± 0%       108 ± 0%    +0.93%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.60%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillGammazeroQueue.txt
name       old time/op    new time/op    delta
/1-4         3.79µs ± 1%    3.85µs ± 1%    +1.59%  (p=0.000 n=9+10)
/10-4        37.8µs ± 7%    36.3µs ± 1%    -4.10%  (p=0.001 n=10+10)
/100-4        361µs ± 7%     511µs ± 2%   +41.48%  (p=0.000 n=9+10)
/1000-4      3.75ms ± 4%    4.50ms ± 1%   +20.16%  (p=0.000 n=10+10)
/10000-4     36.5ms ± 3%    51.8ms ± 2%   +41.98%  (p=0.000 n=10+10)
/100000-4     380ms ± 1%     695ms ± 1%   +82.59%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +335.99%  (p=0.000 n=8+10)
/1000-4      2.42MB ± 0%    6.44MB ± 0%  +166.51%  (p=0.000 n=10+10)
/10000-4     17.0MB ± 0%    94.6MB ± 0%  +454.79%  (p=0.000 n=8+8)
/100000-4     162MB ± 0%     789MB ± 0%  +387.84%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +0.80%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    +0.15%  (p=0.000 n=8+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.02%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullGammazeroQueue.txt
name       old time/op    new time/op    delta
/1-4         3.75µs ± 4%    3.64µs ± 2%    -2.98%  (p=0.001 n=10+10)
/10-4        37.8µs ± 4%    38.1µs ± 9%      ~     (p=0.684 n=10+10)
/100-4        371µs ± 3%     365µs ± 2%      ~     (p=0.063 n=10+10)
/1000-4      4.02ms ± 5%    3.58ms ± 4%   -10.97%  (p=0.000 n=10+10)
/10000-4     39.4ms ± 5%    37.1ms ± 3%    -6.02%  (p=0.000 n=10+9)
/100000-4     392ms ± 3%     818ms ±34%  +108.73%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      2.40MB ± 0%    1.60MB ± 0%   -33.41%  (p=0.000 n=10+10)
/10000-4     16.6MB ± 0%    16.0MB ± 0%    -3.89%  (p=0.000 n=10+9)
/100000-4     161MB ± 0%     632MB ± 0%  +292.50%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    -0.39%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    -0.03%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    -0.00%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4           244ns ± 1%     221ns ± 4%    -9.55%  (p=0.000 n=10+9)
/10-4         1.86µs ± 1%    0.85µs ± 5%   -54.46%  (p=0.000 n=10+9)
/100-4        8.02µs ± 1%    8.91µs ± 8%   +11.09%  (p=0.000 n=10+8)
/1000-4       73.5µs ± 1%    80.4µs ± 4%    +9.39%  (p=0.000 n=10+10)
/10000-4       725µs ± 1%     924µs ± 7%   +27.34%  (p=0.000 n=10+9)
/100000-4     8.20ms ± 0%   10.97ms ± 5%   +33.73%  (p=0.000 n=9+9)
/1000000-4    86.4ms ± 1%   122.9ms ± 9%   +42.36%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      336B ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4         4.91kB ± 0%    0.62kB ± 0%   -87.30%  (p=0.000 n=10+10)
/100-4        7.79kB ± 0%    8.88kB ± 0%   +13.96%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    80.7kB ± 0%   +49.02%  (p=0.000 n=10+10)
/10000-4       491kB ± 0%    1106kB ± 0%  +125.10%  (p=0.000 n=10+10)
/100000-4     4.83MB ± 0%    9.49MB ± 0%   +96.67%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.83%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      4.00 ± 0%   -33.33%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      22.0 ± 0%   -18.52%  (p=0.000 n=10+10)
/100-4           207 ± 0%       208 ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.01k ± 0%    -0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.32%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4          40.8ns ± 3%    35.1ns ± 1%  -13.95%  (p=0.000 n=9+10)
/10-4          413ns ± 1%     355ns ± 1%  -14.10%  (p=0.000 n=8+10)
/100-4        4.05µs ± 5%    3.65µs ±16%   -9.75%  (p=0.003 n=10+9)
/1000-4       42.7µs ± 9%    36.4µs ± 3%  -14.94%  (p=0.000 n=10+10)
/10000-4       436µs ± 6%     363µs ± 4%  -16.76%  (p=0.000 n=9+9)
/100000-4     3.92ms ± 2%    3.61ms ± 4%   -8.03%  (p=0.000 n=9+10)
/1000000-4    39.2ms ± 1%    35.0ms ± 3%  -10.58%  (p=0.000 n=9+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs gammazero - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    35.5ns ± 3%   +2.10%  (p=0.001 n=10+10)
/10-4          353ns ± 2%     370ns ± 3%   +5.00%  (p=0.000 n=10+9)
/100-4        3.45µs ± 1%    3.77µs ± 8%   +9.33%  (p=0.000 n=10+9)
/1000-4       34.5µs ± 1%    36.6µs ± 6%   +6.09%  (p=0.000 n=10+9)
/10000-4       346µs ± 2%     357µs ± 4%   +3.18%  (p=0.010 n=10+9)
/100000-4     3.43ms ± 1%    3.59ms ± 4%   +4.67%  (p=0.000 n=10+9)
/1000000-4    34.4ms ± 1%    38.8ms ±25%  +12.70%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs gammazero - LIFO stack

deque vs gammazero - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillGammazeroStack.txt
name        old time/op    new time/op    delta
/0-4          39.0ns ± 6%    40.3ns ± 7%    +3.51%  (p=0.033 n=9+9)
/1-4           145ns ± 2%     186ns ±10%   +28.68%  (p=0.000 n=8+10)
/10-4          615ns ± 0%     524ns ± 7%   -14.86%  (p=0.000 n=8+10)
/100-4        4.42µs ± 1%    5.47µs ±11%   +23.85%  (p=0.000 n=10+9)
/1000-4       36.2µs ± 1%    46.9µs ± 6%   +29.50%  (p=0.000 n=10+10)
/10000-4       383µs ± 7%     542µs ± 7%   +41.63%  (p=0.000 n=10+10)
/100000-4     4.10ms ± 9%    7.21ms ± 7%   +75.78%  (p=0.000 n=10+9)
/1000000-4    44.5ms ± 4%    80.4ms ± 4%   +80.45%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%      320B ± 0%  +122.22%  (p=0.000 n=10+10)
/10-4           608B ± 0%      464B ± 0%   -23.68%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    7.28kB ± 0%   +17.57%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    64.7kB ± 0%   +95.88%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.72%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.00%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.05%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      12.0 ± 0%   -20.00%  (p=0.000 n=10+10)
/100-4           107 ± 0%       108 ± 0%    +0.93%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.60%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillGammazeroStack.txt
name       old time/op    new time/op    delta
/1-4         3.62µs ± 1%    3.67µs ± 1%    +1.47%  (p=0.000 n=10+10)
/10-4        34.1µs ± 5%    36.0µs ± 2%    +5.44%  (p=0.000 n=10+10)
/100-4        332µs ± 2%     506µs ± 1%   +52.55%  (p=0.000 n=10+10)
/1000-4      3.28ms ± 2%    4.45ms ± 1%   +35.68%  (p=0.000 n=10+10)
/10000-4     34.0ms ± 3%    51.1ms ± 1%   +50.09%  (p=0.000 n=9+10)
/100000-4     370ms ± 5%     694ms ± 2%   +87.75%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +336.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.44MB ± 0%  +302.39%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    94.6MB ± 0%  +490.92%  (p=0.000 n=10+10)
/100000-4     161MB ± 0%     789MB ± 0%  +391.50%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.20%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    +0.20%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.02%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullGammazeroStack.txt
name       old time/op    new time/op    delta
/1-4         3.90µs ± 6%    3.77µs ± 7%      ~     (p=0.075 n=10+9)
/10-4        34.6µs ± 2%    38.6µs ± 5%   +11.40%  (p=0.000 n=9+9)
/100-4        337µs ± 4%     365µs ± 3%    +8.13%  (p=0.000 n=9+10)
/1000-4      3.58ms ± 8%    3.57ms ± 1%      ~     (p=0.968 n=10+9)
/10000-4     34.9ms ± 7%    37.5ms ± 7%    +7.45%  (p=0.001 n=10+9)
/100000-4     339ms ± 2%     691ms ± 3%  +104.01%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%      ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%    -0.00%  (p=0.022 n=10+9)
/100000-4     160MB ± 0%     632MB ± 0%  +294.91%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%      ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.00%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4           240ns ± 0%     216ns ± 9%   -10.07%  (p=0.000 n=9+9)
/10-4          929ns ± 1%     846ns ± 6%    -8.87%  (p=0.000 n=10+10)
/100-4        8.65µs ± 1%    8.69µs ± 5%      ~     (p=0.393 n=10+10)
/1000-4       66.7µs ± 1%    80.1µs ± 5%   +20.08%  (p=0.000 n=10+9)
/10000-4       666µs ± 1%     944µs ± 6%   +41.78%  (p=0.000 n=10+10)
/100000-4     7.78ms ± 1%   11.52ms ± 8%   +48.07%  (p=0.000 n=9+9)
/1000000-4    81.9ms ± 1%   124.7ms ± 3%   +52.18%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      336B ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           768B ± 0%      624B ± 0%   -18.75%  (p=0.000 n=10+10)
/100-4        12.9kB ± 0%     8.9kB ± 0%   -31.23%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    80.7kB ± 0%   +61.38%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%    1106kB ± 0%  +127.01%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.84%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      4.00 ± 0%   -33.33%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      22.0 ± 0%   -12.00%  (p=0.000 n=10+10)
/100-4           209 ± 0%       208 ± 0%    -0.48%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%    +0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.31%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4          35.0ns ± 1%    39.0ns ±24%  +11.46%  (p=0.000 n=8+9)
/10-4          352ns ± 2%     430ns ±34%  +22.06%  (p=0.000 n=10+10)
/100-4        3.45µs ± 2%    3.61µs ± 3%   +4.87%  (p=0.000 n=10+10)
/1000-4       34.4µs ± 1%    36.9µs ± 4%   +7.10%  (p=0.000 n=9+9)
/10000-4       343µs ± 1%     371µs ± 3%   +8.13%  (p=0.000 n=9+9)
/100000-4     3.43ms ± 1%    3.70ms ± 4%   +7.97%  (p=0.000 n=10+10)
/1000000-4    34.3ms ± 1%    37.0ms ± 3%   +7.99%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs gammazero - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 2%    37.2ns ± 7%   +5.44%  (p=0.015 n=10+10)
/10-4          355ns ± 0%     358ns ± 3%     ~     (p=0.515 n=7+10)
/100-4        3.46µs ± 1%    3.55µs ± 7%     ~     (p=0.388 n=10+9)
/1000-4       34.6µs ± 1%    35.5µs ± 4%   +2.47%  (p=0.009 n=8+10)
/10000-4       406µs ±15%     351µs ± 4%  -13.61%  (p=0.003 n=10+10)
/100000-4     3.90ms ±12%    3.67ms ±13%   -5.98%  (p=0.035 n=10+9)
/1000000-4    38.3ms ±18%    35.5ms ± 5%     ~     (p=0.063 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs phf - FIFO queue
deque vs phf - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillPhfQueue.txt
name        old time/op    new time/op    delta
/0-4          39.9ns ± 9%    67.5ns ±10%   +69.32%  (p=0.000 n=10+10)
/1-4           142ns ± 1%     107ns ± 5%   -24.54%  (p=0.000 n=10+10)
/10-4          636ns ± 1%     869ns ±16%   +36.74%  (p=0.000 n=9+9)
/100-4        4.74µs ± 3%    6.55µs ±29%   +38.07%  (p=0.000 n=9+10)
/1000-4       43.0µs ±23%    53.0µs ±14%   +23.17%  (p=0.001 n=10+9)
/10000-4       450µs ±19%     577µs ±14%   +28.22%  (p=0.000 n=10+9)
/100000-4     4.24ms ± 4%    7.65ms ±10%   +80.43%  (p=0.000 n=10+9)
/1000000-4    46.8ms ± 1%    78.6ms ± 6%   +68.02%  (p=0.000 n=8+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            144B ± 0%       80B ± 0%   -44.44%  (p=0.000 n=10+10)
/10-4           608B ± 0%      832B ± 0%   +36.84%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    7.65kB ± 0%   +23.51%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    65.1kB ± 0%   +97.00%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.83%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.02%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.06%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      17.0 ± 0%   +13.33%  (p=0.000 n=10+10)
/100-4           107 ± 0%       113 ± 0%    +5.61%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.02k ± 0%    +0.59%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.56%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs phf - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillPhfQueue.txt
name       old time/op    new time/op    delta
/1-4         3.79µs ± 1%    3.87µs ± 1%    +2.18%  (p=0.000 n=9+10)
/10-4        37.8µs ± 7%    57.1µs ± 2%   +51.04%  (p=0.000 n=10+10)
/100-4        361µs ± 7%     581µs ± 6%   +60.78%  (p=0.000 n=9+10)
/1000-4      3.75ms ± 4%    4.79ms ± 4%   +27.96%  (p=0.000 n=10+10)
/10000-4     36.5ms ± 3%    54.1ms ± 3%   +48.33%  (p=0.000 n=10+10)
/100000-4     380ms ± 1%     721ms ± 3%   +89.57%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    54.4kB ± 0%  +240.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     736kB ± 0%  +359.99%  (p=0.000 n=8+10)
/1000-4      2.42MB ± 0%    6.48MB ± 0%  +168.10%  (p=0.000 n=10+10)
/10000-4     17.0MB ± 0%    94.6MB ± 0%  +455.02%  (p=0.000 n=8+10)
/100000-4     162MB ± 0%     789MB ± 0%  +387.87%  (p=0.000 n=8+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.00%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    +0.17%  (p=0.000 n=8+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.02%  (p=0.000 n=10+9)
```

deque vs phf - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullPhfQueue.txt
name       old time/op    new time/op    delta
/1-4         3.75µs ± 4%    3.94µs ± 2%    +5.05%  (p=0.000 n=10+9)
/10-4        37.8µs ± 4%    39.2µs ± 3%    +3.56%  (p=0.010 n=10+9)
/100-4        371µs ± 3%     384µs ± 5%    +3.54%  (p=0.015 n=10+10)
/1000-4      4.02ms ± 5%    3.86ms ± 8%    -4.08%  (p=0.023 n=10+10)
/10000-4     39.4ms ± 5%    38.8ms ± 2%      ~     (p=0.515 n=10+8)
/100000-4     392ms ± 3%     707ms ± 2%   +80.33%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      2.40MB ± 0%    1.60MB ± 0%   -33.41%  (p=0.000 n=10+10)
/10000-4     16.6MB ± 0%    16.0MB ± 0%    -3.89%  (p=0.000 n=10+10)
/100000-4     161MB ± 0%     632MB ± 0%  +292.50%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    -0.39%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    -0.03%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    -0.00%  (p=0.000 n=10+10)
```

deque vs phf - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreasePhfQueue.txt
name        old time/op    new time/op    delta
/1-4           244ns ± 1%     215ns ± 7%   -11.96%  (p=0.000 n=10+10)
/10-4         1.86µs ± 1%    1.36µs ± 7%   -26.64%  (p=0.000 n=10+10)
/100-4        8.02µs ± 1%    9.66µs ± 5%   +20.47%  (p=0.000 n=10+10)
/1000-4       73.5µs ± 1%    83.0µs ± 2%   +12.90%  (p=0.000 n=10+9)
/10000-4       725µs ± 1%     956µs ± 2%   +31.88%  (p=0.000 n=10+10)
/100000-4     8.20ms ± 0%   11.87ms ± 6%   +44.71%  (p=0.000 n=9+10)
/1000000-4    86.4ms ± 1%   129.3ms ± 4%   +49.67%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      128B ± 0%   -42.86%  (p=0.000 n=10+10)
/10-4         4.91kB ± 0%    0.99kB ± 0%   -79.80%  (p=0.000 n=10+10)
/100-4        7.79kB ± 0%    9.25kB ± 0%   +18.69%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    81.1kB ± 0%   +49.70%  (p=0.000 n=10+10)
/10000-4       491kB ± 0%    1106kB ± 0%  +125.17%  (p=0.000 n=10+10)
/100000-4     4.83MB ± 0%    9.49MB ± 0%   +96.68%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.83%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      27.0 ± 0%      ~     (all equal)
/100-4           207 ± 0%       213 ± 0%    +2.90%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.02k ± 0%    +0.20%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.30%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs phf - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreasePhfQueue.txt
name        old time/op    new time/op    delta
/1-4          40.8ns ± 3%    37.1ns ± 4%   -9.17%  (p=0.000 n=9+10)
/10-4          413ns ± 1%     372ns ± 2%  -10.01%  (p=0.000 n=8+9)
/100-4        4.05µs ± 5%    3.62µs ± 3%  -10.52%  (p=0.000 n=10+10)
/1000-4       42.7µs ± 9%    36.4µs ± 3%  -14.79%  (p=0.000 n=10+10)
/10000-4       436µs ± 6%     368µs ± 2%  -15.52%  (p=0.000 n=9+10)
/100000-4     3.92ms ± 2%    3.70ms ± 4%   -5.61%  (p=0.000 n=9+10)
/1000000-4    39.2ms ± 1%    37.4ms ± 3%   -4.66%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs phf - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStablePhfQueue.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    37.4ns ± 3%  +7.44%  (p=0.000 n=10+10)
/10-4          353ns ± 2%     379ns ± 5%  +7.39%  (p=0.000 n=10+9)
/100-4        3.45µs ± 1%    3.67µs ± 2%  +6.39%  (p=0.000 n=10+10)
/1000-4       34.5µs ± 1%    36.2µs ± 3%  +5.02%  (p=0.000 n=10+10)
/10000-4       346µs ± 2%     371µs ± 4%  +7.23%  (p=0.000 n=10+10)
/100000-4     3.43ms ± 1%    3.71ms ± 3%  +8.09%  (p=0.000 n=10+10)
/1000000-4    34.4ms ± 1%    37.5ms ± 4%  +9.00%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```


#### deque vs phf - LIFO stack

deque vs phf - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillPhfStack.txt
name        old time/op    new time/op    delta
/0-4          39.0ns ± 6%    69.4ns ± 8%   +78.08%  (p=0.000 n=9+10)
/1-4           145ns ± 2%     105ns ± 5%   -27.54%  (p=0.000 n=8+10)
/10-4          615ns ± 0%     852ns ± 2%   +38.52%  (p=0.000 n=8+9)
/100-4        4.42µs ± 1%    5.85µs ± 5%   +32.52%  (p=0.000 n=10+10)
/1000-4       36.2µs ± 1%    47.5µs ± 5%   +31.28%  (p=0.000 n=10+10)
/10000-4       383µs ± 7%     539µs ± 4%   +40.97%  (p=0.000 n=10+10)
/100000-4     4.10ms ± 9%    7.33ms ± 7%   +78.62%  (p=0.000 n=10+10)
/1000000-4    44.5ms ± 4%    81.9ms ± 6%   +84.00%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            144B ± 0%       80B ± 0%   -44.44%  (p=0.000 n=10+10)
/10-4           608B ± 0%      832B ± 0%   +36.84%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    7.65kB ± 0%   +23.51%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    65.1kB ± 0%   +97.00%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.83%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.02%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.06%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      17.0 ± 0%   +13.33%  (p=0.000 n=10+10)
/100-4           107 ± 0%       113 ± 0%    +5.61%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.02k ± 0%    +0.59%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.56%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillPhfStack.txt
name       old time/op    new time/op    delta
/1-4         3.62µs ± 1%    3.93µs ± 3%    +8.55%  (p=0.000 n=10+10)
/10-4        34.1µs ± 5%    57.8µs ± 3%   +69.55%  (p=0.000 n=10+9)
/100-4        332µs ± 2%     557µs ± 2%   +67.95%  (p=0.000 n=10+10)
/1000-4      3.28ms ± 2%    4.72ms ± 2%   +44.04%  (p=0.000 n=10+10)
/10000-4     34.0ms ± 3%    54.4ms ± 3%   +59.74%  (p=0.000 n=9+10)
/100000-4     370ms ± 5%     750ms ±10%  +102.96%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    54.4kB ± 0%  +240.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     736kB ± 0%  +360.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.48MB ± 0%  +304.79%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    94.6MB ± 0%  +491.16%  (p=0.000 n=10+9)
/100000-4     161MB ± 0%     789MB ± 0%  +391.53%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.40%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    +0.22%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.03%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullPhfStack.txt
name       old time/op    new time/op    delta
/1-4         3.90µs ± 6%    3.87µs ± 6%      ~     (p=0.388 n=10+9)
/10-4        34.6µs ± 2%    37.8µs ± 2%    +9.34%  (p=0.000 n=9+10)
/100-4        337µs ± 4%     372µs ± 4%   +10.36%  (p=0.000 n=9+10)
/1000-4      3.58ms ± 8%    3.69ms ± 2%      ~     (p=0.165 n=10+10)
/10000-4     34.9ms ± 7%    38.3ms ± 3%    +9.77%  (p=0.000 n=10+10)
/100000-4     339ms ± 2%     733ms ± 8%  +116.47%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%      ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%    -0.00%  (p=0.033 n=10+10)
/100000-4     160MB ± 0%     632MB ± 0%  +294.91%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%      ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.00%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreasePhfStack.txt
name        old time/op    new time/op    delta
/1-4           240ns ± 0%     212ns ± 2%   -11.97%  (p=0.000 n=9+9)
/10-4          929ns ± 1%    1274ns ± 4%   +37.21%  (p=0.000 n=10+10)
/100-4        8.65µs ± 1%    9.49µs ± 6%    +9.73%  (p=0.000 n=10+10)
/1000-4       66.7µs ± 1%    81.5µs ± 2%   +22.18%  (p=0.000 n=10+10)
/10000-4       666µs ± 1%     943µs ± 3%   +41.62%  (p=0.000 n=10+10)
/100000-4     7.78ms ± 1%   11.64ms ± 5%   +49.70%  (p=0.000 n=9+9)
/1000000-4    81.9ms ± 1%   138.1ms ±13%   +68.57%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      128B ± 0%   -42.86%  (p=0.000 n=10+10)
/10-4           768B ± 0%      992B ± 0%   +29.17%  (p=0.000 n=10+10)
/100-4        12.9kB ± 0%     9.2kB ± 0%   -28.38%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    81.1kB ± 0%   +62.11%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%    1106kB ± 0%  +127.09%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.85%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      27.0 ± 0%    +8.00%  (p=0.000 n=10+10)
/100-4           209 ± 0%       213 ± 0%    +1.91%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.02k ± 0%    +0.30%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.29%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreasePhfStack.txt
name        old time/op    new time/op    delta
/1-4          35.0ns ± 1%    36.7ns ± 3%   +4.76%  (p=0.000 n=8+9)
/10-4          352ns ± 2%     387ns ± 4%   +9.82%  (p=0.000 n=10+10)
/100-4        3.45µs ± 2%    3.79µs ± 5%   +9.93%  (p=0.000 n=10+10)
/1000-4       34.4µs ± 1%    37.4µs ± 4%   +8.56%  (p=0.000 n=9+10)
/10000-4       343µs ± 1%     378µs ± 4%  +10.33%  (p=0.000 n=9+9)
/100000-4     3.43ms ± 1%    3.77ms ± 6%  +10.11%  (p=0.000 n=10+9)
/1000000-4    34.3ms ± 1%    37.1ms ± 1%   +8.14%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs phf - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStablePhfStack.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 2%    39.2ns ±13%  +11.11%  (p=0.000 n=10+10)
/10-4          355ns ± 0%     370ns ± 8%   +4.38%  (p=0.000 n=7+10)
/100-4        3.46µs ± 1%    3.84µs ± 6%  +10.84%  (p=0.000 n=10+10)
/1000-4       34.6µs ± 1%    37.0µs ± 5%   +6.91%  (p=0.000 n=8+10)
/10000-4       406µs ±15%     365µs ± 6%  -10.06%  (p=0.015 n=10+10)
/100000-4     3.90ms ±12%    3.57ms ± 8%   -8.58%  (p=0.005 n=10+10)
/1000000-4    38.3ms ±18%    35.5ms ± 2%     ~     (p=0.105 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs cookiejar - FIFO queue

deque vs cookiejar - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillCookiejarQueue.txt
name        old time/op    new time/op     delta
/0-4          39.9ns ± 9%  10159.8ns ± 2%   +25388.59%  (p=0.000 n=10+8)
/1-4           142ns ± 1%    10443ns ±10%    +7244.02%  (p=0.000 n=10+10)
/10-4          636ns ± 1%    10646ns ±14%    +1574.72%  (p=0.000 n=9+9)
/100-4        4.74µs ± 3%    13.65µs ± 6%     +187.95%  (p=0.000 n=9+9)
/1000-4       43.0µs ±23%     42.2µs ±10%         ~     (p=0.353 n=10+10)
/10000-4       450µs ±19%      342µs ± 6%      -23.93%  (p=0.000 n=10+10)
/100000-4     4.24ms ± 4%     3.87ms ±14%       -8.57%  (p=0.002 n=10+10)
/1000000-4    46.8ms ± 1%     47.5ms ± 8%         ~     (p=0.965 n=8+10)

name        old alloc/op   new alloc/op    delta
/0-4           64.0B ± 0%   65680.0B ± 0%  +102525.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%     65696B ± 0%   +45522.22%  (p=0.000 n=10+10)
/10-4           608B ± 0%     65840B ± 0%   +10728.95%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    67.28kB ± 0%     +986.56%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%     81.7kB ± 0%     +147.34%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%      357kB ± 0%      +10.81%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.25MB ± 0%       +0.80%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%     32.8MB ± 0%       +1.99%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op   delta
/0-4            1.00 ± 0%       3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%       4.00 ± 0%         ~     (all equal)
/10-4           15.0 ± 0%       13.0 ± 0%      -13.33%  (p=0.000 n=10+10)
/100-4           107 ± 0%        103 ± 0%       -3.74%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%      1.00k ± 0%       -0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%      10.0k ± 0%       -0.75%  (p=0.000 n=10+10)
/100000-4       101k ± 0%       100k ± 0%       -0.73%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%      1.00M ± 0%       -0.73%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillCookiejarQueue.txt
name       old time/op    new time/op    delta
/1-4         3.79µs ± 1%    3.51µs ± 5%   -7.41%  (p=0.000 n=9+10)
/10-4        37.8µs ± 7%    32.7µs ± 6%  -13.45%  (p=0.000 n=10+10)
/100-4        361µs ± 7%     311µs ± 2%  -13.93%  (p=0.000 n=9+10)
/1000-4      3.75ms ± 4%    3.05ms ± 3%  -18.44%  (p=0.000 n=10+10)
/10000-4     36.5ms ± 3%    31.3ms ± 4%  -14.11%  (p=0.000 n=10+10)
/100000-4     380ms ± 1%     365ms ± 7%   -3.98%  (p=0.003 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%   +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%   +0.01%  (p=0.000 n=8+10)
/1000-4      2.42MB ± 0%    1.60MB ± 0%  -33.76%  (p=0.000 n=10+10)
/10000-4     17.0MB ± 0%    16.0MB ± 0%   -6.10%  (p=0.000 n=8+10)
/100000-4     162MB ± 0%     161MB ± 0%   -0.73%  (p=0.002 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.39%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%   -0.05%  (p=0.002 n=8+10)
/100000-4     10.0M ± 0%     10.0M ± 0%   -0.01%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullCookiejarQueue.txt
name       old time/op    new time/op    delta
/1-4         3.75µs ± 4%    3.13µs ± 3%  -16.52%  (p=0.000 n=10+9)
/10-4        37.8µs ± 4%    32.2µs ± 7%  -14.80%  (p=0.000 n=10+10)
/100-4        371µs ± 3%     316µs ± 3%  -14.76%  (p=0.000 n=10+9)
/1000-4      4.02ms ± 5%    3.13ms ± 0%  -22.16%  (p=0.000 n=10+7)
/10000-4     39.4ms ± 5%    31.6ms ± 3%  -19.89%  (p=0.000 n=10+10)
/100000-4     392ms ± 3%     366ms ± 7%   -6.49%  (p=0.002 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      2.40MB ± 0%    1.60MB ± 0%  -33.41%  (p=0.000 n=10+10)
/10000-4     16.6MB ± 0%    16.0MB ± 0%   -3.89%  (p=0.000 n=10+10)
/100000-4     161MB ± 0%     160MB ± 0%   -0.61%  (p=0.002 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.39%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%   -0.03%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%   -0.00%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4           244ns ± 1%   10208ns ± 2%   +4080.13%  (p=0.000 n=10+9)
/10-4         1.86µs ± 1%   10.86µs ± 2%    +484.16%  (p=0.000 n=10+10)
/100-4        8.02µs ± 1%   16.48µs ± 4%    +105.61%  (p=0.000 n=10+10)
/1000-4       73.5µs ± 1%    69.0µs ± 2%      -6.10%  (p=0.000 n=10+10)
/10000-4       725µs ± 1%     679µs ±10%      -6.37%  (p=0.002 n=10+10)
/100000-4     8.20ms ± 0%    8.12ms ± 5%        ~     (p=0.447 n=9+10)
/1000000-4    86.4ms ± 1%    83.7ms ± 3%      -3.13%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%    65712B ± 0%  +29235.71%  (p=0.000 n=10+10)
/10-4         4.91kB ± 0%   66.00kB ± 0%   +1243.65%  (p=0.000 n=10+10)
/100-4        7.79kB ± 0%   68.88kB ± 0%    +783.98%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    97.7kB ± 0%     +80.41%  (p=0.000 n=10+10)
/10000-4       491kB ± 0%     583kB ± 0%     +18.56%  (p=0.000 n=10+10)
/100000-4     4.83MB ± 0%    4.91MB ± 0%      +1.80%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.9MB ± 0%      +1.46%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%     -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      23.0 ± 0%     -14.81%  (p=0.000 n=10+10)
/100-4           207 ± 0%       203 ± 0%      -1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.00k ± 0%      -0.60%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%      -0.39%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%      -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%      -0.36%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4          40.8ns ± 3%    36.1ns ±12%  -11.57%  (p=0.000 n=9+10)
/10-4          413ns ± 1%     341ns ± 2%  -17.46%  (p=0.000 n=8+9)
/100-4        4.05µs ± 5%    3.32µs ± 2%  -17.90%  (p=0.000 n=10+10)
/1000-4       42.7µs ± 9%    33.4µs ± 2%  -21.90%  (p=0.000 n=10+10)
/10000-4       436µs ± 6%     332µs ± 1%  -23.92%  (p=0.000 n=9+9)
/100000-4     3.92ms ± 2%    3.30ms ± 2%  -15.83%  (p=0.000 n=9+9)
/1000000-4    39.2ms ± 1%    33.5ms ± 6%  -14.48%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs cookiejar - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    30.8ns ± 3%  -11.53%  (p=0.000 n=10+10)
/10-4          353ns ± 2%     309ns ± 4%  -12.42%  (p=0.000 n=10+10)
/100-4        3.45µs ± 1%    3.03µs ± 2%  -12.14%  (p=0.000 n=10+10)
/1000-4       34.5µs ± 1%    30.1µs ± 3%  -12.75%  (p=0.000 n=10+9)
/10000-4       346µs ± 2%     311µs ±12%   -9.98%  (p=0.000 n=10+10)
/100000-4     3.43ms ± 1%    3.05ms ± 4%  -11.01%  (p=0.000 n=10+10)
/1000000-4    34.4ms ± 1%    30.3ms ± 5%  -12.01%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs cookiejar - LIFO stack

deque vs cookiejar - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillCookiejarStack.txt
name        old time/op    new time/op     delta
/0-4          39.0ns ± 6%  10254.3ns ± 3%   +26208.15%  (p=0.000 n=9+9)
/1-4           145ns ± 2%    10843ns ± 7%    +7397.55%  (p=0.000 n=8+9)
/10-4          615ns ± 0%    11329ns ±12%    +1741.06%  (p=0.000 n=8+10)
/100-4        4.42µs ± 1%    13.91µs ± 8%     +214.88%  (p=0.000 n=10+10)
/1000-4       36.2µs ± 1%     43.8µs ± 5%      +21.18%  (p=0.000 n=10+9)
/10000-4       383µs ± 7%      344µs ± 2%      -10.01%  (p=0.000 n=10+10)
/100000-4     4.10ms ± 9%     3.73ms ± 5%       -8.98%  (p=0.000 n=10+10)
/1000000-4    44.5ms ± 4%     46.7ms ± 3%       +4.89%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op    delta
/0-4           64.0B ± 0%   65680.0B ± 0%  +102525.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%     65696B ± 0%   +45522.22%  (p=0.000 n=10+10)
/10-4           608B ± 0%     65840B ± 0%   +10728.95%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    67.28kB ± 0%     +986.56%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%     81.7kB ± 0%     +147.34%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%      357kB ± 0%      +10.81%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.25MB ± 0%       +0.80%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%     32.8MB ± 0%       +1.99%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op   delta
/0-4            1.00 ± 0%       3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%       4.00 ± 0%         ~     (all equal)
/10-4           15.0 ± 0%       13.0 ± 0%      -13.33%  (p=0.000 n=10+10)
/100-4           107 ± 0%        103 ± 0%       -3.74%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%      1.00k ± 0%       -0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%      10.0k ± 0%       -0.75%  (p=0.000 n=10+10)
/100000-4       101k ± 0%       100k ± 0%       -0.73%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%      1.00M ± 0%       -0.73%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillCookiejarStack.txt
name       old time/op    new time/op    delta
/1-4         3.62µs ± 1%    3.63µs ± 6%    ~     (p=0.780 n=10+10)
/10-4        34.1µs ± 5%    33.9µs ± 5%    ~     (p=0.565 n=10+10)
/100-4        332µs ± 2%     330µs ±11%    ~     (p=0.579 n=10+10)
/1000-4      3.28ms ± 2%    3.38ms ± 6%  +3.17%  (p=0.021 n=10+8)
/10000-4     34.0ms ± 3%    36.5ms ±16%  +7.26%  (p=0.022 n=9+10)
/100000-4     370ms ± 5%     400ms ±11%  +8.30%  (p=0.010 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%  +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%  +0.01%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%  +0.01%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    16.0MB ± 0%  +0.00%  (p=0.000 n=10+9)
/100000-4     161MB ± 0%     161MB ± 0%  +0.00%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%  -0.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%  -0.00%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullCookiejarStack.txt
name       old time/op    new time/op    delta
/1-4         3.90µs ± 6%    3.26µs ± 7%  -16.51%  (p=0.000 n=10+10)
/10-4        34.6µs ± 2%    33.1µs ± 9%   -4.44%  (p=0.035 n=9+10)
/100-4        337µs ± 4%     305µs ± 2%   -9.60%  (p=0.000 n=9+8)
/1000-4      3.58ms ± 8%    3.18ms ± 5%  -11.30%  (p=0.000 n=10+9)
/10000-4     34.9ms ± 7%    31.5ms ± 6%   -9.94%  (p=0.000 n=10+9)
/100000-4     339ms ± 2%     354ms ± 3%   +4.49%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%   -0.00%  (p=0.033 n=10+10)
/100000-4     160MB ± 0%     160MB ± 0%     ~     (all equal)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%     ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%     ~     (all equal)
```

deque vs cookiejar - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4           240ns ± 0%   10375ns ± 3%   +4214.88%  (p=0.000 n=9+10)
/10-4          929ns ± 1%   11122ns ± 3%   +1097.37%  (p=0.000 n=10+10)
/100-4        8.65µs ± 1%   16.70µs ± 4%     +93.11%  (p=0.000 n=10+9)
/1000-4       66.7µs ± 1%    71.1µs ± 3%      +6.61%  (p=0.000 n=10+10)
/10000-4       666µs ± 1%     643µs ± 2%      -3.38%  (p=0.000 n=10+10)
/100000-4     7.78ms ± 1%    8.24ms ± 5%      +5.98%  (p=0.000 n=9+10)
/1000000-4    81.9ms ± 1%    89.0ms ± 3%      +8.58%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%    65712B ± 0%  +29235.71%  (p=0.000 n=10+10)
/10-4           768B ± 0%    66000B ± 0%   +8493.75%  (p=0.000 n=10+10)
/100-4        12.9kB ± 0%    68.9kB ± 0%    +433.46%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    97.7kB ± 0%     +95.36%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     517kB ± 0%      +6.10%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.85MB ± 0%      +0.52%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.8MB ± 0%      +1.32%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%     -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      23.0 ± 0%      -8.00%  (p=0.000 n=10+10)
/100-4           209 ± 0%       203 ± 0%      -2.87%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.00k ± 0%      -0.50%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%      -0.39%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%      -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%      -0.36%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4          35.0ns ± 1%    35.9ns ±10%    ~     (p=0.779 n=8+10)
/10-4          352ns ± 2%     378ns ±17%  +7.18%  (p=0.000 n=10+10)
/100-4        3.45µs ± 2%    3.63µs ±13%  +5.37%  (p=0.043 n=10+10)
/1000-4       34.4µs ± 1%    34.6µs ± 2%    ~     (p=0.236 n=9+8)
/10000-4       343µs ± 1%     344µs ± 1%    ~     (p=0.136 n=9+9)
/100000-4     3.43ms ± 1%    3.50ms ± 5%  +2.14%  (p=0.043 n=10+10)
/1000000-4    34.3ms ± 1%    34.4ms ± 2%    ~     (p=0.739 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs cookiejar - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 2%    30.8ns ± 4%  -12.78%  (p=0.000 n=10+10)
/10-4          355ns ± 0%     304ns ± 2%  -14.26%  (p=0.000 n=7+8)
/100-4        3.46µs ± 1%    2.99µs ± 2%  -13.82%  (p=0.000 n=10+9)
/1000-4       34.6µs ± 1%    30.5µs ± 9%  -11.71%  (p=0.000 n=8+10)
/10000-4       406µs ±15%     300µs ± 4%  -26.13%  (p=0.000 n=10+10)
/100000-4     3.90ms ±12%    2.94ms ± 2%  -24.74%  (p=0.000 n=10+10)
/1000000-4    38.3ms ±18%    30.1ms ± 3%  -21.39%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```
