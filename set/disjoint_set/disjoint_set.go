// 集合相关的数据结构和算法
package set

// 不相交集是解决等价问题的一种有效数据结构：
// 实现简单、可使用简单数组实现、每种操作只需要常数平均时间。

/* -----等价关系------ */
/*
	对任意一对元素 (a, b), a, b∈S, aRb或者为true或为false，则称在集合S上定义关系R。如果aRb为true，那么说a与b有关系。
	等价关系是满足如下三个性质的关系R：
	1.（自反性）对于所有的a∈S，aRa；
	2.（对称性）aRb当且仅当bRa；
	3.（传递性）若aRb且bRc，则aRc；
*/

/* -----动态等价性问题------ */
/*
	已知等价关系“~”，对于任意的a和b，如何确定是否存在a~b。由于关系的定义通常不明显且可能相当隐秘，所以需要找到能快速推断出这些关系的方法。

	等价类：对于一个元素a∈S的等价类是S的一个子集，它包含所有与a有关系的元素。
	为确定是否a~b，只需验证a和b是否都在同一个等价类中。这为解决等价问题提供了方法。
*/

/*************************************************************************************/

// 使用树来表示每一个集合，因为树上每一个元素都有相同的根，所以该根可以用来命名所在的集合。
// 使用树表示每一个集合，集合的名字由根处节点给出；因只需父节点的名字，因此假设树被非显示的存储在一个数组中。
// 数组的每个成员P[i]表示元素i的父亲，若i是根，那么P[i] = 0。

type SetType int
type DisjSet []SetType

// 不相交集的初始化
func Initialize(numSets int, size SetType) DisjSet {
	ds := make(DisjSet, numSets+1)
	for i := numSets; i > 0; i-- {
		ds[i] = size
	}
	return ds
}

// Union(不是最好的方法)
func (ds DisjSet) SetUnion(root1, root2 SetType) {
	ds[root2] = root1
}

/*
// 一个简单的不相交集合Find算法
func (ds DisjSet) Find(x SetType) SetType {
	if ds[x] <= 0 {
		return x
	} else {
		return ds.Find(ds[x])
	}
}
*/

// 使用“路径压缩”的Find算法
func (ds DisjSet) Find(x SetType) SetType {
	if ds[x] <= 0 {
		return x
	} else {
		ds[x] = ds.Find(ds[x])
		return ds[x]
	}
}

// 按大小求并
// 在每个根的数组元素包含它的树的大小的负值
// 连续M次运算需要O(M)平均时间
func (ds DisjSet) UnionBySize(root1, root2 SetType) {
	if ds[root1] <= ds[root2] { // root1的树更大
		ds[root1] += ds[root2]
		ds[root2] = root1
	} else {
		ds[root2] += ds[root1]
		ds[root1] = root2
	}
}

// 按高度（秩）求并
// 在每个根的数组元素包含它的树的高度的负值
// 所有树的深度最多是O(logN)
func (ds DisjSet) UnionByHeight(root1, root2 SetType) {
	if ds[root1] < ds[root2] { //root1的树更深
		ds[root2] = root1
	} else {
		if ds[root1] == ds[root2] {
			ds[root1]--
		}
		ds[root2] = root1
	}
}
