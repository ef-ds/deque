# v1.0.1 vs v1.0.2
## Fill tests
### FIFO queue
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.1.txt testdata/BenchmarkFillDequeQueuev1.0.2.txt
name        old time/op    new time/op    delta
/0-4          38.6ns ± 3%    36.1ns ± 4%   -6.30%  (p=0.000 n=10+10)
/1-4           149ns ± 6%     143ns ± 5%   -4.23%  (p=0.011 n=10+10)
/10-4          631ns ± 1%     664ns ± 9%   +5.37%  (p=0.032 n=8+10)
/100-4        4.68µs ± 5%    4.83µs ± 7%     ~     (p=0.089 n=10+10)
/1000-4       38.0µs ± 1%    42.4µs ±15%  +11.61%  (p=0.000 n=9+10)
/10000-4       385µs ± 3%     417µs ±11%   +8.42%  (p=0.002 n=10+10)
/100000-4     4.00ms ± 1%    4.45ms ±27%  +11.41%  (p=0.002 n=10+10)
/1000000-4    45.2ms ± 2%    51.9ms ±12%  +14.67%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%  -25.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%      128B ± 0%  -11.11%  (p=0.000 n=10+10)
/10-4           608B ± 0%      592B ± 0%   -2.63%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    6.18kB ± 0%   -0.26%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    33.0kB ± 0%   -0.05%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     322kB ± 0%   -0.00%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.22MB ± 0%   -0.00%  (p=0.000 n=10+9)
/1000000-4    32.2MB ± 0%    32.2MB ± 0%   -0.00%  (p=0.000 n=10+10)

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

### LIFO stack
```
benchstat testdata/BenchmarkFillDequeStackv1.0.1.txt testdata/BenchmarkFillDequeStackv1.0.2.txt
name        old time/op    new time/op    delta
/0-4          38.2ns ± 6%    38.3ns ± 3%     ~     (p=0.712 n=10+8)
/1-4           146ns ± 8%     148ns ± 9%     ~     (p=0.645 n=9+10)
/10-4          636ns ± 3%     662ns ±13%     ~     (p=0.288 n=10+10)
/100-4        4.70µs ± 4%    4.99µs ± 6%   +6.13%  (p=0.000 n=10+9)
/1000-4       38.5µs ± 9%    39.5µs ± 8%     ~     (p=0.353 n=10+10)
/10000-4       382µs ± 5%     387µs ± 9%     ~     (p=0.739 n=10+10)
/100000-4     3.95ms ± 3%    4.13ms ± 9%   +4.35%  (p=0.015 n=10+10)
/1000000-4    45.2ms ± 3%    46.8ms ± 4%   +3.41%  (p=0.010 n=8+8)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%  -25.00%  (p=0.000 n=10+10)
/1-4            144B ± 0%      128B ± 0%  -11.11%  (p=0.000 n=10+10)
/10-4           608B ± 0%      592B ± 0%   -2.63%  (p=0.000 n=10+10)
/100-4        6.19kB ± 0%    6.18kB ± 0%   -0.26%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    33.0kB ± 0%   -0.05%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     322kB ± 0%   -0.00%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.22MB ± 0%   -0.00%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    32.2MB ± 0%   -0.00%  (p=0.000 n=10+9)

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


## Microservice tests
### FIFO queue
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.1.txt testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt
name        old time/op    new time/op    delta
/1-4           512ns ± 3%     510ns ± 5%     ~     (p=0.401 n=10+10)
/10-4         3.51µs ± 0%    3.60µs ± 5%     ~     (p=0.062 n=9+10)
/100-4        24.9µs ± 1%    24.8µs ± 7%     ~     (p=0.842 n=9+10)
/1000-4        243µs ± 5%     237µs ± 5%     ~     (p=0.053 n=10+9)
/10000-4      2.38ms ± 1%    2.38ms ± 5%     ~     (p=0.400 n=9+10)
/100000-4     25.7ms ± 3%    25.0ms ± 1%   -2.99%  (p=0.002 n=10+10)
/1000000-4     268ms ± 2%     268ms ± 4%     ~     (p=0.720 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            560B ± 0%      544B ± 0%   -2.86%  (p=0.000 n=10+10)
/10-4         5.71kB ± 0%    5.70kB ± 0%   -0.28%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    15.8kB ± 0%  -24.56%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     133kB ± 0%   -0.74%  (p=0.000 n=10+10)
/10000-4      1.44MB ± 0%    1.43MB ± 0%   -0.64%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    14.4MB ± 0%   -0.04%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     144MB ± 0%   -0.01%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      12.0 ± 0%     ~     (all equal)
/10-4           77.0 ± 0%      77.0 ± 0%     ~     (all equal)
/100-4           709 ± 0%       707 ± 0%   -0.28%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.2k ± 0%     70.2k ± 0%   -0.01%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      702k ± 0%   -0.00%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.02M ± 0%   -0.00%  (p=0.000 n=10+10)
```

