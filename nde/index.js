function addCommasToNumber(numberString) {
    const number = parseFloat(numberString);
    if (!isNaN(number)) {
        return number.toLocaleString();
    } else {
        return "Invalid number format";
    }
}

// Example usage:
const numberWithCommas = addCommasToNumber("1000000");
console.log(numberWithCommas);
