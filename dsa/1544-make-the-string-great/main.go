package main

import "fmt"

func main() {
	fmt.Println(makeGood("leeEtcode"))
	fmt.Println(makeGood("eEeEeECcccCcCCode"))

}

func makeGood(s string) string {
	// var d byte = 'a' - 'A'
	var d uint8 = 'a' - 'A'
	// var d := 'a' - 'A'

	bytes := make([]byte, 0)

	for i := range s {
		n := len(bytes)
		// fmt.Printf("%T", s[i])
		if n > 0 && (bytes[n-1]-s[i] == d || s[i]-bytes[n-1] == d) {
			bytes = bytes[:n-1]
		} else {
			bytes = append(bytes, s[i])
		}
	}
	return string(bytes)
}
