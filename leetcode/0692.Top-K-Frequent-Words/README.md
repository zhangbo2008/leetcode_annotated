# [692. Top K Frequent Words](https://leetcode.com/problems/top-k-frequent-words/)


## 题目

Given a non-empty list of words, return the k most frequent elements.

Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

**Example 1:**

```
Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
Output: ["i", "love"]
Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.
```

**Example 2:**

```
Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
Output: ["the", "is", "sunny", "day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
    with the number of occurrence being 4, 3, 2 and 1 respectively.
```

**Note:**

1. You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
2. Input words contain only lowercase letters.

**Follow up:**

1. Try to solve it in O(n log k) time and O(n) extra space.

## 题目大意

给一非空的单词列表，返回前 *k* 个出现次数最多的单词。返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。

## 解题思路

- 思路很简单的题。维护一个长度为 k 的最大堆，先按照频率排，如果频率相同再按照字母顺序排。最后输出依次将优先队列里面的元素 pop 出来即可。

## 代码

```go
package leetcode

import "container/heap"

func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for _, word := range words {
		m[word]++
	}
	pq := &PQ{}
	heap.Init(pq)
	for w, c := range m {
		heap.Push(pq, &wordCount{w, c})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		wc := heap.Pop(pq).(*wordCount)
		res[i] = wc.word
	}
	return res
}

type wordCount struct {
	word string
	cnt  int
}

type PQ []*wordCount

func (pq PQ) Len() int      { return len(pq) }
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQ) Less(i, j int) bool {
	if pq[i].cnt == pq[j].cnt {
		return pq[i].word > pq[j].word
	}
	return pq[i].cnt < pq[j].cnt
}
func (pq *PQ) Push(x interface{}) {
	tmp := x.(*wordCount)
	*pq = append(*pq, tmp)
}
func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
```