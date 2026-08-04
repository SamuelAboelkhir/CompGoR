// This is the main entry point for the application. It initializes the clients, commands, and configuration, and then starts the REPL.
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SamuelAboelkhir/CompGoR/internal/clients"
	"github.com/SamuelAboelkhir/CompGoR/internal/commands"
	"github.com/SamuelAboelkhir/CompGoR/internal/config"
)

// TODO: Add a way to configure the API client and the scanner from the command line
func main() {
	clientsRegistry := clients.Clients{
		RegisteredClients: make(map[string]clients.Client),
	}
	defaultHTTPClient := clientsRegistry.NewClient(clients.DefaultTimeout, clients.DefaultCacheTimeout, clients.DefaultClientProtocol)
	pubChemHTTPClient := clientsRegistry.NewClient(clients.PubChemTimeout, clients.PubChemCacheTimeout, clients.PubChemClientProtocol)

	clientsRegistry.Register("default", defaultHTTPClient)
	clientsRegistry.Register("pubChem", pubChemHTTPClient)

	newScanner := bufio.NewScanner(os.Stdin)

	c := commands.Commands{
		RegisteredCommands: make(map[string]commands.Command),
	}

	pubChemQueryBuilder := commands.PubChemQueryBuilder{}
	help := commands.Help{
		Commands: &c,
	}
	showTable := commands.ShowTable{}
	fetchData := commands.FetchJSONFromURL{}

	c.Register(pubChemQueryBuilder.Name(), &pubChemQueryBuilder)
	c.Register(help.Name(), &help)
	c.Register(showTable.Name(), &showTable)
	c.Register(fetchData.Name(), &fetchData)

	cfg := config.Config{
		APIClients: clientsRegistry.RegisteredClients,
		Scanner:    newScanner,
	}
	args := os.Args[1:]

	if len(args) > 0 {
		commandName := args[0]
		err := c.Run(&cfg, commandName, args...)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		repl(&cfg, &c)
	}
}
