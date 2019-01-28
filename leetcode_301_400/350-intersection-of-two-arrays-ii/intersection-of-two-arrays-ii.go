package problem350

func intersect(a1, a2 []int) []int {
	m1 := distinct(a1)
	m2 := distinct(a2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	// 题目要求有几个重复，就输出几个
	for n := range m1 {
		m1[n] = min(m1[n], m2[n])
	}

	res := make([]int, 0)
	for n, size := range m1 {
		for i := 0; i < size; i++ {
			res = append(res, n)
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 去重
func distinct(arr []int) map[int]int {
	res := make(map[int]int)
	for _, v := range arr {
		res[v]++
	}
	return res
}
