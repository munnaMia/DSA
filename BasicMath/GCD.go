package basicmath

/*
Problem Statement: Given two integers N1 and N2, find their greatest common divisor.

The Greatest Common Divisor of any two integers is the largest number that divides both integers.

Brute Force approach
Time Complexity: O(min(n,m))
Space Complexity: O(1)
*/

func GCD(n, m int) int {
	gcd := 1
	l := min(n, m)

	for range l {
		if m%l == 0 && n%l == 0 {
			gcd = l
		}
	}
	return gcd
}

/*
Better approach
Time Complexity: O(min(n,m))
Space Complexity: O(1)
*/
func GCDbetter(n, m int) int {

	for i := min(n, m); i > 0; i-- {
		if m%i == 0 && n%i == 0 {
			return i
		}
	}
	return 1
}

/*
The Euclidean Algorithm is a method for finding the greatest common divisor of two numbers. It operates on the principle that the GCD of two numbers remains the same even if the smaller number is subtracted from the larger number.

To find the GCD of n1 and n2 where n1 > n2:

Repeatedly subtract the smaller number from the larger number until one of them becomes 0.
Once one of them becomes 0, the other number is the GCD of the original numbers.
Eg, n1 = 20, n2 = 15:

gcd(20, 15) = gcd(20-15, 15) = gcd(5, 15)

gcd(5, 15) = gcd(15-5, 5) = gcd(10, 5)

gcd(10, 5) = gcd(10-5, 5) = gcd(5, 5)

gcd(5, 5) = gcd(5-5, 5) = gcd(0, 5)

Hence, return 5 as the gcd.
*/

func GCDOptimal(n, m int) int {
	for n > 0 && m > 0 {
		if n > m {
			n = n - m
		} else {
			m = m - n
		}
	}
	if n == 0 {
		return m
	}
	return n
}
