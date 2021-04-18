package domain

type Client interface {
	GetId() string
	Send([]byte) error
}
