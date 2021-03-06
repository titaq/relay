package main

import (
	"context"
	"log"

	"github.com/Azure/go-amqp"
	ceamqp "github.com/cloudevents/sdk-go/protocol/amqp/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"github.com/titaq/relay/pkg/domain"
	"github.com/titaq/relay/pkg/infrastructure"
	"github.com/titaq/relay/pkg/models"
	"github.com/titaq/relay/pkg/presentation"
	"github.com/titaq/relay/pkg/usecases"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Starting titaq relay")

	var cfg models.Configuration
	cleanenv.ReadEnv(&cfg)

	connectionsCache := infrastructure.NewConnections()

	forever := make(chan bool)
	go func() {
		logrus.Info("Starting titaq relay-server")

		p, err := ceamqp.NewProtocol(cfg.AmqpUri, cfg.IncomingTopic, []amqp.ConnOption{}, []amqp.SessionOption{})
		if err != nil {
			log.Fatalf("Failed to create amqp protocol: %v", err)
		}
		defer p.Close(context.Background())

		c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
		if err != nil {
			panic(err)
		}

		publisher := infrastructure.NewMessaging(c)
		relay := usecases.NewRelay(publisher, connectionsCache)

		server := presentation.NewTCPServer(relay)
		server.ListenAndServe("tcp4", ":23000")
	}()

	go func() {
		logrus.Info("Starting titaq relay-receptor")

		p, err := ceamqp.NewProtocol(cfg.AmqpUri, cfg.OutgoingTopic, []amqp.ConnOption{}, []amqp.SessionOption{})
		if err != nil {
			log.Fatalf("Failed to create amqp protocol: %v", err)
		}
		defer p.Close(context.Background())

		c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
		if err != nil {
			panic(err)
		}

		outgoingEvents := make(chan *domain.OutgoingEvent)
		publisher := infrastructure.NewMessaging(c)
		relay := usecases.NewRelay(publisher, connectionsCache)

		go func() {
			logrus.Info("Waiting for events")
			for {
				outgoingEvent := <-outgoingEvents
				relay.RelayOutgoing(context.Background(), outgoingEvent)
			}
		}()

		logrus.Fatal(publisher.Receive(context.Background(), outgoingEvents))
	}()
	<-forever
}
