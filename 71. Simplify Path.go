package leetcode

import "strings"

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if ".." == arr[i] && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		} else if ".." != arr[i] && "." != arr[i] && "" != arr[i] {
			stack = append(stack, arr[i])
		}
	}
	res := strings.Join(stack, "/")
	return "/" + res
}
