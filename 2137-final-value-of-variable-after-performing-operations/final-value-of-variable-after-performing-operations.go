func finalValueAfterOperations(operations []string) int {
    var res int
    for _, operation := range operations {
        if operation[1] == '+' {
            res++
        } else {
            res--
        }
    } 
    return res
}