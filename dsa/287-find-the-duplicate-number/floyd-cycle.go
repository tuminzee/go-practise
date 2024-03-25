package main

import "fmt"

func main() {
	var testcases [][]int
	testcases = [][]int{{1, 2, 3, 4, 5, 5}, {1, 2, 3, 4, 5, 6, 1}}
	for _, testcase := range testcases {
		fmt.Println(findDuplicateFloyd(testcase))
	}
}

func findDuplicateFloyd(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	fast = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

//TC - O(N)
//SC - O(1)
//floyd cycle detection algo
