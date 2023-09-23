package graph

import "fmt"

const N = 7 //图的顶点数
const E = 8 //图的边数

type vextype string  //顶点的数据类型
type adjtype float32 //边权值的数据类型
type GraphMatrix struct {
	Vexs [N]vextype    //顶点数组
	Arcs [N][N]adjtype //邻接矩阵
}

//生成图的邻接矩阵
func CreatAdjMatrix(digraph bool) GraphMatrix {
	var ga GraphMatrix
	var i, j, k int
	var w adjtype
	for i = 0; i < N; i++ {
		fmt.Printf("输入顶点%d:\n", i+1)
		fmt.Scanf("%s\n", &ga.Vexs[i]) //读入顶点信息，建立顶点表
	}
	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			ga.Arcs[i][j] = 0 //邻接矩阵初始化
		}
	}
	for k = 0; k < E; k++ {
		fmt.Println("输入边(Vi,Vj)及其上的权值:")
		fmt.Scanf("%d%d%f\n", &i, &j, &w) //读入边(Vi,Vj)上的权值w
		ga.Arcs[i-1][j-1] = w             //写入邻接矩阵
		if !digraph {                     //是否创建有向图
			ga.Arcs[j-1][i-1] = w
		}
	}
	return ga
}

//打印邻接矩阵
func (ga GraphMatrix) Print() {
	fmt.Println("\n/*******************************/")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%f\t", ga.Arcs[i][j])
		}
		fmt.Println()
	}
	fmt.Println("/*******************************/")
}

// 拓扑排序
// 对有N个顶点的有向图邻接矩阵求拓扑排序
// 邻接矩阵中，顶点的入度与该顶点相对应的列上的1的个数相等；出度与该顶点相对应的行上的1的个数相等；
//
// 算法步骤：
// 1.取1作为第一个新序号。（通过该序号最后是否等于顶点个数，即所有的列是否都得到了新序号，来判断图中是否存在环）
// 2.找到一个没有得到新序号的全0矩阵列，没有则停止。
// 	这时若矩阵中所有的列都得到了新序号，则拓扑排序完成；否则说明该有向图中有环存在。
// 3.把新序号赋给找到的列并将该列对应的顶点输出。
// 4.将找到的列所对应的行置全0。（相当于删除找到的入度为0的顶点及其边）
// 5.新序号加1，重复2~5
//
func (ga GraphMatrix) TopoSort() {
	var m, i, j, k, t, v int
	var D [N]int
	for i = 0; i < N; i++ {
		D[i] = 0
	}
	v = 1 //新序号变量置1
	for k = 0; k < N; k++ {
		for j = 0; j < N; j++ { //寻找全0列
			if D[j] == 0 {
				t = 1
				for i = 0; i < N; i++ {
					if ga.Arcs[i][j] == 1 {
						t = 0
						break
					} //第j列上有1，跳出循环
				}
				if t == 1 {
					m = j
					break
				} //找到第j列为全0列
			}
		}
		if j != N {
			D[m] = v                       //将新序号赋给找到的列
			fmt.Printf("%s\t", ga.Vexs[m]) //将排序结果输出
			for i = 0; i < N; i++ {
				ga.Arcs[m][i] = 0 //将找到的列的相应行置全0
			}
			v++ //新序号加1
		} else {
			break
		}
	}
	if v < N {
		fmt.Println("\n图中存在环。")
	}
}
