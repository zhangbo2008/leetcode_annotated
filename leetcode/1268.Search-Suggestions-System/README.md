# [1268. Search Suggestions System](https://leetcode.com/problems/search-suggestions-system/)

## 题目

Given an array of strings `products` and a string `searchWord`. We want to design a system that suggests at most three product names from `products` after each character of `searchWord` is typed. Suggested products should have common prefix with the searchWord. If there are more than three products with a common prefix return the three lexicographically minimums products.

Return *list of lists* of the suggested `products` after each character of `searchWord` is typed.

**Example 1:**

```
Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
Output: [
["mobile","moneypot","monitor"],
["mobile","moneypot","monitor"],
["mouse","mousepad"],
["mouse","mousepad"],
["mouse","mousepad"]
]
Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"]
After typing m and mo all products match and we show user ["mobile","moneypot","monitor"]
After typing mou, mous and mouse the system suggests ["mouse","mousepad"]

```

**Example 2:**

```
Input: products = ["havana"], searchWord = "havana"
Output: [["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
```

**Example 3:**

```
Input: products = ["bags","baggage","banner","box","cloths"], searchWord = "bags"
Output: [["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"],["bags"]]
```

**Example 4:**

```
Input: products = ["havana"], searchWord = "tatiana"
Output: [[],[],[],[],[],[],[]]
```

**Constraints:**

- `1 <= products.length <= 1000`
- There are no repeated elements in `products`.
- `1 <= Σ products[i].length <= 2 * 10^4`
- All characters of `products[i]` are lower-case English letters.
- `1 <= searchWord.length <= 1000`
- All characters of `searchWord` are lower-case English letters.

## 题目大意

给你一个产品数组 products 和一个字符串 searchWord ，products  数组中每个产品都是一个字符串。请你设计一个推荐系统，在依次输入单词 searchWord 的每一个字母后，推荐 products 数组中前缀与 searchWord 相同的最多三个产品。如果前缀相同的可推荐产品超过三个，请按字典序返回最小的三个。请你以二维列表的形式，返回在输入 searchWord 每个字母后相应的推荐产品的列表。

## 解题思路

- 由于题目要求返回的答案要按照字典序输出，所以先排序。有序字符串又满足了二分搜索的条件，于是可以用二分搜索。sort.SearchStrings 返回的是满足搜索条件的第一个起始下标。末尾不满足条件的字符串要切掉。所以要搜 2 次，第一次二分搜索先将不满足目标串前缀的字符串筛掉。第二次二分搜索再搜索出最终满足题意的字符串。

## 代码

```go
package leetcode

import (
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	searchWordBytes, result := []byte(searchWord), make([][]string, 0, len(searchWord))
	for i := 1; i <= len(searchWord); i++ {
		searchWordBytes[i-1]++
		products = products[:sort.SearchStrings(products, string(searchWordBytes[:i]))]
		searchWordBytes[i-1]--
		products = products[sort.SearchStrings(products, searchWord[:i]):]
		if len(products) > 3 {
			result = append(result, products[:3])
		} else {
			result = append(result, products)
		}
	}
	return result
}
```