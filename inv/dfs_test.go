package inv
import "testing"

var direction = []int{-1, 0, 1, 0, -1}

func maxAreaIsand(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	max :=0 
	for i:=0; i<w; i++ {
		for j:=0; j<h; j++ {
			sum := 0
			dfs(&grid, i, j, w, h, &sum)
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func dfs(grid *[][]int, x, y, w, h int, sum *int) {
	if (*grid)[y][x] == 0 {
		return 
	}
	(*sum)++
	(*grid)[y][x] = 0
	for i:=0; i<4; i++ {
		xx := x + direction[i]
		yy := y + direction[i+1]
		if xx>=0 && yy>=0  && xx < w && yy < h { 
			dfs(grid, xx, yy, w, h, sum)
		}
	}
}

func TestIsland(t *testing.T) {
	grid := [][]int { {0, 1, 0},
			  {1, 1, 1},
			  {0, 0, 0},
		  }
	_ = grid
	grid2 := [][]int{{0, 0, 1,0, 0, 0, 0, 1,0, 0, 0, 0, 0},{0, 0, 0, 0, 0, 0, 0, 1,1,1,0, 0, 0},{0, 1,1,0, 1,0, 0, 0, 0, 0, 0, 0, 0},{0, 1,0, 0, 1,1,0, 0, 1,0, 1,0, 0},{0, 1,0, 0, 1,1,0, 0, 1,1,1,0, 0},{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,0, 0},{0, 0, 0, 0, 0, 0, 0, 1,1,1,0, 0, 0},{0, 0, 0, 0, 0, 0, 0, 1,1,0, 0, 0, 0}}

	res := maxAreaIsand(grid2)
	if res != 6 {
		t.Errorf("expected to be %d, not %d", 6, res)
	}
}

		
	

