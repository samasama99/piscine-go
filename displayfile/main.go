package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		println("Too many arguments")
		return
	}
	filename := os.Args[1]
	content, _ := os.ReadFile(filename)
	println(string(content))
}

// Write a program that displays, on the standard output, the content of a file given as argument.
//
// $ go run .
// File name missing
// $ echo 'Almost there!!' > quest8.txt
// $ go run . quest8.txt main.go
// Too many arguments
// $ go run . quest8.txt
// Almost there!!
