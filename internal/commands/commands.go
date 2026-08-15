// Package commands: This package is where all the commands are defined.
// Commands implement a Command interface to allow for more flexibility with their inputs and outputs
package commands

import (
	"errors"

	"github.com/SamuelAboelkhir/CompGoR/internal/config"
)

// Command interface is implemented by all commands. It allows for more flexibility with their inputs and outputs.
type Command interface {
	Execute(cfg *config.Config, args ...string) error
	Name() string
	Help() string
}

// Commands struct holds a map of all the registered commands.
type Commands struct {
	RegisteredCommands map[string]Command
}

// Run executes the command with the given name and arguments.
func (c *Commands) Run(cfg *config.Config, name string, args ...string) error {
	command, ok := c.RegisteredCommands[name]
	if !ok {
		return errors.New("please provide a valid command name")
	}
	return command.Execute(cfg, args...)
}

// Register adds a command to the map of registered commands.
func (c *Commands) Register(name string, cmd Command) {
	c.RegisteredCommands[name] = cmd
}
