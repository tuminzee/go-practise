package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}

func findRelativeRanks(score []int) []string {
	mp := map[int]int{}
	ans := make([]string, len(score))
	sortedScore := make([]int, len(score))
	copy(sortedScore, score)

	sort.Sort(sort.Reverse(sort.IntSlice(sortedScore)))

	for i := range sortedScore {
		mp[sortedScore[i]] = i + 1
	}

	for i := range score {
		if mp[score[i]] == 1 {
			ans[i] = "Gold Medal"
		} else if mp[score[i]] == 2 {
			ans[i] = "Silver Medal"
		} else if mp[score[i]] == 3 {
			ans[i] = "Bronze Medal"
		} else {
			num := strconv.Itoa(mp[score[i]])
			ans[i] = num
		}
	}

	return ans
}
