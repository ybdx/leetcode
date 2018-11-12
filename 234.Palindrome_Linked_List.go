package leetcode

// 解决思路：对半翻转
func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for nil != fast && nil != fast.Next {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if nil != fast { // if the length is odd
		slow = slow.Next
	}
	var pre *ListNode
	for nil != slow {
		temp := slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}
	fast = head
	slow = pre
	for nil != slow {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}
