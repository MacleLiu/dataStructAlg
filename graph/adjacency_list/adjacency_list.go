package graph

import (
	"container/heap"
	"fmt"

	set "dataStructAlg/set/disjoint_set"

	"gopkg.in/eapache/queue.v1"
)

const N = 7            //图的顶点数
const E = 10           //图的边数
const Infinity = 10000 //路径长度未知或不可达

type vextype string    //定义顶点的数据信息类型
type edgenode struct { //邻接链表节点
	Adjvex int       //邻接点域
	Weight int       //权值
	Next   *edgenode //链域
}
type vexnode struct {
	Vertex vextype   //顶点域
	In     int       //入度域
	Link   *edgenode //指针域
}
type GraphList []vexnode      //顶点表
type book_keeepTable struct { //簿记表数据结构
	Known bool
	Dist  int
	Path  int
}

var visited [N]int //顶点遍历标记

// 生成图的邻接表
func CreatAdjList(digraph bool) GraphList {
	ga := make(GraphList, N)
	var i, j, k, w int
	for i = 0; i < N; i++ {
		fmt.Printf("输入顶点%d及其入度:\n", i+1)
		fmt.Scanf("%s%d\n", &ga[i].Vertex, &ga[i].In) //读入顶点信息和边表头指针初始化
		ga[i].Link = nil
	}
	for k = 0; k < E; k++ { //建立边表
		fmt.Println("输入边(Vi,Vj)的顶点序号和权值:")
		fmt.Scanf("%d%d%d\n", &i, &j, &w) //读入边(Vi,Vj)的顶点序
		ga[i-1].Link = &edgenode{         //将邻接点序号为j的边表节点插入顶点Vi的边表头部
			j - 1,
			w,
			ga[i-1].Link,
		}
		if !digraph { // 是否生成有向图
			ga[j-1].Link = &edgenode{ //将邻接点序号为i的边表节点插入顶点Vj的边表头部
				i - 1,
				w,
				ga[j-1].Link,
			}
		}
	}
	return ga
}

// 打印邻接表
func (ga GraphList) Print() {
	fmt.Println("\n/*******************************/")
	for _, vexnode := range ga {
		fmt.Printf("%s %d", vexnode.Vertex, vexnode.In)
		for vexnode.Link != nil {
			fmt.Printf(" -> %d", vexnode.Link.Adjvex)
			vexnode.Link = vexnode.Link.Next
		}
		fmt.Println()
	}
	fmt.Println("/*******************************/")
}

// 深度优先遍历（DFS）
func (ga GraphList) Dfs(i int) {
	fmt.Println(ga[i].Vertex) //访问顶点Vi
	visited[i] = 1            //标记Vi已被访问
	p := ga[i].Link
	for p != nil { //遍历Vi的邻接顶点
		if visited[p.Adjvex] == 0 {
			ga.Dfs(p.Adjvex)
		}
		p = p.Next
	}
}

// 广度优先搜索遍历（BFS）
func (ga GraphList) Bfs(i int) {
	var m int
	q := queue.New()
	fmt.Println(ga[i].Vertex) //访问出发顶点Vi
	visited[i] = 1
	q.Add(i) //访问过的顶点序号入队列
	for q.Length() != 0 {
		m = q.Remove().(int)
		p := ga[m].Link //取Vm的边表头指针
		for p != nil {
			if visited[p.Adjvex] == 0 {
				fmt.Println(ga[p.Adjvex].Vertex)
				visited[p.Adjvex] = 1
				q.Add(p.Adjvex)
			}
			p = p.Next
		}
	}
}

/*
	拓扑排序是对有向无圈图的顶点的一种排序，使得存在的路径Vi到Vj,那么在排序中Vj出现在Vi后面。
	简单的拓扑排序方法是先找出任意一个没有入边的顶点，然后将它和它的边一起从图中删除，然后对其他部分重复该操作。
*/

