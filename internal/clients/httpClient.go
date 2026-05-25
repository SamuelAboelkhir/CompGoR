package clients

import (
	"fmt"
	"net/http"

	"github.com/SamuelAboelkhir/CompGoR/internal/cache"
)

type HTTPClient struct {
	cache      *cache.Cache
	httpClient http.Client
}

func (c *HTTPClient) GetCompounds(domain, namespace, identifier, output string) (Compounds, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s", commonURL, domain, namespace, identifier, output)
	compound, err := httpAPIHandler[Compounds](c, url)
	if err != nil {
		return Compounds{}, err
	}
	return compound, nil
}

// TODO: Implement Connect
func (c *HTTPClient) Connect() error {
	return nil
}

// TODO: Implement Disconnect
func (c *HTTPClient) Disconnect() error {
	return nil
}

func (c *HTTPClient) GetProtocol() string {
	return "HTTP"
}
