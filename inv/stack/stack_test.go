package inv
import "testing"

type MinStack struct {
	stk []int
	min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stk = append(this.stk, val)
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if this.stk[len(this.stk)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.stk = this.stk[:len(this.stk)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func TestMinstack(t *testing.T) {
	stk := Constructor()
	stk.Push(-2)
	stk.Push(0)
	stk.Push(-3)
	res := stk.GetMin()
	if res != -3 {
		t.Errorf("expected to be %d not %d", -3, res)
	}
	stk.Pop()
	res = stk.Top()
	if res != 0 {
		t.Errorf("expected to be %d not %d", 0, res)
	}
	res = stk.GetMin()
	if res != -2 {
		t.Errorf("expected to be %d not %d", -2, res)
	}
}


