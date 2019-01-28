package problem179

import (
	"sort"
	"strings"
	"strconv"
	"fmt"
)

func join(a, b int) int {
	if b == 0 {
		return a * 10
	}
	tmp := b
	for tmp > 0 {
		a *= 10
		tmp /= 10
	}

	return a + b
}

type IntSort []int

func (a IntSort) Len() int      { return len(a) }
func (a IntSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a IntSort) Less(i, j int) bool {
	j1 := join(a[i], a[j])
	j2 := join(a[j], a[i])
	return j1 < j2
}

func largestNumber(nums []int) string {
	sort.Sort(sort.Reverse(IntSort(nums)))
	fmt.Println(nums)
	b := strings.Builder{}
	skipZero := true
	for _, num := range nums {
		if num != 0 {
			skipZero = false
		} else if skipZero {
			continue
		}

		b.WriteString(strconv.Itoa(num))
	}

	result := b.String()
	if len(result) == 0 {
		return "0"
	}

	return result
}

//import (
//	"strconv"
//	"strings"
//	"fmt"
//	"sort"
//)
//
//type StringSort []string
//
//func (a StringSort) Len() int      { return len(a) }
//func (a StringSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
//func (a StringSort) Less(i, j int) bool {
//	return a[i]+a[j] < a[j]+a[i]
//}
//
//func largestNumber(nums []int) string {
//
//	// 转换
//	var strs []string
//	for _, nums := range nums {
//		strs = append(strs, strconv.Itoa(nums))
//	}
//	strs = StringSort(strs)
//	fmt.Println("排序前：",strs)
//	// 排序
//	quickSort(strs)
//	// 特殊情况
//	if strs[0] == "0" {
//		return "0"
//	}
//	fmt.Println("排序后：",strs)
//	// 汇总
//	flatStr := strings.Builder{}
//	for _, s := range strs {
//		flatStr.WriteString(s)
//	}
//
//	return flatStr.String()
//}
//
//// 快速排序
//func quickSort(strs StringSort) {
//	fmt.Println("000", strs)
//	if len(strs) <= 1 {
//		fmt.Println("000", len(strs))
//		return
//	}
//
//	pivot, i := 0, 1
//	head, tail := 1, strs.Len()-1
//
//	for head < tail {
//		if strs.Less(i, pivot) { // 修改条件
//			strs.Swap(i, tail)
//			tail--
//		} else {
//			strs.Swap(i, head)
//			head ++
//			i++
//		}
//		fmt.Println("111", strs, ", pivot=", strs[pivot])
//
//	}
//	fmt.Println("333", strs)
//	// strs[head] = strs[pivot]
//	strs.Swap(pivot, head)
//	fmt.Println("444", strs)
//	quickSort(strs[:head])
//	quickSort(strs[head+1:])
//	fmt.Println("222", strs)
//}
//
//// 自由排序规则
//func sort(a, b string) bool {
//	if a+b < b+a { // 从大到小排
//		return true
//	}
//	return false
//}
