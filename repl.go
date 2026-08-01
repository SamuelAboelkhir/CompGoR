package main

import (
	"fmt"

	"github.com/SamuelAboelkhir/CompGoR/internal/commands"
	"github.com/SamuelAboelkhir/CompGoR/internal/config"
	"github.com/SamuelAboelkhir/CompGoR/internal/utils"
)

// REPL (Read, Eval, Print, Loop) for the application
// This function provides an interactive command-line interface for the user to interact with the application.
// It continuously reads user input, processes it, and executes the appropriate command.
func repl(cfg *config.Config, cmd *commands.Commands) {
	for {
		fmt.Print("PubChem > ")
		cfg.Scanner.Scan()

		input := utils.CleanInput(cfg.Scanner.Text())
		if len(input) <= 0 {
			fmt.Println("Please provide a command name, or use 'help' for a list of available commands")
			continue
		}
		commandName := input[0]
		args := input[1:]
		err := cmd.Run(cfg, commandName, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
