# [488. Zuma Game](https://leetcode.com/problems/zuma-game/)


## 题目

You are playing a variation of the game Zuma.

In this variation of Zuma, there is a single row of colored balls on a board, where each ball can be colored red 'R', yellow 'Y', blue 'B', green 'G', or white 'W'. You also have several colored balls in your hand.

Your goal is to clear all of the balls from the board. On each turn:

Pick any ball from your hand and insert it in between two balls in the row or on either end of the row.
If there is a group of three or more consecutive balls of the same color, remove the group of balls from the board.
If this removal causes more groups of three or more of the same color to form, then continue removing each group until there are none left.
If there are no more balls on the board, then you win the game.
Repeat this process until you either win or do not have any more balls in your hand.
Given a string board, representing the row of balls on the board, and a string hand, representing the balls in your hand, return the minimum number of balls you have to insert to clear all the balls from the board. If you cannot clear all the balls from the board using the balls in your hand, return -1.

**Example 1**:

```
Input: board = "WRRBBW", hand = "RB"
Output: -1
Explanation: It is impossible to clear all the balls. The best you can do is:
- Insert 'R' so the board becomes WRRRBBW. WRRRBBW -> WBBW.
- Insert 'B' so the board becomes WBBBW. WBBBW -> WW.
There are still balls remaining on the board, and you are out of balls to insert.
```

**Example 2**:
```
Input: board = "WWRRBBWW", hand = "WRBRW"
Output: 2
Explanation: To make the board empty:
- Insert 'R' so the board becomes WWRRRBBWW. WWRRRBBWW -> WWBBWW.
- Insert 'B' so the board becomes WWBBBWW. WWBBBWW -> WWWW -> empty.
2 balls from your hand were needed to clear the board.
```

**Example 3**:
```
Input: board = "G", hand = "GGGGG"
Output: 2
Explanation: To make the board empty:
- Insert 'G' so the board becomes GG.
- Insert 'G' so the board becomes GGG. GGG -> empty.
2 balls from your hand were needed to clear the board.
```

**Example 4**:
```
Input: board = "RBYYBBRRB", hand = "YRBGB"
Output: 3
Explanation: To make the board empty:
- Insert 'Y' so the board becomes RBYYYBBRRB. RBYYYBBRRB -> RBBBRRB -> RRRB -> B.
- Insert 'B' so the board becomes BB.
- Insert 'B' so the board becomes BBB. BBB -> empty.
3 balls from your hand were needed to clear the board.
```

**Constraints**:

- 1 <= board.length <= 16
- 1 <= hand.length <= 5
- board and hand consist of the characters 'R', 'Y', 'B', 'G', and 'W'.
- The initial row of balls on the board will not have any groups of three or more consecutive balls of the same color.

## 题目大意

你正在参与祖玛游戏的一个变种。

在这个祖玛游戏变体中，桌面上有 一排 彩球，每个球的颜色可能是：红色 'R'、黄色 'Y'、蓝色 'B'、绿色 'G' 或白色 'W' 。你的手中也有一些彩球。

你的目标是 清空 桌面上所有的球。每一回合：

从你手上的彩球中选出 任意一颗 ，然后将其插入桌面上那一排球中：两球之间或这一排球的任一端。
接着，如果有出现 三个或者三个以上 且 颜色相同 的球相连的话，就把它们移除掉。
如果这种移除操作同样导致出现三个或者三个以上且颜色相同的球相连，则可以继续移除这些球，直到不再满足移除条件。
如果桌面上所有球都被移除，则认为你赢得本场游戏。
重复这个过程，直到你赢了游戏或者手中没有更多的球。
给你一个字符串 board ，表示桌面上最开始的那排球。另给你一个字符串 hand ，表示手里的彩球。请你按上述操作步骤移除掉桌上所有球，计算并返回所需的 最少 球数。如果不能移除桌上所有的球，返回 -1 。

## 解题思路

- 使用广度优先搜索和剪枝

## 代码

```go

package leetcode

func findMinStep(board string, hand string) int {
	q := [][]string{{board, hand}}
	mp := make(map[string]bool)
	minStep := 0
	for len(q) > 0 {
		length := len(q)
		minStep++
		for length > 0 {
			length--
			cur := q[0]
			q = q[1:]
			curB, curH := cur[0], cur[1]
			for i := 0; i < len(curB); i++ {
				for j := 0; j < len(curH); j++ {
					curB2 := del3(curB[0:i] + string(curH[j]) + curB[i:])
					curH2 := curH[0:j] + curH[j+1:]
					if len(curB2) == 0 {
						return minStep
					}
					if _, ok := mp[curB2+curH2]; ok {
						continue
					}
					mp[curB2+curH2] = true
					q = append(q, []string{curB2, curH2})
				}
			}
		}
	}
	return -1
}

func del3(str string) string {
	cnt := 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			cnt++
		} else {
			if cnt >= 3 {
				return del3(str[0:i-cnt] + str[i:])
			}
			cnt = 1
		}
	}
	if cnt >= 3 {
		return str[0 : len(str)-cnt]
	}
	return str
}
```