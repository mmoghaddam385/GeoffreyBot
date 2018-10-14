package endpoints

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

var endpoints = map[string] HandlerFunc {
	"/" : fallbackHandler,
	"/message" : messageRecieved,
}

// GetAllEndpoints returns all registered endpoints in a map
// The map contains the String path of the endpoint mapped to its handler func
func GetAllEndpoints() map[string] HandlerFunc {
	return endpoints
}