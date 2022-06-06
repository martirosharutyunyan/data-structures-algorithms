const exponentialSearch = (array, elem) => {
    if (array[0] === elem) {
        return 0;
    };

    let i = 1;

    while (i < array.length && array[i] <= elem) {
        i *= 2;
    }

    return binarySearch(array, i/2, Math.min(i, array.length - 1), elem);
}

const binarySearch = (array, left, right, elem) => {
    let middle;

    while (right >= left) {
        middle = left + (right - left) / 2;

        if (array[middle] == elem) {
            return middle;
        }

        if (array[middle] > elem) {
            right = middle - 1;
        } else {
            left = middle + 1
        };
    }

    return -1;
}

const array = [1, ,2 ,3 ,4 ,5, 6, 7];
const elem = 6;
const index = exponentialSearch(array, elem);
console.log(index);