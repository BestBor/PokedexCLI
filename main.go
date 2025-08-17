package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)

	return words
}