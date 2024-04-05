package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	JSON_QUOTE        = '"'
	JSON_WHITESPACE   = " \t\n\r"
	JSON_SYNTAX       = "{}[],:"
	trueString        = "true"
	falseString       = "false"
	nullString        = "null"
	JSON_RIGHTBRACKET = "]"
	JSON_RIGHTBRACE   = "}"
	JSON_LEFTBRACKET  = "["
	JSON_LEFTBRACE    = "{"
	JSON_COMMA        = ","
	JSON_COLON        = ":"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the JSON Object\n")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered the following JSON Object", input)
	fmt.Println()
	tokens := lex(input)
	for _, ch := range tokens {
		fmt.Printf("'%s', ", ch)
	}
	fmt.Println()
	var interfaceTokens []interface{}
	for _, token := range tokens {
		interfaceTokens = append(interfaceTokens, token)
	}

	result, remainingTokens, err := parse(interfaceTokens)

	fmt.Println("Parsed result:", result)

	fmt.Println("Remaining tokens:", remainingTokens)

	if err != nil {
		fmt.Println("Error:", err)
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

func parse(tokens []interface{}) (interface{}, []interface{}, error) {
	if len(tokens) == 0 {
		return nil, nil, errors.New("Unexpected end of input")
	}
	t := tokens[0].(string)
	if t == "[" {
		return parseArray(tokens[1:])
	} else if t == "{" {
		return parseObject(tokens[1:])
	}
	return t, tokens[1:], nil
}

func parseArray(tokens []interface{}) ([]interface{}, []interface{}, error) {
	var jsonArray []interface{}

	if len(tokens) == 0 || tokens[0] == JSON_RIGHTBRACKET {
		return jsonArray, tokens[1:], nil
	}

	for {
		json, remainingTokens, err := parse(tokens)
		if err != nil {
			return nil, nil, err
		}
		jsonArray = append(jsonArray, json)

		if len(remainingTokens) == 0 {
			return nil, nil, errors.New("Expected end-of-array bracket")
		}

		if remainingTokens[0] == JSON_RIGHTBRACKET {
			return jsonArray, remainingTokens[1:], nil
		} else if remainingTokens[0] != JSON_COMMA {
			return nil, nil, errors.New("Expected comma after object in array")
		}

		tokens = remainingTokens[1:]
	}
}

func parseObject(tokens []interface{}) (map[string]interface{}, []interface{}, error) {
	jsonObject := make(map[string]interface{})

	if len(tokens) == 0 || tokens[0] == JSON_RIGHTBRACE {
		return jsonObject, tokens[1:], nil
	}

	for {
		if len(tokens) == 0 {
			return nil, nil, errors.New("Unexpected end of input")
		}

		jsonKey := tokens[0].(string)
		tokens = tokens[1:]

		if len(tokens) == 0 || tokens[0].(string) != JSON_COLON {
			return nil, nil, errors.New("Expected colon after key in object")
		}
		tokens = tokens[1:]

		jsonValue, remainingTokens, err := parse(tokens)
		if err != nil {
			return nil, nil, err
		}
		jsonObject[jsonKey] = jsonValue

		if len(remainingTokens) == 0 {
			return nil, nil, errors.New("Expected end-of-object brace")
		}

		if remainingTokens[0].(string) == JSON_RIGHTBRACE {
			return jsonObject, remainingTokens[1:], nil
		} else if remainingTokens[0].(string) != JSON_COMMA {
			return nil, nil, errors.New("Expected comma after pair in object")
		}

		tokens = remainingTokens[1:]
	}
}
