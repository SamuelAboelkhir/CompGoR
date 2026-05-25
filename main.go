package main

import (
	"bufio"
	"os"
	"time"

	"github.com/SamuelAboelkhir/CompGoR/internal/clients"
	"github.com/SamuelAboelkhir/CompGoR/internal/commands"
	"github.com/SamuelAboelkhir/CompGoR/internal/config"
)

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
	c.Register(builder.Name(), &builder)
	c.Register(help.Name(), &help)

	cfg := config.Config{
		APIClient: clientsRegistry.RegisteredClients["pubChem"],
		Scanner:   newScanner,
	}

	repl(&cfg, &c)
}
