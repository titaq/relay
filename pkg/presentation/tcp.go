package presentation

import (
	"github.com/goinsane/accepter"
	"github.com/titaq/relay/pkg/domain"
)

type RelayServer interface {
	ListenAndServe(network, address string) error
}

func NewTCPServer(relay domain.RelayUsecase) RelayServer {
	server := &tcpServer{
		relay: relay,
	}
	server.a = &accepter.Accepter{
		Handler: server,
	}

	return server
}
