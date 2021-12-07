# [80. Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/)

## 题目

Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

```c
Given nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.

It doesn't matter what you leave beyond the returned length.
```

Example 2:

```c
Given nums = [0,0,1,1,1,1,2,3,3],

Your function should return length = 7, with the first seven elements of nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively.

It doesn't matter what values are set beyond the returned length.
```

Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```c
// nums is passed in by reference. (i.e., without making a copy)
int len = removeElement(nums, val);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

## 题目大意

给定一个有序数组 nums，对数组中的元素进行去重，使得原数组中的每个元素最多暴露 2 个。最后返回去重以后数组的长度值。

## 解题思路

- 问题提示有序数组，一般最容易想到使用双指针的解法，双指针的关键点：移动两个指针的条件。
- 在该题中移动的条件：快指针从头遍历数组，慢指针指向修改后的数组的末端，当慢指针指向倒数第二个数与快指针指向的数不相等时，才移动慢指针，同时赋值慢指针。
- 处理边界条件：当数组小于两个元素时，不做处理。
