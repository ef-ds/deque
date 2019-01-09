# v1.0.0 vs v1.0.1

The performance impact is small with a slightly better performance on the queue tests, but slightly worse in the stack tests.

The overall memory impact is positive with deque now using ~5% less memory for large data sets (>= 10000), but now that the spare slices logic is actually correct, the refill tests for large data sets took a big hit for data sets above 1000 items (deque uses now ~2x as much memory on this scenario). This is expected as with a maxSpareLinks set to 4 and maxInternalSliceSize to 256, the spare slices should be reused only for the first 1000 items or so. Past that, new slices have to be created, which explains the extra memory footprint.

```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillDequeQueue2.txt
name        old time/op    new time/op    delta
/0-4          39.9ns ± 9%    37.4ns ± 1%   -6.06%  (p=0.000 n=10+9)
/1-4           142ns ± 1%     144ns ± 7%     ~     (p=0.896 n=10+9)
/10-4          636ns ± 1%     664ns ± 9%   +4.39%  (p=0.018 n=9+10)
/100-4        4.74µs ± 3%    4.77µs ± 6%     ~     (p=0.922 n=9+10)
/1000-4       43.0µs ±23%    38.2µs ± 2%  -11.16%  (p=0.000 n=10+9)
/10000-4       450µs ±19%     396µs ± 5%  -11.92%  (p=0.004 n=10+10)
/100000-4     4.24ms ± 4%    4.09ms ± 4%   -3.38%  (p=0.011 n=10+10)
/1000000-4    46.8ms ± 1%    48.8ms ±10%   +4.21%  (p=0.021 n=8+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%     ~     (all equal)
/1-4            144B ± 0%      144B ± 0%     ~     (all equal)
/10-4           608B ± 0%      608B ± 0%     ~     (all equal)
/100-4        6.19kB ± 0%    6.19kB ± 0%     ~     (all equal)
/1000-4       33.0kB ± 0%    33.0kB ± 0%     ~     (all equal)
/10000-4       322kB ± 0%     322kB ± 0%     ~     (all equal)
/100000-4     3.22MB ± 0%    3.22MB ± 0%     ~     (all equal)
/1000000-4    32.2MB ± 0%    32.2MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           15.0 ± 0%      15.0 ± 0%     ~     (all equal)
/100-4           107 ± 0%       107 ± 0%     ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%     ~     (all equal)
/10000-4       10.1k ± 0%     10.1k ± 0%     ~     (all equal)
/100000-4       101k ± 0%      101k ± 0%     ~     (all equal)
/1000000-4     1.01M ± 0%     1.01M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillDequeStack2.txt
name        old time/op    new time/op    delta
/0-4          39.0ns ± 6%    39.7ns ± 5%    ~     (p=0.062 n=9+10)
/1-4           145ns ± 2%     148ns ± 4%  +2.10%  (p=0.034 n=8+9)
/10-4          615ns ± 0%     635ns ± 2%  +3.21%  (p=0.000 n=8+8)
/100-4        4.42µs ± 1%    4.53µs ± 2%  +2.61%  (p=0.000 n=10+10)
/1000-4       36.2µs ± 1%    37.9µs ± 4%  +4.85%  (p=0.000 n=10+10)
/10000-4       383µs ± 7%     392µs ± 6%    ~     (p=0.280 n=10+10)
/100000-4     4.10ms ± 9%    4.08ms ± 5%    ~     (p=0.739 n=10+10)
/1000000-4    44.5ms ± 4%    47.6ms ± 7%  +6.96%  (p=0.003 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%    ~     (all equal)
/1-4            144B ± 0%      144B ± 0%    ~     (all equal)
/10-4           608B ± 0%      608B ± 0%    ~     (all equal)
/100-4        6.19kB ± 0%    6.19kB ± 0%    ~     (all equal)
/1000-4       33.0kB ± 0%    33.0kB ± 0%    ~     (all equal)
/10000-4       322kB ± 0%     322kB ± 0%    ~     (all equal)
/100000-4     3.22MB ± 0%    3.22MB ± 0%    ~     (all equal)
/1000000-4    32.2MB ± 0%    32.2MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%    ~     (all equal)
/10-4           15.0 ± 0%      15.0 ± 0%    ~     (all equal)
/100-4           107 ± 0%       107 ± 0%    ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%    ~     (all equal)
/10000-4       10.1k ± 0%     10.1k ± 0%    ~     (all equal)
/100000-4       101k ± 0%      101k ± 0%    ~     (all equal)
/1000000-4     1.01M ± 0%     1.01M ± 0%    ~     (all equal)
```

