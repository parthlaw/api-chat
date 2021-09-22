package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func RoomDeclare(ch *amqp.Channel, name string) {
	ch.ExchangeDeclare(
		name,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	// ch.ExchangeDeclare(
	// 	"coversation.incoming",
	// 	"topic",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	q, e := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	err := ch.ExchangeBind(name, name, "coversation.incoming", false, nil)
	if err != nil {
		fmt.Println(err)
	}
	if e != nil {
		fmt.Println("Failed to declare queue")
	}
	ch.QueueBind(q.Name, "", name, false, nil)
}

func (room *Room) Consume(client *Client) {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error in creating websocket channel")
	}
	msgs, err := ch.Consume(room.Name, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println("Error in reading messages")
	}
	forever := make(chan bool)
	go func() {

		for d := range msgs {
			var message Message
			if err := json.Unmarshal(d.Body, &message); err != nil {
				fmt.Printf("Error on unmarshel JSON message %s", err)
			}
			roomID := message.Target.GetId()
			fmt.Println(message.Target)
			if rom := client.wsServer.findRoomByID(roomID); room != nil {
				rom.broadcast <- &message
			}
		}
	}()
	<-forever
}
