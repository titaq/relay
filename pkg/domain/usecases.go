package domain

import "context"

type RelayUsecase interface {
	RelayIncoming(context.Context, Client, []byte) error
	RelayOutgoing(ctx context.Context, message *OutgoingEvent) error
}
