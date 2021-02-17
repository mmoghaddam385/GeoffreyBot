package commands

import (
	"fmt"
	"geoffrey/types"
)

const CommandNotFoundCode = -100

var commands []types.ConsoleCommand

func RegisterCommand(cmd types.ConsoleCommand) {
	commands = append(commands, cmd)
}

// RunCommand takes an slice of strings and attempts to run the command associated with it.
// The first string in the slice is used to determine the command name, the rest are pass into
// the command as arguments
func RunCommand(commandStrings []string) int {
	if len(commandStrings) > 0 {
		for _, cmd := range commands {
			if cmd.Name() == commandStrings[0] {
				return cmd.Execute(commandStrings[1:])
			}
		}

		fmt.Printf("Command not found! (%v)\n", commandStrings[0])
	} else {
		fmt.Println("Please specify a command:")
	}

	(&HelpCommand{}).Execute(nil)

	return CommandNotFoundCode
}

// ********** HELP COMMAND 

type HelpCommand struct{}

func (*HelpCommand) Name() string { return "help" }
func (*HelpCommand) Usage() string { return "help - List all commands and their usage text\n" + 
											"\thelp ...<cmd> - Print usage text for the given cmds" }
											
func (*HelpCommand) Execute(args []string) int {
	var cmdsToPrint []types.ConsoleCommand

	if len(args) == 0 {
		// No args? print help for all commands!
		cmdsToPrint = commands
	} else {
		// Some args? print help only for the given commands
		for _, cmd := range commands {
			for _, arg := range args {
				if cmd.Name() == arg {
					cmdsToPrint = append(cmdsToPrint, cmd)
				}
			}
		}
	}

	for _, cmd := range cmdsToPrint {
		fmt.Printf("%v usage:\n\t%v\n\n", cmd.Name(), cmd.Usage())
	}

	return 0
}