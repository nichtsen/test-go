package inv

import (
	"fmt"
	"sort"
	"testing"
)
// 455
func children(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi := 0; si := 0
	for ;gi < len(g)&& si < len(s);{
		if g[gi] <=  s[si] {
			gi++
		}
		si++
	}
	return gi
}

func TestChidren(t *testing.T) {
	g := []int{2,1}
	s := []int{1,2,3}
	res := children(g,s)
	fmt.Println(res)
}

func candy(rating []int) int {
	stat := make([]int, len(rating))
	for i:=0;i<len(rating);i++ {
		stat[i]= 1
	}
	// left to right
	for i:=1;i<len(rating);i++ {
		if rating[i] > rating[i-1]  {
			stat[i]=stat[i-1]+1
		}
	}
	// right to left
	for i:=len(rating)-1;i>0;i-- {
		if rating[i-1] > rating[i] && stat[i-1] <= stat[i] {
			stat[i-1] = stat[i]+1
		}
	}
	sum :=0
	for i:=0;i<len(stat);i++ {
		sum += stat[i]
	}
	return sum
}

type Interval [][]int

func (in Interval) Len() int { return len(in) }
func (in Interval) Less(i, j int) bool { return in[i][1] < in[j][1] }
func (in Interval) Swap(i, j int) { tmp := in[i]; in[i] = in[j]; in[j] = tmp;} 
// 435: non-overlapping interval
func interval(in [][]int) int {
	sort.Slice(in, func(i,j int)bool {return in[i][1] < in[j][1]})
	cnt := 0 
	last := in[0][1] 
	for i:=1;i<len(in); i++ {
		// overlap
		if in[i][0] < last {
			cnt++
		}else {
			last = in[i][1]
		}
	}
	return cnt 
}


