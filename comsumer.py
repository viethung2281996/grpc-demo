from confluent_kafka.cimpl import Consumer


class BaseConsumer:
    def __init__(self,
                 topic: object,
                 group_id: object,
                 auto_offset_reset: object = "latest",
                 session_timeout_ms: object = 30000):
        self.topic = topic

        c = Consumer(
            {
                "bootstrap.servers": '127.0.0.1:9092',
                "group.id": group_id,
                "auto.offset.reset": auto_offset_reset,
                "enable.auto.commit": True,
                "session.timeout.ms": session_timeout_ms,
                "broker.address.family": "v4",
                'debug': 'msg'
            }
        )

        c.subscribe([topic])
        print('Start comsume message')
        try:
            while True:
                msg = c.poll(1.0)
                if msg is None:
                    continue
                if msg.error():
                    print('Fail to receive message')
                # json_obj = json.loads(msg.value().decode("utf-8"))
                self.process(msg.value())

                # c.commit()
        except Exception:
            print("Fail to comsume message")
        finally:
            c.close()

    def process(self, message):
        try:
            self.message_handle(message=message)
        except Exception:
            print("Fail to process message")

    def message_handle(self, message):
        pass


class TestComsumer(BaseConsumer):
    def message_handle(self, message):
        # cmd = json_obj.get("cmd")
        print(message)


if __name__ == '__main__':
    comsumer = TestComsumer('test', 'test')
