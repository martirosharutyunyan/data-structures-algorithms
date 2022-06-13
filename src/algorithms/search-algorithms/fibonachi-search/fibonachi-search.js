const fibonachiSearch = (array, elem) => {
    let fibMMm2 = 0;
    let fibMMm1 = 1;
    let fibM = fibMMm2 + fibMMm1;

    while (fibM < array.length) {
        fibMMm2 = fibMMm1;
		fibMMm1 = fibM;
		fibM = fibMMm2 + fibMMm1;
    }

    let offset = -1;

    while (fibM > 1) {
        const i = Math.min(offset + fibMMm2, array.length);

        if (array[i] < elem) {
			fibM = fibMMm1;
			fibMMm1 = fibMMm2;
			fibMMm2 = fibM - fibMMm1;
			offset = i;
		} else if (array[i] > elem) {
			fibM = fibMMm2;
			fibMMm1 = fibMMm1 - fibMMm2;
			fibMMm2 = fibM - fibMMm1
		} else {
			return i;
		}

		if (array[array.length-1] == elem) {
			return array.length - 1;
		}
    }

    return -1;
}