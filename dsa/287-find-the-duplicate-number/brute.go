package main

import "fmt"

func main() {
	var testcases [][]int
	testcases = [][]int{{1, 2, 3, 4, 5, 5}, {1, 2, 3, 4, 5, 6, 1}}
	for _, testcase := range testcases {
		fmt.Println(findDuplicate(testcase))
	}
}

func findDuplicate(nums []int) int {
	n := len(nums)
	arr := make([]bool, n)

	for _, num := range nums {
		if arr[num-1] == true {
			return num
		} else {
			arr[num-1] = true
		}
	}
	return -1
}

//Brute hashmap solution
//TC O(N)
//SC O(N)
