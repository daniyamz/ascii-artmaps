package ascii

import (
	"bufio"
	"errors"
	"os"
)

func LoadFile(str string) (map[rune][]string, error) {
	words := make(map[rune][]string)
	slic := []string{}
	file, err := os.Open(str)
	for err != nil {
		return nil, errors.New("Error")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		slic = append(slic, text)
	}
	for i := 32; i < 127; i++ {
		stat := (i-32)*9 + 1
		words[rune(i)] = slic[stat : stat+8]
	}
	return words, nil
}
