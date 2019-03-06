package problem0020

func isValid(s string) bool {
	size := len(s)
	// 栈
	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++ {
		switch s[i] {
		case '(':
			stack[top] = ')'
			top++
		case '[':
			stack[top] = ']'
			top++
		case '{':
			stack[top] = '}'
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == s[i] {
				top--
			} else {
				return false
			}
		}
	}

	return top == 0
}

// 原始参考
//func isValid(s string) bool {
//	size := len(s)
//	// 栈
//	stack := make([]byte, size)
//	top := 0
//
//	for i := 0; i < size; i++ {
//		c := s[i]
//		switch c {
//		case '(':
//			stack[top] = c + 1 // '('+1 is ')'
//			top++
//		case '[', '{':
//			stack[top] = c + 2
//			top++
//		case ')', ']', '}':
//			if top > 0 && stack[top-1] == c {
//				top--
//			} else {
//				return false
//			}
//		}
//	}
//
//	return top == 0
//}
