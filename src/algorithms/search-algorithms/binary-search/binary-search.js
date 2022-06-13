const binarySearchRecursive = (array, left, right, elem) => {
    if (right >= left) {
        let middle = left + Math.floor((right - left) / 2) ;

        if (array[middle] === elem) {
            return middle;
        }

        if (array[middle] > elem) {
            return binarySearch(array, left, mid - 1, elem);
        }

        return binarySearch(array, mid + 1, right, elem);
    }

    return -1;
}


const array = [1, ,2 ,3 ,4 ,5, 6, 7];
const elem = 6;
const index = binarySearchRecursive(array, 0, array.length, elem);
console.log(index);