// 拓扑排序
// 为便于检查每个顶点的入度，在顶点列表中增加一个入度域（In）。
// 为避免对每个入度为0的顶点进行重复访问，可用一个链栈来存储所有入度为0的顶点。
// 链栈无需额外空间，只需要利用顶点表中入度域值为0的入度域来存放链栈的指针（指向下一个入读域值为0的单元的下标），并用一个栈顶指针top指向栈顶即可。
// 算法步骤：
// 1.扫描顶点表，将入度为0的顶点入栈。
//
//	2.当栈非空时：{
//			弹出栈顶顶点Vi的序号，并输出该顶点；
//			检查Vi的出边表，将每条出边表邻接点域所对应的顶点入度域值减1，若该顶点入度域为0，则将其入栈；
//	}
//
// 3.若输出的顶点数小于N，则“图中存在环”；否则拓扑排序正常结束。
func (ga GraphList) TopoSort() {
	var i, j, k, m int      //m为输出的顶点个数计数器
	var top = -1            //top为栈顶指针
	for i = 0; i < N; i++ { //初始化，建立入度为0的顶点链栈
		if ga[i].In == 0 {
			ga[i].In = top
			top = i
		}
	}
	for top != -1 { //栈非空，执行排序操作
		j = top
		top = ga[top].In                 //第j+1个顶点出栈
		fmt.Printf("%s\t", ga[j].Vertex) //输出出栈的顶点
		m++
		p := ga[j].Link
		for p != nil { //删除所有以Vj+1为起点的出边
			k = p.Adjvex
			ga[k].In--         //将Vk+1的入度减1
			if ga[k].In == 0 { //将入度为0的顶点入栈
				ga[k].In = top
				top = k
			}
			p = p.Next //找Vj+1的下一条边
		}
	}
	if m < N {
		fmt.Println("\n图中存在环")
	}
}

// 无权最短路径
// 入参数为选择的起始顶点在数组中的下标
// 由于双层嵌套for循环该算法的时间复杂度为O(|V|^2)，
// 因为尽管所有的顶点早已成为Known，但是外层循环还是要继续，直到N-1为止。
/*
func (ga GraphList) Unweighted(s int) {
	// 声明并初始化簿记表
	var t [N]book_keeepTable
	for i := 0; i < N; i++ {
		if s == i {
			continue
		}
		t[i].Dist = -1
	}
	for currDist := 0; currDist < N; currDist++ {
		// 遍历每一个顶点
		for i := 0; i < N; i++ {
			// 如果该顶点时未知的，且簿记表中的距离等于当前循环控制的距离
			if !t[i].Known && t[i].Dist == currDist {
				t[i].Known = true //将该顶点标记为已知
				// 遍历当前顶点i的全部邻接顶点
				for ga[i].Link != nil {
					if t[ga[i].Link.Adjvex].Dist == -1 {
						t[ga[i].Link.Adjvex].Dist = currDist + 1
						t[ga[i].Link.Adjvex].Path = ga[i].Vertex
					}
					ga[i].Link = ga[i].Link.Next
				}
			}
		}
	}
	// 打印簿记表
	for i, item := range t {
		fmt.Printf("V%d %v %d V%s\n", i+1, item.Known, item.Dist, item.Path)
	}
}
*/

// 优化后的无权最短路径算法
// 使用了第三方的队列库
func (ga GraphList) Unweighted(s int) {
	// 声明并初始化簿记表
	var t [N]book_keeepTable
	for i := 0; i < N; i++ {
		if s == i {
			continue
		}
		t[i].Dist = Infinity
	}
	q := queue.New()
	q.Add(s) //起始顶点入队

	for q.Length() != 0 {
		v := q.Remove().(int)
		t[v].Known = true //该算法并不是真的需要Known字段
		//遍历邻接顶点
		for ga[v].Link != nil {
			if t[ga[v].Link.Adjvex].Dist == Infinity {
				t[ga[v].Link.Adjvex].Dist = t[v].Dist + 1
				t[ga[v].Link.Adjvex].Path = v
				q.Add(ga[v].Link.Adjvex)
			}
			ga[v].Link = ga[v].Link.Next
		}
	}
	// 打印簿记表
	for i, item := range t {
		fmt.Printf("V%d %v %d V%d\n", i+1, item.Known, item.Dist, item.Path+1)
	}
}

