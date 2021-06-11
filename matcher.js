const OPEN_BRACKET = "{";
const CLOSED_BRACKET = "}";
    
function checkBrackets(stringToTest) {
    let bracketCount = 0;

    for (const char of stringToTest.split('')) {
        if (char === OPEN_BRACKET) 
            bracketCount++;
        else if (char === CLOSED_BRACKET)
            bracketCount--;
            
        if (bracketCount < 0) // Test if we closed a bracket that didn't exist
            break;
    }

    return bracketCount == 0;
}



console.log(checkBrackets("{"))
console.log(checkBrackets("}"))
console.log(checkBrackets("}{"))
console.log(checkBrackets("{}"))
console.log(checkBrackets("a{a}"))