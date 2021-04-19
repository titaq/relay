package domain

import "context"

type RelayUsecase interface {
	RelayIncoming(context.Context, Client, []byte) error
	RelayOutgoing(context.Context, *OutgoingEvent) error
	AddClient(context.Context, Client)
	DeleteClient(context.Context, string)
}
