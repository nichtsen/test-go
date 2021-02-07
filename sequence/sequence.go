package sequence

import "fmt"

type predicate func(int) bool
type Filter func(predicate, []int) []int

type Flow struct {
	seq []int
}

func New(s []int) *Flow {
	return &Flow{
		seq: s,
	}
}

func (f *Flow) Print() {
	fmt.Println(f.seq)
}

func (f *Flow) Map(op func(int) int) {
	for idx, val := range f.seq {
		f.seq[idx] = op(val)
	}
}
func (f *Flow) Filt(p predicate) {
	tmp := make([]int, 0)
	for _, val := range f.seq {
		if p(val) {
			tmp = append(tmp, val)
		}
	}
	f.seq = tmp
}
func (f *Flow) Accumulate(initial int, op func(int, int) int) int {
	for _, val := range f.seq {
		initial = op(initial, val)
	}
	return initial
}

func square(a int) int {
	return a * a
}

func multiply(a int, b int) int {
	return a * b
}

func Fib(n int) []int {
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		seq[i] = fib(i)
	}
	return seq
}

func fib(n int) int {
	var sqn = []int{0, 1}
	switch {
	case n < 0:
		return 0
	case n == 0:
		return sqn[0]
	case n == 1:
		return sqn[1]
	default:
		return fib(n-1) + fib(n-2)
	}
}
func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}
