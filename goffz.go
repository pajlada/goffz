package goffz

import "github.com/dankeroni/jsonapi"

type errorResponse struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// FFZAPI struct
type FFZAPI struct {
	jsonapi.JSONAPI
}

// New instantiates a new FFZAPI object
func New() *FFZAPI {
	return &FFZAPI{
		JSONAPI: jsonapi.JSONAPI{
			BaseURL: "https://api.frankerfacez.com/v1",
		},
	}
}
