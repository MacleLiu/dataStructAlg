package graph

import (
	"fmt"

	"gopkg.in/eapache/queue.v1"
)

const N = 7            //图的顶点数
const E = 12           //图的边数
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
type GraphList []vexnode //顶点表
type pathTable struct {  //簿记表数据结构
	Known bool
	Dist  int
	Path  int
}

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
	var t [N]pathTable
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
// 使用 第三方的队列库
func (ga GraphList) Unweighted(s int) {
	// 声明并初始化簿记表
	var t [N]pathTable
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
	var t [N]pathTable
	for i := 0; i < N; i++ {
		if s == i {
			continue
		}
		t[i].Dist = Infinity
		t[i].Path = -1
	}

	for {
		// 通遍历的方式查找距离最小位置顶点，查找操作的时间复杂的为O(|V|^2)；
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
