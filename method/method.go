package method

type A struct {
	val int
}

// GetCall function as struct's field
type GetCall func() int

// B is the holder of A's method
type B struct {
	call GetCall
}

func New(a int) *A {
	return &A{
		val: a,
	}
}

func (a *A) Get() int {
	return a.val
}

func (a *A) Holder() *B {
	return &B{
		call: a.Get,
	}
}

func (b *B) Add(bb int) int {
	if b.call != nil {
		return b.call() + bb
	}
	return -1
}
