func smallestEvenMultiple(n int) int {
    if n % 2 == 1 {
        return 2 * n
    }
    return n
}