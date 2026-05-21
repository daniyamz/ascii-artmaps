package main

import (
	ascii "ascii-artmaps/asciimaps"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage <input> <filename>")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Argument should not be more than 3")
		return
	}

	input := os.Args[1]
	filename := os.Args[2]
	if input == "" {
		return
	}

	file, err := ascii.LoadFile(filename + ".txt")
	if err != nil {
		fmt.Println("Error", err)
	}
	str := ascii.GenerateBanner(input, file)
	fmt.Print(str)
}
