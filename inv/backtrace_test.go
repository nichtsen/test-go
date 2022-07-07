package inv

import "testing"

func ksubset(a []int, k int) bool {
	sum := 0
	for _, val := range a {
		sum+=val
	}
	if sum % k != 0 {
		return false 
	}
	target := sum/k
	// idx -> volumn
	buk := make([]int, k)
	return btk(a, buk, 0, target)
}
// backtrace the k subset puzzle
func btk(a, b []int, idx, target int) bool {
	if idx == len(a) { return true }
	for i:=0; i<len(b); i++ {
		if b[i] + a[idx] > target { continue }
		if i>0 && b[i] == b[i-1] { continue }
		b[i] += a[idx]
		if btk(a, b, idx+1, target) { return true }
		// withdraw
		b[i] -= a[idx]
		if idx == 0 { break }
	}
	return false
}

func maxSubset(a []int) int {
	for i:=len(a); i>1; i-- {
		if ksubset(a, i) { return i }
	}
	return -1 
}


// 698
func TestKsubset(t *testing.T) {
	a := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	res := ksubset(a, k)
	if !res {
		t.Errorf("got %v", res)
	}
}
// a variant of 698
// find max number of subset we can divide
// if it is less than 2, return -1
func TestMaxSubset(t *testing.T) {
	a:=[]int{3,6,6,3}
	res := maxSubset(a)
	if res != 3 {
		t.Errorf("expected to be %v ,got %v", 3, res)
	}
	a=[]int{3,5}
	res = maxSubset(a)
	if res != -1 {
		t.Errorf("expected to be %v ,got %v", -1, res)
	}
}



