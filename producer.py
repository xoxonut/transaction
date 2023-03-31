from confluent_kafka import Producer
import socket
import json
import random
from time import process_time 
from ksql import KSQLAPI

class p:
    def __init__(self):
        self.__producer = Producer({
            'bootstrap.servers': 'localhost:9092',
            'client.id': socket.gethostname(),
            'queue.buffering.max.kbytes': 2000000,
            'queue.buffering.max.messages': 1000000
        })
        with open('transaction.json','r') as f:
            self.__transactions = json.load(f)['transaction']
        
   
    
    def loop_produce(self):
        

        
        def acked(err,msg):
            if err is not None:
                print("Failed to deliver message: {0}: {1}".format(msg.value(),err.str()))
            else:
                pass
                # print("Message produced: {0}".format(msg.value()))
        
        
        for i in self.__transactions.values():
            self.__producer.poll(0)
            self.__producer.produce(topic='transaction',key=f"{i['giver']}",value=json.dumps({"receiver":i["receiver"],"amount":i["amount"]}).encode('utf-8'),callback=acked)
        self.__producer.flush()
        print('flushed')
        
    # def insert_to_bank(self):
    #     def acked(err,msg):
    #         if err is not None:
    #             print("Failed to deliver message: {0}: {1}".format(msg.value(),err.str()))
    #         else:
    #             pass
    #             print("Message produced: {0}".format(msg.value()))
    #     for i,name in enumerate(self.__names):
    #         self.__producer.poll(0)
    #         self.__producer.produce(topic='balance',key=name,value=json.dumps({'uid':name,'amount':1}).encode('utf-8'),callback=acked)
    #     self.__producer.flush()
    
if __name__ == '__main__':
    p = p()
    s = process_time()
    p.loop_produce()
    print(process_time()-s)
