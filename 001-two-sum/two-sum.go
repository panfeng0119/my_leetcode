package problem001

func twoSum(nums []int, target int) []int {
	// 保存已遍历的数， key = target - pos
	m := make(map[int]int)
	// 开始遍历
	for i, n := range nums {
		// 检查已遍历是否有匹配项，存在则返回结果
		if pos, ok := m[n]; ok {
			return []int{pos, i}
		} else {
			m[target-n] = i
		}
	}
	return nil
}
