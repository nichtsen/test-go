package inv 
import (
	"testing"
	"fmt"
	"math"
)
func TestSqrt(t *testing.T) {
	sqrt := func(a int) int {
		l := 1; r := a-1
		for ;l<=r; {
			// lower bound
			mid := (l+r)/2
			tmp := a/mid
			if tmp > mid {
				l = mid + 1
			} else if tmp < mid {
				r = mid -1
			} else {
				return mid
			}
		}
		return r
	}
	a := 9 
	res := sqrt(a)
	fmt.Println(res)
}

func lower(a []int, target int) int {
	l:=0; r:=len(a)-1;
	for;l<r;{
		mid := (l+r)/2
		if a[mid] >= target {
			r = mid 
		} else {
			l = mid+1
		}
	}
	return l
}

func upper(a []int, target int) int {
	l:=0; r:=len(a)-1;
	for;l<r;{
		mid := int(math.Ceil(float64(l+r)/2))
		if a[mid] > target {
			r = mid -1
		} else {
			l = mid
		}
	}
	return r
}

func searchRange(a []int, target int) []int {
	lo := lower(a, target)
	up := upper(a, target)
	if len(a) == 0 || a[lo] != target {
            return []int{-1,-1}
        }
	return []int{lo, up}
}

func TestSearchRange(t *testing.T){
	a := []int{5,7,7,8,8,10}
	target := 8
	res := searchRange(a, target)
	fmt.Println(res)
}

func TestRotate(t *testing.T) {
	search := func(a []int, target int) bool {
		lo := 0; hi :=len(a)-1;
		for;lo<hi; {
			mid := (lo+hi)/2
			if a[mid] == target {
				return true
			} else if a[hi] == a[mid] {
				hi--
				// replication
			} else if a[hi] > a[mid] {
				// right area is incremental
				if target > a[mid] && target <= a[hi] {
					lo = mid+1
				} else {
					hi = mid-1
				}
			} else {
				// left side area is incremental
				if target < a[mid] && target >= a[lo] {
					hi = mid-1
				} else {
					lo = mid+1
				}
			}
		}
		return false
	}
	a := []int{2,5,6,0,0,1,2}
	target := 0
	res := search(a, target)
	fmt.Println(res)
}
