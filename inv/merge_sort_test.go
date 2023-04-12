package inv

import "fmt"

func MergeSort(a []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	MergeSort(a, l, m)
	MergeSort(a, m+1, r)
	merge(a, l, m, r)
}
func merge(a []int, l, m, r int) {
	ls := make([]int, m-l+1)
	rs := make([]int, r-m)
	copy(ls, a[l:m+1])
	copy(rs, a[m+1:r+1])
	i, j := 0, 0
	for k := l; k <= r; k++ {
		if i == len(ls) {
			a[k] = rs[j]
			j++
			continue
		}
		if j == len(rs) {
			a[k] = ls[i]
			i++
			continue
		}
		if ls[i] < rs[j] {
			a[k] = ls[i]
			i++
		} else {
			a[k] = rs[j]
			j++
		}
	}
}

func ExampleMergeSort() {
	a := []int{7, 3, 0, 6, 1}
	MergeSort(a, 0, len(a)-1)
	fmt.Println(a)
	//  Output:
	//  [0 1 3 6 7]
}
