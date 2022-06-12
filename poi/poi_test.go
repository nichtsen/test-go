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
