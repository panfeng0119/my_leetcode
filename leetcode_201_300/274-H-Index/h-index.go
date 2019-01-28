package problem274

//
//import "sort"
//
//func hIndex(citations []int) int {
//	sort.Ints(citations)
//
//	res := 0
//	for i := len(citations) - 1; i>= 0; i-- {
//		if citations[i] >= len(citations) - i && (i == 0 || citations[i - 1] <= len(citations) - 1 )  {
//			if len(citations) - i > res {
//				res = len(citations) - i
//			}
//		}
//	}
//
//	return res
//}

import (
	"sort"
)

func hIndex(citations []int) int {
	// 排序
	sort.Ints(citations)
	var res int
	for i := len(citations) - 1; i >= 0; i -- {
		if citations[i] >= len(citations)-i {
			res = len(citations)-i
		}
	}
	return res
}
