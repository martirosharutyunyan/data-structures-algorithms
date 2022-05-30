const diapason = 10

const countingSort = (array) => {
    const countArray = new Array(diapason).fill(0)

    for(let i = 0; i < array.length; i++) {
        countArray[array[i]]++
    };

    let z = 0

    for(let i = 0; i < countArray.length; i++) {
        while (countArray[i] > 0) {
            array[z] = i
            countArray[i]--
            z++
        };    
    };
}

const array = [1, 4, 1, 2, 7, 5, 2]

countingSort(array)
console.log(array)