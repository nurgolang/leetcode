import "strings"

func findWordsContaining(words []string, x byte) []int {
    res := []int{}
    for i, val := range words{
        if strings.Contains(val, string(x)) {
            res = append(res, i)
        }
    }
    return res
}