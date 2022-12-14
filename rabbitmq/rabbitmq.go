package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Client struct {
	connection *amqp.Connection
	queue      string
}

func New(url string, queue string) *Client {
	con, err := amqp.Dial(url)
	//defer con.Close()
	if err != nil {
		log.Println("error occurred when connecting to rabbitmq server")
		return nil
	}
	ch, err := con.Channel()
	defer ch.Close()
	_, err = ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return nil
	}
	return &Client{connection: con, queue: queue}
}
func (c *Client) PushMessage(message []byte) error {
	ch, err := c.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	context, cancelFn := context.WithTimeout(context.Background(), time.Second*40)
	defer cancelFn()
	err = ch.PublishWithContext(context,
		"", c.queue, false, false,
		amqp.Publishing{
			Headers:     nil,
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		log.Println("error occurred when publishing the message ")
		return err
	}
	return nil
}

func (c *Client) Consume(messageReceived chan []byte, consumer string) {
	ch, err := c.connection.Channel()
	defer ch.Close()
	if err != nil {
		return
	}
	msgs, err := ch.Consume(
		c.queue,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			messageReceived <- d.Body
		}
	}()
	<-forever
}