```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceDequeQueue2.txt
name        old time/op    new time/op    delta
/1-4           525ns ± 4%     512ns ± 3%  -2.38%  (p=0.041 n=10+10)
/10-4         3.66µs ± 5%    3.65µs ± 3%    ~     (p=0.921 n=10+9)
/100-4        25.9µs ± 2%    25.6µs ± 3%    ~     (p=0.089 n=10+10)
/1000-4        241µs ± 3%     233µs ± 3%  -3.36%  (p=0.003 n=10+9)
/10000-4      2.44ms ± 2%    2.47ms ± 6%    ~     (p=0.356 n=10+9)
/100000-4     27.1ms ± 3%    26.9ms ± 3%    ~     (p=0.460 n=10+8)
/1000000-4     282ms ± 6%     270ms ± 2%  -4.23%  (p=0.003 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            560B ± 0%      560B ± 0%    ~     (all equal)
/10-4         5.71kB ± 0%    5.71kB ± 0%    ~     (all equal)
/100-4        20.9kB ± 0%    20.9kB ± 0%    ~     (all equal)
/1000-4        138kB ± 0%     134kB ± 0%  -3.00%  (p=0.000 n=10+10)
/10000-4      1.54MB ± 0%    1.44MB ± 0%  -6.21%  (p=0.000 n=10+10)
/100000-4     15.3MB ± 0%    14.4MB ± 0%  -5.43%  (p=0.000 n=9+10)
/1000000-4     152MB ± 0%     144MB ± 0%  -5.32%  (p=0.002 n=8+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      12.0 ± 0%    ~     (all equal)
/10-4           77.0 ± 0%      77.0 ± 0%    ~     (all equal)
/100-4           709 ± 0%       709 ± 0%    ~     (all equal)
/1000-4        7.02k ± 0%     7.01k ± 0%  -0.03%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.2k ± 0%  -0.07%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      702k ± 0%  -0.06%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.02M ± 0%  -0.06%  (p=0.000 n=10+10)
```

```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceDequeStack2.txt
name        old time/op    new time/op    delta
/1-4           419ns ± 8%     413ns ± 7%     ~     (p=0.446 n=9+10)
/10-4         2.61µs ± 6%    2.70µs ± 6%   +3.45%  (p=0.034 n=10+8)
/100-4        25.2µs ± 5%    26.7µs ± 2%   +5.83%  (p=0.000 n=10+10)
/1000-4        234µs ± 6%     234µs ± 5%     ~     (p=0.842 n=10+9)
/10000-4      2.36ms ± 7%    2.34ms ± 4%     ~     (p=0.780 n=10+9)
/100000-4     24.2ms ± 2%    26.7ms ± 8%  +10.19%  (p=0.000 n=9+10)
/1000000-4     249ms ± 2%     257ms ± 2%   +3.51%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            304B ± 0%      304B ± 0%     ~     (all equal)
/10-4         1.57kB ± 0%    1.57kB ± 0%     ~     (all equal)
/100-4        20.9kB ± 0%    20.9kB ± 0%     ~     (all equal)
/1000-4        130kB ± 0%     130kB ± 0%     ~     (all equal)
/10000-4      1.29MB ± 0%    1.29MB ± 0%     ~     (all equal)
/100000-4     12.8MB ± 0%    12.8MB ± 0%     ~     (p=0.173 n=10+9)
/1000000-4     128MB ± 0%     128MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      11.0 ± 0%     ~     (all equal)
/10-4           75.0 ± 0%      75.0 ± 0%     ~     (all equal)
/100-4           709 ± 0%       709 ± 0%     ~     (all equal)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.1k ± 0%     70.1k ± 0%     ~     (all equal)
/100000-4       701k ± 0%      701k ± 0%     ~     (all equal)
/1000000-4     7.01M ± 0%     7.01M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillDequeQueue2.txt
name       old time/op    new time/op    delta
/1-4         3.79µs ± 1%    3.78µs ± 4%     ~     (p=0.780 n=9+10)
/10-4        37.8µs ± 7%    39.9µs ±23%     ~     (p=0.912 n=10+10)
/100-4        361µs ± 7%     362µs ± 4%     ~     (p=0.931 n=9+9)
/1000-4      3.75ms ± 4%    3.72ms ± 4%     ~     (p=0.481 n=10+10)
/10000-4     36.5ms ± 3%    41.1ms ± 7%  +12.76%  (p=0.000 n=10+10)
/100000-4     380ms ± 1%     423ms ± 6%  +11.18%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   -0.00%  (p=0.002 n=8+10)
/1000-4      2.42MB ± 0%    1.60MB ± 0%  -33.77%  (p=0.000 n=10+9)
/10000-4     17.0MB ± 0%    30.5MB ± 0%  +78.96%  (p=0.000 n=8+10)
/100000-4     162MB ± 0%     320MB ± 0%  +97.98%  (p=0.000 n=8+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.39%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.01M ± 0%   +0.65%  (p=0.000 n=8+10)
/100000-4     10.0M ± 0%     10.1M ± 0%   +0.76%  (p=0.000 n=10+10)
```

