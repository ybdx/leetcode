package leetcode

func buildTree1(inorder []int, postorder []int) *TreeNode {
	index := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		index[inorder[i]] = i
	}
	return helper(inorder, postorder, 0, len(inorder), 0, len(postorder), index)
}

func helper(inorder, postorder []int, inStart, inEnd, posStart, posEnd int, index map[int]int) *TreeNode {
	if posStart == posEnd {
		return nil
	}
	root := &TreeNode{Val: postorder[posEnd-1]}
	cur := index[postorder[posEnd-1]]
	root.Left = helper(inorder, postorder, inStart, cur, posStart, posStart + cur - inStart, index)
	root.Right = helper(inorder, postorder, cur+1, inEnd, posStart+cur-inStart, posEnd-1, index)
	return root
}
