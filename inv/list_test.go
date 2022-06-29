package inv

type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	prev := head
	cur := head.Next
	prev.Next = nil
	for ;cur!=nil;{
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}
func swapList(head *ListNode) *ListNode {
	if head == nil { return nil}
	if head.Next == nil { return head}

	prev := head
	s := prev.Next
	prev.Next = s.Next
	s.Next = prev
	head = s

	for ;prev.Next!=nil && prev.Next.Next !=nil; {
		old := prev
		prev = prev.Next
		s = prev.Next
		prev.Next = s.Next
		s.Next = prev
		old.Next = s
	}
	return head
}

func intersectionList(headA, headB *ListNode) *ListNode {
	l1 := headA; l2 := headB
	for ;l1!=l2; {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}
	return l1 
}

func parlindromList(head *ListNode) bool {
	l1 := head               
	s := make([]int, 0, 8)
	// collect values into slice
	for ;l1!=nil; {    
		s = append(s, l1.Val)
		l1 = l1.Next
	}
	// upper bound
	hi := len(s)/2
	var lo int
	if len(s) % 2 == 0 {
		lo = hi - 1
	} else {
		lo = hi
	}

	for ;lo>=0; {
		if s[lo] != s[hi] {
			return false
		}
		lo--
		hi++
	}
	
	return true
}
func parlindromList02(head *ListNode) bool {
	l1 := head               
	length := 0
	for ;l1!=nil; {
		length++
		l1 = l1.Next
	}
	// upper bound
	mid := length/2
	
	l1 = head 
	for ;mid>0;mid-- {
		l1 = l1.Next 
	}
	if length % 2 != 0 && length > 2 {
		l1 = l1.Next
	}
	// revere half list
	cur := l1.Next
	l1.Next = nil
	for ;cur!=nil; {
		tmp := cur.Next
		cur.Next = l1
		l1 = cur
		cur = tmp
	}
	// check 
	for ;l1!=nil; {
		if l1.Val != head.Val {
			return false
		}
		l1 = l1.Next
		head = head.Next
	}
	return true
}
