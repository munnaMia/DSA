package basicmath

import (
	"math"
	"strconv"
)

func Armstrong(n int) int {
	power := len(strconv.Itoa(n))

	result := 0
	for 0 < n {
		result += int(math.Pow(float64(n % 10 ), float64(power)))
		n /= 10
	}
	return result
}
