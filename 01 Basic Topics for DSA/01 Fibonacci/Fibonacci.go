package fibonacci

import "fmt"

var Count int = 2 // Count data for fiboRecu() function.

func test() {
	// ---- Loop test ----
	fiboLoop(20)

	// ---- Recursion test ----
	fmt.Println(0)
	fmt.Println(1)
	fiboRecu(0, 1)

	// ---- Recursion test Formula version ----
	fmt.Println(fiboRecuFormula(19))
}

// ==== Implementation using LOOP.
func fiboLoop(n int) {
	n1 := 0
	n2 := 1
	fmt.Println(n1)
	fmt.Println(n2)
	for i := 0; i <= n-2; i++ {
		next := n1 + n2
		fmt.Println(next)
		n1 = n2
		n2 = next

	}
}

// ==== Implementation using RECURSION.
func fiboRecu(n1, n2 int) {
	next := n1 + n2
	fmt.Println(next)
	if Count <= 20 {
		n1 = n2
		n2 = next
		Count++
		fiboRecu(n1, n2)
	} else {
		return
	}
}

// ==== Implementation using RECURSION (F(n) = F(n−1) + F(n−2)) using this formula
// ---- This formula uses a 0-based index. This means that to generate the 20th Fibonacci number, we must write F(19).
func fiboRecuFormula(n int) int {
	if n <= 1 {
		return n
	} else {
		return fiboRecuFormula(n-1) + fiboRecuFormula(n-2)
	}
}
