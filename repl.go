package main

import (
	"fmt"

	"github.com/SamuelAboelkhir/CompGoR/internal/commands"
	"github.com/SamuelAboelkhir/CompGoR/internal/config"
	"github.com/SamuelAboelkhir/CompGoR/internal/utils"
	"github.com/chzyer/readline"
)

// REPL (Read, Eval, Print, Loop) for the application
// This function provides an interactive command-line interface for the user to interact with the application.
// It continuously reads user input, processes it, and executes the appropriate command.
func repl(cfg *config.Config, cmd *commands.Commands) {
	config := &readline.Config{
		Prompt:          "CompGoR > ",
		HistoryFile:     ".history",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	rl, err := readline.NewEx(config)
	if err != nil {
		fmt.Println("Error initializing readline:", err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		input := utils.CleanInput(line)
		if len(input) <= 0 {
			fmt.Println("Please provide a command name, or use 'help' for a list of available commands")
			continue
		}

		commandName := input[0]
		if commandName == "exit" {
			break
		}
		args := input[1:]
		err = cmd.Run(cfg, commandName, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("help"),
	readline.PcItem("buildQuery"),
	readline.PcItem("showTable"),
	readline.PcItem("exit"),
)
