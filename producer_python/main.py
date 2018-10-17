from kafka import KafkaProducer
import json

def main():
    print("Yay Producer")
    producer = KafkaProducer(bootstrap_servers=['kafka:9092'],
                            value_serializer=lambda x:json.dumps(x).encode('utf-8'))
    for e in range(1000):
        data = {'number': e}
        producer.send('numtest',value=data)
    return

if __name__ == "__main__":
    #Call main function
    main()
