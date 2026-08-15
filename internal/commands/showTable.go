package commands

import (
	"os/exec"

	"github.com/SamuelAboelkhir/CompGoR/internal/config"
)

// ShowTable is a command that prints the help menu.
type ShowTable struct{}

// Execute executes the command.
func (h *ShowTable) Execute(cfg *config.Config, args ...string) error {
	cmd := exec.Command("xdg-open", "./assets/ColorLargeTypePeriodicTable.png")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// Name returns the name of the command.
func (h *ShowTable) Name() string {
	return "showTable"
}

// Help returns the help text for the command.
func (h *ShowTable) Help() string {
	return "Renders a periodic table using the system's default image renderer."
}
