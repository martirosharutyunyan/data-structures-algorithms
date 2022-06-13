function swap(arr, xp, yp) {
  const temp = arr[xp];
  arr[xp] = arr[yp];
  arr[yp] = temp;
}

// An optimized version of Bubble Sort
function bubbleSort(arr, n) {
  let i, j;
  for (i = 0; i < n - 1; i++) {
    for (j = 0; j < n - i - 1; j++) {
      if (arr[j] > arr[j + 1]) {
        swap(arr, j, j + 1);
      }
    }
  }
}

// Driver program to test above functions
var arr = [64, 34, 25, 12, 22, 11, 90];

bubbleSort(arr, n);
console.log(arr)