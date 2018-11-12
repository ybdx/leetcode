package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if nil == head {
		return head
	}
	helper := &ListNode{}
	pre := helper
	cur := head
	for nil != cur {
		next := cur.Next
		for nil != pre.Next && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		cur = next
		pre = helper
	}
	return helper.Next
}
