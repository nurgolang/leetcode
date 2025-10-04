import (
    "strconv"
    "strings"
)

func convertDateToBinary(date string) string {
	year, _ := strconv.ParseInt(date[0:4], 10, 64)
	month, _ := strconv.ParseInt(date[5:7], 10, 64)
	day, _ := strconv.ParseInt(date[8:], 10, 64)

    var res strings.Builder 
    res.WriteString(strconv.FormatInt(year, 2))
    res.WriteString("-")
    res.WriteString(strconv.FormatInt(month, 2))
    res.WriteString("-")
    res.WriteString(strconv.FormatInt(day, 2))
    
    return res.String()
}