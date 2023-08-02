package utils

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaInstance struct {
	readerConfig kafka.ReaderConfig
}

func CreateKafkaInstance(broker []string, groupId string, topic string) KafkaInstance {
	return KafkaInstance{
		readerConfig: kafka.ReaderConfig{
			Brokers:   broker,
			GroupID:   groupId,
			Topic:     topic,
			Partition: KafkaPartition,
			MinBytes:  KafkaMinBytes,
			MaxBytes:  KafkaMaxBytes,
			MaxWait:   time.Duration(KafkaMaxWait) * time.Second,
		},
	}
}

func (ki KafkaInstance) StartReader(data chan<- kafka.Message, stop <-chan bool, done chan<- bool) {
	// create reader kafka
	reader := kafka.NewReader(ki.readerConfig)
	go ki.readerLoop(reader, data, stop, done)
}

func (ki KafkaInstance) readerLoop(reader *kafka.Reader, data chan<- kafka.Message, stop <-chan bool, done chan<- bool) {
	for {
		select {
		// stop loop if channel detect termination
		case <-stop:
			Warn("readerLoop", "function stopped")
			err := reader.Close()
			if err != nil {
				Fatal(err, "readerLoop", "")
				return
			}

			// create success return by send value to channel
			done <- true
			return
		default:
			// read message kafka
			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				Error(err, "readerLoop", "error read message")
				return
			}

			// send kafka message to channel
			data <- message
			Debug("readerLoop", message)
		}
	}
}
