package poi

import (
	"fmt"
	"math"
	"testing"
)


func TestP(t *testing.T) {
	var qlist = []float64{0.1, 0.2, 0.3, 0.4}
	var zlist = []int {1,2,3,4,5,6,7,8}
	for _, q := range qlist {
		for _, z := range zlist {
			p := attackerSuccessProbability(q, z)
			fmt.Printf("q:%.2f, z:%d:  %.5f\n", q, z, p)
		}
	}
}

func attackerSuccessProbability(q float64, z int) float64{
	lambda := float64(z)*q/(1.0-q)
	var sum float64 
	for k:=0; k<=z; k++ {
		kp := 1
		for i:=k; i>0; i-- {
			kp *= i
		}
		poission := math.Exp(-lambda)
		poission /= float64(kp) 
		poission *= math.Pow(lambda,float64(k))
		sum += poission*(1 - math.Pow((q/(1-q)), float64(z-k)))
	}
	return 1-sum
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
