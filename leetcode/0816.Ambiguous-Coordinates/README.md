# [816. Ambiguous Coordinates](https://leetcode.com/problems/ambiguous-coordinates/)


## 题目

We had some 2-dimensional coordinates, like `"(1, 3)"` or `"(2, 0.5)"`.  Then, we removed all commas, decimal points, and spaces, and ended up with the string `s`.  Return a list of strings representing all possibilities for what our original coordinates could have been.

Our original representation never had extraneous zeroes, so we never started with numbers like "00", "0.0", "0.00", "1.0", "001", "00.01", or any other number that can be represented with less digits.  Also, a decimal point within a number never occurs without at least one digit occuring before it, so we never started with numbers like ".1".

The final answer list can be returned in any order.  Also note that all coordinates in the final answer have exactly one space between them (occurring after the comma.)

```
Example 1:Input: s = "(123)"
Output: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]

```

```
Example 2:Input: s = "(00011)"
Output:  ["(0.001, 1)", "(0, 0.011)"]
Explanation:
0.0, 00, 0001 or 00.01 are not allowed.

```

```
Example 3:Input: s = "(0123)"
Output: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"]

```

```
Example 4:Input: s = "(100)"
Output: [(10, 0)]
Explanation:
1.0 is not allowed.

```

**Note:**

- `4 <= s.length <= 12`.
- `s[0]` = "(", `s[s.length - 1]` = ")", and the other elements in `s` are digits.

## 题目大意

我们有一些二维坐标，如 "(1, 3)" 或 "(2, 0.5)"，然后我们移除所有逗号，小数点和空格，得到一个字符串S。返回所有可能的原始字符串到一个列表中。原始的坐标表示法不会存在多余的零，所以不会出现类似于"00", "0.0", "0.00", "1.0", "001", "00.01"或一些其他更小的数来表示坐标。此外，一个小数点前至少存在一个数，所以也不会出现“.1”形式的数字。

最后返回的列表可以是任意顺序的。而且注意返回的两个数字中间（逗号之后）都有一个空格。

## 解题思路

- 本题没有什么算法思想，纯暴力题。先将原始字符串一分为二，分为的两个子字符串再移动坐标点，最后将每种情况组合再一次，这算完成了一次切分。将原始字符串每一位都按此规律完成切分，此题便得解。
- 这道题有 2 处需要注意的。第一处是最终输出的字符串，请注意，**两个数字中间（逗号之后）都有一个空格**。不遵守输出格式的要求也会导致 `Wrong Answer`。另外一处是切分数字时，有 2 种违法情况，一种是带前导 0 的，另外一种是末尾带 0 的。带前导 0 的也分为 2 种情况，一种是只有一位，即只有一个 0，这种情况直接返回，因为这一个 0 怎么切分也只有一种切分方法。另外一种是长度大于 1，即 `0xxx` 这种情况。`0xxx` 这种情况只有一种切分方法，即 `0.xxx`。末尾带 0 的只有一种切分方法，即 `xxx0`，不可切分，因为 `xxx.0`，`xx.x0`，`x.xx0` 这些都是违法情况，所以末尾带 0 的也可以直接返回。具体的实现见代码和注释。

## 代码

```go
package leetcode

func ambiguousCoordinates(s string) []string {
	res := []string{}
	s = s[1 : len(s)-1]
	for i := range s[:len(s)-1] {
		a := build(s[:i+1])
		b := build(s[i+1:])
		for _, ta := range a {
			for _, tb := range b {
				res = append(res, "("+ta+", "+tb+")")
			}
		}
	}
	return res
}

func build(s string) []string {
	res := []string{}
	if len(s) == 1 || s[0] != '0' {
		res = append(res, s)
	}
	// 结尾带 0 的情况
	if s[len(s)-1] == '0' {
		return res
	}
	// 切分长度大于一位且带前导 0 的情况
	if s[0] == '0' {
		res = append(res, "0."+s[1:])
		return res
	}
	for i := range s[:len(s)-1] {
		res = append(res, s[:i+1]+"."+s[i+1:])
	}
	return res
}
```