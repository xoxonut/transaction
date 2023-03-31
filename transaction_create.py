import random
import json
st_uid = {i:random.randint(1,100) for i in range(50)}
transaction = {}
for i in range(100000):
    giver,receiver = random.sample(list(st_uid.keys()),2)
    amount = random.randint(1,70)
    transaction[i] = {'giver':giver,'receiver':receiver,'amount':amount}
ret = {'init_state':st_uid,'transaction':transaction}
obj = json.dumps(ret,indent=4)
with open('transaction.json','w') as f:
    f.write(obj)
