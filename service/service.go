package service

import (
	"fmt"
	"github.com/norunners/hellovecty/api"
)

// Service satisfies the api.
// NOTE: Exported to be registered for rpc.
type Service struct {
}

// New creats a service.
func New() api.Service {
	return &Service{}
}

// Add performs the add operation.
func (service *Service) Add(in *api.AddIn, out *api.AddOut) error {
	out.Sum = in.A + in.B
	fmt.Printf("%v + %v = %v\n", in.A, in.B, out.Sum)
	return nil
}
