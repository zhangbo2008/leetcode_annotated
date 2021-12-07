# [1310. XOR Queries of a Subarray](https://leetcode.com/problems/xor-queries-of-a-subarray/)


## 题目

Given the array `arr` of positive integers and the array `queries` where `queries[i] = [Li,Ri]`, for each query `i` compute the **XOR** of elements from `Li` to `Ri` (that is, `arr[Li]xor arr[Li+1]xor ...xor arr[Ri]`). Return an array containing the result for the given `queries`.

**Example 1:**

```
Input: arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
Output: [2,7,14,8]
Explanation:
The binary representation of the elements in the array are:
1 = 0001
3 = 0011
4 = 0100
8 = 1000
The XOR values for queries are:
[0,1] = 1 xor 3 = 2
[1,2] = 3 xor 4 = 7
[0,3] = 1 xor 3 xor 4 xor 8 = 14
[3,3] = 8

```

**Example 2:**

```
Input: arr = [4,8,2,10], queries = [[2,3],[1,3],[0,0],[0,3]]
Output: [8,0,4,4]

```

**Constraints:**

- `1 <= arr.length <= 3 * 10^4`
- `1 <= arr[i] <= 10^9`
- `1 <= queries.length <= 3 * 10^4`
- `queries[i].length == 2`
- `0 <= queries[i][0] <= queries[i][1] < arr.length`

## 题目大意

有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。并返回一个包含给定查询 queries 所有结果的数组。

## 解题思路

- 此题求区间异或，很容易让人联想到区间求和。区间求和利用前缀和，可以使得 query 从 O(n) 降为 O(1)。区间异或能否也用类似前缀和的思想呢？答案是肯定的。利用异或的两个性质，x ^ x = 0，x ^ 0 = x。那么有：（由于 LaTeX 中异或符号 ^ 是特殊字符，笔者用 $\oplus$ 代替异或）

    $$\begin{aligned}Query(left,right) &=arr[left] \oplus \cdots  \oplus arr[right]\\&=(arr[0] \oplus \cdots  \oplus arr[left-1]) \oplus (arr[0] \oplus \cdots  \oplus arr[left-1]) \oplus (arr[left] \oplus \cdots  \oplus arr[right])\\ &=(arr[0] \oplus \cdots  \oplus arr[left-1]) \oplus (arr[0] \oplus \cdots  \oplus arr[right])\\ &=xors[left] \oplus xors[right+1]\\ \end{aligned}$$

    按照这个思路解题，便可以将 query 从 O(n) 降为 O(1)，总的时间复杂度为 O(n)。

## 代码

```go
package leetcode

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr))
	xors[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		xors[i] = arr[i] ^ xors[i-1]
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = xors[q[1]]
		if q[0] > 0 {
			res[i] ^= xors[q[0]-1]
		}
	}
	return res
}
```