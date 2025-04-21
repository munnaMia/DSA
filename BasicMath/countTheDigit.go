package basicmath

import "math"

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

// ==============================================================================
// ==============================================================================

/*
	Brute Force Approach
	Time Complexity : O(log10 N)
	Space Complexity : O(1)

	Algorithm :
	S1-> Initialise a counter to store the number of digits.
	S2-> While N is greater than 0, execute the following:
		=> Increment the counter by 1
		=> Update N by removing its last digit by performing a modulo 10 (%10) operation on it.
	S3-> After exiting the while loop, we return the counter as the number of digits.
*/
func CountTheDigit(number int) int {
	count := 0
	for number > 0 {
		number /= 10
		count++
	}
	return count
}
// ==============================================================================
// ==============================================================================

/*	Optimal Solution
	Time Complexity : O(1)
	Space Complexity : O(1)

	=>log10 N operation gives the logarithmic base 10 of which returns the power to which 10 must be raised to, to be equal to N.
	=>We add 1 to the result which accounts for the possibility that N itself is a power of 10.
	=>Financially cast the result to an integer to ensure that it is rounded down to the nearest whole number.
*/
func CountTheDigitOptimalSol(number int) int {
	return int(math.Log10(float64(number))+1)
}
