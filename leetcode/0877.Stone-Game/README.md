# [877. Stone Game](https://leetcode.com/problems/stone-game/)

## 题目

Alex and Lee play a game with piles of stones.  There are an even number of piles **arranged in a row**, and each pile has a positive integer number of stones `piles[i]`.

The objective of the game is to end with the most stones.  The total number of stones is odd, so there are no ties.

Alex and Lee take turns, with Alex starting first.  Each turn, a player takes the entire pile of stones from either the beginning or the end of the row.  This continues until there are no more piles left, at which point the person with the most stones wins.

Assuming Alex and Lee play optimally, return `True` if and only if Alex wins the game.

**Example 1:**

```
Input: piles = [5,3,4,5]
Output: true
Explanation:
Alex starts first, and can only take the first 5 or the last 5.
Say he takes the first 5, so that the row becomes [3, 4, 5].
If Lee takes 3, then the board is [4, 5], and Alex takes 5 to win with 10 points.
If Lee takes the last 5, then the board is [3, 4], and Alex takes 4 to win with 9 points.
This demonstrated that taking the first 5 was a winning move for Alex, so we return true.

```

**Constraints:**

- `2 <= piles.length <= 500`
- `piles.length` is even.
- `1 <= piles[i] <= 500`
- `sum(piles)` is odd.

## 题目大意

亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。

## 解题思路

- 一遇到石子问题，很容易让人联想到是否和奇偶数相关。此题指定了石子堆数一定是偶数。所以从这里为突破口试试。Alex 先拿，要么取行首下标为 0 的石子，要么取行尾下标为 n-1 的石子。假设取下标为 0 的石子，剩下的石子堆下标从 1 ~ n-1，即 Lee 只能从奇数下标的石子堆 1 或者 n-1。假设 Alex 第一次取下标为 n-1 的石子，剩下的石子堆下标从 0 ~ n-2，即 Lee 只能取偶数下标的石子堆。于是 Alex 的必胜策略是每轮取石子，取此轮奇数下标堆石子数总和，偶数下标堆石子数总和，两者大者。那么下一轮 Lee 只能取石子堆数相对少的那一堆，并且 Lee 取的石子堆下标奇偶性是完全受到上一轮 Alex 控制的。所以只要是 Alex 先手，那么每轮都可以压制 Lee，从而必胜。

## 代码

```go
package leetcode

func stoneGame(piles []int) bool {
	return true
}
```