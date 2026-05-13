package main

import (
	ascii "ascii-artmaps/asciimaps"
	"fmt"
)

func main() {
	data, _ := ascii.LoadFile("thinkertoy.txt")
	fmt.Println(data)
}
