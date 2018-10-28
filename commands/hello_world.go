package commands

import (
	"fmt"
)

// HelloCommands is a types.ConsoleCommand
type HelloCommand struct {}

func (*HelloCommand) Name() string { return "hello" }
func (*HelloCommand) Usage() string {return "Just run it"}
 
func (*HelloCommand) Execute(args []string) int {
	fmt.Println("Hello, World")

	return 0
}
