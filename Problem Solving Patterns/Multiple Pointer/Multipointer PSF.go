/*
	Multiple Pointer Problem solving patterns
	=========================================
	Creating pointers or values that correspond to an index or position and move towards the
	beginning, end or middle based on a certain codition.

	Very efficient for solving problems with minimal space complexity as well

*/

/*
	An Example:
	===========

	Write a function called sumZero()  which accepts a sorted array of integers. the function should
	find the first pair where the sum is ). return an array that includes both value that
	sum to zero or undefined if a pair does not exist.

	Samples:
	========
	sumZero([-3, -2, -1, 0, 1, 2, 3]) ---> [-3, 3]
	sumZero([-2, 0, 1, 3]) ---> undefined
	sumZero([1, 2, 3]) ---> undefined
*/

package main

import "fmt"

func main() {
	testValue1 := []int{-5, -1, 0, 1, 2, 3, 4}
	fmt.Println(refectorSolution(testValue1))
}

/*
	=> Simple solution
	=> time complexity : O(n^2)
	=> space complexity: o(1)
	=> will run two loop on the array and search the match.
*/
func easySolution(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i]+s[j] == 0 {
				return []int{s[i], s[j]}
			}
		}
	}
	return nil
}

/*
	=> This is MULTIPOINTER SOLUTION
	================================

	=> Now let solve the problem in other way
	=> Let's say we have two pointer one on the start of the array and other
	=> one on the end of the array. so now we will run a loop and we gonna
	=> make come closer two pointer and in a mean time we will check all the item
	=> and if we fine a match that sum is 0 then we will return the match in a array.
	=> Here time complextiy will be O(n)
	=> space complexity: o(1)
*/
func refectorSolution(s []int) []int {
	startIndex := 0
	lastIndex := len(s) - 1

	for lastIndex > startIndex {
		sum := s[lastIndex] + s[startIndex]
		if sum == 0 {
			return []int{s[startIndex], s[lastIndex]}
		} else if sum > 0 {
			lastIndex--
		} else {
			startIndex++
		}
	}

	return nil
}
