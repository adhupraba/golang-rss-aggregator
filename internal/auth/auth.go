package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts an api key from the headers of an http request
// example header -> Authorization: ApiKey <api key>
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("No authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("Malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of auth header")
	}

	return vals[1], nil
}
