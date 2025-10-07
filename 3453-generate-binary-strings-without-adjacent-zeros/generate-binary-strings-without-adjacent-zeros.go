func validStrings(n int) []string {
    res := []string{}
    
    var backtrack func(currentString string, previousChar rune)
    backtrack = func(currentString string, previousChar rune) {
        if len(currentString) == n {
            res = append(res, currentString)
            return // stop recursion
            }
        if previousChar == '0' {
            backtrack(currentString + "1", '1')
        } else {
            backtrack(currentString + "0", '0')
            backtrack(currentString + "1", '1')
        }
    }

    backtrack("0", '0')
    backtrack("1", '1')
   
    return res
}