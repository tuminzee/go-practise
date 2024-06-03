package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))

}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var sf []float64
	var m map[float64][]int
	m = make(map[float64][]int)
	n := len(arr)

	for i := 0; i < n; i += 1 {
		for j := i + 1; j < n; j += 1 {
			fr := float64(arr[i]) / float64(arr[j])
			sf = append(sf, fr)
			m[fr] = []int{arr[i], arr[j]}
		}
	}

	sort.Float64Slice(sf).Sort()

	val := sf[k-1]

	return m[val]

}
