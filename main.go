// This is the main entry point for the application. It initializes the clients, commands, and configuration, and then starts the REPL.
package main

import (
	"bufio"
	"os"
	"time"

	"github.com/SamuelAboelkhir/CompGoR/internal/clients"
	"github.com/SamuelAboelkhir/CompGoR/internal/commands"
	"github.com/SamuelAboelkhir/CompGoR/internal/config"
)

// TODO: Add a way to configure the API client and the scanner from the command line
func main() {
	clientsRegistry := clients.Clients{
		RegisteredClients: make(map[string]clients.Client),
	}
	httpClient := clientsRegistry.NewClient(5*time.Second, 5*time.Minute, "HTTP")
	clientsRegistry.Register("pubChem", httpClient)

	newScanner := bufio.NewScanner(os.Stdin)

	c := commands.Commands{
		RegisteredCommands: make(map[string]commands.Command),
	}

	builder := commands.QueryBuilder{}
	help := commands.Help{
		Commands: &c,
	}
	showTable := commands.ShowTable{}

	c.Register(builder.Name(), &builder)
	c.Register(help.Name(), &help)
	c.Register(showTable.Name(), &showTable)

	cfg := config.Config{
		APIClient: clientsRegistry.RegisteredClients["pubChem"],
		Scanner:   newScanner,
	}

	repl(&cfg, &c)
}
