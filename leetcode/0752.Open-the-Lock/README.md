# [752. Open the Lock](https://leetcode.com/problems/open-the-lock/)


## 题目

You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: `'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'`. The wheels can rotate freely and wrap around: for example we can turn `'9'` to be `'0'`, or `'0'` to be `'9'`. Each move consists of turning one wheel one slot.

The lock initially starts at `'0000'`, a string representing the state of the 4 wheels.

You are given a list of `deadends` dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

Given a `target` representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.

**Example 1:**

```
Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation:
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".

```

**Example 2:**

```
Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation:
We can turn the last wheel in reverse to move from "0000" -> "0009".

```

**Example 3:**

```
Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
Output: -1
Explanation:
We can't reach the target without getting stuck.

```

**Example 4:**

```
Input: deadends = ["0000"], target = "8888"
Output: -1

```

**Constraints:**

- `1 <= deadends.length <= 500`
- `deadends[i].length == 4`
- `target.length == 4`
- target **will not be** in the list `deadends`.
- `target` and `deadends[i]` consist of digits only.

## 题目大意

你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

## 解题思路

- 此题可以转化为从起始点到终点的最短路径。采用广度优先搜索。每次广搜枚举转动一次数字的状态，并且用 visited 记录是否被搜索过，如果没有被搜索过，便加入队列，下一轮继续搜索。如果搜索到了 target，便返回对应的旋转次数。如果搜索完成后，仍没有搜索到 target，说明无法解锁，返回 -1。特殊情况，如果 target 就是初始数字 0000，那么直接返回答案 0。
- 在广搜之前，先将 deadends 放入 map 中，搜索中判断是否搜到了 deadends。如果初始数字 0000 出现在 deadends 中，可以直接返回答案 −1。

## 代码

```go
package leetcode

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	targetNum, visited := strToInt(target), make([]bool, 10000)
	visited[0] = true
	for _, deadend := range deadends {
		num := strToInt(deadend)
		if num == 0 {
			return -1
		}
		visited[num] = true
	}
	depth, curDepth, nextDepth := 0, []int16{0}, make([]int16, 0)
	var nextNum int16
	for len(curDepth) > 0 {
		nextDepth = nextDepth[0:0]
		for _, curNum := range curDepth {
			for incrementer := int16(1000); incrementer > 0; incrementer /= 10 {
				digit := (curNum / incrementer) % 10
				if digit == 9 {
					nextNum = curNum - 9*incrementer
				} else {
					nextNum = curNum + incrementer
				}
				if nextNum == targetNum {
					return depth + 1
				}
				if !visited[nextNum] {
					visited[nextNum] = true
					nextDepth = append(nextDepth, nextNum)
				}
				if digit == 0 {
					nextNum = curNum + 9*incrementer
				} else {
					nextNum = curNum - incrementer
				}
				if nextNum == targetNum {
					return depth + 1
				}
				if !visited[nextNum] {
					visited[nextNum] = true
					nextDepth = append(nextDepth, nextNum)
				}
			}
		}
		curDepth, nextDepth = nextDepth, curDepth
		depth++
	}
	return -1
}

func strToInt(str string) int16 {
	return int16(str[0]-'0')*1000 + int16(str[1]-'0')*100 + int16(str[2]-'0')*10 + int16(str[3]-'0')
}
```