```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillDequeStack2.txt
name       old time/op    new time/op    delta
/1-4         3.62µs ± 1%    3.77µs ± 8%  +4.14%  (p=0.002 n=10+10)
/10-4        34.1µs ± 5%    36.5µs ± 8%  +6.99%  (p=0.000 n=10+10)
/100-4        332µs ± 2%     346µs ± 2%  +4.47%  (p=0.000 n=10+10)
/1000-4      3.28ms ± 2%    3.41ms ± 1%  +4.12%  (p=0.000 n=10+9)
/10000-4     34.0ms ± 3%    34.4ms ± 2%    ~     (p=0.065 n=9+10)
/100000-4     370ms ± 5%     378ms ± 4%  +2.29%  (p=0.028 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.173 n=10+9)
/100000-4     161MB ± 0%     161MB ± 0%  -0.00%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%    ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%  -0.00%  (p=0.000 n=10+10)
```

```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullDequeQueue2.txt
name       old time/op    new time/op    delta
/1-4         3.75µs ± 4%    3.78µs ± 5%     ~     (p=0.565 n=10+9)
/10-4        37.8µs ± 4%    38.9µs ± 7%     ~     (p=0.053 n=10+9)
/100-4        371µs ± 3%     377µs ± 9%     ~     (p=0.579 n=10+10)
/1000-4      4.02ms ± 5%    3.68ms ± 6%   -8.46%  (p=0.000 n=10+10)
/10000-4     39.4ms ± 5%    40.6ms ± 3%     ~     (p=0.053 n=10+9)
/100000-4     392ms ± 3%     485ms ±11%  +23.83%  (p=0.000 n=9+8)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      2.40MB ± 0%    1.60MB ± 0%  -33.41%  (p=0.000 n=10+10)
/10000-4     16.6MB ± 0%    30.5MB ± 0%  +83.39%  (p=0.000 n=10+10)
/100000-4     161MB ± 0%     320MB ± 0%  +98.91%  (p=0.000 n=8+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%   -0.39%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.01M ± 0%   +0.67%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.1M ± 0%   +0.77%  (p=0.000 n=10+7)
```

