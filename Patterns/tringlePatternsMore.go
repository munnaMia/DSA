package patterns

import "fmt"

func Tringle(n int) {
	for i := 0; i < n; i++ {
		for j := n - i - 1; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < i*2+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Output:
//     *
//    ***
//   *****
//  *******
// *********

func TringleInverse(n int) {
	for i := n-1; i >= 0 ; i-- {
		for j := n-i-1; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := i*2 + 1; k > 0; k-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
// Output:
// *********
//  *******
//   *****
//    ***
//     *