package rmq

import (
	"context"
	"fmt"
	"go-rabbitmq/config"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewAMQPConnection(conf config.Configuration) (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial(conf.Rabbit.URI)
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = c
			break
		}
		if counts > 3 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}
	return connection, nil
}

func PublishMessage(ctx context.Context, ch *amqp.Channel, q amqp.Queue, msg []byte) error {
	return ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})
}
