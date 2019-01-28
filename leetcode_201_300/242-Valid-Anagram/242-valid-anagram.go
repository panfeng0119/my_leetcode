package _42_Valid_Anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int,26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		m[t[i]-'a']--
		if m[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

//
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	m1 := make(map[byte]int)
//	m2 := make(map[byte]int)
//	for i := 0; i < len(s); i++ {
//		m1[s[i]]++
//	}
//	for i := 0; i < len(t); i++ {
//		m2[t[i]]++
//	}
//	fmt.Println(m1, m2)
//	for k, v := range m2 {
//		if m1[k] != v{
//			return false
//		}
//	}
//	return true
//}
