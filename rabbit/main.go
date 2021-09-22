package rabbit

import (
	"encoding/json"
	"fmt"
	"reflect"

	"api-chat/websocket"

	"github.com/streadway/amqp"
)

type MessageType struct {
	Recieved uint //userId or RoomId
	Sent     uint
	Text     string
	Attach   string
}

var BufferPersonal []websocket.Message

type Conn struct {
	conn *amqp.Connection
}

func NewConn(conn *amqp.Connection) *Conn {
	return &Conn{
		conn: conn,
	}
}
func (conn *Conn) HandleMessages() {
	//conn, err := Connection()
	//utils.FailOnError(err, "Failed to connected to RabbitMQ")
	msg := Recieve(conn.conn)
	forever := make(chan bool)
	go func() {
		for d := range msg {
			fmt.Printf("Recieved Message: %s\n", reflect.TypeOf(d.Body))
			var message websocket.Message
			json.Unmarshal(d.Body, &message)
			BufferPersonal = append(BufferPersonal, message)
			if len(BufferPersonal) >= 10 {
				//gorm bulk create
				fmt.Println("Bulk Create")
				BufferPersonal = []websocket.Message{}
			}
			name := message.Target
			k := amqp.Publishing{ContentType: d.ContentType, Body: d.Body}
			fmt.Println("room", *name, "room")
			Emit(conn.conn, k, message.Target.GetName())
		}
	}()
	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
