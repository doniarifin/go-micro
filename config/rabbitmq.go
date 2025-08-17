package rabbitmq

import (
	"context"
	"encoding/json"
	"go-micro/model"
	"go-micro/service"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var Connection *amqp.Connection

func Connect() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	Connection = conn

	return err
}

func Publish(queue string, msg string) error {
	ch, err := Connection.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := msg
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

	return err
}

func Consume(queue string, db *gorm.DB) error {
	ch, err := Connection.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	// var forever chan struct{}
	goChan := make(chan string)

	go func() {
		for d := range msgs {
			var message *model.Message
			if err := json.Unmarshal(d.Body, &message); err != nil {
				continue
			}
			msg := model.NewMsgRepository(db)
			msgService := service.NewMsgService(msg)
			if err := msgService.Insert(message); err != nil {
				continue
			}
			log.Printf(" [x] %s", d.Body)
			goChan <- string(d.Body)
		}
	}()

	// log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-goChan
	return nil
}

func CloseConnection() {
	if Connection != nil {
		Connection.Close()
	}
}
