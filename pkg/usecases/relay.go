package usecases

import (
	"context"
	"fmt"

	"github.com/titaq/relay/pkg/domain"
	"github.com/titaq/relay/pkg/infrastructure"
)

type relay struct {
	publisher   infrastructure.Messaging
	connections infrastructure.Connections
}

func (r *relay) RelayIncoming(ctx context.Context, client domain.Client, message []byte) error {
	return r.publisher.Publish(ctx, &domain.IncomingEvent{
		ClientId: client.GetId(),
		Data:     message,
	})
}

func (r *relay) RelayOutgoing(ctx context.Context, message *domain.OutgoingEvent) error {
	for _, client := range message.GetClientsList() {
		fmt.Println(client)
		conn, err := r.connections.Get(client)
		fmt.Println(conn, err)
		if err != nil {
			continue
		}
		conn.Send(message.GetData())
	}
	return nil
}

func (r *relay) AddClient(ctx context.Context, client domain.Client) {
	r.connections.Put(client)
}

func (r *relay) DeleteClient(_ context.Context, id string) {
	r.connections.Delete(id)
}

func NewRelay(publisher infrastructure.Messaging, connections infrastructure.Connections) domain.RelayUsecase {
	return &relay{
		publisher:   publisher,
		connections: connections,
	}
}
