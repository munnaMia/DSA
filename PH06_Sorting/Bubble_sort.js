function BubbleSort(array) {
    for (let i = array.length - 1; i >= 0; i--) {
        for (let j = 0; j <= i - 1; j++) {
            if (array[j] > array[j + 1]) {
                let temp = array[j];
                array[j] = array[j + 1];
                array[j + 1] = temp;
            }
        }
    }

    return array; // return sorted array
}

console.log(BubbleSort([2, 9, 1, 3, 5, 3, 6, 3, 4, 5, 6]));
