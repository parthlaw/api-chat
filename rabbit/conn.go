package rabbit

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func Connection() (*amqp.Connection, error) {
	fmt.Println(os.Getenv("RABBIT_URI"))
	conn, err := amqp.Dial(os.Getenv("RABBIT_URI"))
	return conn, err
}
