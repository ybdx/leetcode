package leetcode

func guessNumber(n, pick int) int {
	low, high := 1, n
	for low <= high {
		mid :=low + (high - low) / 2
		res := guess(mid, pick)
		if 0 == res {
			return mid
		} else if res < 0 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// num > pick  return -1
// num < pick return 1
// num == pick return 0
func guess(num, pick int) int {
	if num == pick {
		return 0
	} else if num > pick {
		return -1
	} else {
		return 1
	}
}