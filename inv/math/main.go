package main

import (
	"fmt"
)

func factor(n int) {
	for i:=2; i* i <= n; i++ {
		if n % i == 0 {	
			fmt.Printf("%d ", i)
			factor(n/i)
			return 
		}
	}
	fmt.Printf("%d", n)
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	factor(num)
}
var direction = []int{-1, 0 , 1, 0, -1}

func main02() {
    var m, n int
    fmt.Scanf("%d %d", &m, &n)
    grid := make([][]int, m)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            fmt.Scanf("%d", &grid[i][j])
        }
    }

    if ok, res := search(&grid, 0, 0, m, n); ok {
        fmt.Println("(0,0)")
        for i:=len(res)-1;i>=0;i-- {
            fmt.Printf("(%d,%d)\n", res[i][0], res[i][1])
        }
    }

}
// find path of maze
func search(grid *[][]int,xx, yy, m, n int) (bool, [][2]int) {
    (*grid)[yy][xx] = -1
    for i:=0; i<4; i++ {
        x :=  xx + direction[i]
        y :=  yy + direction[i+1]
            // exit
            if x == n-1 && y == m-1 {
                return true,  [][2]int{{y,x}}
            }
            if x < n && x >=0 && y < m && y >=0 && (*grid)[y][x] == 0  {
                if ok, res := search(grid, x, y, m, n); ok {
                    return true, append(res, [2]int{y, x})
                }
            }
        }
    return false, nil
}

func search1(grid *[][]int, m, n int) (bool, []int) {
    stk := [][2]int{{0, 0}}
    x, y := 0, 0
    for ;len(stk)> 0; {
        x, y = stk[len(stk)-1][0],stk[len(stk)-1][1]
        stk = stk[:len(stk)-1]
        (*grid)[y][x] = -1
        for i:=0; i<4; i++ {
            xx := x + direction[i]
            yy := y + direction[i+1]
            // exit
            if xx == n-1 && yy == m-1 {
                return  true, []int{i}
            }
            if xx < n && xx >=0 && yy < m && yy >=0 && (*grid)[yy][xx] == 0  {
                stk = append(stk,[2]int{xx, yy})
                fmt.Printf("%d, %d \n ", xx, yy)
            }
        }
    }
    return false, nil
}
