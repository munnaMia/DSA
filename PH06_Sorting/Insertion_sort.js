function InsertionSort(nums) {
    for (let i = 0; i < nums.length; i++) {
        let j = i;
        while (j > 0 && nums[j - 1] > nums[j]) {
            let temp = nums[j - 1];
            nums[j - 1] = nums[j];
            nums[j] = temp;
            j--;
        }
    }
    return nums;
}

console.log(InsertionSort([2, 3, 1, 4, 6, 3, 2, 1]));
