package types

import "net/http"

type Route struct {
	Method      string
	Name        string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
