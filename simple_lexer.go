package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Keyword string
type Operator rune

const (
	SelectKeyword Keyword  = "SELECT"
	FromKeyword   Keyword  = "FROM"
	InsertKeyword Keyword  = "INSERT"
	WhereKeyword  Keyword  = "WHERE"
	AsteriskChar  Operator = '*'
	SemicolonChar Operator = ';'
	EqualstoChar  Operator = '='
)

type Token struct {
	Type  string
	Value string
}

func main() {

	i := 0
	for i < 10 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the SQL query\n")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println("You entered the following SQL Query", input)
		fmt.Println()
		tokens := lexer(input)
		fmt.Println("Tokens")
		for _, token := range tokens {
			fmt.Printf("%s: %s\n", token.Type, token.Value)
		}
		fmt.Println()
		i++
	}

}

func lexer(input string) []Token {
	var tokens []Token
	substrings := strings.Split(input, " ")
	for _, word := range substrings {
		switch word {
		case string(SelectKeyword), string(FromKeyword), string(InsertKeyword), string(WhereKeyword):
			tokens = append(tokens, Token{Type: "KEYWORD", Value: word})
		case string(AsteriskChar), string(SemicolonChar), string(EqualstoChar):
			tokens = append(tokens, Token{Type: "OPERATOR", Value: word})
		default:
			tokens = append(tokens, Token{Type: "IDENTIFIER", Value: word})
		}
	}
	return tokens
}
