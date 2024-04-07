package main

import (
	"bufio"
	"fmt"
	"os"
)

type Keyword string
type Symbol rune

const (
	SelectKeyword Keyword = " SELECT"
	FromKeyword   Keyword = "FROM"
	SemicolonChar Symbol  = ';'
	AsteriskChar  Symbol  = '*'
	Whitespace            = " \t"
)

func main() {
	fmt.Println("Enter the SQL Query")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("you entered the following SQL Query  " + input)
}
