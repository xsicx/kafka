package consumer

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

func Run(server, topic, group string) {
	bootstrapServers := []string{server}

	consumer := getKafkaReader(bootstrapServers, topic, group)
	defer consumer.Close()

	log.Printf("Created Consumer %+v\n", consumer.Stats())

	for {
		message, err := consumer.ReadMessage(context.Background())
		if err != nil {
			handleErr(err, "Failed to consume message")
		}
		log.Printf("[MESSAGE] %s\n", string(message.Value))
	}
}
func getKafkaReader(bootstrapServers []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  bootstrapServers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func handleErr(err error, msg string) {
	err = errors.Wrap(err, msg)
	panic(err)
}
