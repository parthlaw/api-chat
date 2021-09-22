package rabbit

import (
	"api-chat/utils"

	"github.com/streadway/amqp"
)

func QueueDeclare(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"backend_messages",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Fail to declare the queue")
	return q
}
