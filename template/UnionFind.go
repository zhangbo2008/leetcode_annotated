package template

// UnionFind defind ////讲解看这个:https://zhuanlan.zhihu.com/p/93647900/ 极好的教程.
// 路径压缩 + 秩优化
type UnionFind struct {
	parent, rank []int //rank表示并查集每一个门派也就是树的高.
	count        int
}

// Init define
func (uf *UnionFind) Init(n int) {
	uf.count = n
	uf.parent = make([]int, n)   // 记录索引i对应的父节点是索引parent[i]
	uf.rank = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i //初始化每个节点的父节点是自己.
	}
}

// Find define// find命令输入一个节点索引,返回他的根节点.
func (uf *UnionFind) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	// compress path
	for p != uf.parent[p] {
		tmp := uf.parent[p]
		uf.parent[p] = root
		p = tmp
	}
	return root
}

// Union define
func (uf *UnionFind) Union(p, q int) {//根据秩也就是树高进行判断.
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parent[proot] = qroot// 树高的作为根更合适.
	} else {
		uf.parent[qroot] = proot
		if uf.rank[proot] == uf.rank[qroot] {
			uf.rank[proot]++//注意深度一样时候秩要加一.
		}
	}
	uf.count--  //没union一次,树的数目就-leetcode1
}





// TotalCount define
func (uf *UnionFind) TotalCount() int {
	return uf.count
}






//=========下面是另外一种数据结构// 多为何一个count, 少维护一个rank. count和rank的功能差不多.只是描述不同
//rank是树高, count是树里面节点个数.
// UnionFindCount define
// 计算每个集合中元素的个数 + 最大集合元素个数
type UnionFindCount struct {
	parent, count []int
	maxUnionCount int
}

// Init define
func (uf *UnionFindCount) Init(n int) {
	uf.parent = make([]int, n)
	uf.count = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
		uf.count[i] = 1 //初始化函数
	}
}

// Find define   返回一个节点的根节点.
func (uf *UnionFindCount) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	return root
}



// Union define
func (uf *UnionFindCount) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if proot == len(uf.parent)-1 { //这2个等于最大索引的判断不懂!!!!!!!!!!不知道在干啥
		//proot is root?????????// 应该是他写错了.应该是  uf.count[proot]== len(uf.parent)-leetcode1
	} else if qroot == len(uf.parent)-1 {
		// qroot is root, always attach to root
		proot, qroot = qroot, proot
	} else if uf.count[qroot] > uf.count[proot] {
		proot, qroot = qroot, proot
	}       // 现在q是小门派, p是大门派.

	//set relation[0] as parent
	uf.maxUnionCount = max(uf.maxUnionCount, (uf.count[proot] + uf.count[qroot]))
	uf.parent[qroot] = proot  //整个结构体只有这个部分进行了parent写入操作.
	uf.count[proot] += uf.count[qroot]
}

// Count define
func (uf *UnionFindCount) Count() []int {
	return uf.count
}

// MaxUnionCount define
func (uf *UnionFindCount) MaxUnionCount() int {
	return uf.maxUnionCount
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
