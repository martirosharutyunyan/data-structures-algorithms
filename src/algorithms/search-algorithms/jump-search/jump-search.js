const jumpoSearch = (array, elem) => {
    let step = parseInt(Math.sqrt(array.length));
    let prev = 0;

    while (array[Math.min(step, array.length) - 1] < elem) {
        prev = step;

        step += parseInt(Math.sqrt(array.length));
        
        if (prev >= array.length) return -1;
    }

    while(array[prev] < elem) {
        prev++;
        if (prev == Math.min(step, array.length)) {
            return -1;
        }
    }

    if (array[prev] === elem) {
        return prev;
    }

    return -1;
}


const array = [1, 2, 3, 4, 5, 6, 7, 8];
const elem = 8
const index = jumpoSearch(array, 8)
console.log(index)
