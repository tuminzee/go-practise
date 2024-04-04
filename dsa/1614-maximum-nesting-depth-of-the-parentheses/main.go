package main

import "fmt"

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	fmt.Println(maxDepth(s))

}

func maxDepth(s string) int {
	var depth, maxDepth int
	for _, c := range s {
		if c == '(' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if c == ')' {
			depth--
		}
	}
	return maxDepth
}

// TC - O(n)
// SC - O(1)

// The solution is done without using stack. We can use stack to solve this problem. The idea is to push the opening bracket to the stack and pop the stack when we encounter closing bracket. The maximum depth of the stack will be the maximum depth of the parentheses.
// If we do not use stack we can save space. We can use two variables to keep track of the depth and maximum depth. We can iterate through the string and increment the depth when we encounter opening bracket and decrement the depth when we encounter closing bracket. We can keep track of the maximum depth and return it at the end.
