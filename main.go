package main

import (
	"fmt"
	"os"
)

const (
	exitStatusOK = iota
	exitStatusCLIError
	exitStatusConvertError
)

func Convert(inp string) string {
	return inp
}

func main() {
	inputText := ""

	if inputText != "" {
		inputText := Convert(inputText)
		fmt.Println(inputText)
		os.Exit(exitStatusOK)
	}
}
