// Package clients: This package provides a generic client interface and management system for various protocol clients.
package clients

import (
	"time"
)

// Client is a generic client interface for all types of clients
type Client interface {
	Connect() error
	Disconnect() error
	GetProtocol() string
}

// Clients struct holds a map of registered clients
type Clients struct {
	// RegisteredClients is a map of client names to client instances
	RegisteredClients map[string]Client
}

// Register adds a new client to the RegisteredClients map
func (c *Clients) Register(name string, client Client) {
	c.RegisteredClients[name] = client
}

// NewClient creates a new client based on the protocol specified
func (c *Clients) NewClient(timeout, cacheTimeout time.Duration, protocol string) Client {
	switch protocol {
	case "HTTP":
		return NewHTTPClient(timeout, cacheTimeout)
	}
	return nil
}
