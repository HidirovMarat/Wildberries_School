package nats

import (
	"log"

	"github.com/nats-io/stan.go"
)

func New(clusterID, clientID, natsURL string) stan.Conn {
	// Подключение к NATS Streaming Server

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Ошибка при подключении к NATS Streaming Server: %v", err)
	}
	log.Print("good nats")
	return sc
}
