package usecases

import (
	"context"

	"github.com/titaq/relay/pkg/domain"
	"github.com/titaq/relay/pkg/infrastructure"
)

type relay struct {
	publisher infrastructure.Messaging
}

func (r *relay) RelayIncoming(ctx context.Context, client domain.Client, message []byte) error {
	return r.publisher.Publish(ctx, &domain.IncomingEvent{
		ClientId: client.GetId(),
		Data:     message,
	})
}

func NewRelay(publisher infrastructure.Messaging) domain.RelayUsecase {
	return &relay{
		publisher: publisher,
	}
}
