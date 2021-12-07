# [1720. Decode XORed Array](https://leetcode.com/problems/decode-xored-array/)


## 题目

There is a **hidden** integer array `arr` that consists of `n` non-negative integers.

It was encoded into another integer array `encoded` of length `n - 1`, such that `encoded[i] = arr[i] XOR arr[i + 1]`. For example, if `arr = [1,0,2,1]`, then `encoded = [1,2,3]`.

You are given the `encoded` array. You are also given an integer `first`, that is the first element of `arr`, i.e. `arr[0]`.

Return *the original array* `arr`. It can be proved that the answer exists and is unique.

**Example 1:**

```
Input: encoded = [1,2,3], first = 1
Output: [1,0,2,1]
Explanation: If arr = [1,0,2,1], then first = 1 and encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]

```

**Example 2:**

```
Input: encoded = [6,2,7,3], first = 4
Output: [4,2,0,7,4]

```

**Constraints:**

- `2 <= n <= 104`
- `encoded.length == n - 1`
- `0 <= encoded[i] <= 105`
- `0 <= first <= 10^5`

## 题目大意

未知 整数数组 arr 由 n 个非负整数组成。经编码后变为长度为 n - 1 的另一个整数数组 encoded ，其中 encoded[i] = arr[i] XOR arr[i + 1] 。例如，arr = [1,0,2,1] 经编码后得到 encoded = [1,2,3] 。给你编码后的数组 encoded 和原数组 arr 的第一个元素 first（arr[0]）。请解码返回原数组 arr 。可以证明答案存在并且是唯一的。

## 解题思路

- 简单题。按照题意，求返回解码以后的原数组，即将编码后的数组前后两两元素依次做异或 `XOR` 运算。

## 代码

```go
package leetcode

func decode(encoded []int, first int) []int {
	arr := make([]int, len(encoded)+1)
	arr[0] = first
	for i, val := range encoded {
		arr[i+1] = arr[i] ^ val
	}
	return arr
}
```