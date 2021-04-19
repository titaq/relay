package infrastructure

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	ceclient "github.com/cloudevents/sdk-go/v2/client"
	"github.com/sirupsen/logrus"
	"github.com/titaq/relay/pkg/domain"
	"google.golang.org/protobuf/proto"
)

type Messaging interface {
	Publish(context.Context, *domain.IncomingEvent) error
	Receive(context.Context, chan<- *domain.OutgoingEvent) error
}

func NewMessaging(cloudEventsClient ceclient.Client) Messaging {
	return &messaging{
		cloudEventsClient: cloudEventsClient,
	}
}

type messaging struct {
	cloudEventsClient ceclient.Client
}

func (m *messaging) Publish(ctx context.Context, e *domain.IncomingEvent) error {
	data, err := proto.Marshal(e)
	if err != nil {
		return err
	}

	event := cloudevents.NewEvent()
	event.SetSource("/titaq/relay")
	event.SetData(cloudevents.TextPlain, data)
	event.SetType("titaq.incomingevent")

	return m.cloudEventsClient.Send(ctx, event)
}

func (m *messaging) Receive(ctx context.Context, ch chan<- *domain.OutgoingEvent) error {
	return m.cloudEventsClient.StartReceiver(ctx, func(e cloudevents.Event) {
		ie := new(domain.OutgoingEvent)
		if err := proto.Unmarshal(e.Data(), ie); err != nil {
			logrus.Warnf("Couldn't unmarshal proto : %s", err.Error())
			return
		}

		ch <- ie
	})
}
