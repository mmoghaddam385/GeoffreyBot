package api

import (
	"strconv"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"errors"
)

type StringResult struct {
	Result string
	Err error
}

// getBodyJson takes an http.Response pointer and reads all of its data into a map
func getBodyJson(response *http.Response) (map[string] string, error) {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	
	if (err != nil) {
		return nil, err
	}
	
	bodyMap := make(map[string] string)

	json.Unmarshal(bodyBytes, &bodyMap)

	return bodyMap, nil
}

// buildStatusError takes an http response and builds an error with the status code and message in it
func buildStatusError(message string, response *http.Response) error {
	return errors.New(message + "; Status " + strconv.Itoa(response.StatusCode) + ": " + response.Status)
}