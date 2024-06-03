package main

import "fmt"

func appendCharacters(s string, t string) int {
	i := 0

	for _, c := range s {
		if i < len(t) && c == rune(t[i]) {
			i++
		}
	}
	return len(t) - i
}

func main() {
	fmt.Println(appendCharacters("coaching", "coding"))
}
