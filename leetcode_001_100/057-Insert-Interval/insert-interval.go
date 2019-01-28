package problem057

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	// 前提是输入的区间是有序
	var res []Interval
	var (
		break_idx = -1
	)
	// 合并器
	merge := Interval{
		newInterval.Start,
		newInterval.End,
	}
	// 遍历
	for idx, i := range intervals {
		if merge.Start > i.End {
			res = append(res, i)
			continue
		}
		if i.Start > merge.End {
			break_idx = idx
			break
		}
		// 求交叉
		merge.Start = min(merge.Start, i.Start)
		merge.End = max(merge.End, i.End)
	}
	res = append(res, merge)
	if break_idx >= 0 {
		res = append(res, intervals[break_idx:]...)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
