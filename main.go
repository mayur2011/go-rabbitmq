package main

import (
	"context"
	"encoding/json"
	"go-rabbitmq/config"
	"go-rabbitmq/model"
	"go-rabbitmq/rmq"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conf := config.GetConfig()
	// try to connect to rabbitmq
	rmqConn, err := rmq.NewAMQPConnection(conf)
	failOnError(err, "failed to connect to rmq")
	defer rmqConn.Close()
	log.Println("connected to rmq..!")

	// open a channel
	ch, err := rmqConn.Channel()
	failOnError(err, "fail to open a channel")
	defer ch.Close()

	// declare queue
	q, err := ch.QueueDeclare("gdq.order", false, false, false, false, nil)
	failOnError(err, "fail to declare a queue")

	message := model.Message{
		Sequence: 1,
		Payload:  "cedd is missing",
	}
	jsonByte, err := json.Marshal(message)
	if err != nil {
		log.Panicln(err)
	}

	ctx := context.TODO()
	if err = rmq.PublishMessage(ctx, ch, q, jsonByte); err != nil {
		failOnError(err, "fail to publish message to queue")
	}
	log.Println(message, ".. is published")

	// start listening for messages

	// create consumer

	// watch the queue and consume events
}
