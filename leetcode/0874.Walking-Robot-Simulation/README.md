# [874. Walking Robot Simulation](https://leetcode.com/problems/walking-robot-simulation/)


## 题目

A robot on an infinite XY-plane starts at point `(0, 0)` and faces north. The robot can receive one of three possible types of `commands`:

- `2`: turn left `90` degrees,
- `1`: turn right `90` degrees, or
- `1 <= k <= 9`: move forward `k` units.

Some of the grid squares are `obstacles`. The `ith` obstacle is at grid point `obstacles[i] = (xi, yi)`.

If the robot would try to move onto them, the robot stays on the previous grid square instead (but still continues following the rest of the route.)

Return *the maximum Euclidean distance that the robot will be from the origin **squared** (i.e. if the distance is* `5`*, return* `25`*)*.

**Note:**

- North means +Y direction.
- East means +X direction.
- South means -Y direction.
- West means -X direction.

**Example 1:**

```
Input: commands = [4,-1,3], obstacles = []
Output: 25
Explanation: The robot starts at (0, 0):
1. Move north 4 units to (0, 4).
2. Turn right.
3. Move east 3 units to (3, 4).
The furthest point away from the origin is (3, 4), which is 32 + 42 = 25 units away.

```

**Example 2:**

```
Input: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
Output: 65
Explanation: The robot starts at (0, 0):
1. Move north 4 units to (0, 4).
2. Turn right.
3. Move east 1 unit and get blocked by the obstacle at (2, 4), robot is at (1, 4).
4. Turn left.
5. Move north 4 units to (1, 8).
The furthest point away from the origin is (1, 8), which is 12 + 82 = 65 units away.

```

**Constraints:**

- `1 <= commands.length <= 104`
- `commands[i]` is one of the values in the list `[-2,-1,1,2,3,4,5,6,7,8,9]`.
- `0 <= obstacles.length <= 104`
- `3 * 104 <= xi, yi <= 3 * 104`
- The answer is guaranteed to be less than `231`.

## 题目大意

机器人在一个无限大小的 XY 网格平面上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 commands ：

- 2 ：向左转 90 度
- -1 ：向右转 90 度
- 1 <= x <= 9 ：向前移动 x 个单位长度

在网格上有一些格子被视为障碍物 obstacles 。第 i 个障碍物位于网格点 obstacles[i] = (xi, yi) 。机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）

注意：

- 北表示 +Y 方向。
- 东表示 +X 方向。
- 南表示 -Y 方向。
- 西表示 -X 方向。

## 解题思路

- 这个题的难点在于，怎么用编程语言去描述机器人的行为，可以用以下数据结构表达机器人的行为：

    ```go
    direct:= 0                    // direct表示机器人移动方向：0 1 2 3 4 （北东南西），默认朝北
    x, y := 0, 0                  // 表示当前机器人所在横纵坐标位置，默认为(0,0)
    directX := []int{0, 1, 0, -1}
    directY := []int{1, 0, -1, 0}
    // 组合directX directY和direct，表示机器人往某一个方向移动
    nextX := x + directX[direct]
    nextY := y + directY[direct]
    ```

    其他代码按照题意翻译即可

## 代码

```go
package leetcode

func robotSim(commands []int, obstacles [][]int) int {
	m := make(map[[2]int]struct{})
	for _, v := range obstacles {
		if len(v) != 0 {
			m[[2]int{v[0], v[1]}] = struct{}{}
		}
	}
	directX := []int{0, 1, 0, -1}
	directY := []int{1, 0, -1, 0}
	direct, x, y := 0, 0, 0
	result := 0
	for _, c := range commands {
		if c == -2 {
			direct = (direct + 3) % 4
			continue
		}
		if c == -1 {
			direct = (direct + 1) % 4
			continue
		}
		for ; c > 0; c-- {
			nextX := x + directX[direct]
			nextY := y + directY[direct]
			if _, ok := m[[2]int{nextX, nextY}]; ok {
				break
			}
			tmpResult := nextX*nextX + nextY*nextY
			if tmpResult > result {
				result = tmpResult
			}
			x = nextX
			y = nextY
		}
	}
	return result
}
```