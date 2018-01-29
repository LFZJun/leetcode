package algorithms

import "github.com/ljun20160606/go-lib/algorithms/stack"

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

// 送分题 - -, 考察栈的应用
func isValidV0(s string) bool {
	lenS := len(s)
	if lenS == 0 {
		return true
	} else if lenS&1 == 1 {
		return false
	}
	i := 0
	st := stack.NewSimpleStack()
	for i < lenS {
		switch b := s[i]; b {
		case '(', '[', '{':
			st.Append(b)
			i++
		case ')':
			if st.Len < 1 || st.PopByte() != '(' {
				return false
			}
			i++
		case ']':
			if st.Len < 1 || st.PopByte() != '[' {
				return false
			}
			i++
		case '}':
			if st.Len < 1 || st.PopByte() != '{' {
				return false
			}
			i++
		}
	}
	return st.Len == 0
}

// 看到的另一种简易栈的实现方式，很有趣
// 申请len(s)长度的[]byte skt
// 栈顶top=-1
// 入栈top+=1
// 出栈top-=1
// 取栈顶stk[top]
func isValid(s string) bool {
	stk := make([]byte, len(s))
	top := -1
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stk[top+1] = s[i]
			top++
		case ')':
			if top == -1 || stk[top] != '(' {
				return false
			}
			top--
		case ']':
			if top == -1 || stk[top] != '[' {
				return false
			}
			top--
		case '}':
			if top == -1 || stk[top] != '{' {
				return false
			}
			top--
		}
	}
	return top == -1
}
