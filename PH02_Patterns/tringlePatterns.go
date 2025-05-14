package patterns

import "fmt"

func StarTringle(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Output:
// *
// **
// ***
// ****
// *****

func StarTringleInverse(n int) {
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Output :
// *****
// ****
// ***
// **
// *

func NumberTringleCol(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

// Output:
// 1
// 12
// 123
// 1234
// 12345

func NumberTringleRow(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

//Output :
// 1
// 22
// 333
// 4444
// 55555

func NumberTringleColInverse(n int) {
	for i := n; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

// Output
// 12345
// 1234
// 123
// 12
// 1

func BinaryTringle(n int) {
	var start int
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			start = 0
		} else {
			start = 1
		}

		for j := 1; j <= i; j++ {
			fmt.Print(start)
			start = 1 - start
		}
		fmt.Println()
	}
}

// output :
// 1
// 01
// 101
// 0101
// 10101

func NumberTringle(n int) {
	start := 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(start, " ")
			start++
		}
		fmt.Println()
	}
}

// Output:
// 1
// 2 3
// 4 5 6
// 7 8 9 10
// 11 12 13 14 15

func DualBtSpTringle(n int) {
	space := 2 * (n - 1)
	for i := 1; i <= n; i++ {
		//number
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		//space
		for k := 1; k <= space; k++ {
			fmt.Print(" ")
		}
		//number
		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
		space -= 2
	}
}

// Output
// 1        1
// 12      21
// 123    321
// 1234  4321
// 1234554321
