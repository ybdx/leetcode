package leetcode

import "sort"

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})
	cur := intervals[0]
	res := make([]Interval, 0)
	for i:=1; i<len(intervals); i++ {
		check := intervals[i]
		if cur.End >= check.Start {
			cur.End = max(cur.End, check.End)
		} else {
			res = append(res, cur)
			cur = check
		}
	}
	res = append(res, cur)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
