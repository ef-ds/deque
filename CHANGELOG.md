# 1.0.0

* First stable release, production ready, certified to use as a FIFO queue or LIFO stack. Mixed Push/Pop/Front/Back is pending more testing and so is not recommended for use in a production setting.


# 1.0.1

* Fixed bug related to spare slices. The bug where the deque was not eliminating reused slices correctly caused it to cache more slices than maxSpareLinks (4), inflating memory unnecessarily.
- Commit 1: https://github.com/ef-ds/deque/commit/5cda9cbd756b5001cd8bb5e2b33675d65c61149d
- Commit 2: https://github.com/ef-ds/deque/commit/4de25e4de16dfe904669fe0a4ad3b0d189095fad
- Benchmark tests: [v1.0.0 vs v1.0.1](testdata/release_v1.0.1.md)

* Mixed Push/Pop/Front/Back is pending more testing and so is not recommended for use in production environment.


# 1.0.2

* Many improvements to make the code and the tests more readable and easier to maintain; the deque is also faster and uses less memory now. Amazing job, [Roger](https://github.com/rogpeppe)!
- Benchmark tests: [v1.0.1 vs v1.0.2](testdata/release_v1.0.2.md)

* Mixed Push/Pop/Front/Back is pending more testing and so is not recommended for use in production environment.


# 1.0.3
- Optimized deque: [here](https://github.com/ef-ds/deque/pull/13) and [here](https://github.com/ef-ds/deque/pull/14)
- Improved mixed tests: [here](https://github.com/ef-ds/deque/pull/15)
- Moved comparison benchmark tests to separate repo: [here](https://github.com/ef-ds/deque/pull/16)
- Benchmark tests: [v1.0.2 vs v1.0.3](testdata/release_v1.0.3.md); [comparison tests](https://github.com/ef-ds/deque-bench-tests/blob/master/PERFORMANCE.md)
- Mixed Push/Pop/Front/Back is now fully tested and should be fine to be used in production environments

# 1.0.4
- Fixed bug related to PushFront/PopBack: [here](https://github.com/ef-ds/deque/pull/21)
- The minor change had no significant performance impact

# 2.0.0
- Updated Deque to support generics!
- This is a breaking change and will require users to specify the desired data type used with Deque (such as Deque[int]). We recommend the users to update their existing code to use the specific data type the Deque needs to handle (such as int, string, etc), but the interface type can still be used. Just declare Deque as "var d deque.Deque[interface{}]" and use it normally. No other changes should be required to use Deque v2.0.0.
- Deque v2 is up to 10% faster and uses up to 35% less memory it used to use in most test cases! Benchmark tests: [v1.0.3 vs v2.0.0](testdata/release_v2.0.0.md); [comparison tests](https://github.com/ef-ds/deque-bench-tests/blob/master/PERFORMANCE.md)
