# Roman to Integer

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Given a roman numeral, convert it to an integer.

`Time: O(n)`

`Space: O(1)`

```Bash
$ go test -v -cover ./...
=== RUN   TestCase1
Expected: 3, Output: 3
--- PASS: TestCase1 (0.00s)
=== RUN   TestCase2
Expected: 4, Output: 4
--- PASS: TestCase2 (0.00s)
=== RUN   TestCase3
Expected: 9, Output: 9
--- PASS: TestCase3 (0.00s)
=== RUN   TestCase4
Expected: 58, Output: 58
--- PASS: TestCase4 (0.00s)
PASS
coverage: 83.3% of statements
ok      github.com/mxssl/LeetCodeRomanToIntegerGolang   0.093s  coverage: 83.3% of statements
```
