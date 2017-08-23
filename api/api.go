// Package api defines state and behavior of the api.
package api

// Service defines operations in an rpc compatible format.
type Service interface {
	Add(in *AddIn, out *AddOut) error
}

// AddIn is the input for the add operation.
type AddIn struct {
	A, B int
}

// AddOut is the output for the add operation.
type AddOut struct {
	Sum int
}
