package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		for {
			bytes := make([]byte, 1024)
			os.Stdin.Read(bytes)
			print(string(bytes))
		}
	}
	files := os.Args[1:]
	for _, filename := range files {
		content, err := os.ReadFile(filename)
		if err != nil {
			println(err.Error())
			continue
		}
		fmt.Print(string(content))
		println()
	}
}
