package problem053

import "fmt"

//
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划  时间复杂度：O(n)，空间复杂度：O(1)
//func maxSubArray(nums []int) int {
//	dp := make([]int, len(nums))
//	var mx int
//	for i := range nums {
//		if i == 0 {
//			dp[0] = nums[0]
//			mx = dp[0]
//			continue
//		}
//		tmp := dp[i-1]
//		if dp[i-1] < 0 {
//			tmp = 0  // sum+n < n，那就还不如直接从 n 开始统计。
//		}
//		dp[i] = nums[i] + tmp
//		mx = max(mx, dp[i])
//	}
//
//	return mx
//}

// 贪心算法  时间复杂度：O(n)，空间复杂度：O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	currMax := nums[0]
	maxSoFar := nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		currMax = max(currMax+curr, curr)
		maxSoFar = max(maxSoFar, currMax)
	}
	return maxSoFar
}

// 带坐标

func maxSubArrayWithLocation(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var start, end int
	currMax := nums[0]
	maxSoFar := nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if currMax+curr < curr {
			currMax = curr
			start = i
		}

		currMax = currMax + curr

		if maxSoFar < currMax {
			maxSoFar = currMax
			end = i
		}
	}
	fmt.Println("start", start, "end", end)
	return maxSoFar
}
