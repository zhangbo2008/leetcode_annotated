# [1656. Design an Ordered Stream](https://leetcode.com/problems/design-an-ordered-stream/)

## 题目

There is a stream of `n` `(id, value)` pairs arriving in an **arbitrary** order, where `id` is an integer between `1` and `n` and `value` is a string. No two pairs have the same `id`.

Design a stream that returns the values in **increasing order of their IDs** by returning a **chunk** (list) of values after each insertion. The concatenation of all the **chunks** should result in a list of the sorted values.

Implement the `OrderedStream` class:

- `OrderedStream(int n)` Constructs the stream to take `n` values.
- `String[] insert(int id, String value)` Inserts the pair `(id, value)` into the stream, then returns the **largest possible chunk** of currently inserted values that appear next in the order.

**Example:**

![https://assets.leetcode.com/uploads/2020/11/10/q1.gif](https://assets.leetcode.com/uploads/2020/11/10/q1.gif)

```
Input
["OrderedStream", "insert", "insert", "insert", "insert", "insert"]
[[5], [3, "ccccc"], [1, "aaaaa"], [2, "bbbbb"], [5, "eeeee"], [4, "ddddd"]]
Output
[null, [], ["aaaaa"], ["bbbbb", "ccccc"], [], ["ddddd", "eeeee"]]

Explanation
// Note that the values ordered by ID is ["aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee"].
OrderedStream os = new OrderedStream(5);
os.insert(3, "ccccc"); // Inserts (3, "ccccc"), returns [].
os.insert(1, "aaaaa"); // Inserts (1, "aaaaa"), returns ["aaaaa"].
os.insert(2, "bbbbb"); // Inserts (2, "bbbbb"), returns ["bbbbb", "ccccc"].
os.insert(5, "eeeee"); // Inserts (5, "eeeee"), returns [].
os.insert(4, "ddddd"); // Inserts (4, "ddddd"), returns ["ddddd", "eeeee"].
// Concatentating all the chunks returned:
// [] + ["aaaaa"] + ["bbbbb", "ccccc"] + [] + ["ddddd", "eeeee"] = ["aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee"]
// The resulting order is the same as the order above.

```

**Constraints:**

- `1 <= n <= 1000`
- `1 <= id <= n`
- `value.length == 5`
- `value` consists only of lowercase letters.
- Each call to `insert` will have a unique `id.`
- Exactly `n` calls will be made to `insert`.

## 题目大意

有 n 个 (id, value) 对，其中 id 是 1 到 n 之间的一个整数，value 是一个字符串。不存在 id 相同的两个 (id, value) 对。

设计一个流，以 任意 顺序获取 n 个 (id, value) 对，并在多次调用时 按 id 递增的顺序 返回一些值。

实现 OrderedStream 类：

- OrderedStream(int n) 构造一个能接收 n 个值的流，并将当前指针 ptr 设为 1 。
- String[] insert(int id, String value) 向流中存储新的 (id, value) 对。存储后：
如果流存储有 id = ptr 的 (id, value) 对，则找出从 id = ptr 开始的 最长 id 连续递增序列 ，并 按顺序 返回与这些 id 关联的值的列表。然后，将 ptr 更新为最后那个 id + 1 。
否则，返回一个空列表。

## 解题思路

- 设计一个具有插入操作的 Ordered Stream。insert 操作先在指定位置插入 value，然后返回当前指针 ptr 到最近一个空位置的最长连续递增字符串。如果字符串不为空，ptr 移动到非空 value 的后一个下标位置处。
- 简单题。按照题目描述模拟即可。注意控制好 ptr 的位置。

## 代码

```go
package leetcode

type OrderedStream struct {
	ptr    int
	stream []string
}

func Constructor(n int) OrderedStream {
	ptr, stream := 1, make([]string, n+1)
	return OrderedStream{ptr: ptr, stream: stream}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.stream[id] = value
	res := []string{}
	if this.ptr == id || this.stream[this.ptr] != "" {
		res = append(res, this.stream[this.ptr])
		for i := id + 1; i < len(this.stream); i++ {
			if this.stream[i] != "" {
				res = append(res, this.stream[i])
			} else {
				this.ptr = i
				return res
			}
		}
	}
	if len(res) > 0 {
		return res
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */
```