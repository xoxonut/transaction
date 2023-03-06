from confluent_kafka import Producer
import socket
import json
import random

class p:
    def __init__(self):
        self.__producer = Producer({
            'bootstrap.servers': 'localhost:9092',
            'client.id': socket.gethostname()
        })
        self.__names = set(['Alice','Bob','Charlie','David','Eve',
                        'Frank','Grace','Hannah','Ivan','Judy',
                        'Kevin','Linda','Mike','Nancy','Oscar',
                        'Pam','Quinn','Ralph','Sara','Tom','Uma',
                        'Victor','Wendy','Xavier','Yvonne','Zach'])
   
    
    def loop_produce(self,t:int = 10):
        
        def make_transaction(self):
            giver = random.choice(self.__names)
            reciver = random.choice(self.__names-set([giver]))
            amount = random.randint(1,100)
            transcation = {'giver':giver,'reciver':reciver,'amount':amount}
            return json.dump(transcation)
        
        def acked(err,msg):
            if err is not None:
                print("Failed to deliver message: {0}: {1}".format(msg.value(),err.str()))
            else:
                print("Message produced: {0}".format(msg.value()))
        
        
        for i in range(t):
            self.__producer.produce(topic='quickstart-events',value=make_transaction(),callback=acked)
            self.__producer.poll(1)
        self.__producer.flush()
        print('flushed')
    
if __name__ == '__main__':
    p = p()
    p.loop_produce()
    print('done')