package leetcode

import (
	"testing"
	"fmt"
	"reflect"
	"strings"
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
	fmt.Println(wordPattern("abba", "dog cat cat do"))
	fmt.Println(-1 * (1 << 32))
	fmt.Println(addStrings("1", "1"))
	root := buildTree1([]int{9,3,15,20,7}, []int{9,15,7,20,3})
	fmt.Println(root)
	fmt.Println(generateMatrix(3))
	fmt.Println(spiralOrder([][]int{{6, 9, 7}}))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(strings.Split("/a/./b/../../c/", "/")[0:2])
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(make([]int, 0, 10))
}



