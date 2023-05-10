# Go test 
## fixed giver and receiver
### 100000 payments
- 100 rps : 64tps
- 1000 rps : 64tps
### 50000 payments
- 10000 rps : 63tps
## random from 0 to 1000000 (500000 payments)
- 100 rps : 61tps
- 1000 rps : 64tps
- 10000 rps : 63tps

# transaction
![](https://github.com/xoxonut/transaction/blob/main/transaction.drawio.png)
## one consumer

285.930852
1737 98263

## two consumer
sometimes time outed
### 1

268.860491
31864 15947

### 2

283.20981489999997
33383 18806

# Next week 
deal with transaction api :
* kafka python have transaction api for sure
    * find if ksql curd has transaction api or not
* where is the begin of transaction?
## GO
- 1sec sleep and 100 rps
    -   1400 payments timeout 1m0.000884419s
- 2sec sleep and 100rps
    - 1200 payments timeout 1m0.509036772s
- 1sec sleep and 10rps
    - 460 payments timeout 1m0.000289041s

##
好像是我grpc沒弄好 client沒有弄到async

client 的ctx要重弄 -> 建立新的rpc連線
https://stackoverflow.com/questions/31800692/multiple-http-requests-gives-cant-assign-requested-address-unless-sped-up