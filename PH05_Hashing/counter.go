package hashing

// Problem statement:
// Given an array of integers: [1, 2, 1, 3, 2]
// and we are given some queries: [1, 3, 4, 2, 10].
// For each query, we need to find out how many times the number
// appears in the array. For example, if the query is 1 our answer
// would be 2, and if the query is 4 the answer will be 0.
func Counter(arr, query []int) map[int]int {
	result := make(map[int]int)

	// initialize the key's from query array and assign the with 0
	for i := range query {
		result[query[i]] = 0
	}

	// counting there repeations
	for i := range arr {
		result[arr[i]]++
	}
	return result
}

//Problem statement
// Given the string: “abcdabefc” we are given some queries: [a, c, z].
// For each query, we need to find out how many times the character appears
// in the string. For example, if the query is ‘a’ our answer would be 2,
// and if the query is ‘z’ the answer will be 0.
func CharCounter(str string, query []rune) map[rune]int {
	result := make(map[rune]int)
	runes := []rune(str)

	// initialize the key's from query array and assign the with 0
	for i := range query {
		result[query[i]] = 0
	}

	// counting there repeations
	for i := range runes {
		result[runes[i]]++
	}
	return result
}

/*  hashing.CharCounter("test case with another", []rune{'t','a', 'e','j','w','i','s'})
Output:
e => 3
j => 0
  => 3
c => 1
h => 2
o => 1
w => 1
i => 1
s => 2
n => 1
r => 1
t => 4
a => 2
*/

func CounterHighandLow(arr []int) (int, int) {
	result := make(map[int]int)

	// counting there repeations
	for i := range arr {
		result[arr[i]]++
	}

	var maxKey, maxRep, minKey, minRep int

	minRep = len(arr) - 1

	for key, value := range result {
		if maxRep < value {
			maxRep = value
			maxKey = key
		}
		if minRep > value {
			minRep = value
			minKey = key

		}
	}
	return maxKey, minKey
}
