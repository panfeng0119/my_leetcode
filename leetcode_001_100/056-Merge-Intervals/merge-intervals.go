package problem056

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	// 按 start 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// 前提是输入的区间是有序
	var res []Interval

	// 合并器
	tmp := Interval{
		intervals[0].Start,
		intervals[0].End,
	}
	// 遍历
	for _, i := range intervals[1:] {
		if tmp.End < i.Start {
			res = append(res, tmp)
			tmp = i
			continue
		}

		// 合并
		tmp.End = max(tmp.End, i.End)
	}
	res = append(res, tmp)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
