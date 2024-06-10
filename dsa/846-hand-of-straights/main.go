package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}

func isNStraightHand(hand []int, groupSize int) bool {
	fmt.Println(hand, groupSize)
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

	hc := make(map[int]int)

	for _, v := range hand {
		hc[v]++
	}

	var keys []int
	for k := range hc {
		keys = append(keys, k)
	}

	fmt.Println(keys)
	sort.Ints(keys)

	fmt.Println(keys)

	return false
}
