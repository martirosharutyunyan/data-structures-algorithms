const interpolationSearch = (array, elem) => {
    const size = array.length;

    let low = 0;
    let high = size - 1;
    let middle;

    while ((array[high] !== array[low]) && (elem >= array[low]) && (elem <= array[high])) {
        middle = low + ((elem - array[low]) * (high - low) / (array[high] - array[low]));

        if (array[middle] < elem) {
            low = middle + 1
        } else if (elem < array[middle]) {
            high = middle - 1
        } else {
            return middle;
        };
    }

    if (elem === array[low]) {
        return low;
    }

    return -1;
}

const array = [10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47];
const elem = 22;

const index = interpolationSearch(array, elem);
console.log(index);