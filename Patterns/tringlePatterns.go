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
		for j := 1; j <= i; j++  {
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
