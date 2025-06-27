package ph06sorting

/*
	What is Merge Sort:	It is a sorting algorithm. (Divide and merge)
	-------------------

	array -> [3, 1, 2, 4, 1, 5, 2, 6, 4] ---> 9 elements
		- This algo hypothitically divide the array into two parts
			- 5/4 or 4/5 part for the given array.
				- [3, 1, 2, 4, 1]--[5, 2, 6, 4] ====> 5 and 4 element after divide
						/	   \		/   \
				- [3, 1, 2]--[4,1]--[5, 2]--[6, 4] ====> divide into more smaller part
					/	 \    /  \    /  \    /  \
				- [3, 1] [2] [4] [1] [5] [2] [6] [4] ====> divide into more smaller parts
				   /   \
				- [3]  [1]	====> now we can't divide further more we have to stop here.

		- After divide the array now merge single element in a sorted way
			- Let,s say we have [3] and [1] to merge them in a sorted way we can merge them in a array like
				- [1, 3] this way.
					- [3] [1]
					    \ /
					- [1, 3]
			- Now we have [1, 3] & [2] imaginarray array and we want to merge them.
				- Here we can say the left and right array both are sorted in their position
					- like [1,3]& [2]
					- let's merge them 
						- [1, 3] [2] 	====> to merge the in a single array 
						- [_, _, _] 	====> we have this array.
						- 1, & 2		====> let's say from the left array we select 1, and from the right array 2
						- 1 > 2 or 1 < 2 ====> compare this both and select the smaller one and push it to the sorted array
						- [1, _, _] 	====> push 1 into the sorted array
						- 3, & 2 		====> now again select 3 & 2 from both array and compare them
						- 3 < 2 or 2 < 3 ====> select the smaller one and push it to the sorted array
						- [1, 2, _] 	====> 2 pushed to the sorted array
						- 3 			====> now we only left with 3 
						- [1, 2, 3] 	====> push 3 also into the sorted array
					
						- like : 
							- [1, 3] [2]
								\	 /
							- [1, 2, 3] ====> sorted one



*/
func MergeSort(nums []int) []int {
	return nums
}
