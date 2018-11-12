package leetcode

import (
	"testing"
	"fmt"
	"reflect"
)

func TestSort(t *testing.T) {
	//nums := []int{5,54,52,67,68,5,52,17,93,53}
	//fmt.Println(largestNumber(nums))
	fmt.Println(reverseVowels("leotcede"))
	fmt.Println(countPrimes(10))
}

func TestPalindRomeLinkedList(t *testing.T) {
	head := &ListNode{Val:1}
	fmt.Println(isPalindrome(head))
	var head1 *ListNode
	fmt.Println(isPalindrome(head1))
	s := "hjhj"
	fmt.Println(reflect.TypeOf(s[0]))
	fmt.Println(isIsomorphic("title", "paper"))
}



