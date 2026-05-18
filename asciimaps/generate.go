package ascii

import "strings"

func GenerateBanner(str string, banner map[rune][]string) string {
	sstr := SplitInput(str)
	var words strings.Builder

	for _, cha := range sstr {
		line := RenderLine(cha, banner)
		for _, ch := range line {
			words.WriteString(ch + "\n")
		}
	}
	return words.String()
}
