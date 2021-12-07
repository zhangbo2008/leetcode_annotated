# [423. Reconstruct Original Digits from English](https://leetcode.com/problems/reconstruct-original-digits-from-english/)


## 题目

Given a **non-empty** string containing an out-of-order English representation of digits `0-9`, output the digits in ascending order.

**Note:**

1. Input contains only lowercase English letters.
2. Input is guaranteed to be valid and can be transformed to its original digits. That means invalid inputs such as "abc" or "zerone" are not permitted.
3. Input length is less than 50,000.

**Example 1:**

```
Input: "owoztneoer"
Output: "012"
```

**Example 2:**

```
Input: "fviefuro"
Output: "45"
```

## 题目大意

给定一个非空字符串，其中包含字母顺序打乱的英文单词表示的数字0-9。按升序输出原始的数字。

注意:

- 输入只包含小写英文字母。
- 输入保证合法并可以转换为原始的数字，这意味着像 "abc" 或 "zerone" 的输入是不允许的。
- 输入字符串的长度小于 50,000。

## 解题思路

- 这道题是一道找规律的题目。首先观察 0-9 对应的英文单词，找到特殊规律：所有的偶数都包含一个独特的字母：

    `z` 只在 `zero` 中出现。

    `w` 只在 `two` 中出现。

    `u` 只在 `four` 中出现。

    `x` 只在 `six` 中出现。

    `g` 只在 `eight` 中出现。

- 所以先排除掉这些偶数。然后在看剩下来几个数字对应的英文字母，这也是计算 3，5 和 7 的关键，因为有些单词只在一个奇数和一个偶数中出现（而且偶数已经被计算过了）：

    `h` 只在 `three` 和 `eight` 中出现。

    `f` 只在 `five` 和 `four` 中出现。

    `s` 只在 `seven` 和 `six` 中出现。

- 接下来只需要处理 9 和 0，思路依然相同。

    `i` 在 `nine`，`five`，`six` 和 `eight` 中出现。

    `n` 在 `one`，`seven` 和 `nine` 中出现。

- 最后按照上述的优先级，依次消耗对应的英文字母，生成最终的原始数字。注意按照优先级换算数字的时候，注意有多个重复数字的情况，比如多个 `1`，多个 `5` 等等。

## 代码

```go
package leetcode

import (
	"strings"
)

func originalDigits(s string) string {
	digits := make([]int, 26)
	for i := 0; i < len(s); i++ {
		digits[int(s[i]-'a')]++
	}
	res := make([]string, 10)
	res[0] = convert('z', digits, "zero", "0")
	res[6] = convert('x', digits, "six", "6")
	res[2] = convert('w', digits, "two", "2")
	res[4] = convert('u', digits, "four", "4")
	res[5] = convert('f', digits, "five", "5")
	res[1] = convert('o', digits, "one", "1")
	res[7] = convert('s', digits, "seven", "7")
	res[3] = convert('r', digits, "three", "3")
	res[8] = convert('t', digits, "eight", "8")
	res[9] = convert('i', digits, "nine", "9")
	return strings.Join(res, "")
}

func convert(b byte, digits []int, s string, num string) string {
	v := digits[int(b-'a')]
	for i := 0; i < len(s); i++ {
		digits[int(s[i]-'a')] -= v
	}
	return strings.Repeat(num, v)
}
```