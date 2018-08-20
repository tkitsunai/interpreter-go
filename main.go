package main

import (
	"fmt"
	"os/user"

	"os"

	"github.com/tkitsunai/interpreter-go/repl"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!, This is the itnerpreter-go programming language!\n enjoy it!", user.Username)

	fmt.Printf("Feed free to type in commands \n")

	repl.Start(os.Stdin, os.Stdout)

}
