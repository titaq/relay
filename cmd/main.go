package main

import (
	"context"
	"log"

	"github.com/Azure/go-amqp"
	ceamqp "github.com/cloudevents/sdk-go/protocol/amqp/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/sirupsen/logrus"
	"github.com/titaq/relay/pkg/infrastructure"
	"github.com/titaq/relay/pkg/presentation"
	"github.com/titaq/relay/pkg/usecases"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Starting titaq relay")

	p, err := ceamqp.NewProtocol("localhost", "titaq", []amqp.ConnOption{}, []amqp.SessionOption{})
	if err != nil {
		log.Fatalf("Failed to create amqp protocol: %v", err)
	}
	defer p.Close(context.Background())

	c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		panic(err)
	}

	publisher := infrastructure.NewMessaging(c)
	relay := usecases.NewRelay(publisher)

	server := presentation.NewTCPServer(relay)
	server.ListenAndServe("tcp4", ":4589")
}
