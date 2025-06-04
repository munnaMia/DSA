
/*
    Algorithm : Selection Sort
    Time Complexity : O(n^2)
    Space Complexity : O(1)
*/

function selectionSort(nums) {
    for (let i = 0; i < nums.length; i++) {
        let minIdx = i; // select the minimum idx
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[j] < nums[minIdx]) {
                minIdx = j;
            }
        }

        // Swap data
        let temp = nums[i];
        nums[i] = nums[minIdx];
        nums[minIdx] = temp;
    }
    return nums;
}

console.log(selectionSort([3, 4, 1, 34, 21, 1, 0]));
