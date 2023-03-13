from confluent_kafka import Producer
import socket
import json
import random
import asyncio
from time import process_time 

class p:
    def __init__(self):
        self.__producer = Producer({
            'bootstrap.servers': 'localhost:9092',
            'client.id': socket.gethostname(),
            'queue.buffering.max.kbytes': 2000000,
            'queue.buffering.max.messages': 1000000
        })
        self.__names = ['Alice','Bob','Charlie','David','Eve',
                        'Frank','Grace','Hannah','Ivan','Judy',
                        'Kevin','Linda','Mike','Nancy','Oscar',
                        'Pam','Quinn','Ralph','Sara','Tom','Uma',
                        'Victor','Wendy','Xavier','Yvonne','Zach']
   
    
    async def __loop_produce(self,t:int = 10):
        
        def make_transaction():
            giver,receiver = random.sample(self.__names,2)
            amount = random.randint(1,100)
            transaction = {'giver':giver,'receiver':receiver,'amount':amount}
            return json.dumps(transaction).encode('utf-8')
        
        def acked(err,msg):
            if err is not None:
                print("Failed to deliver message: {0}: {1}".format(msg.value(),err.str()))
            else:
                pass
                # print("Message produced: {0}".format(msg.value()))
        
        
        for i in range(t):
            print(i)
            self.__producer.produce(topic='quickstart-events',value=make_transaction(),callback=acked)
            self.__producer.poll(0)
            await asyncio.sleep(0)
        self.__producer.flush()
        print('flushed')
        
    async def loop_produce(self,t:int = 10):
        p1 = asyncio.create_task(self.__loop_produce(t//2))
        p2 = asyncio.create_task(self.__loop_produce(t//2))
        await asyncio.gather(p1,p2)
    
if __name__ == '__main__':
    p = p()
    s = process_time()
    asyncio.run(p.loop_produce(10000000))
    print(process_time()-s)