
# v1.0.3 vs v2.0.0
## Fill tests
### FIFO queue
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkFillDequeQueuev1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkFillDequeQueue.txt
name        old time/op    new time/op    delta
/0-4          41.1ns ± 1%    40.5ns ± 0%   -1.42%  (p=0.000 n=8+8)
/1-4           168ns ± 2%     160ns ± 1%   -4.62%  (p=0.000 n=10+10)
/10-4          579ns ± 1%     621ns ±24%     ~     (p=0.730 n=9+9)
/100-4        4.65µs ± 1%    4.15µs ± 6%  -10.67%  (p=0.000 n=9+9)
/1000-4       36.6µs ± 0%    34.3µs ± 2%   -6.39%  (p=0.000 n=9+10)
/10000-4       368µs ± 1%     346µs ± 1%   -6.02%  (p=0.000 n=10+9)
/100000-4     3.78ms ± 0%    3.84ms ± 0%   +1.54%  (p=0.000 n=8+9)
/1000000-4    44.3ms ± 1%    42.1ms ± 2%   -4.92%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%     ~     (all equal)
/1-4            192B ± 0%      160B ± 0%  -16.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%      432B ± 0%  -27.03%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    4.48kB ± 0%  -37.78%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    25.2kB ± 0%  -26.05%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     243kB ± 0%  -24.93%  (p=0.002 n=8+10)
/100000-4     3.22MB ± 0%    2.42MB ± 0%  -24.88%  (p=0.000 n=9+10)
/1000000-4    32.2MB ± 0%    24.2MB ± 0%  -24.85%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           14.0 ± 0%      14.0 ± 0%     ~     (all equal)
/100-4           107 ± 0%       107 ± 0%     ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%     ~     (all equal)
/10000-4       10.1k ± 0%     10.1k ± 0%     ~     (all equal)
/100000-4       101k ± 0%      101k ± 0%     ~     (all equal)
/1000000-4     1.01M ± 0%     1.01M ± 0%     ~     (all equal)
```
### LIFO stack
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkFillDequeStackv1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkFillDequeStack.txt
name        old time/op    new time/op    delta
/0-4          40.4ns ± 0%    40.0ns ± 1%   -1.09%  (p=0.000 n=10+9)
/1-4           163ns ± 0%     158ns ± 2%   -3.27%  (p=0.000 n=9+9)
/10-4          578ns ± 0%     555ns ± 2%   -4.04%  (p=0.000 n=10+9)
/100-4        4.59µs ± 0%    4.13µs ± 2%  -10.10%  (p=0.000 n=8+9)
/1000-4       36.1µs ± 0%    35.0µs ± 7%   -3.25%  (p=0.003 n=8+10)
/10000-4       365µs ± 1%     349µs ± 2%   -4.38%  (p=0.000 n=10+10)
/100000-4     3.77ms ± 0%    3.86ms ± 1%   +2.44%  (p=0.000 n=10+9)
/1000000-4    43.9ms ± 1%    42.3ms ± 3%   -3.70%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%     ~     (all equal)
/1-4            192B ± 0%      160B ± 0%  -16.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%      432B ± 0%  -27.03%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    4.48kB ± 0%  -37.78%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    25.2kB ± 0%  -26.05%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     243kB ± 0%  -24.93%  (p=0.002 n=8+10)
/100000-4     3.22MB ± 0%    2.42MB ± 0%  -24.88%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    24.2MB ± 0%  -24.85%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           14.0 ± 0%      14.0 ± 0%     ~     (all equal)
/100-4           107 ± 0%       107 ± 0%     ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%     ~     (all equal)
/10000-4       10.1k ± 0%     10.1k ± 0%     ~     (all equal)
/100000-4       101k ± 0%      101k ± 0%     ~     (all equal)
/1000000-4     1.01M ± 0%     1.01M ± 0%     ~     (all equal)
```

