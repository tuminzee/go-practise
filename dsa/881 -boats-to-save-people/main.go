package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	n := len(people) - 1
	l, r := 0, n

	for l <= r {
		if people[l]+people[r] <= limit {
			l += 1
		}
		r -= 1
	}
	return n - r
}

func main() {
	people := []int{3, 2, 2, 1}
	limit := 3
	fmt.Print(numRescueBoats(people, limit))

}

// TC - O(N)
// SC - O(1)
