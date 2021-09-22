package rabbit

import "github.com/streadway/amqp"

func DeclareExchange(ch *amqp.Channel) {
	ch.ExchangeDeclare(
		"coversation.incoming",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}
