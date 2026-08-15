package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/SamuelAboelkhir/CompGoR/internal/clients"
	"github.com/SamuelAboelkhir/CompGoR/internal/config"
)

// FetchJSONFromURL is a command to fetch generic data of JSON format.
type FetchJSONFromURL struct{}

// Execute executes the fetchdata command.
func (g *FetchJSONFromURL) Execute(cfg *config.Config, args ...string) error {
	if len(args) <= 0 {
		return errors.New("usage: fetchdata <url>")
	}
	url := args[0]

	defaultClient, ok := cfg.APIClients["default"]

	if !ok {
		return errors.New("pubChem client not found")
	}

	c, ok := defaultClient.(*clients.HTTPClient)
	if !ok {
		return errors.New("client couldn't connect")
	}
	data, err := c.FetchJSONData(url)
	if err != nil {
		return err
	}
	formattedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(formattedData))
	return nil
}

// Name returns the name of the command.
func (g *FetchJSONFromURL) Name() string {
	return "fetchJSONData"
}

// Help returns the help text for the command.
func (g *FetchJSONFromURL) Help() string {
	return "Fetches generic data of JSON format."
}
