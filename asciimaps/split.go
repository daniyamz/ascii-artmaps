package ascii

import "strings"

func SplitInput(input string) []string {

	input = strings.ReplaceAll(input, "\\n", "\n")
	inputsplit := strings.Split(input, "\n")
	return inputsplit

}
