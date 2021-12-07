# [1678. Goal Parser Interpretation](https://leetcode.com/problems/goal-parser-interpretation/)

## 题目

You own a **Goal Parser** that can interpret a string `command`. The `command` consists of an alphabet of `"G"`, `"()"` and/or `"(al)"` in some order. The Goal Parser will interpret `"G"` as the string `"G"`, `"()"` as the string `"o"`, and `"(al)"` as the string `"al"`. The interpreted strings are then concatenated in the original order.

Given the string `command`, return *the **Goal Parser**'s interpretation of* `command`.

**Example 1:**

```
Input: command = "G()(al)"
Output: "Goal"
Explanation: The Goal Parser interprets the command as follows:
G -> G
() -> o
(al) -> al
The final concatenated result is "Goal".
```

**Example 2:**

```
Input: command = "G()()()()(al)"
Output: "Gooooal"
```

**Example 3:**

```
Input: command = "(al)G(al)()()G"
Output: "alGalooG"
```

**Constraints:**

- `1 <= command.length <= 100`
- `command` consists of `"G"`, `"()"`, and/or `"(al)"` in some order.

## 题目大意

请你设计一个可以解释字符串 command 的 Goal 解析器 。command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成。Goal 解析器会将 "G" 解释为字符串 "G"、"()" 解释为字符串 "o" ，"(al)" 解释为字符串 "al" 。然后，按原顺序将经解释得到的字符串连接成一个字符串。给你字符串 command ，返回 Goal 解析器 对 command 的解释结果。

## 解题思路

- 简单题，按照题意修改字符串即可。由于是简单题，这一题也不用考虑嵌套的情况。

## 代码

```go
package leetcode

func interpret(command string) string {
	if command == "" {
		return ""
	}
	res := ""
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res += "G"
		} else {
			if command[i] == '(' && command[i+1] == 'a' {
				res += "al"
				i += 3
			} else {
				res += "o"
				i ++
			}
		}
	}
	return res
}
```