
# v1.0.2 vs v1.0.3
## Fill tests
### FIFO queue
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.2.txt testdata/BenchmarkFillDequeQueuev1.0.3.txt
name        old time/op    new time/op    delta
/0-4          36.1ns ± 4%    37.4ns ± 1%   +3.68%  (p=0.000 n=10+10)
/1-4           143ns ± 5%     171ns ± 1%  +19.96%  (p=0.000 n=10+10)
/10-4          664ns ± 9%     577ns ± 1%  -13.20%  (p=0.000 n=10+10)
/100-4        4.83µs ± 7%    4.74µs ± 2%     ~     (p=0.356 n=10+9)
/1000-4       42.4µs ±15%    37.1µs ± 3%  -12.56%  (p=0.000 n=10+10)
/10000-4       417µs ±11%     370µs ± 2%  -11.29%  (p=0.000 n=10+10)
/100000-4     4.45ms ±27%    3.87ms ± 0%  -13.12%  (p=0.000 n=10+8)
/1000000-4    51.9ms ±12%    44.1ms ± 1%  -14.92%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     64.0B ± 0%  +33.33%  (p=0.000 n=10+10)
/1-4            128B ± 0%      192B ± 0%  +50.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%      592B ± 0%     ~     (all equal)
/100-4        6.18kB ± 0%    7.20kB ± 0%  +16.58%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    34.0kB ± 0%   +3.10%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     323kB ± 0%   +0.32%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.22MB ± 0%   +0.03%  (p=0.000 n=9+10)
/1000000-4    32.2MB ± 0%    32.2MB ± 0%   -0.01%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           15.0 ± 0%      14.0 ± 0%   -6.67%  (p=0.000 n=10+10)
/100-4           107 ± 0%       107 ± 0%     ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%     ~     (all equal)
/10000-4       10.1k ± 0%     10.1k ± 0%     ~     (all equal)
/100000-4       101k ± 0%      101k ± 0%     ~     (all equal)
/1000000-4     1.01M ± 0%     1.01M ± 0%   -0.00%  (p=0.000 n=10+10)
```

### LIFO stack
```
benchstat testdata/BenchmarkFillDequeStackv1.0.2.txt testdata/BenchmarkFillDequeStackv1.0.3.txt
name        old time/op    new time/op    delta
/0-4          38.3ns ± 3%    37.5ns ± 1%     ~     (p=0.085 n=8+9)
/1-4           148ns ± 9%     171ns ± 0%  +14.96%  (p=0.000 n=10+10)
/10-4          662ns ±13%     579ns ± 1%  -12.51%  (p=0.000 n=10+10)
/100-4        4.99µs ± 6%    4.75µs ± 2%   -4.96%  (p=0.000 n=9+10)
/1000-4       39.5µs ± 8%    36.7µs ± 2%   -7.02%  (p=0.000 n=10+9)
/10000-4       387µs ± 9%     363µs ± 2%   -6.36%  (p=0.000 n=10+10)
/100000-4     4.13ms ± 9%    3.81ms ± 0%   -7.60%  (p=0.000 n=10+8)
/1000000-4    46.8ms ± 4%    43.6ms ± 2%   -6.78%  (p=0.000 n=8+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     64.0B ± 0%  +33.33%  (p=0.000 n=10+10)
/1-4            128B ± 0%      192B ± 0%  +50.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%      592B ± 0%     ~     (all equal)
/100-4        6.18kB ± 0%    7.20kB ± 0%  +16.58%  (p=0.000 n=10+10)
/1000-4       33.0kB ± 0%    34.0kB ± 0%   +3.10%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     323kB ± 0%   +0.32%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.22MB ± 0%   +0.03%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    32.2MB ± 0%   -0.01%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           15.0 ± 0%      14.0 ± 0%   -6.67%  (p=0.000 n=10+10)
/100-4           107 ± 0%       107 ± 0%     ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%     ~     (all equal)
/10000-4       10.1k ± 0%     10.1k ± 0%     ~     (all equal)
/100000-4       101k ± 0%      101k ± 0%     ~     (all equal)
/1000000-4     1.01M ± 0%     1.01M ± 0%   -0.00%  (p=0.000 n=10+10)
```

## Microservice tests
### FIFO queue
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.2.txt testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt
name        old time/op    new time/op    delta
/1-4           510ns ± 5%     464ns ± 0%   -9.01%  (p=0.000 n=10+10)
/10-4         3.60µs ± 5%    2.77µs ± 1%  -23.12%  (p=0.000 n=10+10)
/100-4        24.8µs ± 7%    24.2µs ± 0%   -2.52%  (p=0.034 n=10+8)
/1000-4        237µs ± 5%     224µs ± 2%   -5.60%  (p=0.000 n=9+10)
/10000-4      2.38ms ± 5%    2.28ms ± 2%   -4.29%  (p=0.000 n=10+10)
/100000-4     25.0ms ± 1%    24.8ms ± 2%     ~     (p=0.165 n=10+10)
/1000000-4     268ms ± 4%     259ms ± 3%   -3.41%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            544B ± 0%      544B ± 0%     ~     (all equal)
/10-4         5.70kB ± 0%    2.58kB ± 0%  -54.78%  (p=0.000 n=10+10)
/100-4        15.8kB ± 0%    20.9kB ± 0%  +32.76%  (p=0.000 n=10+10)
/1000-4        133kB ± 0%     134kB ± 0%   +0.77%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.43MB ± 0%   +0.07%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    14.4MB ± 0%   +0.01%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     144MB ± 0%   +0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            12.0 ± 0%      11.0 ± 0%   -8.33%  (p=0.000 n=10+10)
/10-4           77.0 ± 0%      75.0 ± 0%   -2.60%  (p=0.000 n=10+10)
/100-4           707 ± 0%       709 ± 0%   +0.28%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.2k ± 0%     70.2k ± 0%     ~     (all equal)
/100000-4       702k ± 0%      702k ± 0%     ~     (all equal)
/1000000-4     7.02M ± 0%     7.02M ± 0%     ~     (all equal)
```

### LIFO stack
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.2.txt testdata/BenchmarkMicroserviceDequeStackv1.0.3.txt
name        old time/op    new time/op    delta
/1-4           389ns ± 1%     356ns ± 0%  -8.44%  (p=0.000 n=10+8)
/10-4         2.51µs ± 1%    2.59µs ± 7%    ~     (p=0.239 n=10+10)
/100-4        22.8µs ± 1%    23.4µs ± 1%  +2.39%  (p=0.000 n=10+10)
/1000-4        220µs ± 2%     220µs ± 0%    ~     (p=0.114 n=9+8)
/10000-4      2.28ms ± 2%    2.27ms ± 1%    ~     (p=0.122 n=8+10)
/100000-4     24.6ms ± 1%    24.9ms ± 2%  +1.12%  (p=0.011 n=9+9)
/1000000-4     258ms ± 1%     259ms ± 2%    ~     (p=0.549 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            288B ± 0%      288B ± 0%    ~     (all equal)
/10-4         1.55kB ± 0%    1.55kB ± 0%    ~     (all equal)
/100-4        15.8kB ± 0%    16.8kB ± 0%  +6.49%  (p=0.000 n=10+10)
/1000-4        129kB ± 0%     130kB ± 0%  +0.79%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.42MB ± 0%  -0.51%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    14.4MB ± 0%  +0.01%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     144MB ± 0%  +0.00%  (p=0.000 n=9+8)

name        old allocs/op  new allocs/op  delta
/1-4            11.0 ± 0%      10.0 ± 0%  -9.09%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      74.0 ± 0%  -1.33%  (p=0.000 n=10+10)
/100-4           707 ± 0%       707 ± 0%    ~     (all equal)
/1000-4        7.01k ± 0%     7.01k ± 0%    ~     (all equal)
/10000-4       70.2k ± 0%     70.2k ± 0%  -0.01%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      702k ± 0%    ~     (all equal)
/1000000-4     7.02M ± 0%     7.02M ± 0%    ~     (all equal)
```

## Other tests
### FIFO queue
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.2.txt testdata/BenchmarkRefillDequeQueuev1.0.3.txt
name       old time/op    new time/op    delta
/1-4         3.97µs ± 6%    3.95µs ± 7%    ~     (p=1.000 n=9+10)
/10-4        36.8µs ± 5%    38.9µs ± 8%  +5.70%  (p=0.009 n=10+10)
/100-4        347µs ± 5%     370µs ± 8%  +6.65%  (p=0.029 n=10+10)
/1000-4      3.67ms ±15%    3.85ms ±10%    ~     (p=0.089 n=10+10)
/10000-4     39.7ms ± 5%    41.0ms ±12%    ~     (p=0.796 n=10+10)
/100000-4     413ms ± 6%     438ms ± 7%  +6.08%  (p=0.022 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%  +0.00%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    30.1MB ± 0%  -0.00%  (p=0.005 n=10+10)
/100000-4     320MB ± 0%     320MB ± 0%  -0.00%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%  -0.00%  (p=0.033 n=10+10)
/100000-4     10.1M ± 0%     10.1M ± 0%  -0.00%  (p=0.000 n=9+10)
```
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.2.txt testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt
name       old time/op    new time/op    delta
/1-4         3.70µs ± 8%    3.55µs ± 1%  -3.98%  (p=0.007 n=10+9)
/10-4        37.4µs ± 7%    35.2µs ± 1%  -5.86%  (p=0.000 n=10+10)
/100-4        360µs ± 4%     344µs ± 1%  -4.46%  (p=0.000 n=10+9)
/1000-4      3.50ms ± 3%    3.41ms ± 1%  -2.38%  (p=0.000 n=9+10)
/10000-4     40.0ms ± 5%    40.6ms ±24%    ~     (p=0.258 n=9+9)
/100000-4     426ms ±11%     451ms ±15%    ~     (p=0.113 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/10000-4     30.1MB ± 0%    30.1MB ± 0%    ~     (p=0.147 n=10+10)
/100000-4     320MB ± 0%     320MB ± 0%    ~     (all equal)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%    ~     (all equal)
/100000-4     10.1M ± 0%     10.1M ± 0%    ~     (all equal)
```
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt
name        old time/op    new time/op    delta
/1-4           248ns ± 5%     241ns ±14%     ~     (p=0.952 n=9+10)
/10-4         2.09µs ±13%    1.32µs ±10%  -36.69%  (p=0.000 n=9+10)
/100-4        8.06µs ± 5%    9.07µs ±11%  +12.61%  (p=0.000 n=10+10)
/1000-4       73.5µs ± 5%    78.2µs ±10%     ~     (p=0.053 n=9+10)
/10000-4       703µs ± 1%     805µs ± 7%  +14.57%  (p=0.000 n=8+9)
/100000-4     8.45ms ± 8%    8.74ms ±12%     ~     (p=0.529 n=10+10)
/1000000-4    90.6ms ± 6%    92.0ms ± 9%     ~     (p=0.549 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      208B ± 0%     ~     (all equal)
/10-4         4.90kB ± 0%    1.78kB ± 0%  -63.73%  (p=0.000 n=10+10)
/100-4        7.78kB ± 0%    8.80kB ± 0%  +13.17%  (p=0.000 n=10+10)
/1000-4       53.2kB ± 0%    54.2kB ± 0%   +1.93%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     487kB ± 0%   +0.21%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.82MB ± 0%   -0.06%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.2MB ± 0%   +0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%  -16.67%  (p=0.000 n=10+10)
/10-4           27.0 ± 0%      25.0 ± 0%   -7.41%  (p=0.000 n=10+10)
/100-4           207 ± 0%       207 ± 0%     ~     (all equal)
/1000-4        2.02k ± 0%     2.02k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%     ~     (all equal)
/100000-4       201k ± 0%      201k ± 0%   -0.00%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.01M ± 0%     ~     (all equal)
```
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.2.txt testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt
name        old time/op    new time/op    delta
/1-4          35.3ns ± 3%    39.2ns ± 7%  +11.02%  (p=0.000 n=10+9)
/10-4          353ns ± 2%     391ns ±11%  +10.86%  (p=0.001 n=9+10)
/100-4        3.61µs ±13%    3.71µs ± 8%     ~     (p=0.243 n=9+10)
/1000-4       35.2µs ± 2%    36.1µs ± 7%     ~     (p=0.274 n=8+10)
/10000-4       356µs ± 7%     375µs ±10%   +5.38%  (p=0.035 n=9+10)
/100000-4     3.51ms ± 3%    3.68ms ± 8%     ~     (p=0.123 n=10+10)
/1000000-4    34.3ms ± 2%    37.0ms ± 9%   +7.86%  (p=0.015 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (p=0.137 n=10+8)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.2.txt testdata/BenchmarkStableDequeQueuev1.0.3.txt
name        old time/op    new time/op    delta
/1-4          35.1ns ± 2%    37.2ns ± 9%    ~     (p=0.128 n=9+10)
/10-4          358ns ± 4%     369ns ± 8%    ~     (p=0.138 n=9+10)
/100-4        3.44µs ± 2%    3.52µs ± 8%    ~     (p=0.604 n=9+10)
/1000-4       34.2µs ± 2%    36.9µs ± 7%  +7.82%  (p=0.001 n=9+9)
/10000-4       357µs ± 6%     353µs ± 9%    ~     (p=0.393 n=10+10)
/100000-4     3.49ms ± 4%    3.60ms ± 8%    ~     (p=0.190 n=10+10)
/1000000-4    35.6ms ± 3%    35.5ms ± 7%    ~     (p=0.720 n=9+10)

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
benchstat testdata/BenchmarkRefillDequeStackv1.0.2.txt testdata/BenchmarkRefillDequeStackv1.0.3.txt
name       old time/op    new time/op    delta
/1-4         4.10µs ±15%    3.91µs ± 8%     ~     (p=0.481 n=10+10)
/10-4        35.7µs ± 2%    36.6µs ± 8%     ~     (p=0.631 n=10+10)
/100-4        353µs ±10%     372µs ± 9%     ~     (p=0.089 n=10+10)
/1000-4      3.43ms ± 5%    3.61ms ± 7%   +5.29%  (p=0.010 n=9+10)
/10000-4     38.2ms ±11%    42.0ms ± 7%  +10.10%  (p=0.000 n=9+10)
/100000-4     402ms ± 5%     444ms ±14%  +10.37%  (p=0.002 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   +0.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.00%  (p=0.000 n=9+10)
/10000-4     30.1MB ± 0%    30.1MB ± 0%   +0.00%  (p=0.000 n=8+8)
/100000-4     320MB ± 0%     320MB ± 0%   +0.00%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%     ~     (all equal)
/100000-4     10.1M ± 0%     10.1M ± 0%     ~     (all equal)
```
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.2.txt testdata/BenchmarkRefillFullDequeStackv1.0.3.txt
name       old time/op    new time/op    delta
/1-4         4.33µs ± 5%    3.85µs ± 9%  -11.15%  (p=0.000 n=9+10)
/10-4        35.8µs ±11%    38.2µs ± 9%   +6.82%  (p=0.009 n=10+10)
/100-4        364µs ±17%     378µs ± 8%     ~     (p=0.063 n=10+10)
/1000-4      3.40ms ± 3%    3.76ms ± 8%  +10.57%  (p=0.000 n=10+10)
/10000-4     36.7ms ± 0%    42.7ms ±10%  +16.29%  (p=0.000 n=9+10)
/100000-4     407ms ±14%     446ms ±11%   +9.76%  (p=0.009 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.5MB ± 0%    30.1MB ± 0%   -1.36%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     320MB ± 0%     ~     (all equal)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%   -0.02%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.1M ± 0%     ~     (all equal)
```
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowIncreaseDequeStackv1.0.3.txt
name        old time/op    new time/op    delta
/1-4           252ns ± 5%     226ns ±10%  -10.58%  (p=0.001 n=10+10)
/10-4         1.02µs ± 5%    1.06µs ± 7%     ~     (p=0.165 n=8+10)
/100-4        7.95µs ± 5%    8.71µs ±13%   +9.65%  (p=0.003 n=10+10)
/1000-4       68.8µs ± 4%    75.5µs ± 9%   +9.86%  (p=0.000 n=10+10)
/10000-4       723µs ± 9%     731µs ± 6%     ~     (p=0.739 n=10+10)
/100000-4     8.33ms ± 4%    8.81ms ±10%     ~     (p=0.182 n=9+10)
/1000000-4    89.7ms ± 6%    94.4ms ±11%     ~     (p=0.211 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      208B ± 0%     ~     (all equal)
/10-4           752B ± 0%      752B ± 0%     ~     (all equal)
/100-4        7.78kB ± 0%    8.80kB ± 0%  +13.17%  (p=0.000 n=10+10)
/1000-4       49.0kB ± 0%    50.0kB ± 0%   +2.09%  (p=0.000 n=10+10)
/10000-4       486kB ± 0%     483kB ± 0%   -0.64%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.82MB ± 0%   +0.02%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.2MB ± 0%   +0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            6.00 ± 0%      5.00 ± 0%  -16.67%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      24.0 ± 0%   -4.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       207 ± 0%     ~     (all equal)
/1000-4        2.01k ± 0%     2.01k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%   -0.01%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      201k ± 0%     ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%     ~     (all equal)
```
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.2.txt testdata/BenchmarkSlowDecreaseDequeStackv1.0.3.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    37.3ns ± 7%   +9.94%  (p=0.000 n=8+10)
/10-4          353ns ± 5%     386ns ± 9%   +9.28%  (p=0.001 n=9+10)
/100-4        3.45µs ± 5%    3.76µs ± 8%   +9.10%  (p=0.001 n=9+10)
/1000-4       33.9µs ± 2%    37.0µs ± 6%   +9.17%  (p=0.000 n=9+10)
/10000-4       342µs ± 3%     377µs ± 8%  +10.20%  (p=0.000 n=8+10)
/100000-4     3.56ms ±10%    3.72ms ± 9%   +4.49%  (p=0.035 n=10+10)
/1000000-4    36.3ms ±13%    38.1ms ± 9%     ~     (p=0.105 n=10+10)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.2.txt testdata/BenchmarkStableDequeStackv1.0.3.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 8%    35.6ns ± 7%  -9.26%  (p=0.001 n=10+10)
/10-4          368ns ± 2%     362ns ± 8%    ~     (p=0.985 n=9+10)
/100-4        3.65µs ± 4%    3.54µs ± 7%    ~     (p=0.123 n=10+10)
/1000-4       36.6µs ± 3%    35.9µs ±10%    ~     (p=0.780 n=9+10)
/10000-4       381µs ±12%     360µs ±10%    ~     (p=0.063 n=10+10)
/100000-4     3.55ms ± 2%    3.53ms ± 6%    ~     (p=0.633 n=8+10)
/1000000-4    37.3ms ± 7%    34.4ms ± 7%  -7.60%  (p=0.000 n=9+10)

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
