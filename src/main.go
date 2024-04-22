package main

import (
	"fmt"
	"logios/src/lexer"
	"os"
)

func main() {
	bytes, err := os.ReadFile("00.lang")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	source := string(bytes)
	tokens := lexer.Tokenize(string(source))
	for _, token := range tokens {
		token.Debug()
	}

}
