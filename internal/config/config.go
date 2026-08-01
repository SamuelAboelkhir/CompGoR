// Package config: This internal package defines the config of the app
package config

import (
	"bufio"

	"github.com/SamuelAboelkhir/CompGoR/internal/clients"
)

// Config holds the configuration of the app
type Config struct {
	APIClient clients.Client
	Scanner   *bufio.Scanner
}
