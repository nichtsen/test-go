package sequence

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	f := New(Fib(10))
	f.Map(func(a int) int {
		return a + 1
	})
	f.Print()
	f.Filt(isEven)
	f.Print()
	res := f.Accumulate(1, multiply)
	fmt.Println(res)
}
