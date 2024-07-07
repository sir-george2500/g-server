package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an API key from the Authorization header of an HTTP request.
// The expected format of the header is:
// Authorization: ApiKey {insert API key here}
//
// Parameters:
// - headers (http.Header): The HTTP headers from which to extract the API key.
//
// Returns:
//   - (string, error): Returns the extracted API key as a string, or an error if the key is not found
//     or is improperly formatted.
//
// Example usage:
//
//	headers := http.Header{}
//	headers.Set("Authorization", "ApiKey abc123")
//	apiKey, err := GetApiKey(headers)
//	if err != nil {
//	  // handle error
//	}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) < 2 {

		return "", errors.New("malfound authentication header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malfound first part of auth header")
	}
	return vals[1], nil
}