### LIFO stack
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.1.txt testdata/BenchmarkMicroserviceDequeStackv1.0.2.txt
name        old time/op    new time/op    delta
/1-4           408ns ± 4%     389ns ± 1%   -4.68%  (p=0.000 n=10+10)
/10-4         2.58µs ± 1%    2.51µs ± 1%   -2.60%  (p=0.000 n=9+10)
/100-4        24.9µs ± 4%    22.8µs ± 1%   -8.27%  (p=0.000 n=10+10)
/1000-4        230µs ± 7%     220µs ± 2%   -4.49%  (p=0.001 n=10+9)
/10000-4      2.47ms ± 7%    2.28ms ± 2%   -7.85%  (p=0.000 n=10+8)
/100000-4     27.2ms ±11%    24.6ms ± 1%   -9.57%  (p=0.000 n=10+9)
/1000000-4     282ms ± 6%     258ms ± 1%   -8.27%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            304B ± 0%      288B ± 0%   -5.26%  (p=0.000 n=10+10)
/10-4         1.57kB ± 0%    1.55kB ± 0%   -1.02%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    15.8kB ± 0%  -24.56%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     129kB ± 0%   -0.76%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.43MB ± 0%   -0.07%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    14.4MB ± 0%   -0.01%  (p=0.000 n=9+10)
/1000000-4     144MB ± 0%     144MB ± 0%   -0.00%  (p=0.000 n=9+9)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      11.0 ± 0%     ~     (all equal)
/10-4           75.0 ± 0%      75.0 ± 0%     ~     (all equal)
/100-4           709 ± 0%       707 ± 0%   -0.28%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.2k ± 0%     70.2k ± 0%     ~     (all equal)
/100000-4       702k ± 0%      702k ± 0%     ~     (all equal)
/1000000-4     7.02M ± 0%     7.02M ± 0%     ~     (all equal)
```

## Other tests
### FIFO queue
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.1.txt testdata/BenchmarkRefillDequeQueuev1.0.2.txt
name       old time/op    new time/op    delta
/1-4         3.74µs ± 1%    3.97µs ± 6%  +6.10%  (p=0.000 n=10+9)
/10-4        35.9µs ± 2%    36.8µs ± 5%  +2.45%  (p=0.011 n=10+10)
/100-4        352µs ± 3%     347µs ± 5%    ~     (p=0.165 n=10+10)
/1000-4      3.42ms ± 1%    3.67ms ±15%  +7.04%  (p=0.000 n=9+10)
/10000-4     37.9ms ± 2%    39.7ms ± 5%  +4.55%  (p=0.001 n=9+10)
/100000-4     400ms ± 1%     413ms ± 6%    ~     (p=0.095 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%  -0.00%  (p=0.000 n=10+10)
/10000-4     30.5MB ± 0%    30.1MB ± 0%  -1.27%  (p=0.000 n=8+10)
/100000-4     320MB ± 0%     320MB ± 0%  -0.13%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%  -0.02%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.1M ± 0%  -0.00%  (p=0.000 n=10+9)
```
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.1.txt testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt
name       old time/op    new time/op    delta
/1-4         3.90µs ± 4%    3.70µs ± 8%  -5.21%  (p=0.035 n=10+10)
/10-4        39.4µs ± 8%    37.4µs ± 7%  -5.26%  (p=0.007 n=10+10)
/100-4        386µs ± 9%     360µs ± 4%  -6.83%  (p=0.000 n=10+10)
/1000-4      3.81ms ± 9%    3.50ms ± 3%  -8.36%  (p=0.000 n=10+9)
/10000-4     41.7ms ±10%    40.0ms ± 5%    ~     (p=0.079 n=10+9)
/100000-4     443ms ±16%     426ms ±11%    ~     (p=0.243 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/10000-4     30.5MB ± 0%    30.1MB ± 0%  -1.36%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     320MB ± 0%  -0.13%  (p=0.000 n=10+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%  -0.02%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.1M ± 0%  -0.00%  (p=0.001 n=8+9)
```
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.1.txt testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt
name        old time/op    new time/op    delta
/1-4           254ns ± 4%     248ns ± 5%    ~     (p=0.195 n=9+9)
/10-4         1.95µs ± 3%    2.09µs ±13%  +7.15%  (p=0.006 n=9+9)
/100-4        8.08µs ± 2%    8.06µs ± 5%    ~     (p=0.256 n=10+10)
/1000-4       73.1µs ± 4%    73.5µs ± 5%    ~     (p=0.931 n=9+9)
/10000-4       768µs ± 4%     703µs ± 1%  -8.48%  (p=0.000 n=9+8)
/100000-4     8.89ms ± 4%    8.45ms ± 8%  -4.95%  (p=0.017 n=9+10)
/1000000-4    91.1ms ± 2%    90.6ms ± 6%    ~     (p=0.905 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      208B ± 0%  -7.14%  (p=0.000 n=10+10)
/10-4         4.91kB ± 0%    4.90kB ± 0%  -0.33%  (p=0.000 n=10+10)
/100-4        7.79kB ± 0%    7.78kB ± 0%  -0.21%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    53.2kB ± 0%  -1.83%  (p=0.000 n=10+10)
/10000-4       491kB ± 0%     486kB ± 0%  -1.05%  (p=0.000 n=10+10)
/100000-4     4.83MB ± 0%    4.82MB ± 0%  -0.02%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.2MB ± 0%  -0.01%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%    ~     (all equal)
/10-4           27.0 ± 0%      27.0 ± 0%    ~     (all equal)
/100-4           207 ± 0%       207 ± 0%    ~     (all equal)
/1000-4        2.02k ± 0%     2.02k ± 0%    ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%  -0.01%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      201k ± 0%    ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%  -0.00%  (p=0.000 n=10+10)
```
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.1.txt testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt
name        old time/op    new time/op    delta
/1-4          35.4ns ± 1%    35.3ns ± 3%    ~     (p=0.956 n=10+10)
/10-4          364ns ± 3%     353ns ± 2%  -3.06%  (p=0.001 n=10+9)
/100-4        3.50µs ± 1%    3.61µs ±13%    ~     (p=0.720 n=10+9)
/1000-4       34.9µs ± 1%    35.2µs ± 2%    ~     (p=0.573 n=10+8)
/10000-4       353µs ± 2%     356µs ± 7%    ~     (p=0.905 n=10+9)
/100000-4     3.56ms ± 3%    3.51ms ± 3%  -1.39%  (p=0.035 n=10+10)
/1000000-4    37.5ms ±16%    34.3ms ± 2%  -8.38%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.294 n=8+10)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.1.txt testdata/BenchmarkStableDequeQueuev1.0.2.txt
name        old time/op    new time/op    delta
/1-4          36.7ns ± 1%    35.1ns ± 2%  -4.47%  (p=0.000 n=10+9)
/10-4          370ns ± 1%     358ns ± 4%  -3.28%  (p=0.002 n=10+9)
/100-4        3.62µs ± 1%    3.44µs ± 2%  -4.99%  (p=0.000 n=10+9)
/1000-4       36.7µs ± 3%    34.2µs ± 2%  -6.80%  (p=0.000 n=10+9)
/10000-4       345µs ± 2%     357µs ± 6%  +3.74%  (p=0.028 n=9+10)
/100000-4     3.42ms ± 2%    3.49ms ± 4%  +1.98%  (p=0.015 n=10+10)
/1000000-4    34.2ms ± 1%    35.6ms ± 3%  +4.21%  (p=0.000 n=10+9)

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

### LIFO stack
```
benchstat testdata/BenchmarkRefillDequeStackv1.0.1.txt testdata/BenchmarkRefillDequeStackv1.0.2.txt
name       old time/op    new time/op    delta
/1-4         3.63µs ± 3%    4.10µs ±15%  +13.16%  (p=0.003 n=9+10)
/10-4        36.3µs ± 6%    35.7µs ± 2%     ~     (p=0.218 n=10+10)
/100-4        344µs ± 2%     353µs ±10%     ~     (p=0.481 n=10+10)
/1000-4      3.42ms ± 3%    3.43ms ± 5%     ~     (p=0.604 n=10+9)
/10000-4     37.4ms ± 8%    38.2ms ±11%     ~     (p=0.156 n=10+9)
/100000-4     393ms ± 0%     402ms ± 5%     ~     (p=0.089 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   -0.00%  (p=0.000 n=10+9)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=10+9)
/10000-4     30.5MB ± 0%    30.1MB ± 0%   -1.36%  (p=0.000 n=8+8)
/100000-4     320MB ± 0%     320MB ± 0%   -0.00%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%   -0.02%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.1M ± 0%     ~     (all equal)
```
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.1.txt testdata/BenchmarkRefillFullDequeStackv1.0.2.txt
name       old time/op    new time/op    delta
/1-4         4.09µs ±12%    4.33µs ± 5%  +5.99%  (p=0.016 n=10+9)
/10-4        35.8µs ± 3%    35.8µs ±11%    ~     (p=0.661 n=9+10)
/100-4        342µs ± 0%     364µs ±17%    ~     (p=0.146 n=8+10)
/1000-4      3.43ms ± 1%    3.40ms ± 3%    ~     (p=0.143 n=10+10)
/10000-4     39.1ms ± 1%    36.7ms ± 0%  -5.99%  (p=0.000 n=8+9)
/100000-4     394ms ± 1%     407ms ±14%    ~     (p=0.481 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/10000-4     30.9MB ± 0%    30.5MB ± 0%  -1.34%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     320MB ± 0%  -0.13%  (p=0.000 n=10+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%  -0.02%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.1M ± 0%  -0.00%  (p=0.000 n=10+10)
```
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.1.txt testdata/BenchmarkSlowIncreaseDequeStackv1.0.2.txt
name        old time/op    new time/op    delta
/1-4           258ns ± 1%     252ns ± 5%   -2.38%  (p=0.049 n=9+10)
/10-4         1.04µs ± 2%    1.02µs ± 5%     ~     (p=0.305 n=10+8)
/100-4        9.66µs ± 3%    7.95µs ± 5%  -17.72%  (p=0.000 n=9+10)
/1000-4       75.2µs ± 2%    68.8µs ± 4%   -8.60%  (p=0.000 n=9+10)
/10000-4       739µs ± 1%     723µs ± 9%     ~     (p=0.218 n=10+10)
/100000-4     8.78ms ± 7%    8.33ms ± 4%   -5.08%  (p=0.000 n=10+9)
/1000000-4    92.8ms ± 1%    89.7ms ± 6%   -3.27%  (p=0.027 n=8+9)

name        old alloc/op   new alloc/op   delta
/1-4            224B ± 0%      208B ± 0%   -7.14%  (p=0.000 n=10+10)
/10-4           768B ± 0%      752B ± 0%   -2.08%  (p=0.000 n=10+10)
/100-4        12.9kB ± 0%     7.8kB ± 0%  -39.78%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    49.0kB ± 0%   -1.98%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     486kB ± 0%   -0.20%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.82MB ± 0%   -0.02%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.2MB ± 0%   -0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      6.00 ± 0%     ~     (all equal)
/10-4           25.0 ± 0%      25.0 ± 0%     ~     (all equal)
/100-4           209 ± 0%       207 ± 0%   -0.96%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%     ~     (all equal)
/100000-4       201k ± 0%      201k ± 0%     ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%     ~     (all equal)
```
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.1.txt testdata/BenchmarkSlowDecreaseDequeStackv1.0.2.txt
name        old time/op    new time/op    delta
/1-4          40.4ns ±35%    33.9ns ± 1%  -16.07%  (p=0.000 n=10+8)
/10-4          374ns ±10%     353ns ± 5%   -5.78%  (p=0.003 n=10+9)
/100-4        3.50µs ± 1%    3.45µs ± 5%   -1.52%  (p=0.011 n=8+9)
/1000-4       35.7µs ± 6%    33.9µs ± 2%   -4.96%  (p=0.000 n=9+9)
/10000-4       360µs ± 8%     342µs ± 3%   -5.11%  (p=0.006 n=10+8)
/100000-4     3.49ms ± 2%    3.56ms ±10%     ~     (p=0.720 n=9+10)
/1000000-4    35.1ms ± 1%    36.3ms ±13%     ~     (p=0.447 n=9+10)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.1.txt testdata/BenchmarkStableDequeStackv1.0.2.txt
name        old time/op    new time/op    delta
/1-4          37.1ns ± 1%    39.2ns ± 8%  +5.78%  (p=0.016 n=8+10)
/10-4          374ns ± 1%     368ns ± 2%  -1.74%  (p=0.004 n=10+9)
/100-4        3.68µs ± 3%    3.65µs ± 4%    ~     (p=0.838 n=10+10)
/1000-4       36.5µs ± 1%    36.6µs ± 3%    ~     (p=0.931 n=9+9)
/10000-4       374µs ± 7%     381µs ±12%    ~     (p=0.631 n=10+10)
/100000-4     3.66ms ± 1%    3.55ms ± 2%  -3.01%  (p=0.000 n=10+8)
/1000000-4    36.5ms ± 1%    37.3ms ± 7%    ~     (p=0.053 n=10+9)

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
