package recursion

import (
	"fmt"
)

// Print your Name N times using recursion
func Name(name string, count int) {
	if count == 0 {
		return
	}
	fmt.Println(name)
	Name(name, count-1)
}

// Print from 1 to N using Recursion
func Number(i, count int) {
	if i > count {
		return
	}
	fmt.Println(i)
	Number(i+1, count)
}

// Print from N to 1 using Recursion
func RevNumber(count int) {
	if count == 0 {
		return
	}
	fmt.Println(count)
	RevNumber(count - 1)
}

// Problem statement: Given a number ‘N’, find out the sum of the first N natural numbers.
// Parameterized way
// Time complexity : O(N)
// Space complexity : O(N)
func Sum1(n, sum int) {
	if n < 1 {
		fmt.Println(sum)
		return
	}

	Sum1(n-1, sum+n)
}

// Functional way
// Time complexity : O(N)
// Space complexity : O(N)
func Sum2(n int) int {
	if n == 0 {
		return 0
	}
	return n + Sum2(n-1)
}

/*
func(3) ---> 3 + func(2) ---> 3 + 2 + func(1) ---> 3 + 2 + 1 + func(0) ---> 3 + 2 + 1 + 0
*/

// Problem Statement: Given a number X,  print its factorial.
// Time complexity : O(N)
// Space complexity : O(N)
func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Problem Statement: You are given an array. The task is to reverse the array and print it.
// Using an extra array.
func ReverseArrayUsingLoop(arr []int) []int {
	length := len(arr)
	var result = make([]int, length)
	for i := length - 1; i >= 0; i-- {
		result[length-i-1] = arr[i]
	}
	return result
}

// Problem Statement: You are given an array. The task is to reverse the array and print it.
// Two pointer solution
func ReverseArrayUsingTwoPointers(arr []int) []int {
	p1 := 0
	p2 := len(arr) - 1
	for p1 < p2 {
		temp := arr[p1]
		arr[p1] = arr[p2]
		arr[p2] = temp
		p1++
		p2--

	}
	return arr
}

// Problem Statement: You are given an array. The task is to reverse the array and print it.
// Recursive method
func ReverseArrayUsingRecursion(arr []int, p1, p2 int) []int {
	if p1 < p2 {
		temp := arr[p1]
		arr[p1] = arr[p2]
		arr[p2] = temp
		ReverseArrayUsingRecursion(arr, p1+1, p2-1)
	}
	return arr
}
