package domain

import "context"

type RelayUsecase interface {
	RelayIncoming(context.Context, Client, []byte) error
	// DeleteClient(context.Context, Client)
}
