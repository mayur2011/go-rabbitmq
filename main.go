package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Rabbit *amqp.Connection
}

type MessagePayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func main(){
	fmt.Println("starting...")	
	
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err !=nil{
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	app := Config{
		Rabbit: rabbitConn,
	}

	// start listening for messages
	// create consumer -- which consumes msg from the queue
	// watch the queue and consume events from the topic
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backoff = 1*time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			panic(err)
			counts++
		}else{
			connection = c
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backoff)
		continue
	}
	return connection, nil
}