```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullDequeStack2.txt
name       old time/op    new time/op    delta
/1-4         3.90µs ± 6%    4.53µs ± 4%  +16.03%  (p=0.000 n=10+9)
/10-4        34.6µs ± 2%    36.4µs ± 5%   +5.07%  (p=0.000 n=9+10)
/100-4        337µs ± 4%     358µs ±10%   +6.21%  (p=0.031 n=9+9)
/1000-4      3.58ms ± 8%    3.43ms ± 1%     ~     (p=0.156 n=10+9)
/10000-4     34.9ms ± 7%    34.9ms ± 4%     ~     (p=0.780 n=10+9)
/100000-4     339ms ± 2%     353ms ± 8%   +4.32%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%   -0.00%  (p=0.013 n=10+8)
/100000-4     160MB ± 0%     160MB ± 0%     ~     (all equal)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%     ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseDequeQueue2.txt
name        old time/op    new time/op    delta
/1-4           244ns ± 1%     260ns ± 7%  +6.43%  (p=0.000 n=10+10)
/10-4         1.86µs ± 1%    1.91µs ± 8%  +2.93%  (p=0.045 n=10+10)
/100-4        8.02µs ± 1%    7.90µs ± 2%  -1.43%  (p=0.002 n=10+9)
/1000-4       73.5µs ± 1%    73.5µs ± 8%    ~     (p=0.447 n=10+9)
/10000-4       725µs ± 1%     714µs ± 2%  -1.55%  (p=0.015 n=10+10)
/100000-4     8.20ms ± 0%    8.36ms ± 8%    ~     (p=1.000 n=9+10)
/1000000-4    86.4ms ± 1%    87.6ms ±13%    ~     (p=0.720 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      224B ± 0%    ~     (all equal)
/10-4         4.91kB ± 0%    4.91kB ± 0%    ~     (all equal)
/100-4        7.79kB ± 0%    7.79kB ± 0%    ~     (all equal)
/1000-4       54.1kB ± 0%    54.1kB ± 0%    ~     (all equal)
/10000-4       491kB ± 0%     491kB ± 0%    ~     (all equal)
/100000-4     4.83MB ± 0%    4.83MB ± 0%    ~     (all equal)
/1000000-4    48.2MB ± 0%    48.2MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%    ~     (all equal)
/10-4           27.0 ± 0%      27.0 ± 0%    ~     (all equal)
/100-4           207 ± 0%       207 ± 0%    ~     (all equal)
/1000-4        2.02k ± 0%     2.02k ± 0%    ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%    ~     (all equal)
/100000-4       201k ± 0%      201k ± 0%    ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%    ~     (all equal)
```

