package leetcode

import (
	"math"
)

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	result := make([]Interval, 0)
	flag := true
	for i := 0; i < len(intervals); i++ {
		if intervals[i].End < newInterval.Start {
			result = append(result, intervals[i])
		} else if intervals[i].Start <= newInterval.End {
			newInterval.Start = int(math.Min(float64(intervals[i].Start), float64(newInterval.Start)))
			newInterval.End = int(math.Max(float64(intervals[i].End), float64(newInterval.End)))
		} else {
			if flag {
				result = append(result, newInterval)
				flag = false
			}
			result = append(result, intervals[i])
		}
	}
	return result
}
