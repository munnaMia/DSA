package basicmath

/*
	Problem:  Given an integer N, return the number of digits in N.
	 Example 1:
        Input:N = 12345

        Output:5

        Explanation:  The number 12345 has 5 digits.

    Example 2:
        Input:N = 7789

        Output: 4

        Explanation: The number 7789 has 4 digits.
*/
func CountTheDigit(number int) int {
	count := 0
	for number > 0 {
		number /= 10
		count++
	}
	return count
}

// Implement the solution...
func CountTheDigitOPTAPR(number int) int {
	return 0
}
