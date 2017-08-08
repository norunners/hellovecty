package dao

import (
	"github.com/norunners/hellovecty/api"
	"net/rpc"
)

// Dao abstracts away the rpc client details.
type Dao struct {
	client *rpc.Client
}

// New creats a Dao from a given rpc client.
func New(client *rpc.Client) *Dao {
	return &Dao{client: client}
}

// Add calls the service and returns the result.
func (dao *Dao) Add(a, b int) (int, error) {
	in := &api.AddIn{A: a, B: b}
	out := &api.AddOut{}
	var err error
	err = dao.client.Call("Service.Add", in, out)
	if err != nil {
		return 0, err
	}
	return out.Sum, nil
}
