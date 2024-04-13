package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Keyword string
type Symbol rune

const (
	SelectKeyword   Keyword = "SELECT"
	FromKeyword     Keyword = "FROM"
	AsteriskSymbol  Symbol  = '*'
	SemicolonSymbol Symbol  = ';'
)

type Token struct {
	Type  string
	Value string
}

func main() {
	fmt.Println("Enter the SQL Query:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered the following SQL Query:", input)
	fmt.Println()
	tokens := lexer(input)
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func lexer(input string) []Token {
	var tokens []Token
	var current string

	for i := 0; i < len(input); i++ {
		char := input[i]
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
		return append(tokens, Token{Type: "Keyword", Value: value})
	} else if isSymbol(rune(value[0])) {
		return append(tokens, Token{Type: "Symbol", Value: value})
	} else {
		return append(tokens, Token{Type: "Identifier", Value: value})
	}
}

func isKeyword(token string) bool {
	switch strings.ToUpper(token) {
	case string(SelectKeyword), string(FromKeyword):
		return true
	default:
		return false
	}
}

func isSymbol(char rune) bool {
	switch char {
	case rune(AsteriskSymbol), rune(SemicolonSymbol):
		return true
	default:
		return false
	}
}
