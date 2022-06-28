package inv
import (
	"testing"
	"fmt"
)

func bipartitie(graph [][]int) bool{
	color := make([]int, len(graph))
	colo := &color

	for idx, edges := range graph {
		// find new start point
		if len(edges) > 0 && color[idx] == 0 {
	                // nodes is a queue
			nodes := []int{ idx }
			if !biSearch(graph, nodes, colo) {
				return false
			}
		}
	}
	return true
}

func biSearch(graph [][]int, nodes []int, color *[]int) bool {
	for ;len(nodes) > 0; {
			node := nodes[0]
			nodes = nodes[1:]
			// start point
			if (*color)[node] == 0 {
				(*color)[node] = 1
			}
		
			adj := graph[node]
			for _, anode := range adj {
				if (*color)[anode] == 0 {
					nodes = append(nodes, anode)
					(*color)[anode] = (*color)[node]%2+1
				} else {
					if (*color)[anode] == (*color)[node] {
						return false
					}
				}
			}
		}
	return true
}



func TestBi(t *testing.T) {
	graph := [][]int{{1,3}, {0, 2}, {1,3}, {0, 2}}
	res := bipartitie(graph)
	fmt.Println(res)
}
