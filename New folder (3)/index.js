function removeSubset(inputString, subsetToRemove) {
    // Escape special characters in the subset to ensure proper matching
    const escapedSubset = subsetToRemove.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
    // Create a regular expression to match the subset
    const regex = new RegExp(escapedSubset, 'g');
    // Use the replace method to remove all occurrences of the subset
    const resultString = inputString.replace(regex, '');
    return resultString.trim();
}

// Example usage:
const inputString = ["This", "is", "a" ,"sample" ,"sentence"];
const subsetToRemove = "sample sentence";
const result = removeSubset(inputString.join(" "), subsetToRemove);


a = [1,2,3]

b = [7,8,9]

b = a

console.log(b)
