package main

import "fmt"

func main() {

	fmt.Println(minimumLength("aabccabba"))

}

func minimumLength(s string) int {
	low, high := 0, len(s)-1

	for low < high && s[low] == s[high] {
		for low+1 < high && s[low] == s[low+1] {
			low += 1
		}
		for high-1 > low && s[high] == s[high-1] {
			high -= 1
		}
		low += 1
		high -= 1
	}

	return high - low + 1

}
