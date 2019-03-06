package quicksort

import (
	"testing"
	"fmt"
)

var tcs = []struct {
	nums []int
}{
	//{[]int{1, 5, 2, 1, 6, 4}},
	{[]int{1, 3, 2, 2, 3, 1, 2, 3, 1, 2, 3, 2, 2, 1, 1, 2, 1, 1, 2, 3, 3, 2, 2, 2, 2, 3, 3, 3}},
	//{[]int{1, 1, 2, 1, 2, 2, 1}},
	//{[]int{1, 3, 2, 2, 3, 1}},
	//{[]int{4, 3}},
	// 可以有多个 testcase
}

func TestQuick2Sort(t *testing.T) {
	//ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		QuickSortDG(tc.nums)
		fmt.Println("result:", tc.nums)
		//ast.True(check(tc.nums), "输入:%v", tc)
	}
}
func TestQuickSort(t *testing.T) {
	//ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		QuickSortFZ(tc.nums)
		fmt.Println(tc.nums)
		//ast.True(check(tc.nums), "输入:%v", tc)
	}
}

// 自己手写一遍，检查测试
func TestQS(t *testing.T) {
	// ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		qs(tc.nums)
		fmt.Println("result:", tc.nums)
		answer := append(make([]int, 0), tc.nums...)
		QuickSortFZ(answer)
		fmt.Println("expected:", answer)
		//ast.Equal(answer,tc.nums, "输入:%v", tc)
	}
}

func Benchmark_qs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			qs2(tc.nums)
		}
	}
}
// 重复元素排序测试
func TestQS2(t *testing.T) {
	// ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		qs2(tc.nums)
		fmt.Println("result:", tc.nums)
		answer := append(make([]int, 0), tc.nums...)
		QuickSortFZ(answer)
		fmt.Println("expected:", answer)
		//ast.Equal(answer,tc.nums, "输入:%v", tc)
	}
}