class Data<T> {
    first: T;
    second: T;

    constructor(first: T, second: T) {
        this.first = first;
        this.second = second;
    }
}

// Test cases
const dataInt = new Data<number>(10, 20);
console.log("dataInt ->", dataInt);
console.assert(dataInt.first === 10, "Expected first to be 10");

const dataString = new Data<string>("A", "B");
console.log("dataString ->", dataString);
console.assert(dataString.first === "A", "Expected first to be A");