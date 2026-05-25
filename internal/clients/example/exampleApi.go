// Package example: This is just an example package
package example

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func exampleAPI() error {
	req, err := http.NewRequest("GET", "http://localhost:5000/elements", nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	decode := json.NewDecoder(res.Body)

	var elements PeriodcTable
	if err := decode.Decode(&elements); err != nil {
		return err
	}
	fmt.Println(elements)
	return nil
}
