package stack
import "testing"
// 155
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

// 20
func Valid(s string) bool {
	counter := make([]rune,0)	
	for _, byt := range s {
		switch byt {
		case '{','(','[':
			counter=append(counter, byt)
		case ')',']','}':
			if len(counter) == 0 {
				return false 
			}
			if byt == ')' && counter[len(counter)-1] == '(' {
				counter = counter[:len(counter)-1] 
				continue
			} 			
			if byt == ']' && counter[len(counter)-1] == '[' {
				counter = counter[:len(counter)-1] 
				continue 
			} 
			if byt == '}' && counter[len(counter)-1] == '{' {
				counter = counter[:len(counter)-1] 
				continue
			} 		
			return false 
		}
	}
	if len(counter) == 0 {
		return true
	}
	return false
}

func TestValid(t *testing.T) {
	s := "([)]"
	res := Valid(s)
	if res != false {
		t.Errorf("expected to be %v, but got %v", false, res)
	}
}

// 739
func temperature(temp []int) []int {
	stk := make([]int, 0)
	res := make([]int, len(temp))
	for i:=0; i<len(temp); i++ {
		for ;len(stk)>0; {
			last := stk[len(stk)-1]
			if temp[i] <= temp[last] {
				break
			}
			res[last] = i - last
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}	
	return res
}
