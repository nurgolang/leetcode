func differenceOfSums(n int, m int) int {
   // gauss mass
   totalSum := (n * (n + 1)) / 2
   k := n / m
   divSum := ((k * (k + 1)) / 2) * m
   return totalSum - 2 * divSum
}