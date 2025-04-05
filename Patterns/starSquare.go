package patterns

import "fmt"

func StarSquare1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func StarSquare2(n int) {
	for range n {
		for range n {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Output:
// *****
// *****
// *****
// *****
// *****