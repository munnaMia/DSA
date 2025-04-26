package basicmath

/*
Problem Statement: Given an integer N return the reverse of the given number.
TIME COMPLEXITY : O(log n)
SPACE COMPLEXITY : O(1)
*/
func ReverseDigit(n int) int {
	result := 0

	for n > 0 {
		result = result*10 + n%10
		n /= 10
	}
	return result
}
