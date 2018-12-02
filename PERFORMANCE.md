# Performance

Below compares the deque [benchmark tests](BENCHMARK_TESTS.md) results with the other tested queues.

## Running the Tests
In the "testdata" directory, we have included the result of local test runs for all queues. Below uses this run to compare the queues, but it's possible and we highly encourage you to run the tests yourself to help validate the results.

To run the tests locally, clone the deque repo, cd to the deque main directory and run below command.

```
go test -benchmem -timeout 60m -bench=. -run=^$
```

This command will run all tests for all queues locally once. This should be good enouh to give you a sense of the queues performance, but to
do a proper comparison, elimating test variations, we recommend you to run the tests as detailed [here](BENCHMARK_TESTS.md) by running the tests with multiple counts, splitting the files with [test-splitter](https://github.com/ef-ds/tools/tree/master/testsplitter) and using the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool to aggregate the results.


## Bottom Line
As a general purpose double-ended queue, deque is the queue that displays the most balanced performance, performing either very competitively or besting all other queues in all the different test scenarios.

Having said that, the tests show there's some room for improvements. We're actively working on the deque. We expect to release better performant deque versions in the near future.

## Recommendations
Using a double-ended queue as a stack is possible and works very well. However, given the stack inverted properties (LIFO) when comparing to a FIFO queue, using a deque as a stack is not the most efficient solution.

For a stack solution, we recommend building a stack using a simple slice, such as the [CustomSliceQueue](testdata_test.go).

For all other uses, we recommend using the deque as it performs very well on both low and high load scenarios on all tests.


## Results

Given the enormous amount of test data, it can be difficult and time consuming to find out the net impact of all the tests,
so we generally spend most of the time on the results of the, arguably, most important test: the Microservice test.

Note: Below results is for deque's latest release, not the current code on the master branch, which might contain experimental changes that has the potential to affect performance and efficiency significantly (for the better or worse).

For a list of all released versions and their impact on performance and efficiency, please refer to [CHANGELOG.md](CHANGELOG.md).


### Microservice Test Results
deque vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt testdata/BenchmarkMicroserviceImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4           510ns ± 5%     646ns ±10%  +26.83%  (p=0.000 n=10+9)
/10-4         3.60µs ± 5%    4.81µs ± 5%  +33.60%  (p=0.000 n=10+10)
/100-4        24.8µs ± 7%    32.3µs ± 3%  +29.99%  (p=0.000 n=10+10)
/1000-4        237µs ± 5%     313µs ± 5%  +32.06%  (p=0.000 n=9+10)
/10000-4      2.38ms ± 5%    3.17ms ±11%  +33.50%  (p=0.000 n=10+10)
/100000-4     25.0ms ± 1%    33.1ms ± 2%  +32.73%  (p=0.000 n=10+8)
/1000000-4     268ms ± 4%     348ms ± 8%  +29.60%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            544B ± 0%      432B ± 0%  -20.59%  (p=0.000 n=10+10)
/10-4         5.70kB ± 0%    6.91kB ± 0%  +21.35%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    29.6kB ± 0%  +87.83%  (p=0.000 n=10+10)
/1000-4        133kB ± 0%     261kB ± 0%  +96.01%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.58MB ± 0%  +80.18%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    25.8MB ± 0%  +78.53%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     258MB ± 0%  +78.37%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      17.0 ± 0%  +41.67%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%     109.0 ± 0%  +41.56%  (p=0.000 n=10+10)
/100-4           707 ± 0%       927 ± 0%  +31.12%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     9.11k ± 0%  +29.88%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     91.0k ± 0%  +29.65%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      909k ± 0%  +29.62%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     9.09M ± 0%  +29.62%  (p=0.000 n=10+10)
```

deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt testdata/BenchmarkMicroserviceListQueue.txt
name        old time/op    new time/op    delta
/1-4           510ns ± 5%     511ns ± 6%      ~     (p=0.983 n=10+9)
/10-4         3.60µs ± 5%    4.81µs ± 2%   +33.44%  (p=0.000 n=10+10)
/100-4        24.8µs ± 7%    46.8µs ± 1%   +88.32%  (p=0.000 n=10+8)
/1000-4        237µs ± 5%     500µs ± 5%  +110.97%  (p=0.000 n=9+8)
/10000-4      2.38ms ± 5%    5.39ms ± 4%  +126.89%  (p=0.000 n=10+9)
/100000-4     25.0ms ± 1%    75.0ms ± 8%  +200.57%  (p=0.000 n=10+10)
/1000000-4     268ms ± 4%     815ms ±13%  +203.81%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            544B ± 0%      496B ± 0%    -8.82%  (p=0.000 n=10+10)
/10-4         5.70kB ± 0%    4.53kB ± 0%   -20.51%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    44.8kB ± 0%  +184.28%  (p=0.000 n=10+10)
/1000-4        133kB ± 0%     448kB ± 0%  +236.49%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    4.48MB ± 0%  +213.02%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    44.8MB ± 0%  +210.58%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     448MB ± 0%  +210.33%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      15.0 ± 0%   +25.00%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%     141.0 ± 0%   +83.12%  (p=0.000 n=10+10)
/100-4           707 ± 0%      1401 ± 0%   +98.16%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.59%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%    140.0k ± 0%   +99.56%  (p=0.000 n=10+10)
/100000-4       702k ± 0%     1400k ± 0%   +99.56%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%    14.00M ± 0%   +99.55%  (p=0.000 n=10+10)
```

deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.2.txt testdata/BenchmarkMicroserviceListStack.txt
name        old time/op    new time/op    delta
/1-4           389ns ± 1%     510ns ± 4%   +31.32%  (p=0.000 n=10+9)
/10-4         2.51µs ± 1%    4.87µs ± 6%   +93.88%  (p=0.000 n=10+9)
/100-4        22.8µs ± 1%    47.1µs ± 2%  +106.20%  (p=0.000 n=10+8)
/1000-4        220µs ± 2%     488µs ±14%  +122.15%  (p=0.000 n=9+10)
/10000-4      2.28ms ± 2%    5.28ms ± 7%  +131.46%  (p=0.000 n=8+9)
/100000-4     24.6ms ± 1%    75.2ms ± 3%  +206.06%  (p=0.000 n=9+9)
/1000000-4     258ms ± 1%     941ms ±12%  +264.10%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            288B ± 0%      496B ± 0%   +72.22%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    4.53kB ± 0%  +191.75%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    44.8kB ± 0%  +184.28%  (p=0.000 n=10+10)
/1000-4        129kB ± 0%     448kB ± 0%  +247.30%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    4.48MB ± 0%  +213.02%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    44.8MB ± 0%  +210.67%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     448MB ± 0%  +210.33%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      15.0 ± 0%   +36.36%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     141.0 ± 0%   +88.00%  (p=0.000 n=10+10)
/100-4           707 ± 0%      1401 ± 0%   +98.16%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.64%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%    140.0k ± 0%   +99.56%  (p=0.000 n=10+10)
/100000-4       702k ± 0%     1400k ± 0%   +99.56%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%    14.00M ± 0%   +99.55%  (p=0.000 n=10+10)
```

deque vs [CustomSliceQueue](testdata_test.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt testdata/BenchmarkMicroserviceSliceQueue.txt
name        old time/op    new time/op    delta
/1-4           510ns ± 5%     475ns ± 2%    -6.76%  (p=0.000 n=10+9)
/10-4         3.60µs ± 5%    3.54µs ± 2%      ~     (p=0.353 n=10+10)
/100-4        24.8µs ± 7%    27.0µs ± 5%    +8.55%  (p=0.000 n=10+10)
/1000-4        237µs ± 5%     278µs ± 2%   +17.26%  (p=0.000 n=9+10)
/10000-4      2.38ms ± 5%    3.08ms ± 2%   +29.56%  (p=0.000 n=10+10)
/100000-4     25.0ms ± 1%    47.4ms ± 3%   +89.80%  (p=0.000 n=10+10)
/1000000-4     268ms ± 4%     596ms ± 7%  +122.13%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            544B ± 0%      232B ± 0%   -57.35%  (p=0.000 n=10+10)
/10-4         5.70kB ± 0%    2.17kB ± 0%   -61.94%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    21.3kB ± 0%   +35.24%  (p=0.000 n=10+10)
/1000-4        133kB ± 0%     214kB ± 0%   +61.07%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.95MB ± 0%  +105.95%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    33.1MB ± 0%  +129.74%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     338MB ± 0%  +134.04%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      14.0 ± 0%   +16.67%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%     101.0 ± 0%   +31.17%  (p=0.000 n=10+10)
/100-4           707 ± 0%       822 ± 0%   +16.27%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     8.77k ± 0%   +24.96%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     87.8k ± 0%   +25.20%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      875k ± 0%   +24.79%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     8.86M ± 0%   +26.32%  (p=0.000 n=10+10)
```

deque vs [CustomSliceQueue](testdata_test.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.2.txt testdata/BenchmarkMicroserviceSliceStack.txt
name        old time/op    new time/op    delta
/1-4           389ns ± 1%     366ns ± 2%   -5.94%  (p=0.000 n=10+8)
/10-4         2.51µs ± 1%    2.15µs ± 2%  -14.36%  (p=0.000 n=10+9)
/100-4        22.8µs ± 1%    18.4µs ± 1%  -19.48%  (p=0.000 n=10+10)
/1000-4        220µs ± 2%     177µs ± 1%  -19.69%  (p=0.000 n=9+10)
/10000-4      2.28ms ± 2%    1.88ms ± 2%  -17.32%  (p=0.000 n=8+10)
/100000-4     24.6ms ± 1%    25.7ms ± 4%   +4.66%  (p=0.000 n=9+10)
/1000000-4     258ms ± 1%     283ms ± 5%   +9.46%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            288B ± 0%      200B ± 0%  -30.56%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    1.40kB ± 0%   -9.79%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    13.3kB ± 0%  -15.87%  (p=0.000 n=10+10)
/1000-4        129kB ± 0%     128kB ± 0%   -0.47%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.51MB ± 0%   +5.25%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    15.9MB ± 0%   +9.94%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     157MB ± 0%   +8.89%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      11.0 ± 0%     ~     (all equal)
/10-4           75.0 ± 0%      76.0 ± 0%   +1.33%  (p=0.000 n=10+10)
/100-4           707 ± 0%       709 ± 0%   +0.28%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.01k ± 0%   -0.01%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.19%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.22%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+10)
```

deque vs [phf](https://github.com/phf/go-queue) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt testdata/BenchmarkMicroservicePhfQueue.txt
name        old time/op    new time/op    delta
/1-4           510ns ± 5%     444ns ± 2%  -12.81%  (p=0.000 n=10+10)
/10-4         3.60µs ± 5%    3.39µs ± 7%   -6.02%  (p=0.002 n=10+10)
/100-4        24.8µs ± 7%    29.5µs ± 6%  +18.86%  (p=0.000 n=10+10)
/1000-4        237µs ± 5%     260µs ± 2%   +9.73%  (p=0.000 n=9+9)
/10000-4      2.38ms ± 5%    2.79ms ± 3%  +17.35%  (p=0.000 n=10+10)
/100000-4     25.0ms ± 1%    33.6ms ± 4%  +34.74%  (p=0.000 n=10+10)
/1000000-4     268ms ± 4%     361ms ± 9%  +34.46%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            544B ± 0%      272B ± 0%  -50.00%  (p=0.000 n=10+10)
/10-4         5.70kB ± 0%    2.18kB ± 0%  -61.80%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    23.0kB ± 0%  +45.84%  (p=0.000 n=10+10)
/1000-4        133kB ± 0%     210kB ± 0%  +57.58%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.69MB ± 0%  +88.12%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    23.8MB ± 0%  +64.88%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     213MB ± 0%  +47.31%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      11.0 ± 0%   -8.33%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%      79.0 ± 0%   +2.60%  (p=0.000 n=10+10)
/100-4           707 ± 0%       721 ± 0%   +1.98%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%   +0.26%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.15%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+10)
```

deque vs [phf](https://github.com/phf/go-queue) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.2.txt testdata/BenchmarkMicroservicePhfStack.txt
name        old time/op    new time/op    delta
/1-4           389ns ± 1%     458ns ± 9%  +17.91%  (p=0.000 n=10+10)
/10-4         2.51µs ± 1%    3.22µs ± 2%  +28.12%  (p=0.000 n=10+9)
/100-4        22.8µs ± 1%    29.0µs ± 4%  +27.09%  (p=0.000 n=10+10)
/1000-4        220µs ± 2%     265µs ± 5%  +20.35%  (p=0.000 n=9+10)
/10000-4      2.28ms ± 2%    2.90ms ± 3%  +27.26%  (p=0.000 n=8+9)
/100000-4     24.6ms ± 1%    36.0ms ± 3%  +46.37%  (p=0.000 n=9+7)
/1000000-4     258ms ± 1%     359ms ± 7%  +39.01%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            288B ± 0%      272B ± 0%   -5.56%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    2.18kB ± 0%  +40.21%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    23.0kB ± 0%  +45.84%  (p=0.000 n=10+10)
/1000-4        129kB ± 0%     210kB ± 0%  +62.64%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.69MB ± 0%  +88.12%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    23.8MB ± 0%  +64.92%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     213MB ± 0%  +47.31%  (p=0.000 n=9+9)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      11.0 ± 0%     ~     (all equal)
/10-4           75.0 ± 0%      79.0 ± 0%   +5.33%  (p=0.000 n=10+10)
/100-4           707 ± 0%       721 ± 0%   +1.98%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%   +0.29%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.15%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+10)
```

deque vs [gammazero](https://github.com/gammazero/deque) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt testdata/BenchmarkMicroserviceGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4           510ns ± 5%     379ns ± 8%  -25.58%  (p=0.000 n=10+10)
/10-4         3.60µs ± 5%    2.57µs ±11%  -28.62%  (p=0.000 n=10+9)
/100-4        24.8µs ± 7%    27.0µs ± 3%   +8.65%  (p=0.000 n=10+10)
/1000-4        237µs ± 5%     251µs ± 2%   +5.89%  (p=0.000 n=9+9)
/10000-4      2.38ms ± 5%    2.75ms ±18%  +15.69%  (p=0.000 n=10+10)
/100000-4     25.0ms ± 1%    33.3ms ±10%  +33.23%  (p=0.000 n=10+9)
/1000000-4     268ms ± 4%     341ms ± 5%  +27.01%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            544B ± 0%      416B ± 0%  -23.53%  (p=0.000 n=10+10)
/10-4         5.70kB ± 0%    1.42kB ± 0%  -75.00%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    22.3kB ± 0%  +41.08%  (p=0.000 n=10+10)
/1000-4        133kB ± 0%     209kB ± 0%  +57.02%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.69MB ± 0%  +88.06%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    23.8MB ± 0%  +64.87%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     213MB ± 0%  +47.31%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%       9.0 ± 0%  -25.00%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%      72.0 ± 0%   -6.49%  (p=0.000 n=10+10)
/100-4           707 ± 0%       714 ± 0%   +0.99%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%   +0.16%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.16%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+10)
```

deque vs [gammazero](https://github.com/gammazero/deque) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.2.txt testdata/BenchmarkMicroserviceGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4           389ns ± 1%     375ns ± 4%   -3.64%  (p=0.001 n=10+9)
/10-4         2.51µs ± 1%    2.56µs ± 5%     ~     (p=1.000 n=10+10)
/100-4        22.8µs ± 1%    27.6µs ± 7%  +20.83%  (p=0.000 n=10+10)
/1000-4        220µs ± 2%     270µs ± 4%  +22.77%  (p=0.000 n=9+9)
/10000-4      2.28ms ± 2%    2.67ms ± 2%  +17.10%  (p=0.000 n=8+9)
/100000-4     24.6ms ± 1%    33.5ms ±19%  +36.28%  (p=0.000 n=9+10)
/1000000-4     258ms ± 1%     342ms ± 6%  +32.27%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            288B ± 0%      416B ± 0%  +44.44%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    1.42kB ± 0%   -8.25%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    22.3kB ± 0%  +41.08%  (p=0.000 n=10+10)
/1000-4        129kB ± 0%     209kB ± 0%  +62.06%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.69MB ± 0%  +88.06%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    23.8MB ± 0%  +64.92%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     213MB ± 0%  +47.31%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%       9.0 ± 0%  -18.18%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      72.0 ± 0%   -4.00%  (p=0.000 n=10+10)
/100-4           707 ± 0%       714 ± 0%   +0.99%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%   +0.19%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.16%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+10)
```

deque vs [juju](https://github.com/juju/utils/blob/master/deque/deque.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt testdata/BenchmarkMicroserviceJujuQueue.txt
name        old time/op    new time/op    delta
/1-4           510ns ± 5%     607ns ± 1%   +19.18%  (p=0.000 n=10+9)
/10-4         3.60µs ± 5%    3.12µs ± 1%   -13.51%  (p=0.000 n=10+10)
/100-4        24.8µs ± 7%    26.8µs ± 1%    +7.82%  (p=0.000 n=10+10)
/1000-4        237µs ± 5%     265µs ± 1%   +11.76%  (p=0.000 n=9+10)
/10000-4      2.38ms ± 5%    2.68ms ± 1%   +12.86%  (p=0.000 n=10+9)
/100000-4     25.0ms ± 1%    31.0ms ± 3%   +24.13%  (p=0.000 n=10+10)
/1000000-4     268ms ± 4%     340ms ± 4%   +26.66%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            544B ± 0%     1296B ± 0%  +138.24%  (p=0.000 n=10+10)
/10-4         5.70kB ± 0%    3.41kB ± 0%   -40.17%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    22.3kB ± 0%   +41.48%  (p=0.000 n=10+10)
/1000-4        133kB ± 0%     217kB ± 0%   +62.94%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.16MB ± 0%   +50.69%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    21.6MB ± 0%   +49.41%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     216MB ± 0%   +49.28%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      11.0 ± 0%    -8.33%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%      77.0 ± 0%      ~     (all equal)
/100-4           707 ± 0%       731 ± 0%    +3.39%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.29k ± 0%    +3.86%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     72.8k ± 0%    +3.80%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      728k ± 0%    +3.79%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.28M ± 0%    +3.79%  (p=0.000 n=10+10)
```

deque vs [juju](https://github.com/juju/utils/blob/master/deque/deque.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.2.txt testdata/BenchmarkMicroserviceJujuStack.txt
name        old time/op    new time/op    delta
/1-4           389ns ± 1%     608ns ± 1%   +56.29%  (p=0.000 n=10+10)
/10-4         2.51µs ± 1%    2.75µs ± 1%    +9.46%  (p=0.000 n=10+10)
/100-4        22.8µs ± 1%    26.2µs ± 1%   +14.71%  (p=0.000 n=10+10)
/1000-4        220µs ± 2%     252µs ± 1%   +14.73%  (p=0.000 n=9+10)
/10000-4      2.28ms ± 2%    2.56ms ± 0%   +12.33%  (p=0.000 n=8+10)
/100000-4     24.6ms ± 1%    30.0ms ± 2%   +21.86%  (p=0.000 n=9+10)
/1000000-4     258ms ± 1%     340ms ± 5%   +31.65%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            288B ± 0%     1296B ± 0%  +350.00%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    2.30kB ± 0%   +48.45%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    21.2kB ± 0%   +34.48%  (p=0.000 n=10+10)
/1000-4        129kB ± 0%     184kB ± 0%   +42.50%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.81MB ± 0%   +26.47%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    18.1MB ± 0%   +25.53%  (p=0.000 n=10+8)
/1000000-4     144MB ± 0%     181MB ± 0%   +25.38%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      11.0 ± 0%      ~     (all equal)
/10-4           75.0 ± 0%      74.0 ± 0%    -1.33%  (p=0.000 n=10+10)
/100-4           707 ± 0%       728 ± 0%    +2.97%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.20k ± 0%    +2.61%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     71.9k ± 0%    +2.45%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      719k ± 0%    +2.45%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.19M ± 0%    +2.45%  (p=0.000 n=10+10)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt testdata/BenchmarkMicroserviceCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4           510ns ± 5%   10506ns ± 2%   +1961.23%  (p=0.000 n=10+10)
/10-4         3.60µs ± 5%   12.85µs ± 7%    +256.53%  (p=0.000 n=10+10)
/100-4        24.8µs ± 7%    31.3µs ± 3%     +25.88%  (p=0.000 n=10+10)
/1000-4        237µs ± 5%     226µs ± 7%      -4.46%  (p=0.004 n=9+10)
/10000-4      2.38ms ± 5%    2.08ms ± 4%     -12.64%  (p=0.000 n=10+10)
/100000-4     25.0ms ± 1%    24.3ms ± 4%      -2.52%  (p=0.009 n=10+10)
/1000000-4     268ms ± 4%     242ms ± 4%      -9.72%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            544B ± 0%    65792B ± 0%  +11994.12%  (p=0.000 n=10+10)
/10-4         5.70kB ± 0%   66.80kB ± 0%   +1072.75%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    76.9kB ± 0%    +387.32%  (p=0.000 n=10+10)
/1000-4        133kB ± 0%     243kB ± 0%     +82.70%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.38MB ± 0%      -3.40%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    12.9MB ± 0%     -10.48%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     129MB ± 0%     -10.71%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      10.0 ± 0%     -16.67%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%      73.0 ± 0%      -5.19%  (p=0.000 n=10+10)
/100-4           707 ± 0%       703 ± 0%      -0.57%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.00k ± 0%      -0.14%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%      -0.21%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%      -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%      -0.22%  (p=0.000 n=10+10)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.2.txt testdata/BenchmarkMicroserviceCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4           389ns ± 1%   10783ns ± 8%   +2674.09%  (p=0.000 n=10+10)
/10-4         2.51µs ± 1%   13.09µs ± 8%    +421.05%  (p=0.000 n=10+10)
/100-4        22.8µs ± 1%    30.6µs ± 4%     +34.26%  (p=0.000 n=10+8)
/1000-4        220µs ± 2%     220µs ± 6%        ~     (p=0.842 n=9+10)
/10000-4      2.28ms ± 2%    2.07ms ± 5%      -9.37%  (p=0.000 n=8+10)
/100000-4     24.6ms ± 1%    24.0ms ± 5%      -2.27%  (p=0.028 n=9+10)
/1000000-4     258ms ± 1%     250ms ± 5%      -3.14%  (p=0.028 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            288B ± 0%    65792B ± 0%  +22744.44%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%   66.80kB ± 0%   +4204.12%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    76.9kB ± 0%    +387.32%  (p=0.000 n=10+10)
/1000-4        129kB ± 0%     178kB ± 0%     +37.73%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.32MB ± 0%      -7.99%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    12.8MB ± 0%     -10.91%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     129MB ± 0%     -10.76%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      10.0 ± 0%      -9.09%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      73.0 ± 0%      -2.67%  (p=0.000 n=10+10)
/100-4           707 ± 0%       703 ± 0%      -0.57%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.00k ± 0%      -0.14%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%      -0.21%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%      -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%      -0.22%  (p=0.000 n=10+10)
```


### Other Test Results
#### deque vs impl7 - FIFO queue
deque vs impl7 - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.2.txt testdata/BenchmarkFillImpl7Queue.txt
name        old time/op    new time/op    delta
/0-4          36.1ns ± 4%    35.9ns ± 3%     ~     (p=0.642 n=10+9)
/1-4           143ns ± 5%     133ns ± 1%   -7.00%  (p=0.000 n=10+10)
/10-4          664ns ± 9%     764ns ± 7%  +15.04%  (p=0.000 n=10+9)
/100-4        4.83µs ± 7%    4.28µs ± 3%  -11.29%  (p=0.000 n=10+9)
/1000-4       42.4µs ±15%    38.8µs ± 7%   -8.59%  (p=0.002 n=10+10)
/10000-4       417µs ±11%     388µs ± 5%   -6.91%  (p=0.043 n=10+10)
/100000-4     4.45ms ±27%    3.95ms ± 2%  -11.31%  (p=0.001 n=10+8)
/1000000-4    51.9ms ±12%    45.9ms ± 4%  -11.50%  (p=0.001 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%     ~     (all equal)
/1-4            128B ± 0%      112B ± 0%  -12.50%  (p=0.000 n=10+10)
/10-4           592B ± 0%      736B ± 0%  +24.32%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    4.26kB ± 0%  -31.09%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    33.2kB ± 0%   +0.63%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     323kB ± 0%   +0.24%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.23MB ± 0%   +0.20%  (p=0.000 n=9+10)
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
benchstat testdata/BenchmarkRefillDequeQueuev1.0.2.txt testdata/BenchmarkRefillImpl7Queue.txt
name       old time/op    new time/op    delta
/1-4         3.97µs ± 6%   10.03µs ± 6%  +152.80%  (p=0.000 n=9+10)
/10-4        36.8µs ± 5%    74.7µs ± 5%  +103.13%  (p=0.000 n=10+10)
/100-4        347µs ± 5%     442µs ± 4%   +27.16%  (p=0.000 n=10+10)
/1000-4      3.67ms ±15%    3.97ms ± 3%    +8.39%  (p=0.002 n=10+9)
/10000-4     39.7ms ± 5%    39.1ms ± 6%      ~     (p=0.218 n=10+10)
/100000-4     413ms ± 6%     421ms ±24%      ~     (p=0.730 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    68.8kB ± 0%  +330.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     421kB ± 0%  +163.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    3.32MB ± 0%  +107.29%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    32.3MB ± 0%    +7.17%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     323MB ± 0%    +0.90%  (p=0.000 n=9+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.60k ± 0%   +60.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +2.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.02M ± 0%    +0.93%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.79%  (p=0.000 n=9+10)
```

deque vs impl7 - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt testdata/BenchmarkRefillFullImpl7Queue.txt
name       old time/op    new time/op    delta
/1-4         3.70µs ± 8%    4.24µs ± 4%   +14.69%  (p=0.000 n=10+10)
/10-4        37.4µs ± 7%    43.3µs ±24%   +15.92%  (p=0.000 n=10+10)
/100-4        360µs ± 4%     398µs ± 5%   +10.65%  (p=0.000 n=10+10)
/1000-4      3.50ms ± 3%    4.01ms ± 7%   +14.75%  (p=0.000 n=9+10)
/10000-4     40.0ms ± 5%    39.9ms ± 7%      ~     (p=0.605 n=9+9)
/100000-4     426ms ±11%     412ms ±15%      ~     (p=0.133 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     322kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.23MB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    32.3MB ± 0%    +7.09%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     322MB ± 0%    +0.84%  (p=0.000 n=8+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       101 ± 0%    +1.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.01k ± 0%    +1.50%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.2k ± 0%    +1.56%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +1.56%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.02M ± 0%    +0.88%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.79%  (p=0.000 n=9+10)
```

deque vs impl7 - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowIncreaseImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4           248ns ± 5%     230ns ± 3%   -7.34%  (p=0.000 n=9+10)
/10-4         2.09µs ±13%    1.61µs ± 6%  -22.92%  (p=0.000 n=9+10)
/100-4        8.06µs ± 5%    7.87µs ± 2%     ~     (p=0.156 n=10+9)
/1000-4       73.5µs ± 5%    71.5µs ± 1%   -2.64%  (p=0.031 n=9+9)
/10000-4       703µs ± 1%     723µs ± 3%   +2.80%  (p=0.001 n=8+9)
/100000-4     8.45ms ± 8%    8.50ms ± 9%     ~     (p=0.739 n=10+10)
/1000000-4    90.6ms ± 6%    93.4ms ± 4%     ~     (p=0.113 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      160B ± 0%  -23.08%  (p=0.000 n=10+10)
/10-4         4.90kB ± 0%    2.98kB ± 0%  -39.22%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    7.94kB ± 0%   +2.06%  (p=0.000 n=10+10)
/1000-4       53.2kB ± 0%    65.9kB ± 0%  +23.90%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     647kB ± 0%  +33.10%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    6.45MB ± 0%  +33.72%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    64.5MB ± 0%  +33.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%     ~     (all equal)
/10-4           27.0 ± 0%      29.0 ± 0%   +7.41%  (p=0.000 n=10+10)
/100-4           207 ± 0%       211 ± 0%   +1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.04k ± 0%   +1.19%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.3k ± 0%   +1.18%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      203k ± 0%   +1.17%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.03M ± 0%   +1.17%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowDecreaseImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 3%    99.0ns ±12%  +180.64%  (p=0.000 n=10+10)
/10-4          353ns ± 2%    1013ns ± 8%  +187.04%  (p=0.000 n=9+10)
/100-4        3.61µs ±13%    9.67µs ± 5%  +168.09%  (p=0.000 n=9+9)
/1000-4       35.2µs ± 2%    93.4µs ± 3%  +165.60%  (p=0.000 n=8+8)
/10000-4       356µs ± 7%     983µs ±12%  +175.98%  (p=0.000 n=9+10)
/100000-4     3.51ms ± 3%   10.70ms ±22%  +205.00%  (p=0.000 n=10+10)
/1000000-4    34.3ms ± 2%    92.9ms ± 6%  +170.39%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.2.txt testdata/BenchmarkStableImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4          35.1ns ± 2%    38.1ns ± 5%    +8.55%  (p=0.000 n=9+9)
/10-4          358ns ± 4%     391ns ± 7%    +9.17%  (p=0.000 n=9+10)
/100-4        3.44µs ± 2%    4.13µs ±19%   +20.14%  (p=0.000 n=9+10)
/1000-4       34.2µs ± 2%    39.6µs ±12%   +15.95%  (p=0.000 n=9+10)
/10000-4       357µs ± 6%     385µs ± 4%    +7.84%  (p=0.000 n=10+10)
/100000-4     3.49ms ± 4%    3.85ms ±11%   +10.29%  (p=0.000 n=10+10)
/1000000-4    35.6ms ± 3%    38.3ms ± 6%    +7.43%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     32.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      322B ± 0%  +101.25%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     322kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    3.23MB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    32.3MB ± 0%  +101.56%  (p=0.000 n=8+10)

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
benchstat testdata/BenchmarkFillDequeQueuev1.0.2.txt testdata/BenchmarkFillListQueue.txt
name        old time/op    new time/op    delta
/0-4          36.1ns ± 4%    39.1ns ± 2%    +8.31%  (p=0.000 n=10+9)
/1-4           143ns ± 5%     107ns ± 1%   -25.30%  (p=0.000 n=10+9)
/10-4          664ns ± 9%     726ns ± 1%    +9.30%  (p=0.000 n=10+10)
/100-4        4.83µs ± 7%    6.74µs ± 1%   +39.56%  (p=0.000 n=10+9)
/1000-4       42.4µs ±15%    69.1µs ± 1%   +62.84%  (p=0.000 n=10+8)
/10000-4       417µs ±11%     712µs ± 1%   +70.79%  (p=0.000 n=10+10)
/100000-4     4.45ms ±27%   20.61ms ±10%  +363.05%  (p=0.000 n=10+10)
/1000000-4    51.9ms ±12%   148.7ms ± 2%  +186.61%  (p=0.000 n=10+7)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%      ~     (all equal)
/1-4            128B ± 0%      112B ± 0%   -12.50%  (p=0.000 n=10+10)
/10-4           592B ± 0%      688B ± 0%   +16.22%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    6.45kB ± 0%    +4.40%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    64.0kB ± 0%   +94.04%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     640kB ± 0%   +98.74%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +98.71%  (p=0.000 n=9+10)
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
benchstat testdata/BenchmarkRefillDequeQueuev1.0.2.txt testdata/BenchmarkRefillListQueue.txt
name       old time/op    new time/op    delta
/1-4         3.97µs ± 6%    7.36µs ± 4%   +85.38%  (p=0.000 n=9+9)
/10-4        36.8µs ± 5%    74.2µs ±10%  +101.65%  (p=0.000 n=10+10)
/100-4        347µs ± 5%     716µs ± 3%  +106.23%  (p=0.000 n=10+9)
/1000-4      3.67ms ±15%    7.13ms ± 4%   +94.45%  (p=0.000 n=10+8)
/10000-4     39.7ms ± 5%    75.9ms ± 9%   +91.31%  (p=0.000 n=10+10)
/100000-4     413ms ± 6%    2066ms ±11%  +400.19%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=10+9)
/10000-4     30.1MB ± 0%    64.0MB ± 0%  +112.51%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     640MB ± 0%  +100.11%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=9+10)
```

deque vs list - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt testdata/BenchmarkRefillFullListQueue.txt
name       old time/op    new time/op    delta
/1-4         3.70µs ± 8%    8.21µs ± 5%  +122.14%  (p=0.000 n=10+10)
/10-4        37.4µs ± 7%    81.5µs ± 4%  +118.12%  (p=0.000 n=10+10)
/100-4        360µs ± 4%     803µs ± 3%  +123.22%  (p=0.000 n=10+10)
/1000-4      3.50ms ± 3%    8.38ms ± 9%  +139.72%  (p=0.000 n=9+10)
/10000-4     40.0ms ± 5%    86.6ms ± 7%  +116.19%  (p=0.000 n=9+10)
/100000-4     426ms ±11%    1608ms ± 6%  +277.62%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    64.0MB ± 0%  +112.52%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     640MB ± 0%  +100.12%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=9+10)
```

deque vs list - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowIncreaseListQueue.txt
name        old time/op    new time/op    delta
/1-4           248ns ± 5%     179ns ± 5%   -27.68%  (p=0.000 n=9+10)
/10-4         2.09µs ±13%    1.44µs ± 4%   -31.10%  (p=0.000 n=9+10)
/100-4        8.06µs ± 5%   14.12µs ± 7%   +75.28%  (p=0.000 n=10+9)
/1000-4       73.5µs ± 5%   139.3µs ± 5%   +89.65%  (p=0.000 n=9+9)
/10000-4       703µs ± 1%    1473µs ± 5%  +109.57%  (p=0.000 n=8+8)
/100000-4     8.45ms ± 8%   22.99ms ± 7%  +171.98%  (p=0.000 n=10+10)
/1000000-4    90.6ms ± 6%   268.3ms ± 8%  +196.11%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      176B ± 0%   -15.38%  (p=0.000 n=10+10)
/10-4         4.90kB ± 0%    1.33kB ± 0%   -72.88%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%   12.85kB ± 0%   +65.23%  (p=0.000 n=10+10)
/1000-4       53.2kB ± 0%   128.0kB ± 0%  +140.91%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%    1280kB ± 0%  +163.28%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   12.80MB ± 0%  +165.29%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   128.0MB ± 0%  +165.61%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      41.0 ± 0%   +51.85%  (p=0.000 n=10+10)
/100-4           207 ± 0%       401 ± 0%   +93.72%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     4.00k ± 0%   +98.56%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     40.0k ± 0%   +99.16%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      400k ± 0%   +99.21%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     4.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowDecreaseListQueue.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 3%    70.4ns ±12%   +99.70%  (p=0.000 n=10+9)
/10-4          353ns ± 2%     703ns ± 4%   +99.15%  (p=0.000 n=9+9)
/100-4        3.61µs ±13%    6.73µs ± 1%   +86.41%  (p=0.000 n=9+8)
/1000-4       35.2µs ± 2%    70.8µs ± 6%  +101.32%  (p=0.000 n=8+9)
/10000-4       356µs ± 7%     696µs ± 3%   +95.40%  (p=0.000 n=9+8)
/100000-4     3.51ms ± 3%    7.25ms ± 6%  +106.53%  (p=0.000 n=10+8)
/1000000-4    34.3ms ± 2%    69.4ms ± 8%  +101.97%  (p=0.000 n=10+9)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.2.txt testdata/BenchmarkStableListQueue.txt
name        old time/op    new time/op    delta
/1-4          35.1ns ± 2%    86.9ns ± 6%  +147.77%  (p=0.000 n=9+10)
/10-4          358ns ± 4%     821ns ± 4%  +129.02%  (p=0.000 n=9+9)
/100-4        3.44µs ± 2%    7.76µs ± 1%  +125.43%  (p=0.000 n=9+8)
/1000-4       34.2µs ± 2%    90.0µs ±37%  +163.14%  (p=0.000 n=9+10)
/10000-4       357µs ± 6%     791µs ± 7%  +121.43%  (p=0.000 n=10+9)
/100000-4     3.49ms ± 4%    9.03ms ±26%  +158.80%  (p=0.000 n=10+9)
/1000000-4    35.6ms ± 3%    83.4ms ± 9%  +134.15%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+9)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=8+9)

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
benchstat testdata/BenchmarkFillDequeStackv1.0.2.txt testdata/BenchmarkFillListStack.txt
name        old time/op    new time/op    delta
/0-4          38.3ns ± 3%    39.0ns ± 1%    +1.76%  (p=0.009 n=8+9)
/1-4           148ns ± 9%     107ns ± 1%   -27.63%  (p=0.000 n=10+10)
/10-4          662ns ±13%     725ns ± 1%    +9.60%  (p=0.002 n=10+10)
/100-4        4.99µs ± 6%    6.75µs ± 1%   +35.17%  (p=0.000 n=9+10)
/1000-4       39.5µs ± 8%    68.4µs ± 0%   +73.18%  (p=0.000 n=10+10)
/10000-4       387µs ± 9%     700µs ± 0%   +80.66%  (p=0.000 n=10+9)
/100000-4     4.13ms ± 9%   21.25ms ±10%  +414.86%  (p=0.000 n=10+10)
/1000000-4    46.8ms ± 4%   157.6ms ±10%  +236.92%  (p=0.000 n=8+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%      ~     (all equal)
/1-4            128B ± 0%      112B ± 0%   -12.50%  (p=0.000 n=10+10)
/10-4           592B ± 0%      688B ± 0%   +16.22%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    6.45kB ± 0%    +4.40%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    64.0kB ± 0%   +94.04%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     640kB ± 0%   +98.74%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +98.71%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    64.0MB ± 0%   +98.81%  (p=0.000 n=9+8)

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
benchstat testdata/BenchmarkRefillDequeStackv1.0.2.txt testdata/BenchmarkRefillListStack.txt
name       old time/op    new time/op    delta
/1-4         4.10µs ±15%    7.28µs ± 5%   +77.39%  (p=0.000 n=10+10)
/10-4        35.7µs ± 2%    71.6µs ± 4%  +100.48%  (p=0.000 n=10+8)
/100-4        353µs ±10%     718µs ± 6%  +103.33%  (p=0.000 n=10+10)
/1000-4      3.43ms ± 5%    7.04ms ± 5%  +105.00%  (p=0.000 n=9+10)
/10000-4     38.2ms ±11%    77.3ms ±16%  +102.49%  (p=0.000 n=9+9)
/100000-4     402ms ± 5%    2156ms ±11%  +435.82%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=9+9)
/10000-4     30.1MB ± 0%    64.0MB ± 0%  +112.70%  (p=0.000 n=8+9)
/100000-4     320MB ± 0%     640MB ± 0%  +100.02%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=10+9)
```

deque vs list - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.2.txt testdata/BenchmarkRefillFullListStack.txt
name       old time/op    new time/op    delta
/1-4         4.33µs ± 5%    8.28µs ± 4%   +91.10%  (p=0.000 n=9+9)
/10-4        35.8µs ±11%    83.8µs ± 2%  +134.37%  (p=0.000 n=10+9)
/100-4        364µs ±17%     831µs ± 2%  +128.43%  (p=0.000 n=10+10)
/1000-4      3.40ms ± 3%    8.55ms ±11%  +151.63%  (p=0.000 n=10+9)
/10000-4     36.7ms ± 0%    96.3ms ±10%  +162.20%  (p=0.000 n=9+9)
/100000-4     407ms ±14%    1619ms ± 9%  +298.03%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4     30.5MB ± 0%    64.0MB ± 0%  +109.81%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     640MB ± 0%  +100.03%  (p=0.000 n=8+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.61%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=10+8)
```

deque vs list - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowIncreaseListStack.txt
name        old time/op    new time/op    delta
/1-4           252ns ± 5%     180ns ± 6%   -28.54%  (p=0.000 n=10+10)
/10-4         1.02µs ± 5%    1.46µs ± 9%   +43.66%  (p=0.000 n=8+9)
/100-4        7.95µs ± 5%   13.99µs ± 9%   +75.99%  (p=0.000 n=10+9)
/1000-4       68.8µs ± 4%   139.6µs ± 2%  +103.04%  (p=0.000 n=10+8)
/10000-4       723µs ± 9%    1520µs ±11%  +110.15%  (p=0.000 n=10+10)
/100000-4     8.33ms ± 4%   27.93ms ± 9%  +235.20%  (p=0.000 n=9+10)
/1000000-4    89.7ms ± 6%   357.6ms ±29%  +298.50%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      176B ± 0%   -15.38%  (p=0.000 n=10+10)
/10-4           752B ± 0%     1328B ± 0%   +76.60%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%   12.85kB ± 0%   +65.23%  (p=0.000 n=10+10)
/1000-4       49.0kB ± 0%   128.0kB ± 0%  +161.28%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%    1280kB ± 0%  +163.28%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   12.80MB ± 0%  +165.52%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   128.0MB ± 0%  +165.61%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      41.0 ± 0%   +64.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       401 ± 0%   +93.72%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     4.00k ± 0%   +98.76%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     40.0k ± 0%   +99.16%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      400k ± 0%   +99.22%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     4.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

deque vs list - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowDecreaseListStack.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    72.6ns ± 3%  +114.12%  (p=0.000 n=8+8)
/10-4          353ns ± 5%     717ns ± 3%  +103.19%  (p=0.000 n=9+10)
/100-4        3.45µs ± 5%    7.01µs ± 4%  +103.16%  (p=0.000 n=9+9)
/1000-4       33.9µs ± 2%    81.4µs ±36%  +140.31%  (p=0.000 n=9+10)
/10000-4       342µs ± 3%     699µs ± 5%  +104.56%  (p=0.000 n=8+8)
/100000-4     3.56ms ±10%    7.22ms ± 6%  +102.63%  (p=0.000 n=10+8)
/1000000-4    36.3ms ±13%    71.9ms ± 7%   +98.31%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+8)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.2.txt testdata/BenchmarkStableListStack.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 8%    87.1ns ±12%  +122.19%  (p=0.000 n=10+9)
/10-4          368ns ± 2%     920ns ±31%  +150.08%  (p=0.000 n=9+10)
/100-4        3.65µs ± 4%    8.71µs ±17%  +138.29%  (p=0.000 n=10+10)
/1000-4       36.6µs ± 3%    87.1µs ±31%  +138.29%  (p=0.000 n=9+9)
/10000-4       381µs ±12%     801µs ± 6%  +110.17%  (p=0.000 n=10+9)
/100000-4     3.55ms ± 2%    8.66ms ±23%  +143.72%  (p=0.000 n=8+9)
/1000000-4    37.3ms ± 7%    85.9ms ±13%  +130.51%  (p=0.000 n=9+10)

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

#### deque vs slice - FIFO queue
deque vs slice - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.2.txt testdata/BenchmarkFillSliceQueue.txt
name        old time/op    new time/op    delta
/0-4          36.1ns ± 4%    41.1ns ± 4%   +13.87%  (p=0.000 n=10+10)
/1-4           143ns ± 5%     101ns ± 7%   -29.24%  (p=0.000 n=10+8)
/10-4          664ns ± 9%     662ns ±14%      ~     (p=0.661 n=10+9)
/100-4        4.83µs ± 7%    4.64µs ±30%      ~     (p=0.404 n=10+10)
/1000-4       42.4µs ±15%    33.1µs ± 3%   -21.93%  (p=0.000 n=10+9)
/10000-4       417µs ±11%     410µs ± 9%      ~     (p=0.912 n=10+10)
/100000-4     4.45ms ±27%    9.28ms ±10%  +108.57%  (p=0.000 n=10+9)
/1000000-4    51.9ms ±12%   115.2ms ±33%  +122.04%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     32.0B ± 0%   -33.33%  (p=0.000 n=10+10)
/1-4            128B ± 0%       56B ± 0%   -56.25%  (p=0.000 n=10+10)
/10-4           592B ± 0%      440B ± 0%   -25.68%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    3.67kB ± 0%   -40.54%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    32.4kB ± 0%    -1.82%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     546kB ± 0%   +69.64%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.19%  (p=0.000 n=9+10)
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
benchstat testdata/BenchmarkRefillDequeQueuev1.0.2.txt testdata/BenchmarkRefillSliceQueue.txt
name       old time/op    new time/op    delta
/1-4         3.97µs ± 6%    6.00µs ± 9%   +51.17%  (p=0.000 n=9+10)
/10-4        36.8µs ± 5%    45.8µs ±11%   +24.49%  (p=0.000 n=10+10)
/100-4        347µs ± 5%     330µs ± 4%    -4.96%  (p=0.000 n=10+10)
/1000-4      3.67ms ±15%    3.77ms ±26%      ~     (p=0.481 n=10+10)
/10000-4     39.7ms ± 5%    38.6ms ±13%      ~     (p=0.247 n=10+10)
/100000-4     413ms ± 6%     813ms ±10%   +96.76%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.0kB ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     314kB ± 0%   +96.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    3.15MB ± 0%   +96.80%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    48.8MB ± 0%   +61.93%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     548MB ± 0%   +71.31%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.1k ± 0%    +1.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    +0.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.62%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=9+10)
```

deque vs slice - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt testdata/BenchmarkRefillFullSliceQueue.txt
name       old time/op    new time/op    delta
/1-4         3.70µs ± 8%    4.16µs ±10%   +12.54%  (p=0.000 n=10+9)
/10-4        37.4µs ± 7%    40.4µs ± 7%    +8.12%  (p=0.000 n=10+10)
/100-4        360µs ± 4%     406µs ±27%   +13.03%  (p=0.000 n=10+9)
/1000-4      3.50ms ± 3%    4.11ms ± 5%   +17.65%  (p=0.000 n=9+8)
/10000-4     40.0ms ± 5%    40.4ms ± 6%      ~     (p=0.489 n=9+9)
/100000-4     426ms ±11%     929ms ±18%  +118.10%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    48.2kB ± 0%  +201.08%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     483kB ± 0%  +201.70%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    5.15MB ± 0%  +221.87%  (p=0.000 n=10+8)
/10000-4     30.1MB ± 0%    52.0MB ± 0%   +72.82%  (p=0.000 n=10+8)
/100000-4     320MB ± 0%     551MB ± 0%   +72.32%  (p=0.000 n=8+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    +0.03%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=9+9)
```

deque vs slice - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowIncreaseSliceQueue.txt
name        old time/op    new time/op    delta
/1-4           248ns ± 5%     160ns ± 7%   -35.33%  (p=0.000 n=9+8)
/10-4         2.09µs ±13%    1.10µs ± 4%   -47.14%  (p=0.000 n=9+9)
/100-4        8.06µs ± 5%    7.23µs ± 5%   -10.25%  (p=0.000 n=10+9)
/1000-4       73.5µs ± 5%    66.7µs ± 4%    -9.22%  (p=0.000 n=9+9)
/10000-4       703µs ± 1%     729µs ± 3%    +3.68%  (p=0.000 n=8+10)
/100000-4     8.45ms ± 8%   15.10ms ±23%   +78.63%  (p=0.000 n=10+9)
/1000000-4    90.6ms ± 6%   169.3ms ± 9%   +86.87%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%       88B ± 0%   -57.69%  (p=0.000 n=10+10)
/10-4         4.90kB ± 0%    0.78kB ± 0%   -84.15%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    6.66kB ± 0%   -14.30%  (p=0.000 n=10+10)
/1000-4       53.2kB ± 0%    74.0kB ± 0%   +39.21%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     991kB ± 0%  +103.75%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   11.42MB ± 0%  +136.60%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   114.6MB ± 0%  +137.75%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      29.0 ± 0%    +7.41%  (p=0.000 n=10+10)
/100-4           207 ± 0%       214 ± 0%    +3.38%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.02k ± 0%    +0.25%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.25%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs slice - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowDecreaseSliceQueue.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 3%    56.3ns ± 3%   +59.55%  (p=0.000 n=10+8)
/10-4          353ns ± 2%     574ns ± 4%   +62.67%  (p=0.000 n=9+8)
/100-4        3.61µs ±13%    5.50µs ± 3%   +52.38%  (p=0.000 n=9+9)
/1000-4       35.2µs ± 2%    55.6µs ± 8%   +58.06%  (p=0.000 n=8+9)
/10000-4       356µs ± 7%     547µs ± 2%   +53.52%  (p=0.000 n=9+8)
/100000-4     3.51ms ± 3%    6.84ms ±61%   +94.91%  (p=0.000 n=10+10)
/1000000-4    34.3ms ± 2%    56.0ms ± 5%   +63.18%  (p=0.000 n=10+8)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.2.txt testdata/BenchmarkStableSliceQueue.txt
name        old time/op    new time/op    delta
/1-4          35.1ns ± 2%    39.2ns ± 5%   +11.75%  (p=0.000 n=9+8)
/10-4          358ns ± 4%     452ns ±37%   +26.06%  (p=0.000 n=9+10)
/100-4        3.44µs ± 2%    3.77µs ± 8%    +9.58%  (p=0.000 n=9+9)
/1000-4       34.2µs ± 2%    39.7µs ±25%   +16.02%  (p=0.000 n=9+9)
/10000-4       357µs ± 6%     400µs ± 4%   +11.86%  (p=0.000 n=10+9)
/100000-4     3.49ms ± 4%    3.91ms ± 7%   +12.16%  (p=0.000 n=10+9)
/1000000-4    35.6ms ± 3%    39.4ms ± 5%   +10.57%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     48.0B ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      481B ± 0%  +200.62%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    48.2kB ± 0%  +200.96%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     482kB ± 0%  +200.97%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    4.82MB ± 0%  +200.97%  (p=0.000 n=10+8)
/1000000-4    16.0MB ± 0%    48.2MB ± 0%  +200.97%  (p=0.000 n=8+7)

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
benchstat testdata/BenchmarkFillDequeStackv1.0.2.txt testdata/BenchmarkFillSliceStack.txt
name        old time/op    new time/op    delta
/0-4          38.3ns ± 3%    43.4ns ± 3%   +13.24%  (p=0.000 n=8+9)
/1-4           148ns ± 9%     100ns ± 4%   -32.72%  (p=0.000 n=10+9)
/10-4          662ns ±13%     623ns ± 5%      ~     (p=0.064 n=10+8)
/100-4        4.99µs ± 6%    4.79µs ±37%      ~     (p=0.156 n=9+10)
/1000-4       39.5µs ± 8%    36.0µs ± 7%    -8.80%  (p=0.001 n=10+10)
/10000-4       387µs ± 9%     426µs ± 6%    +9.99%  (p=0.000 n=10+9)
/100000-4     4.13ms ± 9%   10.34ms ±14%  +150.46%  (p=0.000 n=10+10)
/1000000-4    46.8ms ± 4%   116.0ms ± 9%  +148.09%  (p=0.000 n=8+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     32.0B ± 0%   -33.33%  (p=0.000 n=10+10)
/1-4            128B ± 0%       56B ± 0%   -56.25%  (p=0.000 n=10+10)
/10-4           592B ± 0%      440B ± 0%   -25.68%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    3.67kB ± 0%   -40.54%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    32.4kB ± 0%    -1.82%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     546kB ± 0%   +69.64%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.19%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    61.2MB ± 0%   +90.08%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkRefillDequeStackv1.0.2.txt testdata/BenchmarkRefillSliceStack.txt
name       old time/op    new time/op    delta
/1-4         4.10µs ±15%    3.01µs ± 5%  -26.64%  (p=0.000 n=10+10)
/10-4        35.7µs ± 2%    31.7µs ±24%     ~     (p=0.143 n=10+10)
/100-4        353µs ±10%     286µs ±17%  -18.90%  (p=0.000 n=10+9)
/1000-4      3.43ms ± 5%    2.66ms ± 2%  -22.57%  (p=0.000 n=9+9)
/10000-4     38.2ms ±11%    29.4ms ±16%  -22.98%  (p=0.000 n=9+9)
/100000-4     402ms ± 5%     343ms ± 3%  -14.69%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=9+9)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.80%  (p=0.000 n=8+9)
/100000-4     320MB ± 0%     162MB ± 0%  -49.51%  (p=0.000 n=9+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+8)
```

deque vs slice - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.2.txt testdata/BenchmarkRefillFullSliceStack.txt
name       old time/op    new time/op    delta
/1-4         4.33µs ± 5%    3.11µs ± 9%  -28.19%  (p=0.000 n=9+9)
/10-4        35.8µs ±11%    29.6µs ± 9%  -17.29%  (p=0.000 n=10+9)
/100-4        364µs ±17%     288µs ± 6%  -20.73%  (p=0.000 n=10+9)
/1000-4      3.40ms ± 3%    2.96ms ±12%  -12.73%  (p=0.000 n=10+9)
/10000-4     36.7ms ± 0%    29.3ms ± 7%  -20.29%  (p=0.000 n=9+10)
/100000-4     407ms ±14%     336ms ± 2%  -17.32%  (p=0.000 n=10+8)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.5MB ± 0%    16.0MB ± 0%  -47.55%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     160MB ± 0%  -49.99%  (p=0.002 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.70%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs slice - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowIncreaseSliceStack.txt
name        old time/op    new time/op    delta
/1-4           252ns ± 5%     158ns ± 5%  -37.50%  (p=0.000 n=10+10)
/10-4         1.02µs ± 5%    0.89µs ± 6%  -12.66%  (p=0.000 n=8+10)
/100-4        7.95µs ± 5%    6.52µs ± 2%  -18.01%  (p=0.000 n=10+8)
/1000-4       68.8µs ± 4%    65.0µs ±25%     ~     (p=0.143 n=10+10)
/10000-4       723µs ± 9%     639µs ± 3%  -11.66%  (p=0.000 n=10+9)
/100000-4     8.33ms ± 4%   11.38ms ± 5%  +36.52%  (p=0.000 n=9+9)
/1000000-4    89.7ms ± 6%   138.6ms ± 8%  +54.43%  (p=0.000 n=9+8)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%       88B ± 0%  -57.69%  (p=0.000 n=10+10)
/10-4           752B ± 0%      600B ± 0%  -20.21%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    5.27kB ± 0%  -32.20%  (p=0.000 n=10+10)
/1000-4       49.0kB ± 0%    48.4kB ± 0%   -1.22%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     706kB ± 0%  +45.28%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    7.85MB ± 0%  +62.93%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    77.2MB ± 0%  +60.17%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%  -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      26.0 ± 0%   +4.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       209 ± 0%   +0.97%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%   -0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%   -0.32%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%   -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%   -0.39%  (p=0.000 n=10+10)
```

deque vs slice - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowDecreaseSliceStack.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    32.3ns ± 6%     ~     (p=0.102 n=8+8)
/10-4          353ns ± 5%     307ns ± 2%  -12.88%  (p=0.000 n=9+9)
/100-4        3.45µs ± 5%    3.04µs ± 4%  -11.97%  (p=0.000 n=9+10)
/1000-4       33.9µs ± 2%    31.8µs ±11%   -6.27%  (p=0.004 n=9+9)
/10000-4       342µs ± 3%     317µs ± 6%   -7.19%  (p=0.000 n=8+10)
/100000-4     3.56ms ±10%    3.08ms ± 5%  -13.42%  (p=0.000 n=10+10)
/1000000-4    36.3ms ±13%    30.8ms ± 5%  -14.96%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.2.txt testdata/BenchmarkStableSliceStack.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 8%    27.4ns ± 5%  -30.20%  (p=0.000 n=10+9)
/10-4          368ns ± 2%     282ns ± 6%  -23.37%  (p=0.000 n=9+8)
/100-4        3.65µs ± 4%    2.70µs ± 3%  -26.23%  (p=0.000 n=10+8)
/1000-4       36.6µs ± 3%    26.8µs ± 6%  -26.61%  (p=0.000 n=9+8)
/10000-4       381µs ±12%     265µs ± 6%  -30.46%  (p=0.000 n=10+9)
/100000-4     3.55ms ± 2%    2.64ms ± 4%  -25.67%  (p=0.000 n=8+8)
/1000000-4    37.3ms ± 7%    26.9ms ± 7%  -27.74%  (p=0.000 n=9+8)

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
benchstat testdata/BenchmarkFillDequeQueuev1.0.2.txt testdata/BenchmarkFillPhfQueue.txt
name        old time/op    new time/op    delta
/0-4          36.1ns ± 4%    67.5ns ±10%   +86.85%  (p=0.000 n=10+10)
/1-4           143ns ± 5%     107ns ± 5%   -24.86%  (p=0.000 n=10+10)
/10-4          664ns ± 9%     869ns ±16%   +30.81%  (p=0.000 n=10+9)
/100-4        4.83µs ± 7%    6.55µs ±29%   +35.59%  (p=0.000 n=10+10)
/1000-4       42.4µs ±15%    53.0µs ±14%   +24.89%  (p=0.000 n=10+9)
/10000-4       417µs ±11%     577µs ±14%   +38.30%  (p=0.000 n=10+9)
/100000-4     4.45ms ±27%    7.65ms ±10%   +71.77%  (p=0.000 n=10+9)
/1000000-4    51.9ms ±12%    78.6ms ± 6%   +51.53%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     64.0B ± 0%   +33.33%  (p=0.000 n=10+10)
/1-4            128B ± 0%       80B ± 0%   -37.50%  (p=0.000 n=10+10)
/10-4           592B ± 0%      832B ± 0%   +40.54%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    7.65kB ± 0%   +23.83%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    65.1kB ± 0%   +97.09%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.85%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.02%  (p=0.000 n=9+10)
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
benchstat testdata/BenchmarkRefillDequeQueuev1.0.2.txt testdata/BenchmarkRefillPhfQueue.txt
name       old time/op    new time/op    delta
/1-4         3.97µs ± 6%    3.87µs ± 1%      ~     (p=0.113 n=9+10)
/10-4        36.8µs ± 5%    57.1µs ± 2%   +55.28%  (p=0.000 n=10+10)
/100-4        347µs ± 5%     581µs ± 6%   +67.25%  (p=0.000 n=10+10)
/1000-4      3.67ms ±15%    4.79ms ± 4%   +30.72%  (p=0.000 n=10+10)
/10000-4     39.7ms ± 5%    54.1ms ± 3%   +36.31%  (p=0.000 n=10+10)
/100000-4     413ms ± 6%     721ms ± 3%   +74.59%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    54.4kB ± 0%  +240.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     736kB ± 0%  +360.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.48MB ± 0%  +304.79%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    94.6MB ± 0%  +214.13%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     789MB ± 0%  +146.74%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.40%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.46%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=9+9)
```

deque vs phf - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt testdata/BenchmarkRefillFullPhfQueue.txt
name       old time/op    new time/op    delta
/1-4         3.70µs ± 8%    3.94µs ± 2%   +6.60%  (p=0.023 n=10+9)
/10-4        37.4µs ± 7%    39.2µs ± 3%   +4.84%  (p=0.008 n=10+9)
/100-4        360µs ± 4%     384µs ± 5%   +6.86%  (p=0.000 n=10+10)
/1000-4      3.50ms ± 3%    3.86ms ± 8%  +10.32%  (p=0.000 n=9+10)
/10000-4     40.0ms ± 5%    38.8ms ± 2%   -3.11%  (p=0.011 n=9+8)
/100000-4     426ms ±11%     707ms ± 2%  +65.93%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.87%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     632MB ± 0%  +97.58%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.76%  (p=0.000 n=9+10)
```

deque vs phf - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowIncreasePhfQueue.txt
name        old time/op    new time/op    delta
/1-4           248ns ± 5%     215ns ± 7%   -13.23%  (p=0.000 n=9+10)
/10-4         2.09µs ±13%    1.36µs ± 7%   -34.70%  (p=0.000 n=9+10)
/100-4        8.06µs ± 5%    9.66µs ± 5%   +19.85%  (p=0.000 n=10+10)
/1000-4       73.5µs ± 5%    83.0µs ± 2%   +12.98%  (p=0.000 n=9+9)
/10000-4       703µs ± 1%     956µs ± 2%   +36.08%  (p=0.000 n=8+10)
/100000-4     8.45ms ± 8%   11.87ms ± 6%   +40.42%  (p=0.000 n=10+10)
/1000000-4    90.6ms ± 6%   129.3ms ± 4%   +42.64%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      128B ± 0%   -38.46%  (p=0.000 n=10+10)
/10-4         4.90kB ± 0%    0.99kB ± 0%   -79.74%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    9.25kB ± 0%   +18.93%  (p=0.000 n=10+10)
/1000-4       53.2kB ± 0%    81.1kB ± 0%   +52.50%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%    1106kB ± 0%  +127.55%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.72%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      27.0 ± 0%      ~     (all equal)
/100-4           207 ± 0%       213 ± 0%    +2.90%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.02k ± 0%    +0.20%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.29%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs phf - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowDecreasePhfQueue.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 3%    37.1ns ± 4%  +5.10%  (p=0.000 n=10+10)
/10-4          353ns ± 2%     372ns ± 2%  +5.42%  (p=0.000 n=9+9)
/100-4        3.61µs ±13%    3.62µs ± 3%    ~     (p=0.661 n=9+10)
/1000-4       35.2µs ± 2%    36.4µs ± 3%  +3.59%  (p=0.002 n=8+10)
/10000-4       356µs ± 7%     368µs ± 2%  +3.43%  (p=0.022 n=9+10)
/100000-4     3.51ms ± 3%    3.70ms ± 4%  +5.51%  (p=0.000 n=10+10)
/1000000-4    34.3ms ± 2%    37.4ms ± 3%  +8.81%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.173 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs phf - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueuev1.0.2.txt testdata/BenchmarkStablePhfQueue.txt
name        old time/op    new time/op    delta
/1-4          35.1ns ± 2%    37.4ns ± 3%  +6.53%  (p=0.000 n=9+10)
/10-4          358ns ± 4%     379ns ± 5%  +5.71%  (p=0.000 n=9+9)
/100-4        3.44µs ± 2%    3.67µs ± 2%  +6.67%  (p=0.000 n=9+10)
/1000-4       34.2µs ± 2%    36.2µs ± 3%  +5.93%  (p=0.000 n=9+10)
/10000-4       357µs ± 6%     371µs ± 4%    ~     (p=0.052 n=10+10)
/100000-4     3.49ms ± 4%    3.71ms ± 3%  +6.33%  (p=0.000 n=10+10)
/1000000-4    35.6ms ± 3%    37.5ms ± 4%  +5.37%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkFillDequeStackv1.0.2.txt testdata/BenchmarkFillPhfStack.txt
name        old time/op    new time/op    delta
/0-4          38.3ns ± 3%    69.4ns ± 8%   +81.11%  (p=0.000 n=8+10)
/1-4           148ns ± 9%     105ns ± 5%   -29.38%  (p=0.000 n=10+10)
/10-4          662ns ±13%     852ns ± 2%   +28.81%  (p=0.000 n=10+9)
/100-4        4.99µs ± 6%    5.85µs ± 5%   +17.23%  (p=0.000 n=9+10)
/1000-4       39.5µs ± 8%    47.5µs ± 5%   +20.32%  (p=0.000 n=10+10)
/10000-4       387µs ± 9%     539µs ± 4%   +39.28%  (p=0.000 n=10+10)
/100000-4     4.13ms ± 9%    7.33ms ± 7%   +77.49%  (p=0.000 n=10+10)
/1000000-4    46.8ms ± 4%    81.9ms ± 6%   +75.19%  (p=0.000 n=8+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     64.0B ± 0%   +33.33%  (p=0.000 n=10+10)
/1-4            128B ± 0%       80B ± 0%   -37.50%  (p=0.000 n=10+10)
/10-4           592B ± 0%      832B ± 0%   +40.54%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    7.65kB ± 0%   +23.83%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    65.1kB ± 0%   +97.09%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.85%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.02%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.06%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkRefillDequeStackv1.0.2.txt testdata/BenchmarkRefillPhfStack.txt
name       old time/op    new time/op    delta
/1-4         4.10µs ±15%    3.93µs ± 3%      ~     (p=0.631 n=10+10)
/10-4        35.7µs ± 2%    57.8µs ± 3%   +61.80%  (p=0.000 n=10+9)
/100-4        353µs ±10%     557µs ± 2%   +57.72%  (p=0.000 n=10+10)
/1000-4      3.43ms ± 5%    4.72ms ± 2%   +37.51%  (p=0.000 n=9+10)
/10000-4     38.2ms ±11%    54.4ms ± 3%   +42.42%  (p=0.000 n=9+10)
/100000-4     402ms ± 5%     750ms ±10%   +86.42%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    54.4kB ± 0%  +240.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     736kB ± 0%  +360.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.48MB ± 0%  +304.79%  (p=0.000 n=9+10)
/10000-4     30.1MB ± 0%    94.6MB ± 0%  +214.41%  (p=0.000 n=8+9)
/100000-4     320MB ± 0%     789MB ± 0%  +146.62%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.40%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.46%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.2.txt testdata/BenchmarkRefillFullPhfStack.txt
name       old time/op    new time/op    delta
/1-4         4.33µs ± 5%    3.87µs ± 6%  -10.80%  (p=0.000 n=9+9)
/10-4        35.8µs ±11%    37.8µs ± 2%   +5.83%  (p=0.002 n=10+10)
/100-4        364µs ±17%     372µs ± 4%     ~     (p=0.280 n=10+10)
/1000-4      3.40ms ± 3%    3.69ms ± 2%   +8.64%  (p=0.000 n=10+10)
/10000-4     36.7ms ± 0%    38.3ms ± 3%   +4.37%  (p=0.000 n=9+10)
/100000-4     407ms ±14%     733ms ± 8%  +80.23%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.5MB ± 0%    16.0MB ± 0%  -47.55%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     632MB ± 0%  +97.48%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.70%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.76%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowIncreasePhfStack.txt
name        old time/op    new time/op    delta
/1-4           252ns ± 5%     212ns ± 2%   -16.11%  (p=0.000 n=10+9)
/10-4         1.02µs ± 5%    1.27µs ± 4%   +24.98%  (p=0.000 n=8+10)
/100-4        7.95µs ± 5%    9.49µs ± 6%   +19.37%  (p=0.000 n=10+10)
/1000-4       68.8µs ± 4%    81.5µs ± 2%   +18.48%  (p=0.000 n=10+10)
/10000-4       723µs ± 9%     943µs ± 3%   +30.38%  (p=0.000 n=10+10)
/100000-4     8.33ms ± 4%   11.64ms ± 5%   +39.71%  (p=0.000 n=9+9)
/1000000-4    89.7ms ± 6%   138.1ms ±13%   +53.91%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      128B ± 0%   -38.46%  (p=0.000 n=10+10)
/10-4           752B ± 0%      992B ± 0%   +31.91%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    9.25kB ± 0%   +18.93%  (p=0.000 n=10+10)
/1000-4       49.0kB ± 0%    81.1kB ± 0%   +65.39%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%    1106kB ± 0%  +127.55%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.89%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%   -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      27.0 ± 0%    +8.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       213 ± 0%    +2.90%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.02k ± 0%    +0.30%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.29%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowDecreasePhfStack.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    36.7ns ± 3%   +8.16%  (p=0.000 n=8+9)
/10-4          353ns ± 5%     387ns ± 4%   +9.67%  (p=0.000 n=9+10)
/100-4        3.45µs ± 5%    3.79µs ± 5%   +9.76%  (p=0.000 n=9+10)
/1000-4       33.9µs ± 2%    37.4µs ± 4%  +10.25%  (p=0.000 n=9+10)
/10000-4       342µs ± 3%     378µs ± 4%  +10.58%  (p=0.000 n=8+9)
/100000-4     3.56ms ±10%    3.77ms ± 6%   +5.94%  (p=0.035 n=10+9)
/1000000-4    36.3ms ±13%    37.1ms ± 1%     ~     (p=0.173 n=10+8)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.2.txt testdata/BenchmarkStablePhfStack.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 8%    39.2ns ±13%    ~     (p=0.754 n=10+10)
/10-4          368ns ± 2%     370ns ± 8%    ~     (p=0.794 n=9+10)
/100-4        3.65µs ± 4%    3.84µs ± 6%  +5.08%  (p=0.015 n=10+10)
/1000-4       36.6µs ± 3%    37.0µs ± 5%    ~     (p=0.549 n=9+10)
/10000-4       381µs ±12%     365µs ± 6%    ~     (p=0.315 n=10+10)
/100000-4     3.55ms ± 2%    3.57ms ± 8%    ~     (p=0.173 n=8+10)
/1000000-4    37.3ms ± 7%    35.5ms ± 2%  -4.82%  (p=0.002 n=9+10)

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

#### deque vs gammazero - FIFO queue
deque vs gammazero - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.2.txt testdata/BenchmarkFillGammazeroQueue.txt
name        old time/op    new time/op    delta
/0-4          36.1ns ± 4%    38.7ns ± 6%    +7.17%  (p=0.000 n=10+10)
/1-4           143ns ± 5%     176ns ± 5%   +23.25%  (p=0.000 n=10+9)
/10-4          664ns ± 9%     510ns ± 7%   -23.30%  (p=0.000 n=10+9)
/100-4        4.83µs ± 7%    5.47µs ±11%   +13.32%  (p=0.000 n=10+10)
/1000-4       42.4µs ±15%    44.7µs ± 3%    +5.40%  (p=0.028 n=10+9)
/10000-4       417µs ±11%     509µs ± 2%   +22.02%  (p=0.000 n=10+8)
/100000-4     4.45ms ±27%    7.22ms ±20%   +62.17%  (p=0.000 n=10+9)
/1000000-4    51.9ms ±12%    81.8ms ±15%   +57.60%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%      ~     (all equal)
/1-4            128B ± 0%      320B ± 0%  +150.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%      464B ± 0%   -21.62%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    7.28kB ± 0%   +17.88%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    64.7kB ± 0%   +95.98%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.74%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.01%  (p=0.000 n=9+10)
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
benchstat testdata/BenchmarkRefillDequeQueuev1.0.2.txt testdata/BenchmarkRefillGammazeroQueue.txt
name       old time/op    new time/op    delta
/1-4         3.97µs ± 6%    3.85µs ± 1%      ~     (p=0.065 n=9+10)
/10-4        36.8µs ± 5%    36.3µs ± 1%      ~     (p=0.305 n=10+10)
/100-4        347µs ± 5%     511µs ± 2%   +47.17%  (p=0.000 n=10+10)
/1000-4      3.67ms ±15%    4.50ms ± 1%   +22.76%  (p=0.000 n=10+10)
/10000-4     39.7ms ± 5%    51.8ms ± 2%   +30.48%  (p=0.000 n=10+10)
/100000-4     413ms ± 6%     695ms ± 1%   +68.16%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +336.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.44MB ± 0%  +302.39%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    94.6MB ± 0%  +214.00%  (p=0.000 n=10+8)
/100000-4     320MB ± 0%     789MB ± 0%  +146.72%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.48%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=9+10)
```

deque vs gammazero - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt testdata/BenchmarkRefillFullGammazeroQueue.txt
name       old time/op    new time/op    delta
/1-4         3.70µs ± 8%    3.64µs ± 2%     ~     (p=0.644 n=10+10)
/10-4        37.4µs ± 7%    38.1µs ± 9%     ~     (p=0.247 n=10+10)
/100-4        360µs ± 4%     365µs ± 2%     ~     (p=0.190 n=10+10)
/1000-4      3.50ms ± 3%    3.58ms ± 4%   +2.40%  (p=0.001 n=9+10)
/10000-4     40.0ms ± 5%    37.1ms ± 3%   -7.45%  (p=0.000 n=9+9)
/100000-4     426ms ±11%     818ms ±34%  +92.06%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.87%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     632MB ± 0%  +97.58%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.76%  (p=0.000 n=9+10)
```

deque vs gammazero - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowIncreaseGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4           248ns ± 5%     221ns ± 4%   -10.85%  (p=0.000 n=9+9)
/10-4         2.09µs ±13%    0.85µs ± 5%   -59.47%  (p=0.000 n=9+9)
/100-4        8.06µs ± 5%    8.91µs ± 8%   +10.51%  (p=0.000 n=10+8)
/1000-4       73.5µs ± 5%    80.4µs ± 4%    +9.46%  (p=0.000 n=9+10)
/10000-4       703µs ± 1%     924µs ± 7%   +31.39%  (p=0.000 n=8+9)
/100000-4     8.45ms ± 8%   10.97ms ± 5%   +29.77%  (p=0.000 n=10+9)
/1000000-4    90.6ms ± 6%   122.9ms ± 9%   +35.67%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      336B ± 0%   +61.54%  (p=0.000 n=10+10)
/10-4         4.90kB ± 0%    0.62kB ± 0%   -87.25%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    8.88kB ± 0%   +14.20%  (p=0.000 n=10+10)
/1000-4       53.2kB ± 0%    80.7kB ± 0%   +51.81%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%    1106kB ± 0%  +127.48%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.71%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      4.00 ± 0%   -33.33%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      22.0 ± 0%   -18.52%  (p=0.000 n=10+10)
/100-4           207 ± 0%       208 ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.01k ± 0%    -0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.31%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowDecreaseGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 3%    35.1ns ± 1%    ~     (p=0.381 n=10+10)
/10-4          353ns ± 2%     355ns ± 1%    ~     (p=0.359 n=9+10)
/100-4        3.61µs ±13%    3.65µs ±16%    ~     (p=0.287 n=9+9)
/1000-4       35.2µs ± 2%    36.4µs ± 3%  +3.42%  (p=0.002 n=8+10)
/10000-4       356µs ± 7%     363µs ± 4%    ~     (p=0.222 n=9+9)
/100000-4     3.51ms ± 3%    3.61ms ± 4%  +2.81%  (p=0.023 n=10+10)
/1000000-4    34.3ms ± 2%    35.0ms ± 3%  +2.06%  (p=0.012 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.173 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs gammazero - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueuev1.0.2.txt testdata/BenchmarkStableGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4          35.1ns ± 2%    35.5ns ± 3%    ~     (p=0.117 n=9+10)
/10-4          358ns ± 4%     370ns ± 3%  +3.35%  (p=0.003 n=9+9)
/100-4        3.44µs ± 2%    3.77µs ± 8%  +9.61%  (p=0.000 n=9+9)
/1000-4       34.2µs ± 2%    36.6µs ± 6%  +7.01%  (p=0.000 n=9+9)
/10000-4       357µs ± 6%     357µs ± 4%    ~     (p=0.968 n=10+9)
/100000-4     3.49ms ± 4%    3.59ms ± 4%  +2.97%  (p=0.013 n=10+9)
/1000000-4    35.6ms ± 3%    38.8ms ±25%    ~     (p=0.156 n=9+10)

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

#### deque vs gammazero - LIFO stack
deque vs gammazero - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStackv1.0.2.txt testdata/BenchmarkFillGammazeroStack.txt
name        old time/op    new time/op    delta
/0-4          38.3ns ± 3%    40.3ns ± 7%    +5.27%  (p=0.007 n=8+9)
/1-4           148ns ± 9%     186ns ±10%   +25.40%  (p=0.000 n=10+10)
/10-4          662ns ±13%     524ns ± 7%   -20.84%  (p=0.000 n=10+10)
/100-4        4.99µs ± 6%    5.47µs ±11%    +9.56%  (p=0.000 n=9+9)
/1000-4       39.5µs ± 8%    46.9µs ± 6%   +18.69%  (p=0.000 n=10+10)
/10000-4       387µs ± 9%     542µs ± 7%   +39.93%  (p=0.000 n=10+10)
/100000-4     4.13ms ± 9%    7.21ms ± 7%   +74.66%  (p=0.000 n=10+9)
/1000000-4    46.8ms ± 4%    80.4ms ± 4%   +71.81%  (p=0.000 n=8+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%      ~     (all equal)
/1-4            128B ± 0%      320B ± 0%  +150.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%      464B ± 0%   -21.62%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    7.28kB ± 0%   +17.88%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    64.7kB ± 0%   +95.98%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.74%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.01%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.05%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkRefillDequeStackv1.0.2.txt testdata/BenchmarkRefillGammazeroStack.txt
name       old time/op    new time/op    delta
/1-4         4.10µs ±15%    3.67µs ± 1%   -10.49%  (p=0.003 n=10+10)
/10-4        35.7µs ± 2%    36.0µs ± 2%      ~     (p=0.353 n=10+10)
/100-4        353µs ±10%     506µs ± 1%   +43.26%  (p=0.000 n=10+10)
/1000-4      3.43ms ± 5%    4.45ms ± 1%   +29.53%  (p=0.000 n=9+10)
/10000-4     38.2ms ±11%    51.1ms ± 1%   +33.82%  (p=0.000 n=9+10)
/100000-4     402ms ± 5%     694ms ± 2%   +72.45%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +336.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.44MB ± 0%  +302.39%  (p=0.000 n=9+10)
/10000-4     30.1MB ± 0%    94.6MB ± 0%  +214.28%  (p=0.000 n=8+10)
/100000-4     320MB ± 0%     789MB ± 0%  +146.61%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.48%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.2.txt testdata/BenchmarkRefillFullGammazeroStack.txt
name       old time/op    new time/op    delta
/1-4         4.33µs ± 5%    3.77µs ± 7%  -13.09%  (p=0.000 n=9+9)
/10-4        35.8µs ±11%    38.6µs ± 5%   +7.83%  (p=0.001 n=10+9)
/100-4        364µs ±17%     365µs ± 3%     ~     (p=0.481 n=10+10)
/1000-4      3.40ms ± 3%    3.57ms ± 1%   +5.10%  (p=0.000 n=10+9)
/10000-4     36.7ms ± 0%    37.5ms ± 7%     ~     (p=0.258 n=9+9)
/100000-4     407ms ±14%     691ms ± 3%  +69.85%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.5MB ± 0%    16.0MB ± 0%  -47.55%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     632MB ± 0%  +97.48%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.70%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.76%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowIncreaseGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4           252ns ± 5%     216ns ± 9%   -14.30%  (p=0.000 n=10+9)
/10-4         1.02µs ± 5%    0.85µs ± 6%   -16.99%  (p=0.000 n=8+10)
/100-4        7.95µs ± 5%    8.69µs ± 5%    +9.37%  (p=0.000 n=10+10)
/1000-4       68.8µs ± 4%    80.1µs ± 5%   +16.44%  (p=0.000 n=10+9)
/10000-4       723µs ± 9%     944µs ± 6%   +30.53%  (p=0.000 n=10+10)
/100000-4     8.33ms ± 4%   11.52ms ± 8%   +38.18%  (p=0.000 n=9+9)
/1000000-4    89.7ms ± 6%   124.7ms ± 3%   +38.95%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      336B ± 0%   +61.54%  (p=0.000 n=10+10)
/10-4           752B ± 0%      624B ± 0%   -17.02%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    8.88kB ± 0%   +14.20%  (p=0.000 n=10+10)
/1000-4       49.0kB ± 0%    80.7kB ± 0%   +64.64%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%    1106kB ± 0%  +127.48%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.88%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      4.00 ± 0%   -33.33%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      22.0 ± 0%   -12.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       208 ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%    +0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.31%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowDecreaseGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    39.0ns ±24%  +15.08%  (p=0.000 n=8+9)
/10-4          353ns ± 5%     430ns ±34%  +21.89%  (p=0.000 n=9+10)
/100-4        3.45µs ± 5%    3.61µs ± 3%   +4.71%  (p=0.001 n=9+10)
/1000-4       33.9µs ± 2%    36.9µs ± 4%   +8.77%  (p=0.000 n=9+9)
/10000-4       342µs ± 3%     371µs ± 3%   +8.38%  (p=0.000 n=8+9)
/100000-4     3.56ms ±10%    3.70ms ± 4%     ~     (p=0.052 n=10+10)
/1000000-4    36.3ms ±13%    37.0ms ± 3%     ~     (p=0.190 n=10+10)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.2.txt testdata/BenchmarkStableGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 8%    37.2ns ± 7%  -5.05%  (p=0.041 n=10+10)
/10-4          368ns ± 2%     358ns ± 3%  -2.83%  (p=0.002 n=9+10)
/100-4        3.65µs ± 4%    3.55µs ± 7%    ~     (p=0.053 n=10+9)
/1000-4       36.6µs ± 3%    35.5µs ± 4%  -3.01%  (p=0.013 n=9+10)
/10000-4       381µs ±12%     351µs ± 4%  -8.04%  (p=0.001 n=10+10)
/100000-4     3.55ms ± 2%    3.67ms ±13%    ~     (p=0.321 n=8+9)
/1000000-4    37.3ms ± 7%    35.5ms ± 5%  -4.67%  (p=0.006 n=9+10)

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

#### deque vs Juju - FIFO queue
deque vs juju - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.2.txt testdata/BenchmarkFillJujuQueue.txt
name        old time/op    new time/op    delta
/0-4          36.1ns ± 4%   376.2ns ± 5%   +941.53%  (p=0.000 n=10+10)
/1-4           143ns ± 5%     400ns ± 1%   +180.18%  (p=0.000 n=10+10)
/10-4          664ns ± 9%     741ns ± 0%    +11.56%  (p=0.000 n=10+9)
/100-4        4.83µs ± 7%    4.57µs ± 0%     -5.26%  (p=0.002 n=10+9)
/1000-4       42.4µs ±15%    40.7µs ± 0%       ~     (p=0.173 n=10+8)
/10000-4       417µs ±11%     403µs ± 2%       ~     (p=0.853 n=10+10)
/100000-4     4.45ms ±27%    4.34ms ± 1%       ~     (p=0.481 n=10+10)
/1000000-4    51.9ms ±12%    59.1ms ± 4%    +13.96%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%   1184.0B ± 0%  +2366.67%  (p=0.000 n=10+10)
/1-4            128B ± 0%     1200B ± 0%   +837.50%  (p=0.000 n=10+10)
/10-4           592B ± 0%     1344B ± 0%   +127.03%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    4.99kB ± 0%    -19.17%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    34.8kB ± 0%     +5.57%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     333kB ± 0%     +3.53%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.33MB ± 0%     +3.26%  (p=0.000 n=9+10)
/1000000-4    32.2MB ± 0%    33.3MB ± 0%     +3.29%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      4.00 ± 0%   +300.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      5.00 ± 0%    +25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      14.0 ± 0%     -6.67%  (p=0.000 n=10+10)
/100-4           107 ± 0%       110 ± 0%     +2.80%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.05k ± 0%     +3.85%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.5k ± 0%     +3.86%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      105k ± 0%     +3.87%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.05M ± 0%     +3.88%  (p=0.000 n=10+10)
```

deque vs juju - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.2.txt testdata/BenchmarkRefillJujuQueue.txt
name       old time/op    new time/op    delta
/1-4         3.97µs ± 6%    4.07µs ± 6%      ~     (p=0.133 n=9+10)
/10-4        36.8µs ± 5%    45.4µs ± 9%   +23.39%  (p=0.000 n=10+10)
/100-4        347µs ± 5%     464µs ±11%   +33.70%  (p=0.000 n=10+10)
/1000-4      3.67ms ±15%    4.53ms ±10%   +23.54%  (p=0.000 n=10+10)
/10000-4     39.7ms ± 5%    49.5ms ±12%   +24.81%  (p=0.000 n=10+10)
/100000-4     413ms ± 6%     493ms ±10%   +19.26%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    29.8kB ± 0%   +86.25%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     326kB ± 0%  +103.50%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    3.31MB ± 0%  +106.94%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    33.2MB ± 0%   +10.31%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     332MB ± 0%    +3.95%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.04k ± 0%    +3.70%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.4k ± 0%    +4.50%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      105k ± 0%    +4.65%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.05M ± 0%    +3.97%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.5M ± 0%    +3.88%  (p=0.000 n=9+10)
```

deque vs juju - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt testdata/BenchmarkRefillFullJujuQueue.txt
name       old time/op    new time/op    delta
/1-4         3.70µs ± 8%    4.31µs ± 1%   +16.54%  (p=0.000 n=10+10)
/10-4        37.4µs ± 7%    43.7µs ± 1%   +16.86%  (p=0.000 n=10+10)
/100-4        360µs ± 4%     421µs ± 1%   +17.01%  (p=0.000 n=10+10)
/1000-4      3.50ms ± 3%    4.21ms ± 1%   +20.32%  (p=0.000 n=9+10)
/10000-4     40.0ms ± 5%    41.5ms ± 1%    +3.54%  (p=0.004 n=9+9)
/100000-4     426ms ±11%     428ms ± 1%      ~     (p=0.258 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    3.33kB ± 0%  +107.81%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    33.2kB ± 0%  +107.81%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     332kB ± 0%  +107.81%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.33MB ± 0%  +107.81%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    33.2MB ± 0%   +10.41%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     332MB ± 0%    +3.97%  (p=0.000 n=8+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       104 ± 0%    +4.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.05k ± 0%    +4.60%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.5k ± 0%    +4.68%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      105k ± 0%    +4.69%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.05M ± 0%    +3.98%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.5M ± 0%    +3.89%  (p=0.000 n=9+10)
```

deque vs juju - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowIncreaseJujuQueue.txt
name        old time/op    new time/op    delta
/1-4           248ns ± 5%     549ns ±19%  +121.49%  (p=0.000 n=9+10)
/10-4         2.09µs ±13%    1.22µs ±10%   -41.37%  (p=0.000 n=9+10)
/100-4        8.06µs ± 5%    9.57µs ±17%   +18.80%  (p=0.000 n=10+10)
/1000-4       73.5µs ± 5%    93.7µs ±15%   +27.62%  (p=0.000 n=9+10)
/10000-4       703µs ± 1%     878µs ±11%   +24.89%  (p=0.000 n=8+10)
/100000-4     8.45ms ± 8%   10.67ms ± 6%   +26.16%  (p=0.000 n=10+9)
/1000000-4    90.6ms ± 6%   118.5ms ±13%   +30.74%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%     1216B ± 0%  +484.62%  (p=0.000 n=10+10)
/10-4         4.90kB ± 0%    1.50kB ± 0%   -69.28%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    7.70kB ± 0%    -1.03%  (p=0.000 n=10+10)
/1000-4       53.2kB ± 0%    67.4kB ± 0%   +26.82%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     666kB ± 0%   +36.91%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    6.65MB ± 0%   +37.85%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    66.5MB ± 0%   +37.99%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%      ~     (all equal)
/10-4           27.0 ± 0%      24.0 ± 0%   -11.11%  (p=0.000 n=10+10)
/100-4           207 ± 0%       213 ± 0%    +2.90%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.10k ± 0%    +4.07%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.9k ± 0%    +4.26%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      209k ± 0%    +4.28%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.09M ± 0%    +4.28%  (p=0.000 n=10+10)
```

deque vs juju - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowDecreaseJujuQueue.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 3%    38.9ns ± 8%  +10.35%  (p=0.000 n=10+10)
/10-4          353ns ± 2%     394ns ± 9%  +11.71%  (p=0.000 n=9+10)
/100-4        3.61µs ±13%    3.95µs ± 8%   +9.34%  (p=0.008 n=9+10)
/1000-4       35.2µs ± 2%    38.6µs ± 9%   +9.68%  (p=0.001 n=8+10)
/10000-4       356µs ± 7%     406µs ±14%  +13.87%  (p=0.000 n=9+10)
/100000-4     3.51ms ± 3%    3.83ms ± 8%   +9.13%  (p=0.001 n=10+10)
/1000000-4    34.3ms ± 2%    38.6ms ±14%  +12.34%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (p=0.173 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs juju - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueuev1.0.2.txt testdata/BenchmarkStableJujuQueue.txt
name        old time/op    new time/op    delta
/1-4          35.1ns ± 2%    46.0ns ±11%   +30.98%  (p=0.000 n=9+10)
/10-4          358ns ± 4%     488ns ±15%   +36.30%  (p=0.000 n=9+10)
/100-4        3.44µs ± 2%    4.70µs ±11%   +36.49%  (p=0.000 n=9+10)
/1000-4       34.2µs ± 2%    45.1µs ±11%   +31.98%  (p=0.000 n=9+10)
/10000-4       357µs ± 6%     453µs ±10%   +26.64%  (p=0.000 n=10+10)
/100000-4     3.49ms ± 4%    4.59ms ±12%   +31.56%  (p=0.000 n=10+10)
/1000000-4    35.6ms ± 3%    45.1ms ± 9%   +26.69%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     33.0B ± 0%  +106.25%  (p=0.000 n=10+10)
/10-4           160B ± 0%      332B ± 0%  +107.50%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    3.33kB ± 0%  +107.81%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    33.2kB ± 0%  +107.81%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     332kB ± 0%  +107.81%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    3.33MB ± 0%  +107.81%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    33.2MB ± 0%  +107.81%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       104 ± 0%    +4.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     1.05k ± 0%    +4.60%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.5k ± 0%    +4.68%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      105k ± 0%    +4.69%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.05M ± 0%    +4.69%  (p=0.000 n=10+10)
```

#### deque vs juju - LIFO stack
deque vs juju - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStackv1.0.2.txt testdata/BenchmarkFillJujuStack.txt
name        old time/op    new time/op    delta
/0-4          38.3ns ± 3%   359.9ns ± 1%   +839.04%  (p=0.000 n=8+9)
/1-4           148ns ± 9%     404ns ± 0%   +172.24%  (p=0.000 n=10+9)
/10-4          662ns ±13%     737ns ± 0%    +11.40%  (p=0.002 n=10+8)
/100-4        4.99µs ± 6%    4.58µs ± 1%     -8.35%  (p=0.000 n=9+10)
/1000-4       39.5µs ± 8%    40.8µs ± 1%       ~     (p=0.353 n=10+10)
/10000-4       387µs ± 9%     401µs ± 0%       ~     (p=0.075 n=10+10)
/100000-4     4.13ms ± 9%    4.37ms ± 1%     +5.92%  (p=0.028 n=10+9)
/1000000-4    46.8ms ± 4%    58.8ms ± 2%    +25.70%  (p=0.000 n=8+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%   1184.0B ± 0%  +2366.67%  (p=0.000 n=10+10)
/1-4            128B ± 0%     1200B ± 0%   +837.50%  (p=0.000 n=10+10)
/10-4           592B ± 0%     1344B ± 0%   +127.03%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    4.99kB ± 0%    -19.17%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    34.8kB ± 0%     +5.57%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     333kB ± 0%     +3.53%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.33MB ± 0%     +3.26%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    33.3MB ± 0%     +3.29%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      4.00 ± 0%   +300.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      5.00 ± 0%    +25.00%  (p=0.000 n=10+10)
/10-4           15.0 ± 0%      14.0 ± 0%     -6.67%  (p=0.000 n=10+10)
/100-4           107 ± 0%       110 ± 0%     +2.80%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.05k ± 0%     +3.85%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.5k ± 0%     +3.86%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      105k ± 0%     +3.87%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.05M ± 0%     +3.88%  (p=0.000 n=10+10)
```

deque vs juju - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStackv1.0.2.txt testdata/BenchmarkRefillJujuStack.txt
name       old time/op    new time/op    delta
/1-4         4.10µs ±15%    4.21µs ± 9%      ~     (p=0.493 n=10+10)
/10-4        35.7µs ± 2%    40.1µs ± 8%   +12.19%  (p=0.000 n=10+10)
/100-4        353µs ±10%     504µs ±14%   +42.67%  (p=0.000 n=10+10)
/1000-4      3.43ms ± 5%    4.57ms ±10%   +33.12%  (p=0.000 n=9+10)
/10000-4     38.2ms ±11%    48.2ms ±13%   +26.16%  (p=0.000 n=9+10)
/100000-4     402ms ± 5%     490ms ±11%   +21.81%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     381kB ± 0%  +138.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    3.37MB ± 0%  +110.40%  (p=0.000 n=9+9)
/10000-4     30.1MB ± 0%    33.2MB ± 0%   +10.41%  (p=0.000 n=8+10)
/100000-4     320MB ± 0%     332MB ± 0%    +3.90%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      105k ± 0%    +4.80%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.05M ± 0%    +3.97%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.5M ± 0%    +3.88%  (p=0.000 n=10+10)
```

deque vs juju - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.2.txt testdata/BenchmarkRefillFullJujuStack.txt
name       old time/op    new time/op    delta
/1-4         4.33µs ± 5%    4.46µs ±19%      ~     (p=0.150 n=9+10)
/10-4        35.8µs ±11%    40.6µs ± 8%   +13.62%  (p=0.000 n=10+10)
/100-4        364µs ±17%     517µs ± 4%   +42.08%  (p=0.000 n=10+8)
/1000-4      3.40ms ± 3%    4.71ms ±10%   +38.64%  (p=0.000 n=10+10)
/10000-4     36.7ms ± 0%    48.9ms ±12%   +33.17%  (p=0.000 n=9+10)
/100000-4     407ms ±14%     505ms ±14%   +24.23%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     381kB ± 0%  +138.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.37MB ± 0%  +110.40%  (p=0.000 n=10+10)
/10000-4     30.5MB ± 0%    33.2MB ± 0%    +8.91%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     333MB ± 0%    +3.94%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      105k ± 0%    +4.80%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.05M ± 0%    +3.95%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.5M ± 0%    +3.89%  (p=0.000 n=10+10)
```

deque vs juju - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowIncreaseJujuStack.txt
name        old time/op    new time/op    delta
/1-4           252ns ± 5%     546ns ±21%  +116.41%  (p=0.000 n=10+10)
/10-4         1.02µs ± 5%    1.19µs ±14%   +17.05%  (p=0.000 n=8+10)
/100-4        7.95µs ± 5%    9.92µs ±12%   +24.84%  (p=0.000 n=10+10)
/1000-4       68.8µs ± 4%    92.3µs ±15%   +34.19%  (p=0.000 n=10+10)
/10000-4       723µs ± 9%     882µs ± 9%   +21.92%  (p=0.000 n=10+10)
/100000-4     8.33ms ± 4%   10.95ms ±12%   +31.44%  (p=0.000 n=9+10)
/1000000-4    89.7ms ± 6%   129.9ms ±10%   +44.75%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%     1216B ± 0%  +484.62%  (p=0.000 n=10+10)
/10-4           752B ± 0%     1504B ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    8.80kB ± 0%   +13.17%  (p=0.000 n=10+10)
/1000-4       49.0kB ± 0%    68.5kB ± 0%   +39.80%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     666kB ± 0%   +36.91%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    6.65MB ± 0%   +37.97%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    66.5MB ± 0%   +37.99%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%      ~     (all equal)
/10-4           25.0 ± 0%      24.0 ± 0%    -4.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       216 ± 0%    +4.35%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.10k ± 0%    +4.32%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.9k ± 0%    +4.26%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      209k ± 0%    +4.28%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.09M ± 0%    +4.28%  (p=0.000 n=10+10)
```

deque vs juju - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowDecreaseJujuStack.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    39.0ns ± 8%  +14.93%  (p=0.000 n=8+10)
/10-4          353ns ± 5%     412ns ±14%  +16.82%  (p=0.000 n=9+10)
/100-4        3.45µs ± 5%    3.87µs ± 7%  +12.00%  (p=0.000 n=9+10)
/1000-4       33.9µs ± 2%    37.7µs ± 8%  +11.17%  (p=0.000 n=9+10)
/10000-4       342µs ± 3%     379µs ± 8%  +10.80%  (p=0.000 n=8+10)
/100000-4     3.56ms ±10%    3.88ms ± 9%   +8.97%  (p=0.002 n=10+10)
/1000000-4    36.3ms ±13%    38.6ms ± 8%   +6.55%  (p=0.029 n=10+10)

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

deque vs juju - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStackv1.0.2.txt testdata/BenchmarkStableJujuStack.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 8%    38.1ns ± 9%     ~     (p=0.383 n=10+10)
/10-4          368ns ± 2%     406ns ±13%  +10.35%  (p=0.004 n=9+10)
/100-4        3.65µs ± 4%    3.81µs ± 8%   +4.13%  (p=0.023 n=10+10)
/1000-4       36.6µs ± 3%    39.4µs ±13%   +7.85%  (p=0.017 n=9+10)
/10000-4       381µs ±12%     388µs ±15%     ~     (p=0.631 n=10+10)
/100000-4     3.55ms ± 2%    3.80ms ±10%     ~     (p=0.055 n=8+10)
/1000000-4    37.3ms ± 7%    37.3ms ± 9%     ~     (p=0.968 n=9+10)

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
benchstat testdata/BenchmarkFillDequeQueuev1.0.2.txt testdata/BenchmarkFillCookiejarQueue.txt
name        old time/op    new time/op     delta
/0-4          36.1ns ± 4%  10159.8ns ± 2%   +28027.77%  (p=0.000 n=10+8)
/1-4           143ns ± 5%    10443ns ±10%    +7213.17%  (p=0.000 n=10+10)
/10-4          664ns ± 9%    10646ns ±14%    +1502.06%  (p=0.000 n=10+9)
/100-4        4.83µs ± 7%    13.65µs ± 6%     +182.79%  (p=0.000 n=10+9)
/1000-4       42.4µs ±15%     42.2µs ±10%         ~     (p=0.912 n=10+10)
/10000-4       417µs ±11%      342µs ± 6%      -17.94%  (p=0.000 n=10+10)
/100000-4     4.45ms ±27%     3.87ms ±14%      -12.95%  (p=0.001 n=10+10)
/1000000-4    51.9ms ±12%     47.5ms ± 8%       -8.53%  (p=0.035 n=10+10)

name        old alloc/op   new alloc/op    delta
/0-4           48.0B ± 0%   65680.0B ± 0%  +136733.33%  (p=0.000 n=10+10)
/1-4            128B ± 0%     65696B ± 0%   +51225.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%     65840B ± 0%   +11021.62%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    67.28kB ± 0%     +989.38%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%     81.7kB ± 0%     +147.46%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%      357kB ± 0%      +10.82%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.25MB ± 0%       +0.80%  (p=0.000 n=9+10)
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
benchstat testdata/BenchmarkRefillDequeQueuev1.0.2.txt testdata/BenchmarkRefillCookiejarQueue.txt
name       old time/op    new time/op    delta
/1-4         3.97µs ± 6%    3.51µs ± 5%  -11.70%  (p=0.000 n=9+10)
/10-4        36.8µs ± 5%    32.7µs ± 6%  -11.02%  (p=0.000 n=10+10)
/100-4        347µs ± 5%     311µs ± 2%  -10.47%  (p=0.000 n=10+10)
/1000-4      3.67ms ±15%    3.05ms ± 3%  -16.68%  (p=0.000 n=10+10)
/10000-4     39.7ms ± 5%    31.3ms ± 4%  -21.07%  (p=0.000 n=10+10)
/100000-4     413ms ± 6%     365ms ± 7%  -11.57%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%   +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%   +0.02%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.01%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.86%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     161MB ± 0%  -49.79%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=9+10)
```

deque vs cookiejar - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt testdata/BenchmarkRefillFullCookiejarQueue.txt
name       old time/op    new time/op    delta
/1-4         3.70µs ± 8%    3.13µs ± 3%  -15.29%  (p=0.000 n=10+9)
/10-4        37.4µs ± 7%    32.2µs ± 7%  -13.75%  (p=0.000 n=10+10)
/100-4        360µs ± 4%     316µs ± 3%  -12.03%  (p=0.000 n=10+9)
/1000-4      3.50ms ± 3%    3.13ms ± 0%  -10.47%  (p=0.000 n=9+7)
/10000-4     40.0ms ± 5%    31.6ms ± 3%  -21.11%  (p=0.000 n=9+10)
/100000-4     426ms ±11%     366ms ± 7%  -13.96%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.87%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     160MB ± 0%  -49.97%  (p=0.002 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=9+10)
```

deque vs cookiejar - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowIncreaseCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4           248ns ± 5%   10208ns ± 2%   +4019.78%  (p=0.000 n=9+9)
/10-4         2.09µs ±13%   10.86µs ± 2%    +419.98%  (p=0.000 n=9+10)
/100-4        8.06µs ± 5%   16.48µs ± 4%    +104.55%  (p=0.000 n=10+10)
/1000-4       73.5µs ± 5%    69.0µs ± 2%      -6.04%  (p=0.000 n=9+10)
/10000-4       703µs ± 1%     679µs ±10%        ~     (p=0.055 n=8+10)
/100000-4     8.45ms ± 8%    8.12ms ± 5%      -4.00%  (p=0.019 n=10+10)
/1000000-4    90.6ms ± 6%    83.7ms ± 3%      -7.68%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%    65712B ± 0%  +31492.31%  (p=0.000 n=10+10)
/10-4         4.90kB ± 0%   66.00kB ± 0%   +1248.04%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%   68.88kB ± 0%    +785.80%  (p=0.000 n=10+10)
/1000-4       53.2kB ± 0%    97.7kB ± 0%     +83.77%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     583kB ± 0%     +19.81%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.91MB ± 0%      +1.82%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.9MB ± 0%      +1.47%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%     -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      23.0 ± 0%     -14.81%  (p=0.000 n=10+10)
/100-4           207 ± 0%       203 ± 0%      -1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.00k ± 0%      -0.60%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%      -0.38%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%      -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%      -0.36%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowDecreaseCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 3%    36.1ns ±12%    ~     (p=0.896 n=10+10)
/10-4          353ns ± 2%     341ns ± 2%  -3.31%  (p=0.000 n=9+9)
/100-4        3.61µs ±13%    3.32µs ± 2%  -7.95%  (p=0.000 n=9+10)
/1000-4       35.2µs ± 2%    33.4µs ± 2%  -5.04%  (p=0.000 n=8+10)
/10000-4       356µs ± 7%     332µs ± 1%  -6.86%  (p=0.000 n=9+9)
/100000-4     3.51ms ± 3%    3.30ms ± 2%  -5.91%  (p=0.000 n=10+9)
/1000000-4    34.3ms ± 2%    33.5ms ± 6%  -2.40%  (p=0.008 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.211 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs cookiejar - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueuev1.0.2.txt testdata/BenchmarkStableCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4          35.1ns ± 2%    30.8ns ± 3%  -12.28%  (p=0.000 n=9+10)
/10-4          358ns ± 4%     309ns ± 4%  -13.80%  (p=0.000 n=9+10)
/100-4        3.44µs ± 2%    3.03µs ± 2%  -11.91%  (p=0.000 n=9+10)
/1000-4       34.2µs ± 2%    30.1µs ± 3%  -11.99%  (p=0.000 n=9+9)
/10000-4       357µs ± 6%     311µs ±12%  -12.86%  (p=0.000 n=10+10)
/100000-4     3.49ms ± 4%    3.05ms ± 4%  -12.45%  (p=0.000 n=10+10)
/1000000-4    35.6ms ± 3%    30.3ms ± 5%  -14.94%  (p=0.000 n=9+9)

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
benchstat testdata/BenchmarkFillDequeStackv1.0.2.txt testdata/BenchmarkFillCookiejarStack.txt
name        old time/op    new time/op     delta
/0-4          38.3ns ± 3%  10254.3ns ± 3%   +26656.25%  (p=0.000 n=8+9)
/1-4           148ns ± 9%    10843ns ± 7%    +7206.83%  (p=0.000 n=10+9)
/10-4          662ns ±13%    11329ns ±12%    +1611.91%  (p=0.000 n=10+10)
/100-4        4.99µs ± 6%    13.91µs ± 8%     +178.56%  (p=0.000 n=9+10)
/1000-4       39.5µs ± 8%     43.8µs ± 5%      +11.07%  (p=0.000 n=10+9)
/10000-4       387µs ± 9%      344µs ± 2%      -11.09%  (p=0.000 n=10+10)
/100000-4     4.13ms ± 9%     3.73ms ± 5%       -9.56%  (p=0.000 n=10+10)
/1000000-4    46.8ms ± 4%     46.7ms ± 3%         ~     (p=0.963 n=8+9)

name        old alloc/op   new alloc/op    delta
/0-4           48.0B ± 0%   65680.0B ± 0%  +136733.33%  (p=0.000 n=10+10)
/1-4            128B ± 0%     65696B ± 0%   +51225.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%     65840B ± 0%   +11021.62%  (p=0.000 n=10+10)
/100-4        6.18kB ± 0%    67.28kB ± 0%     +989.38%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%     81.7kB ± 0%     +147.46%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%      357kB ± 0%      +10.82%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.25MB ± 0%       +0.80%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%     32.8MB ± 0%       +1.99%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkRefillDequeStackv1.0.2.txt testdata/BenchmarkRefillCookiejarStack.txt
name       old time/op    new time/op    delta
/1-4         4.10µs ±15%    3.63µs ± 6%  -11.62%  (p=0.003 n=10+10)
/10-4        35.7µs ± 2%    33.9µs ± 5%   -5.17%  (p=0.000 n=10+10)
/100-4        353µs ±10%     330µs ±11%   -6.53%  (p=0.009 n=10+10)
/1000-4      3.43ms ± 5%    3.38ms ± 6%     ~     (p=0.743 n=9+8)
/10000-4     38.2ms ±11%    36.5ms ±16%     ~     (p=0.079 n=9+10)
/100000-4     402ms ± 5%     400ms ±11%     ~     (p=0.912 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%   +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%   +0.01%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.01%  (p=0.000 n=9+10)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.81%  (p=0.000 n=8+9)
/100000-4     320MB ± 0%     161MB ± 0%  -49.82%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.2.txt testdata/BenchmarkRefillFullCookiejarStack.txt
name       old time/op    new time/op    delta
/1-4         4.33µs ± 5%    3.26µs ± 7%  -24.81%  (p=0.000 n=9+10)
/10-4        35.8µs ±11%    33.1µs ± 9%   -7.50%  (p=0.003 n=10+10)
/100-4        364µs ±17%     305µs ± 2%  -16.12%  (p=0.000 n=10+8)
/1000-4      3.40ms ± 3%    3.18ms ± 5%   -6.46%  (p=0.000 n=10+9)
/10000-4     36.7ms ± 0%    31.5ms ± 6%  -14.37%  (p=0.000 n=9+9)
/100000-4     407ms ±14%     354ms ± 3%  -13.00%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.5MB ± 0%    16.0MB ± 0%  -47.55%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     160MB ± 0%  -49.99%  (p=0.002 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.70%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowIncreaseCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4           252ns ± 5%   10375ns ± 3%   +4012.13%  (p=0.000 n=10+10)
/10-4         1.02µs ± 5%   11.12µs ± 3%    +990.70%  (p=0.000 n=8+10)
/100-4        7.95µs ± 5%   16.70µs ± 4%    +110.07%  (p=0.000 n=10+9)
/1000-4       68.8µs ± 4%    71.1µs ± 3%      +3.38%  (p=0.002 n=10+10)
/10000-4       723µs ± 9%     643µs ± 2%     -11.05%  (p=0.000 n=10+10)
/100000-4     8.33ms ± 4%    8.24ms ± 5%        ~     (p=0.400 n=9+10)
/1000000-4    89.7ms ± 6%    89.0ms ± 3%        ~     (p=0.842 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%    65712B ± 0%  +31492.31%  (p=0.000 n=10+10)
/10-4           752B ± 0%    66000B ± 0%   +8676.60%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%   68.88kB ± 0%    +785.80%  (p=0.000 n=10+10)
/1000-4       49.0kB ± 0%    97.7kB ± 0%     +99.31%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     517kB ± 0%      +6.31%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.85MB ± 0%      +0.54%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.8MB ± 0%      +1.33%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%     -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      23.0 ± 0%      -8.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       203 ± 0%      -1.93%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.00k ± 0%      -0.50%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%      -0.39%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%      -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%      -0.36%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowDecreaseCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    35.9ns ±10%  +5.78%  (p=0.000 n=8+10)
/10-4          353ns ± 5%     378ns ±17%  +7.04%  (p=0.004 n=9+10)
/100-4        3.45µs ± 5%    3.63µs ±13%    ~     (p=0.053 n=9+10)
/1000-4       33.9µs ± 2%    34.6µs ± 2%  +2.04%  (p=0.001 n=9+8)
/10000-4       342µs ± 3%     344µs ± 1%    ~     (p=0.277 n=8+9)
/100000-4     3.56ms ±10%    3.50ms ± 5%    ~     (p=0.971 n=10+10)
/1000000-4    36.3ms ±13%    34.4ms ± 2%  -5.23%  (p=0.023 n=10+10)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.2.txt testdata/BenchmarkStableCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 8%    30.8ns ± 4%  -21.46%  (p=0.000 n=10+10)
/10-4          368ns ± 2%     304ns ± 2%  -17.39%  (p=0.000 n=9+8)
/100-4        3.65µs ± 4%    2.99µs ± 2%  -18.30%  (p=0.000 n=10+9)
/1000-4       36.6µs ± 3%    30.5µs ± 9%  -16.43%  (p=0.000 n=9+10)
/10000-4       381µs ±12%     300µs ± 4%  -21.37%  (p=0.000 n=10+10)
/100000-4     3.55ms ± 2%    2.94ms ± 2%  -17.36%  (p=0.000 n=8+10)
/1000000-4    37.3ms ± 7%    30.1ms ± 3%  -19.19%  (p=0.000 n=9+10)

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
