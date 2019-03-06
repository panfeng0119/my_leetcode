package heapsort

import (
	"testing"
	"fmt"
)

var tcs = []struct {
	nums []int
}{
	{[]int{1, 5, 2, 1, 6, 4}},
	//{[]int{1, 3, 2, 2, 3, 1, 2, 3, 1, 2, 3, 2, 2, 1, 1, 2, 1, 1, 2, 3, 3, 2, 2, 2, 2, 3, 3, 3}},
	//{[]int{1, 1, 2, 1, 2, 2, 1}},
	//{[]int{1, 3, 2, 2, 3, 1}},
	{[]int{4, 3}},
	// 可以有多个 testcase
}

func TestHeapSort(t *testing.T) {
	// ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		HeapSort(tc.nums)
		fmt.Println("HeapSort result:", tc.nums)
		answer := append(make([]int, 0), tc.nums...)
		//QuickSort(answer)
		fmt.Println("expected:", answer)
		//ast.Equal(answer,tc.nums, "输入:%v", tc)
	}
}

