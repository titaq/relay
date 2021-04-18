package infrastructure

import "github.com/titaq/relay/pkg/domain"

type Connections interface {
	Get(string) domain.Client
	Put(domain.Client)
	Delete(string)
}
