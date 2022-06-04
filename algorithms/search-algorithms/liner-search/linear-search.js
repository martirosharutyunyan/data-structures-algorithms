const linearSearch = (array, elem) => {
    for(let i = 0; i < array.length; i++) {
        if (array[i] === elem) return i;
    };

    return -1;
}

const array = [1, 2, 3, 4, 5, 6, 7, 8];
const elem = 7;
const index = linearSearch(array, elem);
console.log(index);