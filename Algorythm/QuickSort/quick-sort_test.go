package quicksort

import (
	"testing"
	"fmt"
)

var tcs = []struct {
	nums []int
}{
	{[]int{1, 5, 2, 1, 6, 4}},
	//{[]int{1, 3, 2, 2, 3, 1}},
	//{[]int{1, 1, 2, 1, 2, 2, 1}},
	//{[]int{1, 3, 2, 2, 3, 1}},
	//{[]int{1, 3, 2, 2, 3, 1}},
	// 可以有多个 testcase
}


func TestQuick2Sort(t *testing.T) {
	//ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		Quick2Sort(tc.nums)
		fmt.Println(tc.nums)
		//ast.True(check(tc.nums), "输入:%v", tc)
	}
}
func TestQuickSort(t *testing.T) {
	//ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		QuickSort(tc.nums)
		fmt.Println(tc.nums)
		//ast.True(check(tc.nums), "输入:%v", tc)
	}
}
