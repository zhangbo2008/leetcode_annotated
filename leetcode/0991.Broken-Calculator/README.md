# [991. Broken Calculator](https://leetcode.com/problems/broken-calculator/)


## 题目

On a broken calculator that has a number showing on its display, we can perform two operations:

- **Double**: Multiply the number on the display by 2, or;
- **Decrement**: Subtract 1 from the number on the display.

Initially, the calculator is displaying the number `X`.

Return the minimum number of operations needed to display the number `Y`.

**Example 1:**

```
Input: X = 2, Y = 3
Output: 2
Explanation: Use double operation and then decrement operation {2 -> 4 -> 3}.
```

**Example 2:**

```
Input: X = 5, Y = 8
Output: 2
Explanation: Use decrement and then double {5 -> 4 -> 8}.
```

**Example 3:**

```
Input: X = 3, Y = 10
Output: 3
Explanation:  Use double, decrement and double {3 -> 6 -> 5 -> 10}.
```

**Example 4:**

```
Input: X = 1024, Y = 1
Output: 1023
Explanation: Use decrement operations 1023 times.
```

**Note:**

1. `1 <= X <= 10^9`
2. `1 <= Y <= 10^9`

## 题目大意

在显示着数字的坏计算器上，我们可以执行以下两种操作：

- 双倍（Double）：将显示屏上的数字乘 2；
- 递减（Decrement）：将显示屏上的数字减 1 。

最初，计算器显示数字 X。返回显示数字 Y 所需的最小操作数。

## 解题思路

- 看到本题的数据规模非常大，`10^9`，算法只能采用 `O(sqrt(n))`、`O(log n)`、`O(1)` 的算法。`O(sqrt(n))` 和 `O(1)` 在本题中是不可能的。所以按照数据规模来估计，本题只能尝试 `O(log n)` 的算法。`O(log n)` 的算法有二分搜索，不过本题不太符合二分搜索算法背景。题目中明显出现乘 2，这很明显是可以达到 `O(log n)` 的。最终确定解题思路是数学方法，循环中会用到乘 2 或者除 2 的计算。
- 既然出现了乘 2 和减一的操作，很容易考虑到奇偶性上。题目要求最小操作数，贪心思想，应该尽可能多的使用除 2 操作，使得 Y 和 X 大小差不多，最后再利用加一操作微调。只要 Y 比 X 大就执行除法操作。当然这里要考虑一次奇偶性，如果 Y 是奇数，先加一变成偶数再除二；如果 Y 是偶数，直接除二。如此操作直到 Y 不大于 X，最后执行 `X-Y` 次加法操作微调即可。

## 代码

```go
package leetcode

func brokenCalc(X int, Y int) int {
	res := 0
	for Y > X {
		res++
		if Y&1 == 1 {
			Y++
		} else {
			Y /= 2
		}
	}
	return res + X - Y
}
```