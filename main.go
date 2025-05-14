package main

import (
	"fmt"

	hashing "github.com/munnaMia/DSA/PH05_Hashing"
)

func main() {
	fmt.Println(hashing.Counter([]int{3, 2, 3, 4, 2, 1,43, 1, 1, 1}, []int{3, 2, 1, 4, 5}))
}
