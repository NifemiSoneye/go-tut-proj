package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GETAPIKEY extract an API Key from headers of an api key
//Example:
//Authorization : APIKey (insert apiKey here)
func GetApiKey (headers http.Header) (string,error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "" , errors.New(("no authentication info found"))


	}
	vals:= strings.Split(val , " ")
	if len(vals) != 2 {
		return "" , errors.New("malformed with auth header")
	}

	if vals[0] != "ApiKey" {
		return "" , errors.New("malformed first part of auth header")
	}

	return vals[1] , nil
}