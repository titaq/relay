package usecases

import (
	"context"

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
		conn, err := r.connections.Get(client)
		if err != nil {
			continue
		}
		conn.Send(message.GetData())
	}
	return nil
}

func (r *relay) AddClient(ctx context.Context, client domain.Client) {
	r.connections.Put(client)
	r.publisher.Publish(ctx, &domain.IncomingEvent{
		ClientId: client.GetId(),
		Type:     domain.IncomingEvent_CONNECT,
	})
}

func (r *relay) DeleteClient(ctx context.Context, id string) {
	r.connections.Delete(id)
	r.publisher.Publish(ctx, &domain.IncomingEvent{
		ClientId: id,
		Type:     domain.IncomingEvent_DISCONNECT,
	})
}

func NewRelay(publisher infrastructure.Messaging, connections infrastructure.Connections) domain.RelayUsecase {
	return &relay{
		publisher:   publisher,
		connections: connections,
	}
}
