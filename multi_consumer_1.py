from confluent_kafka import Consumer
import json 
from time import process_time
from ksql import KSQLAPI
import ast

class c:
    
    def __init__(self) -> None:
        self.s=0
        self.f=0
        self.__consumer = Consumer({
            'bootstrap.servers': 'localhost:9092',
            'group.id': f'hello',
            'auto.offset.reset': 'earliest',
            'enable.partition.eof': True
        })
        def on_assign(c,ps):
            for p in ps:
                p.offset = -2
            c.assign(ps)
            return
        self.__consumer.subscribe(['transaction'])
        self.__client = KSQLAPI('http://localhost:8088')

    def loop_consume(self):
        while True :
            msg = self.__consumer.poll(1)
            if msg is None:
                print('no message')
                continue
            if msg.error():
                print("Consumer error: {}".format(msg.error()))
                break;
            else :
                self.execute_transaction(msg)
        return 
                
    def execute_transaction(self,msg):
        data = json.loads(msg.value().decode('utf-8'))
        k =int(msg.key().decode('utf-8'),10)
        if data['amount'] > self.get_balance(k):
            self.f+=1
            return
        self.s+=1
        self.insert_balance(k,-data['amount'])
        self.insert_balance(data['receiver'],data['amount'])
        return
    
    def get_balance(self,uid):
        query = self.__client.query(f'select total from bank where uid = {uid}',use_http2=True)
        q = list(query)
        d = ast.literal_eval(q[0])
        self.__client.close_query(d["queryId"])        
        return ast.literal_eval(q[1])[0]
    def insert_balance(self,uid,total):
        # ret = self.__client.query(f'insert into BALANCE_STREAM (uid,total) values ({uid},{total})')
        # print(ret)
        state = [{'uid':uid,'amount':total}]
        r = self.__client.inserts_stream('BALANCE_STREAM',state)
        return
    def cl(self):
        self.__consumer.close()
    
if __name__ == '__main__':
    c = c()
    try:
        s = process_time()
        c.loop_consume()
        print(process_time()-s)
    finally:
        print(c.s,c.f)
        c.cl()
