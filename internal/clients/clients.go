// Package clients: This package handles all the API requests to PubChem
package clients

import (
	"net/http"
	"time"

	"github.com/SamuelAboelkhir/CompGoR/internal/cache"
)

type Client interface {
	Connect() error
	Disconnect() error
	GetProtocol() string
}

type Clients struct {
	RegisteredClients map[string]Client
}

func (c *Clients) Register(name string, client Client) {
	c.RegisteredClients[name] = client
}

func (c *Clients) NewClient(timeout, cacheTimeout time.Duration, protocol string) Client {
	switch protocol {
	case "HTTP":
		return NewHTTPClient(timeout, cacheTimeout)
	}
	return nil
}

func NewHTTPClient(timeout, cacheTimeout time.Duration) *HTTPClient {
	return &HTTPClient{
		cache: cache.NewCache(cacheTimeout),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
