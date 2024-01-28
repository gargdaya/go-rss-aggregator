package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey <key>
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing Authorization header")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey"{
		return "", errors.New("invalid Authorization header")
	}
	
	return vals[1], nil
}