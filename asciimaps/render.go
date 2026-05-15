package ascii

import "strings"

func RenderLine(str string, banner map[rune][]string) []string {

	text_slice := []string{}
	for i := range 8 {
		var words_build strings.Builder
		for _, char := range str {

			words_build.WriteString(banner[char][i])
		}
		text_slice = append(text_slice, words_build.String())
	}
	return text_slice
}
