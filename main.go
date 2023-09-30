package main

import graph "dataStructAlg/graph/adjacency_list"

// graph "dataStructAlg/graph/adjacency_list"

func main() {
	ga := graph.CreatAdjList(true)
	ga.CriticalPath()
}
