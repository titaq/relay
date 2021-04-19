package infrastructure

import (
	"os"

	"github.com/titaq/relay/pkg/domain"
)

type Connections interface {
	Get(string) (domain.Client, error)
	Put(domain.Client)
	Delete(string)
}

func NewConnections() Connections {
	return &connections{
		list: make(map[string]domain.Client),
	}
}

type connections struct {
	list map[string]domain.Client
}

func (c *connections) Get(id string) (domain.Client, error) {
	conn, ok := c.list[id]
	if !ok {
		return nil, os.ErrNotExist
	}
	return conn, nil
}

func (c *connections) Put(client domain.Client) {
	c.list[client.GetId()] = client
}

func (c *connections) Delete(id string) {
	delete(c.list, id)
}
