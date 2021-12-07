# [1734. Decode XORed Permutation](https://leetcode.com/problems/decode-xored-permutation/)


## 题目

There is an integer array `perm` that is a permutation of the first `n` positive integers, where `n` is always **odd**.

It was encoded into another integer array `encoded` of length `n - 1`, such that `encoded[i] = perm[i] XOR perm[i + 1]`. For example, if `perm = [1,3,2]`, then `encoded = [2,1]`.

Given the `encoded` array, return *the original array* `perm`. It is guaranteed that the answer exists and is unique.

**Example 1:**

```
Input: encoded = [3,1]
Output: [1,2,3]
Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]

```

**Example 2:**

```
Input: encoded = [6,5,4,6]
Output: [2,4,1,5,3]

```

**Constraints:**

- `3 <= n < 10^5`
- `n` is odd.
- `encoded.length == n - 1`

## 题目大意

给你一个整数数组 perm ，它是前 n 个正整数的排列，且 n 是个奇数 。它被加密成另一个长度为 n - 1 的整数数组 encoded ，满足 encoded[i] = perm[i] XOR perm[i + 1] 。比方说，如果 perm = [1,3,2] ，那么 encoded = [2,1] 。给你 encoded 数组，请你返回原始数组 perm 。题目保证答案存在且唯一。

## 解题思路

- 这一题与第 136 题和第 137 题思路类似，借用 `x ^ x = 0` 这个性质解题。依题意，原数组 perm 是 n 个正整数，即取值在 `[1,n+1]` 区间内，但是排列顺序未知。可以考虑先将 `[1,n+1]` 区间内的所有数异或得到 total。再将 encoded 数组中奇数下标的元素异或得到 odd：

    $$\begin{aligned}odd &= encoded[1] + encoded[3] + ... + encoded[n-1]\\&= (perm[1] \,\, XOR \,\, perm[2]) + (perm[3] \,\,  XOR  \,\, perm[4]) + ... + (perm[n-1]  \,\, XOR \,\, perm[n])\end{aligned}$$

    total 是 n 个正整数异或全集，odd 是 `n-1` 个正整数异或集。两者异或 `total ^ odd` 得到的值必定是 perm[0]，因为 `x ^ x = 0`，那么重复出现的元素被异或以后消失了。算出 perm[0] 就好办了。

    $$\begin{aligned}encoded[0] &= perm[0] \,\, XOR \,\, perm[1]\\perm[0] \,\, XOR \,\, encoded[0] &= perm[0] \,\, XOR \,\, perm[0] \,\, XOR \,\, perm[1] = perm[1]\\perm[1] \,\, XOR \,\, encoded[1] &= perm[1] \,\, XOR \,\, perm[1] \,\, XOR \,\, perm[2] = perm[2]\\...\\perm[n-1] \,\, XOR \,\, encoded[n-1] &= perm[n-1] \,\, XOR \,\, perm[n-1] \,\, XOR \,\, perm[n] = perm[n]\\\end{aligned}$$

    依次类推，便可以推出原数组 perm 中的所有数。

## 代码

```go
package leetcode

func decode(encoded []int) []int {
	n, total, odd := len(encoded), 0, 0
	for i := 1; i <= n+1; i++ {
		total ^= i
	}
	for i := 1; i < n; i += 2 {
		odd ^= encoded[i]
	}
	perm := make([]int, n+1)
	perm[0] = total ^ odd
	for i, v := range encoded {
		perm[i+1] = perm[i] ^ v
	}
	return perm
}
```