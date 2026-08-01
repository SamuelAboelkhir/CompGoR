package commands

import (
	"fmt"

	"github.com/SamuelAboelkhir/CompGoR/internal/config"
)

// Help is a command that prints the help menu.
type Help struct {
	Commands *Commands
}

// Execute prints the help menu.
func (h *Help) Execute(cfg *config.Config, args ...string) error {
	fmt.Println("Welcome to the CompGoR CLI!")
	fmt.Println("Available Commands:")
	fmt.Println("-------------------")
	for _, command := range h.Commands.RegisteredCommands {
		fmt.Printf("%s: %s\n\n", command.Name(), command.Help())
	}
	return nil
}

// Name returns the name of the command.
func (h *Help) Name() string {
	return "help"
}

// Help returns the help text for the command.
func (h *Help) Help() string {
	return "Prints this help menu"
}
