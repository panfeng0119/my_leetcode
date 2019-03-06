package problem0643

//
func findMaxAverage(nums []int, k int) float64 {
	var max, tmp int
	for i := 0; i < k; i++ {
		tmp += nums[i]
	}
	max = tmp
	for i := k; i < len(nums); i++ {
		// 根据 i 获取 temp
		tmp = tmp - nums[i-k] + nums[i]
		if max < tmp {
			max = tmp
		}
	}

	return float64(max) / float64(k)
}
