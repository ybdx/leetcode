package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	if n < 1 {
		return []string{}
	}
	result := make([]string, 0)
	for i := 0; i < n; i++ {
		div3 := (i+1)%3 == 0
		div5 := (i+1)%5 == 0
		if div3 && div5 {
			result = append(result, "FizzBuzz")
		} else if div3 {
			result = append(result, "Fizz")
		} else if div5 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i+1))
		}
	}
	return result
}
