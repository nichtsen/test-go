package method

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	a := New(7)
	f := a.Get
	val := f()
	fmt.Println(val)
}

func TestMethodHolder(t *testing.T) {
	a := New(7)
	b := a.Holder()
	c := b.Add(5)
	fmt.Println(c)
}
