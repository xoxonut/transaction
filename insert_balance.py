import json
from ksql import KSQLAPI
with open('transaction.json','r') as f:
    init = json.load(f)['init_state']
client = KSQLAPI('http://localhost:8088')

state = [{'uid':int(k),'amount':int(v)} for k,v in init.items()]
r = client.inserts_stream('BALANCE_STREAM',state)
print(r)