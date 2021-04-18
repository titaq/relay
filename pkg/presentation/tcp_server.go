package presentation

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net"

	"github.com/goinsane/accepter"
	"github.com/sirupsen/logrus"
	"github.com/titaq/relay/pkg/domain"
)

type tcpServer struct {
	a     *accepter.Accepter
	relay domain.RelayUsecase
}

func (s *tcpServer) ListenAndServe(network, address string) error {
	return s.a.ListenAndServe(network, address)
}

func (s *tcpServer) Serve(ctx context.Context, conn net.Conn) {
	logrus.Info("Client connected !")

	remoteAddrHash := sha256.Sum256(([]byte)(conn.RemoteAddr().String()))
	clientId := fmt.Sprintf("%x", remoteAddrHash[:])
	logrus.Infof("Assigned id %s", clientId)

	client := &tcpClient{
		conn: conn,
		id:   clientId,
	}

	for {
		var b [32 * 1024]byte
		n, err := conn.Read(b[:])
		if err != nil {
			break
		}

		err = s.relay.RelayIncoming(ctx, client, b[:n])
		if err != nil {
			logrus.Error(err)
		}
		logrus.Debugf("Received a message from client %s", clientId)
	}
}
