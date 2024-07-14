package main

import (
	"fmt"
	"monkey_interpreter_go/repl"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language repl!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
