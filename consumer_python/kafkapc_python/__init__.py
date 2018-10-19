from kafka import KafkaConsumer, KafkaProducer
import json

class Consumer(KafkaConsumer):
    def __init__(self, topicName, kafkaPort, consumerGroup):
        KafkaConsumer.__init__(
            self,
            topicName,
            bootstrap_servers=[kafkaPort],
            auto_offset_reset='earliest',
            enable_auto_commit=True,
            group_id=consumerGroup,
            value_deserializer=lambda x: json.loads(x.decode('utf-8'))
            )

class Producer(KafkaProducer):
    def __init__(self, kafkaPort):
        KafkaProducer.__init__(
            self,
            bootstrap_servers=[kafkaPort],
            value_serializer=lambda x: json.dumps(x).encode('utf-8'),
            key_serializer=str.encode
            )