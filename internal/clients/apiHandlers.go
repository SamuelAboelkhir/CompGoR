package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		var responseObject T
		return responseObject, err
	}

	var responseObject T
	if err := json.Unmarshal(data, &responseObject); err != nil {
		return responseObject, err
	}

	c.cache.Add(url, data)
	return responseObject, nil
}
