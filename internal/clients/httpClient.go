package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/SamuelAboelkhir/CompGoR/internal/cache"
)

// HTTPClient is a client that fetches data from a HTTP API
type HTTPClient struct {
	cache      *cache.Cache
	httpClient http.Client
}

// GetCompounds fetches compounds from the HTTP API
func (c *HTTPClient) GetCompounds(domain, namespace, identifier, output string) (Compounds, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s", BasePubChemURL, domain, namespace, identifier, output)
	compound, err := httpAPIHandler[Compounds](c, url)
	if err != nil {
		return Compounds{}, err
	}
	return compound, nil
}

func (c *HTTPClient) FetchJSONData(url string) (any, error) {
	data, err := httpAPIHandler[any](c, url)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// TODO: Implement Connect

// Connect initializes the HTTP client connection by verifying connectivity to the common URL
func (c *HTTPClient) Connect() error {
	return nil
}

// TODO: Implement Disconnect

// Disconnect closes the HTTP client connection
func (c *HTTPClient) Disconnect() error {
	return nil
}

// GetProtocol returns the protocol used by the client
func (c *HTTPClient) GetProtocol() string {
	return "HTTP"
}

// NewHTTPClient creates a new HTTP client with the specified timeout and cache timeout
func NewHTTPClient(timeout, cacheTimeout time.Duration) *HTTPClient {
	return &HTTPClient{
		cache: cache.NewCache(cacheTimeout),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

// httpAPIHandler is a generic function to handle HTTP API requests and cache the responses.
func httpAPIHandler[T any](c *HTTPClient, url string) (T, error) {
	fmt.Println(url)

	if data, ok := c.cache.Get(url); ok {
		var responseObject T
		err := json.Unmarshal(data, &responseObject)
		if err != nil {
			return responseObject, err
		}
		return responseObject, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		var responseObject T
		return responseObject, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		var responseObject T
		return responseObject, err
	}

	if res.Header.Get("Content-Type") != "application/json" {
		var responseObject T
		return responseObject, errors.New("invalid content type, JSON content expected")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		var responseObject T
		return responseObject, err
	}

	var responseObject T
	if err := json.Unmarshal(data, &responseObject); err != nil {
		return responseObject, err
	}

	checkResponse := func(responseObject T) bool {
		switch v := any(responseObject).(type) {
		case Compounds:
			return v.PCCompounds == nil
		default:
			return false
		}
	}

	if checkResponse(responseObject) {
		return responseObject, errors.New("failed to fill the response object")
	}

	c.cache.Add(url, data)
	return responseObject, nil
}
