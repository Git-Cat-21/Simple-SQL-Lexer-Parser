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
type SelectStatement struct {
	AllColumns bool
	Columns    []string
	Table      string
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
	statement, err := parseSelectStatement(lexer(input))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("SELECT %s FROM %s;\n", strings.Join(statement.Columns, ","), statement.Table)
	fmt.Println("Parsed the statement successfully")
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

func parseSelectStatement(tokens []Token) (*SelectStatement, error) {
	var statement SelectStatement
	var inSelectColumns bool
	var seenFrom bool

	for _, token := range tokens {
		switch token.Type {
		case "Keyword":
			if token.Value == "SELECT" {
				inSelectColumns = true
			} else if token.Value == "FROM" {
				inSelectColumns = false
				seenFrom = true
			} else {
				return nil, fmt.Errorf("unexpected keyword: %s", token.Value)
			}
		case "Identifier":
			if inSelectColumns {
				statement.Columns = append(statement.Columns, token.Value)
			} else if seenFrom {
				statement.Table = token.Value
			} else {
				return nil, fmt.Errorf("unexpected identifier: %s", token.Value)
			}
		case "Symbol":
			if token.Value == "*" {
				if len(statement.Columns) > 0 {
					return nil, fmt.Errorf("cannot mix * with column names")
				}
				statement.AllColumns = true
			} else if token.Value == ";" {
				if len(statement.Columns) == 0 && !statement.AllColumns || statement.Table == "" {
					return nil, fmt.Errorf("incomplete SQL statement")
				}
				return &statement, nil
			} else {
				return nil, fmt.Errorf("unexpected symbol: %s", token.Value)
			}
		default:
			return nil, fmt.Errorf("unexpected token type: %s", token.Type)
		}
	}

	return nil, fmt.Errorf("incomplete SQL statement")
}
