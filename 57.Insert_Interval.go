package leetcode

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	result := make([]Interval, 0)
	flag := true
	for i := 0; i < len(intervals); i++ {
		if intervals[i].End < newInterval.Start {
			result = append(result, intervals[i])
		} else if intervals[i].Start <= newInterval.End {
			newInterval.Start = min(intervals[i].Start, newInterval.Start)
			newInterval.End = max(intervals[i].End, newInterval.End)
		} else {
			if flag {
				result = append(result, newInterval)
				flag = false
			}
			result = append(result, intervals[i])
		}
	}
	if flag {
		result = append(result, newInterval)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
