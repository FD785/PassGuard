package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"passguard/internal/analyzer"
	"passguard/internal/output"
)

func main() {
	var password string

	if len(os.Args) > 1 {
		password = os.Args[1]
	} else {
		fmt.Print("Enter password: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		password = strings.TrimSpace(input)
	}

	result := analyzer.Analyze(password)
	output.PrintJSON(result)
}
