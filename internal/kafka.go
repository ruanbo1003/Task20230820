package internal

import (
	"HeidiTask/config"
	"fmt"
	kafkaSdk "github.com/segmentio/kafka-go"
	"time"
)

func NewWriter() *kafkaSdk.Writer {
	return &kafkaSdk.Writer{
		Addr:  kafkaSdk.TCP(config.KafkaHost),
		Topic: config.KafkaTopic,
	}
}

func NewReader() *kafkaSdk.Reader {
	fmt.Println("NewReader:", config.KafkaHost, config.KafkaTopic, config.KafkaGroupId)

	return kafkaSdk.NewReader(kafkaSdk.ReaderConfig{
		Brokers: []string{config.KafkaHost},
		//Dialer:   &dialer,
		GroupID:  config.KafkaGroupId,
		Topic:    config.KafkaTopic,
		MinBytes: 1,    // 1B
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Second,
	})
}

// CreateKafkaTopic create kafka topic
func CreateKafkaTopic() {
	fmt.Println("CreateKafkaTopic:", config.KafkaTopic)
	conn, err := kafkaSdk.Dial("tcp", config.KafkaHost)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	topics := listTopics(conn)
	if _, ok := topics[config.KafkaTopic]; ok {
		// topic already created
		fmt.Printf("topic %s already existed\n", config.KafkaTopic)
		return
	}

	topicConfig := []kafkaSdk.TopicConfig{
		{
			Topic:             config.KafkaTopic,
			NumPartitions:     2,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(topicConfig...)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("topic %s created successfully.\n", config.KafkaTopic)
}

func listTopics(conn *kafkaSdk.Conn) map[string]struct{} {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	topics := map[string]struct{}{}

	for _, p := range partitions {
		topics[p.Topic] = struct{}{}
	}

	return topics
}
