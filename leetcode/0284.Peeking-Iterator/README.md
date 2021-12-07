# [284. Peeking Iterator](https://leetcode.com/problems/peeking-iterator/)

## 题目

Given an Iterator class interface with methods: `next()` and `hasNext()`, design and implement a PeekingIterator that support the `peek()` operation -- it essentially peek() at the element that will be returned by the next call to next().

**Example:**

```
Assume that the iterator is initialized to the beginning of the list: [1,2,3].

Call next() gets you 1, the first element in the list.
Now you call peek() and it returns 2, the next element. Calling next() after that still return 2. 
You call next() the final time and it returns 3, the last element. 
Calling hasNext() after that should return false.
```

**Follow up**: How would you extend your design to be generic and work with all types, not just integer?

## 题目大意

给定一个迭代器类的接口，接口包含两个方法： next() 和 hasNext()。设计并实现一个支持 peek() 操作的顶端迭代器 -- 其本质就是把原本应由 next() 方法返回的元素 peek() 出来。

> peek() 是偷看的意思。偷偷看一看下一个元素是什么，但是并不是 next() 访问。

## 解题思路

- 简单题。在 PeekingIterator 内部保存 2 个变量，一个是下一个元素值，另一个是是否有下一个元素。在 next() 操作和 hasNext() 操作时，访问保存的这 2 个变量。peek() 操作也比较简单，判断是否有下一个元素，如果有，即返回该元素值。这里实现了迭代指针不移动的功能。如果没有保存下一个元素值，即没有 peek() 偷看，next() 操作继续往后移动指针，读取后一位元素。
- 这里复用了是否有下一个元素值，来判断 hasNext() 和 peek() 操作中不移动指针的逻辑。

## 代码

```go
package leetcode

//Below is the interface for Iterator, which is already defined for you.

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return true
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	nextEl int
	hasEl  bool
	iter   *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	if this.hasEl {
		return true
	}
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.hasEl {
		this.hasEl = false
		return this.nextEl
	} else {
		return this.iter.next()
	}
}

func (this *PeekingIterator) peek() int {
	if this.hasEl {
		return this.nextEl
	}
	this.hasEl = true
	this.nextEl = this.iter.next()
	return this.nextEl
}
```