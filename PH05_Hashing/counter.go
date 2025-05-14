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
	for i:=range query {
		result[query[i]]=0
	}

	// counting there repeations
	for i:=range arr {
		result[arr[i]]++
	}
	return result
}