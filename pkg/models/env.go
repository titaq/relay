package models

type Configuration struct {
	AmqpUri       string `env:"AMQP_URI" env-default:"localhost"`
	IncomingTopic string `env:"TOPIC_INCOMING" env-default:"relay-in"`
	OutgoingTopic string `env:"TOPIC_OUTGOING" env-default:"relay-out"`
}
