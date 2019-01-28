package problem324

import (
	"sort"
	"fmt"
)

func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	sort.Ints(nums)
	mid := len(nums)/2 + len(nums)%2

	var res []int
	for i := len(nums) - 1; i >= mid; i-- {
		fmt.Println(nums, i, i-mid)
		res = append(res, nums[i-mid])
		if i -mid >= 0 {
			res = append(res, nums[i])
		}
	}
	for i := 0; i < len(res); i++ {
		nums[i] = res[i]
	}
	fmt.Println("res", res)
	fmt.Println("res", nums)
}
