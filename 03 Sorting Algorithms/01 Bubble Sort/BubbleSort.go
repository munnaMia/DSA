package BubbleSort

import "fmt"

/*
	Bubble sort is an algorithm that sort a data structure like ARRAY.
		- The word 'Bubble' comes from how this algorithm works,
		- it makes the highest/lowest values 'bubble up'.

		How it Works:
			1. Traverse the Array: Begin by going through each value in the array, one at a time.
			2. Compare Adjacent Values: For each value, compare it with the next one in the sequence.
			3. Swap if Needed: If the current value is greater than the next, swap them so the larger value moves to the end of the list.
			4. Repeat the Process: Continue this process for the entire array as many times as there are values, ensuring that each pass gradually sorts the array.
*/

func BubbleSort() {
	// Array to sort.
	test := []int{12, 10, 4, 19, 20, 5, 3, 1, 11, 10}

	length := len(test) // Length of the Array.

	for i := range length - 1 {
		for j := range length - i - 1 {
			if test[j] > test[j+1] {
				test[j], test[j+1] = test[j+1], test[j]
			}
		}
	}
	fmt.Println(test)

}

/*
	Output:
	--> [12 10 4 19 20 5 3 1 11 10] // UNSORTED
	--> [10 4 12 19 5 3 1 11 10 20]
	--> [4 10 12 5 3 1 11 10 19 20]
	--> [4 10 5 3 1 11 10 12 19 20]
	--> [4 5 3 1 10 10 11 12 19 20]
	--> [4 3 1 5 10 10 11 12 19 20]
	--> [3 1 4 5 10 10 11 12 19 20]
	--> [1 3 4 5 10 10 11 12 19 20]
	--> [1 3 4 5 10 10 11 12 19 20] // SORTED
*/