```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseDequeStack2.txt
name        old time/op    new time/op    delta
/1-4           240ns ± 0%     245ns ± 2%   +1.84%  (p=0.000 n=9+8)
/10-4          929ns ± 1%     946ns ± 0%   +1.79%  (p=0.000 n=10+9)
/100-4        8.65µs ± 1%    8.94µs ± 1%   +3.37%  (p=0.000 n=10+10)
/1000-4       66.7µs ± 1%    69.6µs ± 5%   +4.39%  (p=0.000 n=10+9)
/10000-4       666µs ± 1%     695µs ± 1%   +4.35%  (p=0.000 n=10+9)
/100000-4     7.78ms ± 1%    8.28ms ± 6%   +6.45%  (p=0.000 n=9+10)
/1000000-4    81.9ms ± 1%    94.5ms ± 8%  +15.37%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      224B ± 0%     ~     (all equal)
/10-4           768B ± 0%      768B ± 0%     ~     (all equal)
/100-4        12.9kB ± 0%    12.9kB ± 0%     ~     (all equal)
/1000-4       50.0kB ± 0%    50.0kB ± 0%     ~     (all equal)
/10000-4       487kB ± 0%     487kB ± 0%     ~     (all equal)
/100000-4     4.82MB ± 0%    4.82MB ± 0%     ~     (all equal)
/1000000-4    48.2MB ± 0%    48.2MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%     ~     (all equal)
/10-4           25.0 ± 0%      25.0 ± 0%     ~     (all equal)
/100-4           209 ± 0%       209 ± 0%     ~     (all equal)
/1000-4        2.01k ± 0%     2.01k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%     ~     (all equal)
/100000-4       201k ± 0%      201k ± 0%     ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseDequeQueue2.txt
name        old time/op    new time/op    delta
/1-4          40.8ns ± 3%    35.7ns ± 1%  -12.65%  (p=0.000 n=9+10)
/10-4          413ns ± 1%     362ns ± 3%  -12.37%  (p=0.000 n=8+9)
/100-4        4.05µs ± 5%    3.73µs ± 8%   -7.78%  (p=0.000 n=10+10)
/1000-4       42.7µs ± 9%    36.3µs ± 9%  -15.08%  (p=0.000 n=10+9)
/10000-4       436µs ± 6%     355µs ± 2%  -18.58%  (p=0.000 n=9+10)
/100000-4     3.92ms ± 2%    3.57ms ± 2%   -9.05%  (p=0.000 n=9+10)
/1000000-4    39.2ms ± 1%    35.5ms ± 3%   -9.33%  (p=0.000 n=9+10)

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

```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseDequeStack2.txt
name        old time/op    new time/op    delta
/1-4           240ns ± 0%     245ns ± 2%   +1.84%  (p=0.000 n=9+8)
/10-4          929ns ± 1%     946ns ± 0%   +1.79%  (p=0.000 n=10+9)
/100-4        8.65µs ± 1%    8.94µs ± 1%   +3.37%  (p=0.000 n=10+10)
/1000-4       66.7µs ± 1%    69.6µs ± 5%   +4.39%  (p=0.000 n=10+9)
/10000-4       666µs ± 1%     695µs ± 1%   +4.35%  (p=0.000 n=10+9)
/100000-4     7.78ms ± 1%    8.28ms ± 6%   +6.45%  (p=0.000 n=9+10)
/1000000-4    81.9ms ± 1%    94.5ms ± 8%  +15.37%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      224B ± 0%     ~     (all equal)
/10-4           768B ± 0%      768B ± 0%     ~     (all equal)
/100-4        12.9kB ± 0%    12.9kB ± 0%     ~     (all equal)
/1000-4       50.0kB ± 0%    50.0kB ± 0%     ~     (all equal)
/10000-4       487kB ± 0%     487kB ± 0%     ~     (all equal)
/100000-4     4.82MB ± 0%    4.82MB ± 0%     ~     (all equal)
/1000000-4    48.2MB ± 0%    48.2MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%     ~     (all equal)
/10-4           25.0 ± 0%      25.0 ± 0%     ~     (all equal)
/100-4           209 ± 0%       209 ± 0%     ~     (all equal)
/1000-4        2.01k ± 0%     2.01k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%     ~     (all equal)
/100000-4       201k ± 0%      201k ± 0%     ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%     ~     (all equal)
```

```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableDequeQueue2.txt
name        old time/op    new time/op    delta
/1-4          34.8ns ± 1%    35.6ns ± 4%  +2.24%  (p=0.000 n=10+10)
/10-4          353ns ± 2%     363ns ± 2%  +2.84%  (p=0.000 n=10+10)
/100-4        3.45µs ± 1%    3.48µs ± 3%    ~     (p=0.097 n=10+9)
/1000-4       34.5µs ± 1%    35.5µs ± 4%  +2.89%  (p=0.001 n=10+10)
/10000-4       346µs ± 2%     349µs ± 3%    ~     (p=0.211 n=10+9)
/100000-4     3.43ms ± 1%    3.45ms ± 1%    ~     (p=0.075 n=10+10)
/1000000-4    34.4ms ± 1%    34.6ms ± 3%    ~     (p=0.853 n=10+10)

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

```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableDequeStack2.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 2%    37.8ns ± 2%  +6.99%  (p=0.000 n=10+9)
/10-4          355ns ± 0%     383ns ± 3%  +8.08%  (p=0.000 n=7+9)
/100-4        3.46µs ± 1%    3.72µs ± 2%  +7.33%  (p=0.000 n=10+9)
/1000-4       34.6µs ± 1%    37.9µs ± 5%  +9.42%  (p=0.000 n=8+10)
/10000-4       406µs ±15%     380µs ± 8%    ~     (p=0.079 n=10+9)
/100000-4     3.90ms ±12%    3.80ms ± 4%    ~     (p=0.113 n=10+9)
/1000000-4    38.3ms ±18%    40.2ms ±11%    ~     (p=0.143 n=10+10)

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
