package main

import "fmt"

func main() {
 CountDown(10)
 fmt.Println(sumRange(10))
}

// Problem 01 : Make a simple Counter programe
func CountDown(n int) {
	if n <= 0 {
		fmt.Println("Done!")
		return
	}
	fmt.Println(n)
	n--
	CountDown(n)
}


// Problem 02 : Sum of range
func sumRange(n int)int {
	if n== 1 {return 1}
	return n + sumRange(n-1)
}


// Problem 03 : Factorial
func Factorial(n int) int{
	if n==1 {return 1}
	return n*Factorial(n-1)
}

