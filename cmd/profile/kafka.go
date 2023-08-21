package main

import (
	"HeidiTask/internal"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

var (
	writer *kafka.Writer
)

func init() {
	// kafka topic initialization
	internal.CreateKafkaTopic()

	// init writer
	writer = internal.NewWriter()
	//conn, _ = kafka.DialLeader(context.Background(), "tcp", config.KafkaHost, config.KafkaTopic, 0)
}

// send one message to kafka
func sendMessage(msg interface{}) {
	jsonMsg, _ := json.Marshal(msg)
	err := writer.WriteMessages(
		context.Background(),
		kafka.Message{Value: jsonMsg},
	)
	//_, err := writer.WriteMessages(
	//	kafka.Message{Value: jsonMsg},
	//	//context.Background(),
	//	//kafka.Message{
	//	//	Value: jsonMsg,
	//	//}
	//)
	if err != nil {
		fmt.Println("sendMessage:", err.Error())
	}
}

// send slice of internal.ProfileUpdateData to kafka
func sendUpdateKafkaMessages(messages []internal.ProfileUpdateData) {
	kafkaEvents := internal.ProfileUpdateEvents{
		Events: messages,
	}
	sendMessage(kafkaEvents)
}
