func finalValueAfterOperations(operations []string) int {
    var res int
    for _, operation := range operations {
        if operation == "--X" || operation == "X--" {
            res--
        } else {
            res++
        }
    } 
    return res
}