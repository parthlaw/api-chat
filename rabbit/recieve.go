package rabbit

import (
	"api-chat/utils"

	"github.com/streadway/amqp"
)

func Recieve(conn *amqp.Connection) <-chan amqp.Delivery {
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	q := QueueDeclare(ch)
	e := ch.ExchangeDeclare("connection.outgoing", "fanout", true, false, false, false, nil)
	utils.FailOnError(e, "Fail to declare the exchange")
	er := ch.QueueBind(q.Name, "", "connection.outgoing", false, nil)
	utils.FailOnError(er, "Failed to bind the queue")
	msg, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return msg
}
