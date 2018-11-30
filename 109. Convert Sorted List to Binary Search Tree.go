package leetcode

var head1 *ListNode
func sortedListToBST(head *ListNode) *TreeNode {
	if nil == head {
		return &TreeNode{}
	}
	size := getSize(head)
	head1 = head
	return convertListToBST(0, size-1)
}

func convertListToBST(l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := l + (r-l)/2
	left := convertListToBST(l, mid-1)
	node := &TreeNode{Val: head1.Val}
	node.Left = left
	head1 = head1.Next
	right := convertListToBST(mid+1, r)
	node.Right = right
	return node
}

func getSize(head *ListNode) int {
	temp := head
	res := 0
	for temp != nil {
		res++
		temp = temp.Next
	}
	return res
}
