package main

import (
  "geoffrey/commands"

  "os"
)

func main() {
  commands.RegisterCommand(&commands.ServerCommand{})
  commands.RegisterCommand(&commands.HelpCommand{})

  commands.RunCommand(os.Args[1:])
}
