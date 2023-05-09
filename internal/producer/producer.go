package producer

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func Run(server, topic string) {
	producer := newKafkaWriter(server, topic)
	defer producer.Close()

	log.Printf("Created Producer %v\n", producer.Stats())

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		key := "1"
		msgValue := scanner.Text()
		slc := strings.Split(msgValue, ":")
		if len(slc) > 1 {
			key = slc[0]
			msgValue = slc[1]
		}

		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(msgValue),
		}
		err := producer.WriteMessages(context.Background(), msg)
		if err != nil {
			handleErr(err, "Failed to produce message")
		}
	}
}

func handleErr(err error, msg string) {
	err = errors.Wrap(err, msg)
	panic(err)
}
