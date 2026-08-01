package commands

import (
	"fmt"

	"github.com/SamuelAboelkhir/CompGoR/internal/clients"
	"github.com/SamuelAboelkhir/CompGoR/internal/config"
)

// QueryConstructor represents a query to the PubChem API
type QueryConstructor struct {
	input struct {
		domain string
		// TODO: Expand with <structure search>
		namespace   string
		identifiers string
	}
	operation *string
	// NOTE: JSONP can be a callback. Needs investigation
	output string
}

// QueryBuilder is a command that builds a query to the PubChem API
// and allows the user to interactively build and execute a query to the PubChem API
type QueryBuilder struct{}

// Execute implements the Command interface
func (q *QueryBuilder) Execute(cfg *config.Config, args ...string) error {
	err := queryCommandHandler(cfg, args...)
	if err != nil {
		return err
	}
	return nil
}

// Name returns the name of the command
func (q *QueryBuilder) Name() string {
	return "buildQuery"
}

// Help returns the help text for the command
func (q *QueryBuilder) Help() string {
	helpString := `Takes a domain, namespace, identifier, an optional operation, 
	and an output type, and queries the PubChem API for a matching 
	chemical substance or compound
	Example: domain: compound, 
		 namespace: name, 
		 identifier: hydrogen, 
		 operation: <optional, ENTER to skip>, 
		 format: JSON`
	return helpString
}

// TODO: Add support for multiple identifiers
//
// `queryCommandHandler` handles the query command by prompting the user for input and executing the query
func queryCommandHandler(cfg *config.Config, args ...string) error {
	steps := []string{"domain: ", "namespace: ", "identifier: ", "operation (optional): ", "output format: "}
	query := QueryConstructor{}
	for _, step := range steps {
		buildQuery(cfg, &query, step)
		fmt.Println(query)
	}

	c, ok := cfg.APIClient.(*clients.HTTPClient)
	if ok {
		elements, err := c.GetCompounds(query.input.domain, query.input.namespace, query.input.identifiers, query.output)
		if err != nil {
			return err
		}
		for _, element := range elements.PCCompounds {
			fmt.Println("Element ID: ", element.ID)
			fmt.Println("Props: ", element.Props)
			for _, atom := range element.Atoms.Element {
				fmt.Println("Atom: ", atom)
			}
		}
	}
	return nil
}

// `buildQuery` builds a query to the PubChem API by prompting the user for input
func buildQuery(cfg *config.Config, query *QueryConstructor, step string) {
	fmt.Printf("Please provide a %s", step)
	cfg.Scanner.Scan()
	param := cfg.Scanner.Text()
	switch step {
	case "domain: ":
		query.input.domain = param
	case "namespace: ":
		query.input.namespace = param
	case "identifier: ":
		query.input.identifiers = param
	case "operation (optional): ":
		if len(param) > 0 {
			query.operation = &param
		} else {
			query.operation = nil
		}
	case "output format: ":
		query.output = param
	}
}
