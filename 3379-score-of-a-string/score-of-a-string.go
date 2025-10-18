func scoreOfString(s string) int {
    var score int

    for i:=0; i < len(s) - 1; i++ {
        score += int(math.Abs(float64(int(s[i]) - int(s[i+1]))))
    }
    return score
}