package auth

import "net/http"

// GETAPIKEY extract an API Key from headers of an api key
//Example:
//Authorization : APIKey (insert apiKey here)
func GetApiKey (headers http.Header) (string,error)