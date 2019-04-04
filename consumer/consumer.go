package consumer

import (
	"fmt"
	"sync"

	"bitbucket.org/animor/vetstoria/configuration"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

//ConsumeQueue consume
func ConsumeQueue(rc configuration.RabbitMqConnection, queueName string, wg *sync.WaitGroup) {
	replies, err := rc.Channel.Consume(
		queueName, // queue
		fmt.Sprintf("%s_%s", queueName, "consumer"), // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		wg.Done()
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Error received")
	}

	for r := range replies {
		log.WithFields(log.Fields{
			"data":  string(r.Body),
			"queue": queueName,
		}).Info("Received Message")
	}
	wg.Done()
}

//StartConsumer start consumer
func StartConsumer(rc configuration.RabbitMqConnection, queueName string) (<-chan amqp.Delivery, error) {
	return rc.Channel.Consume(
		queueName, // queue
		fmt.Sprintf("%s_%s", queueName, "consumer"), // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

//GetMessageProcessor start a message processor thread
func GetMessageProcessor(in <-chan string, queueName string, workerID int, wg *sync.WaitGroup) {
	for {
		message := <-in

		log.WithFields(log.Fields{
			"data":   message,
			"queue":  queueName,
			"worker": workerID,
		}).Info("Received Message")

		wg.Done()
	}
}
