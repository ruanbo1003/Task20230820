
* list all topics
```shell
root@f566379aedfd:/# kafka-topics.sh --bootstrap-server localhost:9092 --list
dev-topic
root@f566379aedfd:/#
```
* describe the specific topic
```shell
root@f566379aedfd:/# kafka-topics.sh --describe --bootstrap-server localhost:9092 --topic dev-topic
Topic: dev-topic	TopicId: NE7CPde8SWKWyoJTjMiZVg	PartitionCount: 2	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: dev-topic	Partition: 0	Leader: 1	Replicas: 1	Isr: 1
	Topic: dev-topic	Partition: 1	Leader: 1	Replicas: 1	Isr: 1
root@f566379aedfd:/#
```
* describe the specific consumer group
```shell
root@f566379aedfd:/opt/kafka/bin# kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group dev-group

GROUP           TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID                                                                                                HOST            CLIENT-ID
dev-group       dev-topic       0          1               1               0               consultation@geelydeMacBook-Pro.local (github.com/segmentio/kafka-go)-28aeb830-c082-421c-bb43-6b99c59c05ae /172.18.0.1     consultation@geelydeMacBook-Pro.local (github.com/segmentio/kafka-go)
dev-group       dev-topic       1          -               0               -               consultation@geelydeMacBook-Pro.local (github.com/segmentio/kafka-go)-28aeb830-c082-421c-bb43-6b99c59c05ae /172.18.0.1     consultation@geelydeMacBook-Pro.local (github.com/segmentio/kafka-go)
root@f566379aedfd:/opt/kafka/bin#
```


