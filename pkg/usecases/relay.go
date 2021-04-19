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
		conn := r.connections.Get(client)
		conn.Send(message.GetData())
	}
	return nil
}

func NewRelay(publisher infrastructure.Messaging) domain.RelayUsecase {
	return &relay{
		publisher: publisher,
	}
}
