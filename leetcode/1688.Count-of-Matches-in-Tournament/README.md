# [1688. Count of Matches in Tournament](https://leetcode.com/problems/count-of-matches-in-tournament/)


## 题目

You are given an integer `n`, the number of teams in a tournament that has strange rules:

- If the current number of teams is **even**, each team gets paired with another team. A total of `n / 2` matches are played, and `n / 2` teams advance to the next round.
- If the current number of teams is **odd**, one team randomly advances in the tournament, and the rest gets paired. A total of `(n - 1) / 2` matches are played, and `(n - 1) / 2 + 1` teams advance to the next round.

Return *the number of matches played in the tournament until a winner is decided.*

**Example 1:**

```
Input: n = 7
Output: 6
Explanation: Details of the tournament: 
- 1st Round: Teams = 7, Matches = 3, and 4 teams advance.
- 2nd Round: Teams = 4, Matches = 2, and 2 teams advance.
- 3rd Round: Teams = 2, Matches = 1, and 1 team is declared the winner.
Total number of matches = 3 + 2 + 1 = 6.
```

**Example 2:**

```
Input: n = 14
Output: 13
Explanation: Details of the tournament:
- 1st Round: Teams = 14, Matches = 7, and 7 teams advance.
- 2nd Round: Teams = 7, Matches = 3, and 4 teams advance.
- 3rd Round: Teams = 4, Matches = 2, and 2 teams advance.
- 4th Round: Teams = 2, Matches = 1, and 1 team is declared the winner.
Total number of matches = 7 + 3 + 2 + 1 = 13.
```

**Constraints:**

- `1 <= n <= 200`

## 题目大意

给你一个整数 n ，表示比赛中的队伍数。比赛遵循一种独特的赛制：

- 如果当前队伍数是 偶数 ，那么每支队伍都会与另一支队伍配对。总共进行 n / 2 场比赛，且产生 n / 2 支队伍进入下一轮。
- 如果当前队伍数为 奇数 ，那么将会随机轮空并晋级一支队伍，其余的队伍配对。总共进行 (n - 1) / 2 场比赛，且产生 (n - 1) / 2 + 1 支队伍进入下一轮。

返回在比赛中进行的配对次数，直到决出获胜队伍为止。

## 解题思路

- 简单题，按照题目的规则模拟。
- 这一题还有更加简洁的代码，见解法一。n 个队伍，一个冠军，需要淘汰 n-1 个队伍。每一场比赛淘汰一个队伍，因此进行了 n-1 场比赛。所以共有 n-1 个配对。

## 代码

```go
package leetcode

// 解法一
func numberOfMatches(n int) int {
	return n - 1
}

// 解法二 模拟
func numberOfMatches1(n int) int {
	sum := 0
	for n != 1 {
		if n&1 == 0 {
			sum += n / 2
			n = n / 2
		} else {
			sum += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return sum
}
```