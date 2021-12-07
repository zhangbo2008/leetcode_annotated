# [1423. Maximum Points You Can Obtain from Cards](https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/)


## 题目

There are several cards **arranged in a row**, and each card has an associated number of points The points are given in the integer array `cardPoints`.

In one step, you can take one card from the beginning or from the end of the row. You have to take exactly `k` cards.

Your score is the sum of the points of the cards you have taken.

Given the integer array `cardPoints` and the integer `k`, return the *maximum score* you can obtain.

**Example 1:**

```
Input: cardPoints = [1,2,3,4,5,6,1], k = 3
Output: 12
Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.
```

**Example 2:**

```
Input: cardPoints = [2,2,2], k = 2
Output: 4
Explanation: Regardless of which two cards you take, your score will always be 4.
```

**Example 3:**

```
Input: cardPoints = [9,7,7,9,7,7,9], k = 7
Output: 55
Explanation: You have to take all the cards. Your score is the sum of points of all cards.
```

**Example 4:**

```
Input: cardPoints = [1,1000,1], k = 1
Output: 1
Explanation: You cannot take the card in the middle. Your best score is 1. 
```

**Example 5:**

```
Input: cardPoints = [1,79,80,1,1,1,200,1], k = 3
Output: 202
```

**Constraints:**

- `1 <= cardPoints.length <= 10^5`
- `1 <= cardPoints[i] <= 10^4`
- `1 <= k <= cardPoints.length`

## 题目大意

几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。你的点数就是你拿到手中的所有卡牌的点数之和。给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。

## 解题思路

- 这一题是滑动窗口题的简化题。从卡牌两边取 K 张牌，可以转换成在中间连续取 n-K 张牌。从两边取牌的点数最大，意味着剩下来中间牌的点数最小。扫描一遍数组，在每一个窗口大小为 n-K 的窗口内计算累加和，记录下最小的累加和。题目最终求的最大点数等于牌的总和减去中间最小的累加和。

## 代码

```go
package leetcode

func maxScore(cardPoints []int, k int) int {
	windowSize, sum := len(cardPoints)-k, 0
	for _, val := range cardPoints[:windowSize] {
		sum += val
	}
	minSum := sum
	for i := windowSize; i < len(cardPoints); i++ {
		sum += cardPoints[i] - cardPoints[i-windowSize]
		if sum < minSum {
			minSum = sum
		}
	}
	total := 0
	for _, pt := range cardPoints {
		total += pt
	}
	return total - minSum
}
```