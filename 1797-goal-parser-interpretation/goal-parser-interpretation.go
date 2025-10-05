import "strings"

func interpret(command string) string {
	var replacer = strings.NewReplacer("G", "G", "()", "o", "(al)", "al")
	return replacer.Replace(command)
}