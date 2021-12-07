# [677. Map Sum Pairs](https://leetcode.com/problems/map-sum-pairs/)


## 题目

Design a map that allows you to do the following:

- Maps a string key to a given value.
- Returns the sum of the values that have a key with a prefix equal to a given string.

Implement the `MapSum` class:

- `MapSum()` Initializes the `MapSum` object.
- `void insert(String key, int val)` Inserts the `key-val` pair into the map. If the `key` already existed, the original `key-value` pair will be overridden to the new one.
- `int sum(string prefix)` Returns the sum of all the pairs' value whose `key` starts with the `prefix`.

**Example 1:**

```
Input
["MapSum", "insert", "sum", "insert", "sum"]
[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
Output
[null, null, 3, null, 5]

Explanation
MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);
mapSum.sum("ap");           // return 3 (apple = 3)
mapSum.insert("app", 2);
mapSum.sum("ap");           // return 5 (apple +app = 3 + 2 = 5)

```

**Constraints:**

- `1 <= key.length, prefix.length <= 50`
- `key` and `prefix` consist of only lowercase English letters.
- `1 <= val <= 1000`
- At most `50` calls will be made to `insert` and `sum`.

## 题目大意

实现一个 MapSum 类，支持两个方法，insert 和 sum：

- MapSum() 初始化 MapSum 对象
- void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
- int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。

## 解题思路

- 简单题。用一个 map 存储数据，Insert() 方法即存储 key-value。Sum() 方法即累加满足条件前缀对应的 value。判断是否满足条件，先根据前缀长度来判断，只有长度大于等于 prefix 长度才可能满足要求。如果 key 是具有 prefix 前缀的，那么累加上这个值。最后输出总和即可。

## 代码

```go
package leetcode

type MapSum struct {
	keys map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.keys[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	prefixAsRunes, res := []rune(prefix), 0
	for key, val := range this.keys {
		if len(key) >= len(prefix) {
			shouldSum := true
			for i, char := range key {
				if i >= len(prefixAsRunes) {
					break
				}
				if prefixAsRunes[i] != char {
					shouldSum = false
					break
				}
			}
			if shouldSum {
				res += val
			}
		}
	}
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
```