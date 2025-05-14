package basicmath

import "fmt"

/*
Problem Statement: Given an integer N, return true if it is a palindrome else return
false.

A palindrome is a number that reads the same backward as forward. For example, 121,
1331, and 4554 are palindromes because they remain the same when their digits are
reversed.

Time Complexity: O(logN+1)
Space Complexity: O(1)
*/
func Palindrome(n int) int {
	result := 0
	for n > 0 {
		result = n % 10 + (result * 10)
		fmt.Println(result)
		n/=10
	}
	return result
}
