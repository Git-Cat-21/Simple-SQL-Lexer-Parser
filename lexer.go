package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	JSON_QUOTE      = '"'
	JSON_WHITESPACE = " \t\n\r"
	JSON_SYNTAX     = "{}[],:"
	trueString      = "true"
	falseString     = "false"
	nullString      = "null"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the SQL query\n")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered the following SQL Query", input)
	fmt.Println()
	tokens := lex(input)
	//fmt.Println(tokens)
	for _, ch := range tokens {
		fmt.Printf("'%s', ", ch)
	}

}

func lex(input string) []interface{} {
	var tokens []interface{}
	for len(input) > 0 {
		switch input[0] {
		case '{', '}', '[', ']', ':', ',':
			tokens = append(tokens, string(input[0]))
			input = input[1:]
		case JSON_QUOTE:
			str, rest := lexString(input)
			tokens = append(tokens, str)
			input = rest
		default:
			if strings.ContainsRune(JSON_WHITESPACE, rune(input[0])) {
				input = input[1:]
			} else {
				num, rest := lexNumber(input)
				tokens = append(tokens, num)
				input = rest
			}
		}
	}
	return tokens
}

func lexString(input string) (string, string) {
	if input[0] != JSON_QUOTE {
		return "", input
	}
	input = input[1:]
	str := ""
	for len(input) > 0 && input[0] != JSON_QUOTE {
		str += string(input[0])
		input = input[1:]
	}
	if len(input) == 0 || input[0] != JSON_QUOTE {
		panic("Invalid string format")
	}
	return str, input[1:]
}

func lexNumber(input string) (string, string) {
	num := ""
	for len(input) > 0 && (input[0] == '-' || (input[0] >= '0' && input[0] <= '9') || input[0] == '.') {
		num += string(input[0])
		input = input[1:]
	}
	return num, input
}

func lexBool(input string) (bool, string) {
	stringLen := len(input)
	if stringLen >= len(trueString) && input[:len(trueString)] == trueString {
		return true, input[len(trueString):]
	} else if stringLen >= len(falseString) && input[:len(falseString)] == falseString {
		return false, input[len(falseString):]
	}
	return false, input
}

func lexNull(input string) (bool, string) {
	stringLen := len(input)
	if stringLen >= len(nullString) && input[:len(nullString)] == nullString {
		return true, input[len(nullString):]
	}
	return false, input
}
