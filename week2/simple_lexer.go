// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Keyword string
// type Symbol rune

// const (
// 	SelectKeyword Keyword = "SELECT"
// 	FromKeyword   Keyword = "FROM"
// 	InsertKeyword Keyword = "INSERT"
// 	IntoKeyword   Keyword = "INTO"
// 	ValuesKeyword Keyword = "VALUES"
// 	AsteriskChar  Symbol  = '*'
// 	SemicolonChar Symbol  = ';'
// 	EqualstoChar  Symbol  = '='
// 	OpenParen     Symbol  = '('
// 	CloseParen    Symbol  = ')'
// )

// type Token struct {
// 	Type  string
// 	Value string
// }

// func main() {

// 	i := 0
// 	for i < 3 {
// 		scanner := bufio.NewScanner(os.Stdin)
// 		fmt.Println("Enter the SQL query\n")
// 		scanner.Scan()
// 		input := scanner.Text()
// 		fmt.Println("You entered the following SQL Query", input)
// 		fmt.Println()
// 		tokens := lexer(input)

// 		fmt.Println("Tokens")
// 		for _, token := range tokens {
// 			fmt.Printf("%s: %s\n", token.Type, token.Value)
// 		}
// 		parser(tokens)
// 		fmt.Println()
// 		i++
// 	}

// }

// func lexer(input string) []Token {
// 	var tokens []Token
// 	substrings := strings.Split(input, " ")
// 	for _, word := range substrings {
// 		switch word {
// 		case string(SelectKeyword), string(FromKeyword), string(InsertKeyword), string(IntoKeyword), string(ValuesKeyword):
// 			tokens = append(tokens, Token{Type: "KEYWORD", Value: word})
// 		case string(AsteriskChar), string(SemicolonChar), string(EqualstoChar), string(OpenParen), string(CloseParen):
// 			tokens = append(tokens, Token{Type: "SYMBOL", Value: word})
// 		default:
// 			tokens = append(tokens, Token{Type: "IDENTIFIER", Value: word})
// 		}
// 	}
// 	return tokens
// }

// func parser(tokens []Token) {
// 	if len(tokens) == 0 {
// 		fmt.Println("Invalid SQL Query Empty")
// 		return
// 	}
// 	if len(tokens) <= 4 {
// 		fmt.Println("Invalid SQL Query: Incomplete")
// 		return
// 	}
// 	if tokens[0].Type != "KEYWORD" {
// 		fmt.Println("Invalid SQL Query: Missing the keyword")
// 		return
// 	}
// 	switch tokens[0].Type {
// 	case "KEYWORD":
// 		if tokens[0].Value == string(SelectKeyword) {
// 			parseSelect(tokens)
// 		} else if tokens[0].Value == string(InsertKeyword) {
// 			parseInsert(tokens)
// 		}
// 	}
// }

// func parseSelect(tokens []Token) {
// 	if tokens[1].Type == "SYMBOL" || tokens[1].Value == string(AsteriskChar) || tokens[1].Type != "SYMBOL" {
// 		parseSelectAll(tokens)
// 		return
// 	}
// 	if tokens[1].Type == "IDENTIFIER" {
// 		fmt.Println("Parsed SQL Query")
// 		return
// 	}

// 	if tokens[3].Type == "IDENTIFIER" {
// 		fmt.Println("Invalid SQL query")
// 		return
// 	}
// 	if tokens[4].Type != "SYMBOL" || tokens[4].Value != string(SemicolonChar) {
// 		fmt.Println("Invalid SQL query")
// 		return
// 	} else {
// 		fmt.Println("Parsed SQL Query")
// 		return
// 	}
// }

// func parseSelectAll(tokens []Token) {
// 	if tokens[1].Value != string(AsteriskChar) {
// 		fmt.Println("Invalid SQL query")
// 		return
// 	}
// 	if tokens[2].Type != "KEYWORD" || tokens[2].Value != string(FromKeyword) {
// 		fmt.Println("Invalid SQL query")
// 		return

// 	}
// 	if tokens[3].Type != "IDENTIFIER" {
// 		fmt.Println("Invalid SQL query")
// 		return
// 	}
// 	if tokens[4].Type != "SYMBOL" || tokens[4].Value != string(SemicolonChar) {
// 		fmt.Println("Invalid SQL query")
// 		return
// 	} else {
// 		fmt.Println("Parsed SQL Query")
// 		return
// 	}
// }

// func parseInsert(tokens []Token) {
// 	if tokens[1].Type == "KEYWORD" && tokens[1].Value == string(IntoKeyword) {
// 		if tokens[2].Type == "IDENTIFIER" {
// 			if tokens[3].Type == "KEYWORD" && tokens[3].Value == string(ValuesKeyword) {
// 				if tokens[4].Type == "SYMBOL" && tokens[4].Value == string(OpenParen) {
// 					if tokens[5].Type == "IDENTIFIER" {
// 						if tokens[6].Type == "SYMBOL" && tokens[6].Value == string(CloseParen) {
// 							if tokens[7].Type == "SYMBOL" && tokens[7].Value == string(SemicolonChar) {
// 								fmt.Println("Parsed SQL QUERY")
// 							}

// 						}
// 					}
// 				}

// 			}

// 		}
// 	} else {
// 		fmt.Println("Invalid SQL query")

// 	}

// }
