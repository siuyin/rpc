// Package rpc provide simple rpc structs
package rpc

// Request is an rpc request.
type Request struct {
	Request string
	Params  []string
	ID      string
	Ver     string
}

// Response is an rpc response
type Response []string
