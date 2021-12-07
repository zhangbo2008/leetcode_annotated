# [299. Bulls and Cows](https://leetcode.com/problems/bulls-and-cows/)


## 题目

You are playing the Bulls and Cows game with your friend.

You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:

The number of "bulls", which are digits in the guess that are in the correct position.
The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.
Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.

The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows. Note that both secret and guess may contain duplicate digits.

**Example 1:**

```
Input: secret = "1807", guess = "7810"
Output: "1A3B"
Explanation: Bulls are connected with a '|' and cows are underlined:
"1807"
  |
"7810"
```

**Example 2:**

```
Input: secret = "1123", guess = "0111"
Output: "1A1B"
Explanation: Bulls are connected with a '|' and cows are underlined:
"1123"        "1123"
  |      or     |
"0111"        "0111"
Note that only one of the two unmatched 1s is counted as a cow since the non-bull digits can only be rearranged to allow one 1 to be a bull.
```

**Example 3:**

```
Input: secret = "1", guess = "0"
Output: "0A0B"
```

**Example 4:**

```
Input: secret = "1", guess = "1"
Output: "1A0B"
```

**Constraints:**

- 1 <= secret.length, guess.length <= 1000
- secret.length == guess.length
- secret and guess consist of digits only.

## 题目大意

你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：

写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：

猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls", 公牛），
有多少位属于数字猜对了但是位置不对（称为 "Cows", 奶牛）。也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
给你一个秘密数字secret 和朋友猜测的数字guess ，请你返回对朋友这次猜测的提示。

提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B表示奶牛。

请注意秘密数字和朋友猜测的数字都可能含有重复数字。

## 解题思路

- 计算下标一致并且对应下标的元素一致的个数，即 x
- secret 和 guess 分别去除 x 个公牛的元素,剩下 secret 和 guess 求共同的元素个数就是 y
- 把 x， y 转换成字符串，分别与 "A" 和 "B" 进行拼接返回结果

## 代码

```go

package leetcode

import "strconv"

func getHint(secret string, guess string) string {
	cntA, cntB := 0, 0
	mpS := make(map[byte]int)
	var strG []byte
	n := len(secret)
	var ans string
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			cntA++
		} else {
			mpS[secret[i]] += 1
			strG = append(strG, guess[i])
		}
	}
	for _, v := range strG {
		if _, ok := mpS[v]; ok {
			if mpS[v] > 1 {
				mpS[v] -= 1
			} else {
				delete(mpS, v)
			}
			cntB++
		}
	}
	ans += strconv.Itoa(cntA) + "A" + strconv.Itoa(cntB) + "B"
	return ans
}
```