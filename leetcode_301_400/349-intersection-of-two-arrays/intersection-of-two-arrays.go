package problem349

func intersection(a1, a2 []int) []int {
	m1 := distinct(a1)
	m2 := distinct(a2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	res := make([]int, 0)
	for n := range m1 {
		if _,ok := m2[n];ok{
			res = append(res, n)
		}
	}
	return res
}

// 去重
func distinct(arr []int) map[int]int {
	res := make(map[int]int)
	for _, v := range arr {
		res[v]++
	}
	return res
}
