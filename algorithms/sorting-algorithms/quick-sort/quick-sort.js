const quickSort = (array) => {
    if (array.length < 2) {
        return array
    }

    const pivot = array[0]
    
    const leftItems = []
    const rightItems = [];

    for(let i = 1; i < array.length; i++) {
        if (array[i] < pivot) {
            leftItems.push(array[i])
        } else {
            rightItems.push(array[i])
        }
    };

    return [...quickSort(leftItems), pivot, ...quickSort(rightItems)]
}

const array = [64, 25, 12, 22, 11]

console.log(quickSort(array))