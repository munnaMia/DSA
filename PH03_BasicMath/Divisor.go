package basicmath

/*
Print all Divisors of a given Number
	Problem Statement: Given an integer N, return all divisors of N.
	A divisor of an integer N is a positive integer that divides N without 
	leaving a remainder. In other words, if N is divisible by another integer 
	without any remainder, then that integer is considered a divisor of N.
*/


// Time Complexity: O(n)
// Space Complexity: O(n)
func Divisor(n int) []int{
	result := []int{}
	for i:=1;i<=n;i++{
		if n%i ==0 {
			result= append(result, i)
		}
	}
	return result
}


func Divisor2(n int) []int{
	result := []int{}
	for i:=1;i*i<=n;i++{
		if n%i ==0 {
			result= append(result, i)
			result= append(result, n/i)
		}
	}
	return result
}