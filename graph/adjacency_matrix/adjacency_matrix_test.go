package graph

import (
	"testing"
)

func TestTopoSort(t *testing.T) {
	ga := GraphMatrix{}
	ga.Vexs = [N]vextype{"1", "2", "3", "4", "5", "6", "7"}
	ga.Arcs = [N][N]adjtype{
		{0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
		{},
		{},
	}
	ga.Print()
	ga.TopoSort()
}
