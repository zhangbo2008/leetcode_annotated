# [609. Find Duplicate File in System](https://leetcode.com/problems/find-duplicate-file-in-system/)


## 题目

Given a list `paths` of directory info, including the directory path, and all the files with contents in this directory, return *all the duplicate files in the file system in terms of their paths*. You may return the answer in **any order**.

A group of duplicate files consists of at least two files that have the same content.

A single directory info string in the input list has the following format:

- `"root/d1/d2/.../dm f1.txt(f1_content) f2.txt(f2_content) ... fn.txt(fn_content)"`

It means there are `n` files `(f1.txt, f2.txt ... fn.txt)` with content `(f1_content, f2_content ... fn_content)` respectively in the directory "`root/d1/d2/.../dm"`. Note that `n >= 1` and `m >= 0`. If `m = 0`, it means the directory is just the root directory.

The output is a list of groups of duplicate file paths. For each group, it contains all the file paths of the files that have the same content. A file path is a string that has the following format:

- `"directory_path/file_name.txt"`

**Example 1:**

```
Input: paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"]
Output: [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]

```

**Example 2:**

```
Input: paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)"]
Output: [["root/a/2.txt","root/c/d/4.txt"],["root/a/1.txt","root/c/3.txt"]]

```

**Constraints:**

- `1 <= paths.length <= 2 * 104`
- `1 <= paths[i].length <= 3000`
- `1 <= sum(paths[i].length) <= 5 * 105`
- `paths[i]` consist of English letters, digits, `'/'`, `'.'`, `'('`, `')'`, and `' '`.
- You may assume no files or directories share the same name in the same directory.
- You may assume each given directory info represents a unique directory. A single blank space separates the directory path and file info.

**Follow up:**

- Imagine you are given a real file system, how will you search files? DFS or BFS?
- If the file content is very large (GB level), how will you modify your solution?
- If you can only read the file by 1kb each time, how will you modify your solution?
- What is the time complexity of your modified solution? What is the most time-consuming part and memory-consuming part of it? How to optimize?
- How to make sure the duplicated files you find are not false positive?

## 题目大意

给定一个目录信息列表，包括目录路径，以及该目录中的所有包含内容的文件，您需要找到文件系统中的所有重复文件组的路径。一组重复的文件至少包括二个具有完全相同内容的文件。输入列表中的单个目录信息字符串的格式如下：`"root/d1/d2/.../dm f1.txt(f1_content) f2.txt(f2_content) ... fn.txt(fn_content)"`。这意味着有 n 个文件（`f1.txt, f2.txt ... fn.txt` 的内容分别是 `f1_content, f2_content ... fn_content`）在目录 `root/d1/d2/.../dm` 下。注意：n>=1 且 m>=0。如果 m=0，则表示该目录是根目录。该输出是重复文件路径组的列表。对于每个组，它包含具有相同内容的文件的所有文件路径。文件路径是具有下列格式的字符串：`"directory_path/file_name.txt"`

## 解题思路

- 这一题算简单题，考察的是字符串基本操作与 map 的使用。首先通过字符串操作获取目录路径、文件名和文件内容。再使用 map 来寻找重复文件，key 是文件内容，value 是存储路径和文件名的列表。遍历每一个文件，并把它加入 map 中。最后遍历 map，如果一个键对应的值列表的长度大于 1，说明找到了重复文件，可以把这个列表加入到最终答案中。
- 这道题有价值的地方在 **Follow up** 中。感兴趣的读者可以仔细想想以下几个问题：
    1. 假设您有一个真正的文件系统，您将如何搜索文件？广度搜索还是宽度搜索？
    2. 如果文件内容非常大（GB级别），您将如何修改您的解决方案？
    3. 如果每次只能读取 1 kb 的文件，您将如何修改解决方案？
    4. 修改后的解决方案的时间复杂度是多少？其中最耗时的部分和消耗内存的部分是什么？如何优化？
    5. 如何确保您发现的重复文件不是误报？

## 代码

```go
package leetcode

import "strings"

func findDuplicate(paths []string) [][]string {
	cache := make(map[string][]string)
	for _, path := range paths {
		parts := strings.Split(path, " ")
		dir := parts[0]
		for i := 1; i < len(parts); i++ {
			bracketPosition := strings.IndexByte(parts[i], '(')
			content := parts[i][bracketPosition+1 : len(parts[i])-1]
			cache[content] = append(cache[content], dir+"/"+parts[i][:bracketPosition])
		}
	}
	res := make([][]string, 0, len(cache))
	for _, group := range cache {
		if len(group) >= 2 {
			res = append(res, group)
		}
	}
	return res
}
```