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
## a group of ksql connections (10000 payments 1000 rps, random giver and receiver)
- 2 connections : 54tps
- 10 connections : 61tps
- 100 connections : 56tps
## one payment one ksql connections (10000 payments 1000 rps, random giver and receiver)
- with connections close : 61tps
- without connections close : 62tps
## sync client and create ksql connection per payment in server (10000 payments 1000 rps, random giver and receiver)
- 34tps
##  without grpc one connection and sync 10000 payments
- python : 83tps
    - 之前比較快是應為只計算user mode kernel的時間 
    - 可能沒有等回傳值，等一段時間就跳掉
```
[2023-05-12 00:32:38,727] INFO stream-thread [_confluent-ksql-default_query_CTAS_BANK_5-d9bd2938-cfa3-4d79-8e40-65f436377c2b-StreamThread-1] Processed 8948 total records, ran 0 punctuators, and committed 57 total tasks since the last update (org.apache.kafka.streams.processor.internals.StreamThread:842)
```
- go : 37.5 tps (async and sync are the same)

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