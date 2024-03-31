// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// const JSON_QUOTE = "\""
// const JSON_WHITESPACE = " \t\n\r"
// const JSON_SYNTAX = "{}[],:"

// type JSONValue struct {
// 	Type  string      // Type of JSON value (string, number, bool, null)
// 	Value interface{} // Value of JSON value
// }

// func lexString(input string) (string, string) {
// 	var jsonString string

// 	if len(input) == 0 || string(input[0]) != JSON_QUOTE {
// 		return "", input
// 	}

// 	input = input[1:] // Skip opening quote

// 	for i, r := range input {
// 		if r == '"' && input[i-1] != '\\' { // Check for unescaped closing quote
// 			return input[:i+1], input[i+1:]
// 		}
// 		jsonString += string(r)
// 	}

// 	panic("Expected end-of-string quote")
// }

// func lexNumber(input string) (JSONValue, string) {
// 	var jsonNumber string

// 	for _, r := range input {
// 		if !unicode.IsDigit(r) && r != '-' && r != 'e' && r != '.' {
// 			break
// 		}
// 		jsonNumber += string(r)
// 	}

// 	rest := input[len(jsonNumber):]

// 	if len(jsonNumber) == 0 {
// 		return JSONValue{}, input
// 	}

// 	if contains(jsonNumber, '.') {
// 		return JSONValue{Type: "number", Value: jsonNumber}, rest
// 	}

// 	return JSONValue{Type: "number", Value: jsonNumber}, rest
// }

// func contains(s string, char rune) bool {
// 	for _, c := range s {
// 		if c == char {
// 			return true
// 		}
// 	}
// 	return false
// }

// func lexBool(input string) (JSONValue, string) {
// 	if len(input) >= 4 && input[:4] == "true" {
// 		return JSONValue{Type: "bool", Value: true}, input[4:]
// 	}
// 	if len(input) >= 5 && input[:5] == "false" {
// 		return JSONValue{Type: "bool", Value: false}, input[5:]
// 	}
// 	return JSONValue{}, input
// }

// func lexNull(input string) (JSONValue, string) {
// 	if len(input) >= 4 && input[:4] == "null" {
// 		return JSONValue{Type: "null", Value: nil}, input[4:]
// 	}
// 	return JSONValue{}, input
// }

// func lex(input string) []JSONValue {
// 	var tokens []JSONValue

// 	for len(input) > 0 {
// 		jsonString, rest := lexString(input)
// 		if jsonString != "" {
// 			tokens = append(tokens, JSONValue{Type: "string", Value: jsonString})
// 			input = rest
// 			continue
// 		}

// 		jsonNumber, rest := lexNumber(input)
// 		if jsonNumber != (JSONValue{}) {
// 			tokens = append(tokens, jsonNumber)
// 			input = rest
// 			continue
// 		}

// 		jsonBool, rest := lexBool(input)
// 		if jsonBool != (JSONValue{}) {
// 			tokens = append(tokens, jsonBool)
// 			input = rest
// 			continue
// 		}

// 		jsonNull, rest := lexNull(input)
// 		if jsonNull != (JSONValue{}) {
// 			tokens = append(tokens, jsonNull)
// 			input = rest
// 			continue
// 		}

// 		if unicode.IsSpace(rune(input[0])) {
// 			input = input[1:]
// 			continue
// 		}

// 		if contains(JSON_SYNTAX, rune(input[0])) {
// 			tokens = append(tokens, JSONValue{Type: "syntax", Value: string(input[0])})
// 			input = input[1:]
// 			continue
// 		}

// 		panic(fmt.Sprintf("Unexpected character: %c", input[0]))
// 	}

// 	return tokens
// }

// func main() {
// 	input := `{"key": "value", "number": 42, "bool": true, "null": null}`
// 	tokens := lex(input)
// 	fmt.Println(tokens)
// }
