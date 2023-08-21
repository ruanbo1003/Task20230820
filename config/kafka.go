package config

var (
	KafkaHost    = getEnv("KAFKA_HOST", "localhost:9092")
	KafkaGroupId = getEnv("KAFKA_GROUP", "dev-group")
	KafkaTopic   = getEnv("KAFKA_TOPIC", "dev-topic")
)
