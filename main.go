package main

import (
	ascii "ascii-artmaps/asciimaps"
	. "fmt"
	. "os"
)

func main() {
	if len(Args) < 3 {
		Print("Usage <input> <filename>")
		return
	}
	if len(Args) > 3 {
		Print("Argument should not be more than 3")
		return
	}

	input := Args[1]
	filename := Args[2]

	if input == "" {
		return
	}

	file, err := ascii.LoadFile(filename + ".txt")
	if err != nil {
		Println("Error occured ")
		return
	}
	str := ascii.GenerateBanner(input, file)
	Print(str)
}
