from kafka import KafkaConsumer, KafkaProducer
import json

class Consumer(KafkaConsumer):
    def __init__(self, topicName, kafkaPort):
        KafkaConsumer.__init__(
            self,
            topicName,
            bootstrap_servers=[kafkaPort],
            group_id='my-group',
            value_deserializer=lambda x: json.loads(x.decode('utf-8'))
            )
        print("completed initialization")    


