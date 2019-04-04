package main

import (
	"sync"

	"bitbucket.org/animor/vetstoria/consumer"

	"bitbucket.org/animor/vetstoria/configuration"
	"bitbucket.org/animor/vetstoria/tokenmanager"

	log "github.com/sirupsen/logrus"
)

func main() {
	var wg sync.WaitGroup
	token := tokenmanager.ReadToken()

	log.WithFields(log.Fields{
		"token": token,
	}).Info("Token in main")

	rabbitClient := configuration.GetRabbitClient()

	fatalErrors := make(chan error, 5)
	log.Info("Start post consumer")
	rabbitPost, errPost := consumer.StartConsumer(*rabbitClient, "vetstoria_post")
	workPost := make(chan string, 1024)
	log.Info("Start put consumer")
	rabbitPut, errPut := consumer.StartConsumer(*rabbitClient, "vetstoria_put")
	workPut := make(chan string, 1024)
	workerCount := 100

	if errPost != nil {
		fatalErrors <- errPost
	}

	if errPut != nil {
		fatalErrors <- errPut
	}

	for i := 0; i < workerCount; i++ {
		go consumer.GetMessageProcessor(workPost, "vetstoria_post", i, &wg)
		go consumer.GetMessageProcessor(workPut, "vetstoria_put", i, &wg)
	}

	for item := range rabbitPost {
		go func() {
			wg.Add(1)
			workPost <- string(item.Body)
		}()
	}

	for item := range rabbitPut {
		go func() {
			wg.Add(1)
			workPut <- string(item.Body)
		}()
	}

	wg.Wait()
}
