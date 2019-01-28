package problem164

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 找出最大值、最小值
	var (
		mx = nums[0]
		mn = nums[0]
	)
	for _, num := range nums {
		mx = max(mx, num)
		mn = min(mn, num)
	}

	// 数字相同，直接返回
	if mx == mn {
		return 0
	}

	// 确定桶的大小 （ +1 向上取整）
	bucketField := (mx-mn)/len(nums) + 1
	// 确定桶的个数
	bucketLen := (mx-mn)/bucketField + 1

	// 创建
	buckets := make(map[int][]int)
	// 导入桶
	for _, num := range nums {
		index := (num - mn) / bucketField
		if _, ok := buckets[index]; ok {
			buckets[index] = insertSort(buckets[index], num)
		} else {
			buckets[index] = []int{num}
		}
	}

	gap := 0
	prev := 0
	// 查找最大值
	for i := 0; i < bucketLen; i ++ {
		if _, ok := buckets[i]; !ok {
			continue
		}

		gap = max(gap, buckets[i][0]-buckets[prev][len(buckets[prev])-1])
		//fmt.Println(i,buckets[i][0]-buckets[prev][len(buckets[prev])-1])
		//fmt.Println(buckets[i])
		prev = i

	}
	return gap
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

func insertSort(arr []int, new int) []int {
	var res []int
	for i, num := range arr {
		if new > num {
			continue
		}
		res = append(res, arr[:i]...)
		res = append(res, new)
		res = append(res, arr[i:]...)
		return res
	}
	res = append(arr, new)

	return []int{res[0], res[len(res)-1]}
}
