package patterns

import (
	"fmt"
)

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
	for i := n - 1; i >= 0; i-- {
		for j := n - i - 1; j > 0; j-- {
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

func SideWallTringle(n int) {
	for i := 1; i <= 2*n-1; i++ {
		star := i
		if i > n {
			star = 2*n - i
		}
		for j := 1; j <= star; j++ {
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
// ****
// ***
// **
// *

func AlpaTringle(n int) {
	for i := 1; i <= n; i++ {
		for j := 'A'; j < 'A'+rune(i); j++ {
			fmt.Print(string(j))
		}
		fmt.Println()
	}
}

// Output
// A
// AB
// ABC
// ABCD
// ABCDE

func ReverseAlpaTringle(n int) {
	for i := n; i > 0; i-- {
		for j := 'A'; j < 'A'+rune(i); j++ {
			fmt.Print(string(j))
		}
		fmt.Println()
	}
}

// Output :
// ABCDE
// ABCD
// ABC
// AB
// A

func AlpaTringleRow(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string('A' + rune(i)))
		}
		fmt.Println()
	}
}

// Output
// A
// BB
// CCC
// DDDD
// EEEEE

func AlpaPiramid(n int) {
	for i := 0; i < n; i++ {
		for j := n - i - 1; j > 0; j-- {
			fmt.Print(" ")
		}
		ch := 'A'
		breakPoint := int((i*2 + 1) / 2)
		for k := 1; k <= i*2+1; k++ {
			fmt.Print(string(ch))
			if breakPoint >= k {
				ch++
			} else {
				ch--
			}
		}
		fmt.Println()
	}
}

// Output
//     A
//    ABA
//   ABCBA
//  ABCDCBA
// ABCDEDCBA

func BackWardAplaTringle(n int) {
	for i := 0; i < n; i++ {
		for ch := 'E' - rune(i); ch <= 'E'; ch++ {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}

// output
// E
// DE
// CDE
// BCDE
// ABCDE
