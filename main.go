package main

import (
	"fmt"

	hashing "github.com/munnaMia/DSA/PH05_Hashing"
)

func main() {
	for key, value := range hashing.CharCounter("test case with another", []rune{'t', 'a', 'e', 'j', 'w', 'i', 's'}) {
		fmt.Printf("%s => %d \n", string(key), value)
	}
}
