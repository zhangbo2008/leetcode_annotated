# [492. Construct the Rectangle](https://leetcode.com/problems/construct-the-rectangle/)


## 题目

A web developer needs to know how to design a web page's size. 
So, given a specific rectangular web page’s area, your job by now is to design a rectangular web page,
whose length L and width W satisfy the following requirements:

    The area of the rectangular web page you designed must equal to the given target area.
    The width W should not be larger than the length L, which means L >= W.
    The difference between length L and width W should be as small as possible.
    Return an array [L, W] where L and W are the length and width of the web page you designed in sequence.

**Example 1:**

    Input: area = 4
    Output: [2,2]
    Explanation: The target area is 4, and all the possible ways to construct it are [1,4], [2,2], [4,1].
    But according to requirement 2, [1,4] is illegal; according to requirement 3,  [4,1] is not optimal compared to [2,2]. So the length L is 2, and the width W is 2.

**Example 2:**

    Input: area = 37
    Output: [37,1]

**Example 3:**

    Input: area = 122122
    Output: [427,286]

**Constraints**

   - 1 <= area <= 10000000

## 题目大意

作为一位 web 开发者， 懂得怎样去规划一个页面的尺寸是很重要的。 现给定一个具体的矩形页面面积，你的任务是设计一个长度为 L 和宽度为 W 且满足以下要求的矩形的页面。要求：

1. 你设计的矩形页面必须等于给定的目标面积。
2. 宽度 W 不应大于长度 L，换言之，要求 L >= W 。
3. 长度 L 和宽度 W 之间的差距应当尽可能小。 

你需要按顺序输出你设计的页面的长度 L 和宽度 W。

## 解题思路

- 令 W 等于根号 area
- 在 W 大于等于 1 的情况下,判断 area%W 是否等于 0,如果不相等 W 就减 1 继续循环,如果相等就返回 [area/W, W]

## 代码

```go

package leetcode

import "math"

func constructRectangle(area int) []int {
	ans := make([]int, 2)
	W := int(math.Sqrt(float64(area)))
	for W >= 1 {
		if area%W == 0 {
			ans[0], ans[1] = area/W, W
			break
		}
		W -= 1
	}
	return ans
}

``