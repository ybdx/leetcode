package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := 0
	// 此部分可以采用map函数替代，用map函数来记录数字在inorder中的位置，这样只需遍历一次就够了
	for index < len(inorder) && inorder[index] != preorder[0] {
		index++
	}

	linorder := make([]int, index)
	rinorder := make([]int, len(inorder)-index-1)
	if 0 != index {
		copy(linorder, inorder[:index])
	}
	if index+1 != len(inorder) {
		copy(rinorder, inorder[index+1:])
	}
	if 1+index == len(inorder) {
		root.Left = buildTree(preorder[1:], linorder)
		root.Right = nil
	} else if 0 == index {
		root.Left = nil
		root.Right = buildTree(preorder[1:], rinorder)
	} else {
		root.Left = buildTree(preorder[1:index+1], linorder)
		root.Right = buildTree(preorder[index+1:], rinorder)
	}
	return root
}
