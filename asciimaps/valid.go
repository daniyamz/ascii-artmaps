package ascii

import "errors"

func ValidateInput(str string) (rune, error) {
	for _, cha := range str {
		if cha < 32 || cha > 126 {
			return cha, errors.New("invalid entery")
		}
	}
	return 0, nil
}
