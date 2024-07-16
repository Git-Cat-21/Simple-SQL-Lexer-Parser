package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Token struct {
	Type  TokenType
	Value string
}
type TokenType int

const (
	Integer TokenType = iota
	Plus
	Minus
	Multiply
	Divide
)

func Lexer(input string) []Token {
	var tokens []Token
	var currentToken string
	for _, char := range input {
		switch {
		case unicode.IsDigit(char):
			currentToken += string(char)
		case char == '+':
			tokens = append(tokens, Token{Type: Plus, Value: "+"})

		case char == '-':
			tokens = append(tokens, Token{Type: Minus, Value: "-"})
		case char == '*':
			tokens = append(tokens, Token{Type: Multiply, Value: "*"})
		case char == '/':
			tokens = append(tokens, Token{Type: Divide, Value: "/"})
		case char == ' ':
			if currentToken != "" {
				tokens = append(tokens, Token{Type: Integer, Value: currentToken})
			}
		}
	}
	if currentToken != "" {
		tokens = append(tokens, Token{Type: Integer, Value: currentToken})
	}
	return tokens
}

func Parser(tokens []Token) (int, error) {
	var result int
	var operator TokenType
	for _, token := range tokens {
		switch token.Type {
		case Integer:
			num, err := strconv.Atoi(token.Value)
			if err != nil {
				return 0, err
			}
			switch operator {
			case Plus:
				result += num
			case Minus:
				result -= num
			case Multiply:
				result *= num
			case Divide:
				result /= num
			default:
				result = num
			}
		default:
			operator = token.Type
		}
	}
	return result, nil
}

func main() {
	expression := "3 + 5 * 2 - 8 / 2"
	tokens := Lexer(expression)
	fmt.Println(tokens)
	// result, err := Parser(tokens)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// } else {
	// 	fmt.Println("Result ", result)
	// }

}
