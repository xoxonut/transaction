from confluent_kafka import Consumer
import random
import json 
import threading


class c:
    
    def __init__(self) -> None:
        
        self.__consumer = Consumer({
            'bootstrap.servers': 'localhost:9092',
            'group.id': 'foo',
            'auto.offset.reset': 'earliest'
        })
        
        self.__consumer.subscribe(['quickstart-events'])
        
        names = ['Alice','Bob','Charlie','David','Eve',
                        'Frank','Grace','Hannah','Ivan','Judy',
                        'Kevin','Linda','Mike','Nancy','Oscar',
                        'Pam','Quinn','Ralph','Sara','Tom','Uma',
                        'Victor','Wendy','Xavier','Yvonne','Zach']
        
        self.__banks = {i:random.randint(1,1000) for i in names}
    
    def loop_consume(self):
        while True :
            msg = self.__consumer.poll(1.0)
            if msg is None:
                print('no message')
                continue
            if msg.error():
                print("Consumer error: {}".format(msg.error()))
            else :
                print('Received message: {}'.format(msg.value().decode('utf-8')))
                data = json.loads(msg.value().decode('utf-8'))
                print(data)

if __name__ == '__main__':
    c = c()
    c.loop_consume()