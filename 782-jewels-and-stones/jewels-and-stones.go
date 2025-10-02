func numJewelsInStones(jewels string, stones string) int {
    var res int
    for _, jewel := range jewels {
        for _, stone := range stones {
            if stone == jewel {
                res++
            } 
        }
    }
    return res
}