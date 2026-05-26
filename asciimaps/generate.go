package ascii

import "strings"

func GenerateBanner(str string, banner map[rune][]string) string {
	sstr := SplitInput(str)
	var words strings.Builder

	for i, cha := range sstr {
		if cha == "" {
			if i == len(sstr)-1 {
				continue
			}
			words.WriteString("\n")
			continue
		}
		runes, err := ValidateInput(cha)
		if err != nil {
			return string(runes) + "\n"
		}
		line := RenderLine(cha, banner)
		for _, ch := range line {
			words.WriteString(ch)
			words.WriteString("\n")
		}
	}
	return words.String()
}
