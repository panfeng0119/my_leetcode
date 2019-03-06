package shortest_lenth

import "fmt"

/**
	一串首尾相连的珠子(m 个)，有 N 种颜色(N<=10)，
	设计一个算法，取出其中一段，要求包含所有 N 中颜色，并使长度最短。
	并分析时间复杂度与空间复杂度。
*/
func shortestLength(s []int, N int) []int {

	s2 := append(s, s...)
	// 结果
	start, end := 0, len(s2)-1

	fmt.Println("s2:", s2)
	//遍历所有可能的起始点
	for i := 0; i < len(s); i++ {
		// 从不重复开始
		for i+1 < len(s2) && s2[i+1] == s2[i] {
			i++
			continue
		}

		// 出现颜色统计；数组用来判断，存放出现的颜色
		cnt := 0
		h := make([]int, N+1)
		for j := range h {
			h[j] = 0
		}
		curLength := 0
		//找到在当前起点下找到所有颜色的结尾
		tmpEnd := i
		for ; cnt < N && tmpEnd < len(s2); tmpEnd ++ {
			h[s2[tmpEnd]]++
			// 用来检查颜色条件是否达到
			if h[s2[tmpEnd]] == 0 {
				cnt++
			}
			// 当前最长
			curLength = tmpEnd - i + 1
			//fmt.Println("h:", h, "cnt:", cnt,"N:",N, "tmpEnd:", tmpEnd, "s2:", s2)
		}

		fmt.Println("当前位置：",i,"，curLength:", curLength, "，h:", h, "，cnt:", cnt)

		if cnt == N && curLength < end - start +1 {
			start, end = i, tmpEnd
		}
		//fmt.Println("minimumLength:", minimumLength, "start:", start, "end:", end)
	}

	return s2[start:end]
}
