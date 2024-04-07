package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Token interface{}

type Keyword string
type Symbol rune

const (
	SelectKeyword Keyword = "SELECT"
	FromKeyword   Keyword = "FROM"
	SemicolonChar Symbol  = ';'
	AsteriskChar  Symbol  = '*'
	Whitespace            = " \t"
)

func main() {
	fmt.Println("Enter the SQL Query:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered the following SQL Query:", input)
	tokens := lexer(input)
	fmt.Println(tokens)
}

func lexer(input string) []Token {
	var tokens []Token
	var current string

	for i := 0; i < len(input); i++ {
		char := input[i]
		// Handle whitespace
		if unicode.IsSpace(rune(char)) {
			if current != "" {
				tokens = appendToken(tokens, current)
				current = ""
			}
			continue
		}

		current += string(char)
	}

	if current != "" {
		tokens = appendToken(tokens, current)
	}

	return tokens
}

func appendToken(tokens []Token, value string) []Token {
	value = strings.TrimSpace(value)
	if isKeyword(value) {
		tokens = append(tokens, Keyword(value))
	} else if isSymbol(rune(value[0])) {
		tokens = append(tokens, value)
	} else {
		tokens = append(tokens, value)
	}
	return tokens
}

func isKeyword(token string) bool {
	switch strings.ToUpper(token) {
	case string(SelectKeyword), string(FromKeyword):
		return true
	}
	return false
}

func isSymbol(char rune) bool {
	switch char {
	case ';', '*':
		return true
	}
	return false
}
