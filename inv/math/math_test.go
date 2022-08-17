package main

import "math"


import (
	"testing"
	"fmt"
)

func countQueue(n int) []int {
	a := make([]int, 100)
	for i:=0;i<len(a);i++ {
		a[i] = i+1
	}
	cnt := 0
	for alive:=100;alive>=n; {
		for i:=0;i<len(a);i++ {
			if a[i] > 0 {
				cnt++
				if cnt % n == 0 {
					a[i] = -1
					alive--
				}
			}
		}

	}
	res := make([]int, 0)
	for i:=0;i<len(a);i++ {
		if a[i] > 0 {
			res = append(res, a[i])
		}
	}
	return res
}

// there are 100 people who sit side by side and form a loop
// they count numbers incrementally, every time a person who counts a number that is multiple of a certain N (e.g 3), 
// that perpson is discarded from loop
// game terminates when there are less than N people still alive
func TestCountQueue(t *testing.T) {
	res := countQueue(3)
	if fmt.Sprint(res) != fmt.Sprint([]int{58, 91}) {
		t.Errorf("got %v",  res)
	}
}

//greedy
func leastCars(cars []int) int {
	cnt:=0
	res := 0 
	for idx, stat := range cars {
		if stat == 1 {
			cnt++
		}
		if stat == 0 || idx == len(cars) -1 {
			switch cnt {
				case 1,2,3:
					res++
				default:
					res += cnt/3 
					if cnt % 3 > 0 {
						res++
					}
			}
			cnt = 0
		}
	}
	return res
}


func TestLeastCars(t *testing.T) {
	cars := []int{1,1,0, 0, 1,1,1,0, 1}
	res := leastCars(cars)
	if res != 3 {
		t.Errorf("expected to %v, got %v", 3, res)
	}
}

func triSqrt(a float64) float64 {
	x := a
	for ;math.Abs(x*x*x - a) > 0.1; {
		x = (x + a/(x*x))/2 
		fmt.Println(x)
	}
	return x
}

func TestTriSqrt(t *testing.T) {
	a := 19.9
	_ = triSqrt(a)
	
}





			
