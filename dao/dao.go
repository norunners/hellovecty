// Package dao provides a data abstraction object for client calls.
package dao

import (
	"github.com/norunners/hellovecty/api"
	"net/rpc"
)

// Dao defines add behavior.
type Dao interface {
	Add(a, b int) (int, error)
}

// dao satisfies dao with an rpc client.
type dao struct {
	client *rpc.Client
}

// New creates a dao from a given rpc client.
func New(client *rpc.Client) Dao {
	return &dao{client: client}
}

// Add calls the service and returns the result.
func (dao *dao) Add(a, b int) (int, error) {
	in := &api.AddIn{A: a, B: b}
	out := &api.AddOut{}
	var err error
	err = dao.client.Call("Service.Add", in, out)
	if err != nil {
		return 0, err
	}
	return out.Sum, nil
}