## Microservice tests
### FIFO queue
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkMicroserviceDequeQueue.txt
name        old time/op    new time/op    delta
/0-4          44.3ns ± 0%    45.9ns ± 2%   +3.52%  (p=0.000 n=9+9)
/1-4           456ns ± 0%     429ns ± 1%   -5.78%  (p=0.000 n=9+9)
/10-4         2.80µs ± 0%    2.70µs ± 1%   -3.61%  (p=0.000 n=9+9)
/100-4        24.2µs ± 0%    23.8µs ± 1%   -1.82%  (p=0.000 n=9+10)
/1000-4        226µs ± 1%     229µs ± 5%     ~     (p=0.247 n=10+10)
/10000-4      2.37ms ± 4%    2.32ms ± 3%     ~     (p=0.052 n=10+10)
/100000-4     26.2ms ± 1%    25.7ms ± 1%   -1.91%  (p=0.000 n=10+10)
/1000000-4     266ms ± 1%     274ms ± 1%   +2.94%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%     ~     (all equal)
/1-4            544B ± 0%      384B ± 0%  -29.41%  (p=0.000 n=10+10)
/10-4         2.58kB ± 0%    1.90kB ± 0%  -26.09%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    16.2kB ± 0%  -22.77%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     123kB ± 0%   -8.13%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.28MB ± 0%  -10.77%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    12.8MB ± 0%  -11.05%  (p=0.000 n=9+10)
/1000000-4     144MB ± 0%     128MB ± 0%  -11.08%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            11.0 ± 0%      11.0 ± 0%     ~     (all equal)
/10-4           75.0 ± 0%      75.0 ± 0%     ~     (all equal)
/100-4           709 ± 0%       709 ± 0%     ~     (all equal)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.2k ± 0%     70.2k ± 0%     ~     (all equal)
/100000-4       702k ± 0%      702k ± 0%     ~     (all equal)
/1000000-4     7.02M ± 0%     7.02M ± 0%     ~     (p=0.332 n=10+10)
```
### LIFO stack
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkMicroserviceDequeStackv1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkMicroserviceDequeStack.txt
name        old time/op    new time/op    delta
/0-4          45.3ns ± 6%    45.3ns ± 1%     ~     (p=0.156 n=10+9)
/1-4           384ns ±12%     356ns ± 0%   -7.37%  (p=0.000 n=10+9)
/10-4         2.60µs ± 2%    2.50µs ± 1%   -3.89%  (p=0.000 n=9+10)
/100-4        29.3µs ±80%    22.8µs ± 1%  -21.95%  (p=0.000 n=9+10)
/1000-4        228µs ± 8%     221µs ± 1%   -3.19%  (p=0.000 n=8+8)
/10000-4      2.70ms ±14%    2.28ms ± 2%  -15.66%  (p=0.000 n=10+9)
/100000-4     27.2ms ±17%    24.8ms ± 0%   -8.94%  (p=0.000 n=9+8)
/1000000-4     265ms ± 1%     267ms ± 1%     ~     (p=0.075 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%     ~     (all equal)
/1-4            288B ± 0%      256B ± 0%  -11.11%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    1.39kB ± 0%  -10.31%  (p=0.000 n=10+10)
/100-4        16.8kB ± 0%    14.1kB ± 0%  -16.19%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     121kB ± 0%   -6.82%  (p=0.000 n=10+10)
/10000-4      1.42MB ± 0%    1.27MB ± 0%  -10.55%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    12.8MB ± 0%  -11.04%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     128MB ± 0%  -11.08%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/10-4           74.0 ± 0%      74.0 ± 0%     ~     (all equal)
/100-4           707 ± 0%       707 ± 0%     ~     (all equal)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.2k ± 0%     70.2k ± 0%     ~     (all equal)
/100000-4       702k ± 0%      702k ± 0%     ~     (all equal)
/1000000-4     7.02M ± 0%     7.02M ± 0%     ~     (p=0.871 n=10+10)
```

