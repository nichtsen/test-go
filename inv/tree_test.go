package inv

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepthLoop(max *int, depth int, root *TreeNode) {
	if root == nil {
		return 
	}
	depth++
	if depth > *max {
		*max = depth
	}
	maxDepthLoop(max, depth, root.Left)
	maxDepthLoop(max, depth, root.Right)
}

func maxDepth(root *TreeNode) int {
	tmp := 0
	max := &tmp
	depth := 0
	maxDepthLoop(max, depth, root)
	return *max
}

func BalancedTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := BalancedTree(root.Left)
	right := BalancedTree(root.Right)
	if left == -1 || right == -1 {
		return -1 
	}
	abs := left - right
	if abs < 0 { abs = -abs}
	if abs > 1 {
		return -1 
	}

	switch {
	case left > right:
		return left + 1
	default:
		return right+1
	}
}

func DiameterTree(root *TreeNode, dia *int) int {
	if root == nil {
		return 0
	}
	left := DiameterTree(root.Left, dia)
	right := DiameterTree(root.Right, dia)
	if left + right  > *dia {
		*dia = left + right
	}
	if left > right {
		return left + 1
	}else {
		return right + 1
	}
}

func PathSumTree(root *TreeNode, sum int) int {
	if root == nil {return 0}
	return pathSumTree(root, sum) + PathSumTree(root.Left, sum) + PathSumTree(root.Right, sum)
}

func pathSumTree(root *TreeNode, sum int) int {
	if root == nil { return 0}
	if root.Val == sum {
		return 1 + pathSumTree(root.Left, sum-root.Val) + pathSumTree(root.Right, sum-root.Val)

	}
	return pathSumTree(root.Left, sum-root.Val) + pathSumTree(root.Right, sum-root.Val)
}

func symmetricTree(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return symmetricTree(left.Left, right.Right) && symmetricTree(left.Right, right.Left)
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	tmp := make([]*TreeNode,0, 1)
	res := &tmp
	del := make(map[int]bool, len(to_delete))
	for _, val := range to_delete {
		del[val] = true
	}
	flg := false
	if !del[root.Val] {
		tmp = append(tmp, root)
	} else {
		flg = true
	}
	delNode(root, del, flg, res)
	return *res
}
func delNode(root *TreeNode, todel map[int]bool, del bool, res *[]*TreeNode)  {
	if root == nil {
		return 
	}
	left := root.Left
	right := root.Right

	var lf, rf bool
	if left !=nil {
		if todel[left.Val] {
			root.Left = nil
			lf = true
		} else if del  {
		// root is removed from its left child but child is not removed
		*res = append(*res, left)
		}
	}

	if right !=nil {
		if todel[right.Val] {
			root.Right = nil
			rf = true
		} else if del  {
			*res = append(*res, right)
		}
	}

	delNode(left, todel, lf, res)
	delNode(right, todel, rf,  res)
}

func average(level []*TreeNode, res *[]float64) {
	nextlevel := make([]*TreeNode, 0)
	sum := 0
	count := 0
	for _, node := range level {
		sum += node.Val
		count ++
		if node.Right != nil { nextlevel = append( nextlevel, node.Right) }
		if node.Left != nil { nextlevel = append( nextlevel, node.Left) }
	}
	if count > 0 {
		*res = append(*res, float64(sum)/float64(count))
	}
	if len(nextlevel) > 0 {
		average(nextlevel, res)
	}
}
func buildTree(preorder, inorder []int) *TreeNode {
	m := make(map[int]int)
	for idx, val := range inorder {
		m[val] = idx 
	}
	return BuildTree(preorder, inorder, m, 0)
}

// hash: inorder.val -> idx
func BuildTree(preorder, inorder []int, hash map[int]int, offset int) *TreeNode {
	if len(preorder) == 0 { return nil }
	root := &TreeNode{
		Val: preorder[0],
	}
	lefts := inorder[:hash[root.Val]-offset]
	rights := inorder[hash[root.Val]+1-offset:]
	plefts := preorder[1:1+len(lefts)]
	prights := preorder[1+len(rights):]
	if len(lefts) > 0 {
		root.Left = BuildTree(plefts, lefts, hash,offset)
	}
	if len(rights) > 0 { 
		root.Right = BuildTree(prights, rights, hash, hash[root.Val]+1)
	}
	return root
}

func iterationPreorder(root *TreeNode) []int {
	if root == nil { return []int{} }
	res := make([]int, 0, 2)
	stk := make([]*TreeNode,0)
	stk = append(stk, root)
	for ;len(stk) > 0; {
		node := stk[len(stk)-1]
		res = append(res, node.Val)
		stk = stk[:len(stk)-1]
		if node.Right != nil {
			stk = append(stk, node.Right)
		}
		if node.Left != nil {
			stk = append(stk, node.Left)
		}
	}
	return res
}	
func InorderRecover(root *TreeNode) {
	node1 := new(*TreeNode)
	node2 := new(*TreeNode)
	inorderRocver(root, node1, node2)
	if *node2 != nil {
		tmp := (*node1).Val
		(*node1).Val = (*node2).Val
		(*node2).Val = tmp

	}
}

func inorderRocver(root *TreeNode, node1, node2 **TreeNode) {
	if root.Left != nil {
		inorderRocver(root.Left, node1, node2)
	}
	if *node1 != nil {
		// a misorder detected
		if root.Val < (*node1).Val {
			*node2 = root
		} else {
			// still try to find first mismatch so just update node1
			if *node2 == nil {
				*node1 = root
			}
			// transfer current ndoe1 
		}
	} else {
		// transfer root as node1
		*node1 = root 
	}
	if root.Right != nil {
		inorderRocver(root.Right, node1, node2)
	}
}

func trimBst(root *TreeNode, l,r int) *TreeNode {
	if root == nil {
		return nil 
	}
	switch {
	case root.Val > r:
		return trimBst(root.Left, l, r)
	case root.Val < l:
		return trimBst(root.Right, l, r)
	}
	root.Left = trimBst(root.Left, l, r)
	root.Right = trimBst(root.Right, l, r)
	return root
}
