// Javascript program for implementation of selection sort
function swap(arr,xp, yp)
{
	const temp = arr[xp];
	arr[xp] = arr[yp];
	arr[yp] = temp;
}

function selectionSort(arr, n)
{
	let i, j, min_idx;

	// One by one move boundary of unsorted subarray
	for (i = 0; i < n-1; i++)
	{
		// Find the minimum element in unsorted array
		min_idx = i;
		for (j = i + 1; j < n; j++)
		if (arr[j] < arr[min_idx])
			min_idx = j;

		// Swap the found minimum element with the first element
		swap(arr,min_idx, i);
	}
}

const arr = [64, 25, 12, 22, 11];
const n = 5;
selectionSort(arr, n);
console.log(arr)
// This code is contributed by akshitsaxenaa09.
