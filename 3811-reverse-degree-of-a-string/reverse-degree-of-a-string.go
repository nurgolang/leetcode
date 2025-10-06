func reverseDegree(s string) int {
    var res int
    for i, letter := range s {
        res += (26 - int(letter - 'a')) * (i + 1) 
    }
    return res
}