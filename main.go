package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to monkey!\n", user.Username)
	fmt.Printf("Monkey is ready\n")
	repl.Start(os.Stdin, os.Stdout)
}
