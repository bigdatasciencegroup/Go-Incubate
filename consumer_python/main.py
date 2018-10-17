import kafka
from kafka import KafkaConsumer
import json

def main():
    print("Yay Consumer")
    print(kafka.__version__)
    consumer = KafkaConsumer(
        'numtest',
        bootstrap_servers=['kafka:9092'],
        auto_offset_reset='earliest',
        enable_auto_commit=True,
        group_id='my-group',
        value_deserializer=lambda x: json.loads(x.decode('utf-8'))
    )

    for message in consumer:
        print(message)
        val = message.value
        print(val)

    return

if __name__ == "__main__":
    #Call main function
    main()
