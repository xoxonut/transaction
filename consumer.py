from confluent_kafka import Consumer
import random
import json 
import threading
from time import process_time

class c:
    
    def __init__(self) -> None:
        
        self.__consumer = Consumer({
            'bootstrap.servers': 'localhost:9092',
            'group.id': 'foo',
            'auto.offset.reset': 'earliest'
        })
        def on_assign(c,ps):
            for p in ps:
                p.offset = -2
            c.assign(ps)
            return
        self.__consumer.subscribe(['quickstart-events'],on_assign=on_assign)
        names = ['Alice','Bob','Charlie','David','Eve',
                        'Frank','Grace','Hannah','Ivan','Judy',
                        'Kevin','Linda','Mike','Nancy','Oscar',
                        'Pam','Quinn','Ralph','Sara','Tom','Uma',
                        'Victor','Wendy','Xavier','Yvonne','Zach']
        self.count = 0
        self.__banks = {i:random.randint(1,70) for i in names}
        self.success=[0,0]
    def loop_consume(self,t):
        while True :
            if self.count == t: 
                break
            msg = self.__consumer.poll(1)
            if msg is None:
                print('no message')
                continue
            if msg.error():
                print("Consumer error: {}".format(msg.error()))
            else :
                self.count+=1
                data = json.loads(msg.value().decode('utf-8'))
                self.execute_transaction(data)
    def execute_transaction(self,data):
        if data['amount'] > self.__banks[data['giver']]:
            self.success[0]+=1
            return
        self.__banks[data['giver']] -= data['amount']
        self.__banks[data['receiver']] += data['amount']
        self.success[1]+=1
        return
    
    
    
if __name__ == '__main__':
    c = c()
    s = process_time()
    c.loop_consume(10000000)
    print(process_time()-s)
    print(c.success)