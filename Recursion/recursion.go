package recursion

import "fmt"

// Print your Name N times using recursion
func Name( name string, count int) {
	if count ==  0 {
		return
	}
	fmt.Println(name)
	Name(name, count-1)
}

// Print from 1 to N using Recursion
func Number(i, count int) {
	if i>count {
		return
	}
	fmt.Println(i)
	Number(i+1, count)
}

// Print from N to 1 using Recursion
func RevNumber(count int){
	if count==0 {
		return
	}
	fmt.Println(count)
	RevNumber(count-1)
}
