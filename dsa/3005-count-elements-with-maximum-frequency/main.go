package main

func main() {
	var nums []int
	nums = []int{1, 2, 3, 4, 5}

	maxFrequencyElements(nums)
}

func maxFrequencyElements(nums []int) int {
	freqMap := make(map[int]int)
	var maxV int
	var count int

	for _, num := range nums {
		freqMap[num] += 1
		maxV = max(maxV, freqMap[num])
	}

	for _, value := range freqMap {
		if value == maxV {
			count += maxV
		}
	}

	return count
}