// Dijkstra算法
// 解决“单源最短路径”问题的一般方法，Dijkstra算法是贪婪算法。
// 普通实现
func (ga GraphList) Dijkstra(s int) {
	// 声明并初始化簿记表
	var t [N]book_keeepTable
	for i := 0; i < N; i++ {
		if s == i {
			continue
		}
		t[i].Dist = Infinity
		t[i].Path = -1
	}

	for {
		// 通遍历的方式查找距离最小位置顶点，整个算法查找最小值花费的时间为O(|V|^2)；
		// todo:
		// 可以使用DeleteMin操作来实现，因为一旦未知的最小值顶点被找到，那么它就不再是未知的，以后不再考虑。
		// 二叉堆DeleteMin操作的时间复杂度为O(logN)
		v := -1
		min := Infinity + 1
		for i, item := range t { //获取未知顶点中距离最小的
			if !item.Known && item.Dist < min {
				min = item.Dist
				v = i
			}
		}
		if v == -1 { //全部顶点已知
			break
		}
		if t[v].Path == -1 {
			break
		}

		t[v].Known = true
		//遍历该顶点的全部邻接顶点
		for ga[v].Link != nil {
			if !t[ga[v].Link.Adjvex].Known { //该邻接点是未知的
				// 对于更新Dw的操作，在图是稠密的情况下，简单实现基本上是最优，但是当图是稀疏的时，这种算法就太慢了；
				// 对于稀疏图，距离需要存储在优先队列中，有两种方式可以实现：
				// 1.把更新处理为DecreaseKey操作，此时查找最小值的时间复杂度为O(log|V|)，即为执行更新的时间，它相当于执行那些DecreaseKey操作的时间。
				// 	由于优先队列不是很好地支持Find操作，因此Di的每个值在优先队列的位置需要保留，并当Di在优先队列中改变时更新。
				// 	如果优先队列时二叉堆实现，这将较难实现；如果使用配对堆(pairing heap)会比较容易实现。
				// 2.在执行Dw更新时，把w（本轮查询到的最小距离未知顶点的邻接点）和新的Dw插入到优先队列中。
				// 	这样优先队列中的每个顶点就可能有多于一个的代表。当DeleteMin操作把最小的顶点从优先队列中删除时，必须检查以确定它不是已知的。
				// 	这样将获取最小距离未知顶点的操作变成一个循环，它执行DeleteMin直到一个未知的顶点合并为止。
				// 	这种方法在软件的观点上是优越的，且编程易实现，但是队列大小可能达到与边数相同大小。
				// 	而且该算法不影响渐进时间界，任是一个O(|E|log|V|)算法，不过空间需求却增加了。
				// 	不仅如此该方法需要|E|次而不仅仅是|V|次DeleteMin，所以它在实践中肯更慢。
				// 如果使用不同的数据结构，那么Dijkstra算法可能会有更好的时间界。比如“斐波那契堆”，使用这种数据结构时间复杂度为O(|E|+|V|log|V|)，
				// 不过需要更高的系统开销。因此，不能确定在实践中使用“斐波那契堆”比使用“二叉堆”的Dijkstra算法更好。
				if t[v].Dist+ga[v].Link.Weight < t[ga[v].Link.Adjvex].Dist {
					t[ga[v].Link.Adjvex].Dist = t[v].Dist + ga[v].Link.Weight
					t[ga[v].Link.Adjvex].Path = v
				}
			}
			ga[v].Link = ga[v].Link.Next
		}
	}
	// 打印簿记表
	for i, item := range t {
		fmt.Printf("V%d %v %d V%d\n", i+1, item.Known, item.Dist, item.Path+1)
	}
}

// 具有负边值的图
func (ga GraphList) WeightedNegative(s int) {
	var counts [N]int //顶点出队次数计数器
	// 声明并初始化簿记表
	var t [N]book_keeepTable
	for i := 0; i < N; i++ {
		if s == i {
			continue
		}
		t[i].Dist = Infinity
	}
	q := queue.New()
	q.Add(s)          //起始顶点入队
	t[s].Known = true //标记顶点入队
	for q.Length() != 0 {
		v := q.Remove().(int)
		t[v].Known = false //标记顶点出队
		counts[v]++        //顶点出队计数器加一
		//针对存在负值圈的情况，使算法能够终止运行
		if counts[v] <= N {
			//遍历邻接顶点，需保留原邻接表
			// 所以不能直接使用ga[v].Link = ga[v].Link.Next，这会导致顶点二次进入队列时，邻接表为空
			link := ga[v].Link
			for link != nil {
				if t[v].Dist+link.Weight < t[link.Adjvex].Dist {
					t[link.Adjvex].Dist = t[v].Dist + link.Weight
					t[link.Adjvex].Path = v
					//如果邻接顶点不在队列中，则将其入队
					if !t[link.Adjvex].Known {
						q.Add(link.Adjvex)
						t[link.Adjvex].Known = true
					}
				}
				link = link.Next
			}
		} else {
			fmt.Println("图中存在负值圈")
			return
		}
	}
	// 打印簿记表
	for i, item := range t {
		fmt.Printf("V%d %v %d V%d\n", i+1, item.Known, item.Dist, item.Path+1)
	}
}

// Prim算法
func (ga GraphList) Prim(s int) {
	// 声明并初始化簿记表
	var t [N]book_keeepTable
	for i := 0; i < N; i++ {
		if s == i {
			continue
		}
		t[i].Dist = Infinity
		t[i].Path = -1
	}

	for {
		v := -1
		min := Infinity + 1
		for i, item := range t { //获取未知顶点中距离最小的
			if !item.Known && item.Dist < min {
				min = item.Dist
				v = i
			}
		}
		if v == -1 { //全部顶点已知
			break
		}
		if t[v].Path == -1 {
			break
		}

		t[v].Known = true
		//遍历该顶点的全部邻接顶点
		for ga[v].Link != nil {
			if !t[ga[v].Link.Adjvex].Known {
				// 对于每个与v邻接的未知顶点w，Dw=min(Dw, Cw,v)
				if ga[v].Link.Weight < t[ga[v].Link.Adjvex].Dist {
					t[ga[v].Link.Adjvex].Dist = ga[v].Link.Weight
					t[ga[v].Link.Adjvex].Path = v
				}
			}
			ga[v].Link = ga[v].Link.Next
		}
	}
	// 打印簿记表
	for i, item := range t {
		fmt.Printf("V%d %v %d V%d\n", i+1, item.Known, item.Dist, item.Path+1)
	}
}

