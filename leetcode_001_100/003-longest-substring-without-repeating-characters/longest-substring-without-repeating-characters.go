package problem003

// 我的提交 4ms
func MylengthOfLongestSubstring(s string) int {
	var i, ret int  // i子串头，ret最长子串
	m := [256]int{} // 假定输入的字符串只有ASCII字符
	for i := range m {
		m[i] = -1
	}
	for j := 0; j < len(s); j++ {
		// 如果当前字符s[j]出现过，那么字典m中就会有对应的位置v
		// 此时，计算一下当前子串的长度len(s[i:j])，并对比已经找到的长度ret
		// 子串头重置，有可能上一个s[j]出现在子串之前
		// 如当前子串在之前出现过，因此i取和v+1比较后较大的值

		v := m[s[j]]
		// 当前字符s[j]出现有两种情况
		// 1.重复字符v在i和j之间，说明上次循环过程中，计算的字符串肯定比这次长
		// 因此此次不做ret更新，直接更新i
		if v >= i {
			i = v + 1
		} else {
			// 2.如果v出现在i之前
			// 说明，本次有可能是最长子串，判断一下进行更新
			if j-i+1 > ret {
				ret = j - i + 1
			}
		}
		// 无论当前字符s[j]是否重复出现，都更新出现位置
		m[s[j]] = j
	}
	return ret
}
