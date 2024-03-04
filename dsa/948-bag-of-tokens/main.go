package main

import (
	"fmt"
	"sort"
)

func main() {
	var tokens []int
	var power int

	tokens = []int{100, 400, 300, 200}
	power = 200

	fmt.Println(bagOfTokensScore(tokens, power))
}

func bagOfTokensScore(tokens []int, power int) int {
	n := len(tokens)
	sort.Ints(tokens)
	fmt.Println(tokens)
	maxScore := 0
	score := 0
	l, r := 0, n-1

	for l <= r {
		if power >= tokens[l] {
			score += 1
			maxScore = max(maxScore, score)
			power -= tokens[l]
			l += 1
		} else if score >= 1 {
			score -= 1
			maxScore = max(maxScore, score)
			power += tokens[r]
			r -= 1
		} else if score == 0 {
			break
		}
	}

	return maxScore
}
