from confluent_kafka import Consumer
import random
import asyncio

conf = {
    'bootstrap.servers': 'localhost:9092',
    'group.id': 'foo',   
    'auto.offset.reset': 'earliest'
}
consumer = Consumer(conf)
consumer.subscribe(['quickstart-events'])
# while True:
#     msg = consumer.poll(1.0)

#     if msg is None:
#         continue
#     if msg.error():
#         print("Consumer error: {}".format(msg.error()))
#         continue

#     print('Received message: {}'.format(msg.value().decode('utf-8')))

class c:
    
    def __init__(self) -> None:
        
        self.__consumer = Consumer({
            'bootstrap.servers': 'localhost:9092',
            'group.id': None,
        })
        
        self.__consumer.subscribe(['quickstart-events'])
        
        names = ['Alice','Bob','Charlie','David','Eve',
                        'Frank','Grace','Hannah','Ivan','Judy',
                        'Kevin','Linda','Mike','Nancy','Oscar',
                        'Pam','Quinn','Ralph','Sara','Tom','Uma',
                        'Victor','Wendy','Xavier','Yvonne','Zach']
        
        self.__banks = {i:random.randint(1,1000) for i in names}
    
    async def loop_consume(self):
        