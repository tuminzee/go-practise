package main

import "fmt"

func main() {
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*))"))
	fmt.Println(checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"))
}

type memo struct {
	i   int
	bal int
}

func checkValidString(s string) bool {
	// fmt.Print(s, " ")
	m := make(map[memo]bool)
	var dfs func(int, int) bool

	dfs = func(i, bal int) bool {
		if i == len(s) {
			return bal == 0
		}
		if bal < 0 {
			return false
		}

		c := s[i]

		if _, ok := m[memo{i, bal}]; !ok {
			if c == '(' {
				m[memo{i, bal}] = dfs(i+1, bal+1)

			} else if c == ')' {
				m[memo{i, bal}] = dfs(i+1, bal-1)
			} else {
				m[memo{i, bal}] = dfs(i+1, bal+1) || dfs(i+1, bal-1) || dfs(i+1, bal)
			}

		}
		return m[memo{i, bal}]
	}

	return dfs(0, 0)
}
