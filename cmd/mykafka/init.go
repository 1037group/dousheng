package mykafka

import (
	"fmt"
	"github.com/1037group/dousheng/pkg/consts"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var KafukaProducerProducer *kafka.Producer

func Init() {
	var err error
	KafukaProducerProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": consts.KafkaHost})
	if err != nil {
		panic(err)
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range KafukaProducerProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}
