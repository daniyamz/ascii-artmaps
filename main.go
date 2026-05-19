package main

import (
	ascii "ascii-artmaps/asciimaps"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	if input == "" {
		return
	}
	file, _ := ascii.LoadFile("thinkertoy.txt")
	str := ascii.GenerateBanner(input, file)
	fmt.Print(str)
}