// Kruskal算法
type edge struct { //定义边
	From   int
	To     int
	Weight int
}
type edgeHeap []edge

func (h edgeHeap) Len() int           { return len(h) }
func (h edgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h edgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *edgeHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(edge))
}

func (h *edgeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 将图的边信息读入堆数组
func (h *edgeHeap) ReadGraphIntoHeapArry(ga GraphList) {
	for i, vex := range ga {
		for vex.Link != nil {
			e := edge{
				i,               //边的起始点
				vex.Link.Adjvex, //边的终点
				vex.Link.Weight, //权
			}
			*h = append(*h, e)
			vex.Link = vex.Link.Next
		}
	}
}

func (ga GraphList) Kruskal() {
	var edgesAccepted int
	var s set.DisjSet
	var h edgeHeap

	s = set.Initialize(N, -1)   //初始化顶点的不相交集
	h.ReadGraphIntoHeapArry(ga) //将边读入优先队列
	heap.Init(&h)               //初始化边的优先队列

	for edgesAccepted < N-1 {
		e := heap.Pop(&h).(edge)
		u := s.Find(set.SetType(e.From))
		v := s.Find(set.SetType(e.To))
		if u != v { //当前两个顶点是非连通的
			edgesAccepted++
			s.UnionBySize(u, v)
			fmt.Printf("v%d->v%d	weight: %d\n", e.From, e.To, e.Weight)
		}
	}
}

// 关键路径算法
func (ga GraphList) CriticalPath() {
	var i, j, k, count int
	var front = -1
	var rear = -1 //顺序队列的首尾指针置初值-1
	var tpord, vl, ve [N]int
	var e, l [E]int
	// for i = 0; i < N; i++ {
	// 	ve[i] = 0 //各事件的最早发生事件均置0
	// }
	for i = 0; i < N; i++ { //扫描顶点表，将入度为0的顶点入队
		if ga[i].In == 0 {
			rear++
			tpord[rear] = i
		}
	}

	// 按拓扑排序顺序求出各顶点的最早发生时间Ve
	for front != rear { //队列非空
		front++
		j = tpord[front] //Vj+1出队，即删去Vj+1
		count++          //对出队顶点个数计数
		p := ga[j].Link
		for p != nil {
			k = p.Adjvex //k是边<Vj+1, Vk+1>终点Vk+1的下标
			ga[k].In--   //Vk+1入度减1
			if ve[j]+p.Weight > ve[k] {
				ve[k] = ve[j] + p.Weight //修改Ve[k]
			}
			if ga[k].In == 0 {
				rear++
				tpord[rear] = k //新的入度为0的顶点Vk+1入队
			}
			p = p.Next
		}
	}
	if count < N { //存在回路，算法终止
		fmt.Println("该AOE网络存在回路")
		return
	}

	// 按拓扑序列的逆序求出各顶点事件的最迟发生时间Vl
	for i = 0; i < N; i++ {
		vl[i] = ve[N-1] //初始化vl数组，赋值为ve的最后一个事件（汇点）的最早发生时间
	}
	for i = N - 2; i >= 0; i-- { //按拓扑序列的逆序取顶点
		j = tpord[i]
		p := ga[j].Link
		for p != nil {
			k = p.Adjvex //为<Vj+1, Vk+1>的终点Vk+1的下标
			if vl[k]-p.Weight < vl[j] {
				vl[j] = vl[k] - p.Weight
			}
			p = p.Next
		}
	}

	i = -1                  //边计数器置初值
	for j = 0; j < N; j++ { //扫描顶点Vj+1的邻接表，计算<Vj+1, Vk+1>所代表的活动ai+1的e[i]和l[i]
		p := ga[j].Link
		for p != nil {
			k = p.Adjvex
			i++
			e[i] = ve[j]
			l[i] = vl[k] - p.Weight
			fmt.Printf("V%s\tV%s\t%d\t%d\t%d\t", ga[j].Vertex, ga[k].Vertex, e[i], l[i], l[i]-e[i])
			if l[i] == e[i] {
				fmt.Println("关键活动")
			}
			fmt.Println()
			p = p.Next
		}
	}
}