## Other tests
### FIFO queue
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkRefillDequeQueuev1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkRefillDequeQueue.txt
name       old time/op    new time/op    delta
/1-4         3.61µs ± 0%    3.86µs ± 7%   +7.04%  (p=0.000 n=10+9)
/10-4        34.9µs ± 1%    34.7µs ± 3%     ~     (p=0.258 n=9+9)
/100-4        346µs ± 3%     334µs ± 2%   -3.37%  (p=0.002 n=10+10)
/1000-4      3.38ms ± 2%    3.28ms ± 2%   -2.84%  (p=0.001 n=10+10)
/10000-4     37.6ms ± 7%    35.5ms ± 5%   -5.60%  (p=0.000 n=10+10)
/100000-4     384ms ± 2%     417ms ±10%   +8.61%  (p=0.000 n=8+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   -0.00%  (p=0.000 n=10+8)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=7+10)
/10000-4     30.1MB ± 0%    23.1MB ± 0%  -23.17%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     241MB ± 0%  -24.70%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%   -0.00%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.1M ± 0%   -0.00%  (p=0.028 n=10+10)
```
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkRefillFullDequeQueue.txt
name       old time/op    new time/op    delta
/1-4         3.41µs ± 0%    3.53µs ± 1%   +3.53%  (p=0.000 n=9+10)
/10-4        34.2µs ± 1%    34.1µs ± 0%     ~     (p=0.887 n=10+7)
/100-4        334µs ± 1%     333µs ± 1%     ~     (p=0.095 n=10+9)
/1000-4      3.33ms ± 1%    3.32ms ± 1%   -0.44%  (p=0.040 n=9+9)
/10000-4     36.9ms ± 0%    35.3ms ± 1%   -4.38%  (p=0.000 n=9+8)
/100000-4     376ms ± 0%     387ms ± 1%   +3.06%  (p=0.000 n=8+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.016 n=10+10)
/10000-4     30.1MB ± 0%    23.1MB ± 0%  -23.16%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     241MB ± 0%  -24.70%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%   -0.00%  (p=0.001 n=9+10)
/100000-4     10.1M ± 0%     10.1M ± 0%   -0.00%  (p=0.000 n=8+9)
```
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkSlowIncreaseDequeQueue.txt
name        old time/op    new time/op    delta
/1-4           234ns ±21%     192ns ± 0%  -18.01%  (p=0.000 n=10+10)
/10-4         1.14µs ± 4%    1.03µs ± 0%  -10.13%  (p=0.000 n=9+9)
/100-4        7.83µs ± 1%    7.28µs ± 0%   -7.06%  (p=0.000 n=10+9)
/1000-4       69.3µs ± 1%    67.4µs ± 1%   -2.78%  (p=0.000 n=10+10)
/10000-4       688µs ± 0%     688µs ± 6%     ~     (p=0.190 n=9+9)
/100000-4     8.11ms ± 1%    7.55ms ± 2%   -7.00%  (p=0.000 n=9+10)
/1000000-4    82.1ms ± 1%    82.6ms ± 1%   +0.58%  (p=0.004 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      176B ± 0%  -15.38%  (p=0.000 n=10+10)
/10-4         1.78kB ± 0%    1.10kB ± 0%  -37.84%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    6.08kB ± 0%  -30.91%  (p=0.000 n=10+10)
/1000-4       54.2kB ± 0%    43.3kB ± 0%  -20.14%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     405kB ± 0%  -16.95%  (p=0.000 n=10+8)
/100000-4     4.82MB ± 0%    4.02MB ± 0%  -16.62%  (p=0.000 n=10+8)
/1000000-4    48.2MB ± 0%    40.2MB ± 0%  -16.60%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           25.0 ± 0%      25.0 ± 0%     ~     (all equal)
/100-4           207 ± 0%       207 ± 0%     ~     (all equal)
/1000-4        2.02k ± 0%     2.02k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%     ~     (all equal)
/100000-4       201k ± 0%      201k ± 0%     ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%     ~     (all equal)
```
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkSlowDecreaseDequeQueue.txt
name        old time/op    new time/op    delta
/1-4          35.8ns ± 4%    34.9ns ± 2%  -2.66%  (p=0.000 n=9+9)
/10-4          375ns ± 6%     352ns ± 1%  -6.22%  (p=0.000 n=10+9)
/100-4        3.60µs ± 2%    3.43µs ± 1%  -4.58%  (p=0.000 n=9+10)
/1000-4       36.2µs ± 5%    34.2µs ± 1%  -5.40%  (p=0.000 n=10+9)
/10000-4       364µs ± 8%     342µs ± 1%  -6.00%  (p=0.000 n=9+10)
/100000-4     3.51ms ± 1%    3.42ms ± 0%  -2.43%  (p=0.000 n=10+9)
/1000000-4    35.2ms ± 1%    34.2ms ± 1%  -2.65%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%  -0.00%  (p=0.011 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  -0.00%  (p=0.011 n=10+10)

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
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkStableDequeQueuev1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkStableDequeQueue.txt
name        old time/op    new time/op    delta
/1-4          33.4ns ± 2%    34.2ns ± 1%  +2.18%  (p=0.002 n=9+9)
/10-4          339ns ± 1%     345ns ± 2%  +1.77%  (p=0.000 n=9+10)
/100-4        3.45µs ± 6%    3.34µs ± 1%    ~     (p=0.122 n=10+8)
/1000-4       33.1µs ± 1%    33.5µs ± 1%  +1.12%  (p=0.001 n=9+10)
/10000-4       332µs ± 1%     335µs ± 0%  +0.72%  (p=0.004 n=10+8)
/100000-4     3.30ms ± 1%    3.35ms ± 1%  +1.34%  (p=0.000 n=10+10)
/1000000-4    33.1ms ± 1%    33.5ms ± 1%  +1.06%  (p=0.006 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (p=0.982 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.065 n=10+10)

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
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkRefillDequeStackv1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkRefillDequeStack.txt
name       old time/op    new time/op    delta
/1-4         4.08µs ±15%    3.98µs ±14%     ~     (p=0.367 n=10+9)
/10-4        36.3µs ± 4%    37.5µs ±22%     ~     (p=0.968 n=9+10)
/100-4        350µs ± 3%     389µs ±21%  +11.04%  (p=0.003 n=8+10)
/1000-4      3.49ms ± 3%    3.28ms ± 1%   -5.99%  (p=0.000 n=10+9)
/10000-4     38.0ms ± 5%    34.6ms ± 1%   -9.09%  (p=0.000 n=10+9)
/100000-4     397ms ± 2%     387ms ± 1%   -2.55%  (p=0.000 n=8+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   -0.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=9+10)
/10000-4     30.1MB ± 0%    23.1MB ± 0%  -23.14%  (p=0.000 n=8+10)
/100000-4     320MB ± 0%     241MB ± 0%  -24.71%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%   -0.00%  (p=0.001 n=9+10)
/100000-4     10.1M ± 0%     10.1M ± 0%   -0.00%  (p=0.001 n=10+10)
```
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkRefillFullDequeStackv1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkRefillFullDequeStack.txt
name       old time/op    new time/op    delta
/1-4         3.40µs ± 1%    3.48µs ± 1%   +2.40%  (p=0.000 n=10+10)
/10-4        33.5µs ± 1%    33.6µs ± 1%     ~     (p=0.720 n=10+9)
/100-4        332µs ± 1%     326µs ± 0%   -1.74%  (p=0.000 n=8+9)
/1000-4      3.31ms ± 1%    3.25ms ± 1%   -1.77%  (p=0.000 n=7+9)
/10000-4     36.6ms ± 1%    34.8ms ± 1%   -4.87%  (p=0.000 n=9+10)
/100000-4     371ms ± 0%     384ms ± 1%   +3.46%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.003 n=10+9)
/10000-4     30.1MB ± 0%    23.1MB ± 0%  -23.14%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     241MB ± 0%  -24.71%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%   -0.00%  (p=0.033 n=10+10)
/100000-4     10.1M ± 0%     10.1M ± 0%   -0.00%  (p=0.000 n=9+10)
```
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkSlowIncreaseDequeStackv1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkSlowIncreaseDequeStack.txt
name        old time/op    new time/op    delta
/1-4           198ns ± 0%     198ns ± 2%     ~     (p=0.493 n=10+10)
/10-4          925ns ±10%     908ns ± 1%     ~     (p=0.481 n=9+8)
/100-4        8.35µs ± 5%    7.60µs ± 2%   -9.06%  (p=0.000 n=10+10)
/1000-4       70.1µs ± 3%    68.7µs ± 1%   -1.92%  (p=0.003 n=10+9)
/10000-4       722µs ± 3%     715µs ± 1%     ~     (p=0.083 n=10+8)
/100000-4     8.08ms ± 3%    7.98ms ±11%     ~     (p=0.143 n=10+10)
/1000000-4    81.0ms ± 1%    83.7ms ± 2%   +3.35%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      176B ± 0%  -15.38%  (p=0.000 n=10+10)
/10-4           752B ± 0%      592B ± 0%  -21.28%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    6.08kB ± 0%  -30.91%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    41.2kB ± 0%  -17.72%  (p=0.000 n=10+10)
/10000-4       483kB ± 0%     403kB ± 0%  -16.67%  (p=0.000 n=8+10)
/100000-4     4.82MB ± 0%    4.02MB ± 0%  -16.62%  (p=0.000 n=10+9)
/1000000-4    48.2MB ± 0%    40.2MB ± 0%  -16.60%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           24.0 ± 0%      24.0 ± 0%     ~     (all equal)
/100-4           207 ± 0%       207 ± 0%     ~     (all equal)
/1000-4        2.01k ± 0%     2.01k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%     ~     (all equal)
/100000-4       201k ± 0%      201k ± 0%     ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%     ~     (all equal)
```
```
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkSlowDecreaseDequeStackv1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkSlowDecreaseDequeStack.txt
name        old time/op    new time/op    delta
/1-4          35.2ns ± 1%    34.7ns ± 1%  -1.44%  (p=0.000 n=10+10)
/10-4          356ns ± 1%     353ns ± 1%  -0.63%  (p=0.019 n=8+10)
/100-4        3.48µs ± 0%    3.42µs ± 0%  -1.69%  (p=0.000 n=8+8)
/1000-4       34.8µs ± 1%    34.0µs ± 1%  -2.19%  (p=0.000 n=9+10)
/10000-4       348µs ± 1%     340µs ± 1%  -2.38%  (p=0.000 n=9+10)
/100000-4     3.55ms ± 5%    3.41ms ± 1%  -3.87%  (p=0.000 n=10+9)
/1000000-4    35.4ms ± 7%    34.2ms ± 1%  -3.56%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%  -0.00%  (p=0.023 n=8+9)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  -0.00%  (p=0.041 n=10+10)

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
benchstat ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkStableDequeStackv1.0.3.txt ./../../../ef-ds/deque-bench-tests/testdata/BenchmarkStableDequeStack.txt
name        old time/op    new time/op    delta
/1-4          32.8ns ± 1%    33.7ns ± 1%  +2.72%  (p=0.000 n=10+8)
/10-4          332ns ± 1%     339ns ± 1%  +2.19%  (p=0.000 n=9+10)
/100-4        3.24µs ± 1%    3.33µs ± 2%  +2.82%  (p=0.000 n=8+9)
/1000-4       33.2µs ± 7%    32.5µs ± 1%    ~     (p=0.360 n=10+8)
/10000-4       322µs ± 2%     325µs ± 1%  +0.69%  (p=0.034 n=10+8)
/100000-4     3.23ms ± 1%    3.26ms ± 1%  +1.05%  (p=0.001 n=10+10)
/1000000-4    32.2ms ± 1%    32.5ms ± 1%  +0.84%  (p=0.003 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (p=0.127 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  -0.00%  (p=0.018 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```
