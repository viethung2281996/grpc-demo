from confluent_kafka.cimpl import Producer

from src.utils import compress


class KafkaProducer(object):
    _instance = None

    @classmethod
    def instance(cls):
        if cls._instance is None:
            cls._instance = Producer(
                {
                    "request.timeout.ms": 20000,
                    "bootstrap.servers": '127.0.0.1:9092',
                    "broker.address.family": "v4",
                    # "compression.type": "snappy",
                    "linger.ms": 5,
                    'debug': 'msg',
                }
            )
        return cls._instance

    def flush_message(self, topic: str, key: str, value):
        self.instance().produce(
            topic=topic,
            key=key,
            value=value,
            callback=self.kafka_delivery_report,
        )
        # result = self.instance().poll(0)
        self.instance().flush()

    def kafka_delivery_report(self, err, msg):
        """ Called once for each message produced to indicate delivery result.
            Triggered by poll() or flush(). """
        if err is not None:
            print("Message delivery failed: {}".format(err))
        print("Delivery message success")


if __name__ == '__main__':
    producer = KafkaProducer()
    with open('data.json', 'rb') as f:
        message = f.read()
    encoded_message = compress(message, 'snappy')
    print(f"Produce message {encoded_message}")
    print(f'Message type {type(encoded_message)}')
    producer.flush_message('test', 'test', encoded_message)
