package main

import (
	"fmt"
	"monkey_interpreter_go/lexer"
	"monkey_interpreter_go/token"
)

func main() {
	l := lexer.New("==!234")

	tok := l.NextToken()
	fmt.Println(tok.Type)
	fmt.Println(tok.Literal)

	for tok.Type != token.EOF {
		tok := l.NextToken()
		fmt.Println(tok.Type)
		fmt.Println(tok.Literal)
	}
}
