package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 2, 3, 4, 9, 6}))
}

func heightChecker(heights []int) int {
	var heightsCopy = make([]int, len(heights))
	copy(heightsCopy, heights)
	sort.Ints(heights)
	cnt := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != heightsCopy[i] {
			cnt += 1
		}
	}

	return cnt
}
