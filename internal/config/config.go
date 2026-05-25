// Package config: This internal package defines the config of the app
package config

import (
	"bufio"

	"github.com/SamuelAboelkhir/CompGoR/internal/clients"
)

type Config struct {
	APIClient clients.Client
	Scanner   *bufio.Scanner
}
