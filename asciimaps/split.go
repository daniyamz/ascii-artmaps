package ascii

import "strings"

func SplitInput(input string) []string {
	inputsplit := strings.Split(input, "\\n")
	return inputsplit

}
