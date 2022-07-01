package inv

import (
	"testing"
	"fmt"
)

func maxprofit(price []int) int {
	dp := make([]int, len(price))
	dp[0] = price[0]
	for i:=1;i<len(price);i++ {
		switch {
		case price[i] < dp[i-1]:
			dp[i] = price[i]
		default:
			dp[i] = dp[i-1]
		}
	}
	max := 0
	profit := 0
	for i:=0;i<len(price);i++ {
		 profit = price[i] - dp[i]
		 if profit > max {
			 max = profit
		 }
	 }
	 return max
}

func TestProfit(t *testing.T) {
	p := []int {7,1,5,3,6,4}
	res := maxprofit(p)
	if res != 5 {
		t.Errorf("expected: %d ,get %d", 5, res)
	}
}

func TestPack(t *testing.T) {
	v := []int{2,4,6}
	w := []int{1,1,5}
	V := 6
	res := solvePack(v,w,V)
	fmt.Println(res)
}

func solvePack(v, w []int, V int) int {
	var dp [][]int
	N := len(v)
	dp = make([][]int, N+1)
	for i:=0; i<N+1; i++ {
		dp[i] = make([]int, V+1)
	}
	for i:=1; i<=N; i++ {
		val := v[i-1]
		wei := w[i-1]
		for j:=1; j<=V; j++ {
			if j>= wei {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wei]+ val)
			} else {
				dp[i][j] = dp[i-1][j]
				
			}
		}
	}
	fmt.Printf("%v", dp)
	return dp[N][V]
}
func max(a, b int) int {
	if a>b {
		return a
	} else {
		return b
	}
}
