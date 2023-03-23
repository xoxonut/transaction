# transaction
![](https://github.com/xoxonut/transaction/blob/main/transaction.drawio.png)
## one consumer
```
Consumer error: KafkaError{code=_PARTITION_EOF,val=-191,str="Fetch from broker 0 reached end of partition at offset 52189 (HighwaterMark 52189)"}
```
269.1394375ms
s:613 f:91576

## 2 consumer
sometimes time outed
### 1

#### first
time out
9681 10084
#### sec
```
Consumer error: KafkaError{code=_PARTITION_EOF,val=-191,str="Fetch from broker 0 reached end of partition at offset 47811 (HighwaterMark 47811)"}
```
92.4908152
9161 8650

### 2
```
Consumer error: KafkaError{code=_PARTITION_EOF,val=-191,str="Fetch from broker 0 reached end of partition at offset 47811 (HighwaterMark 47811)"}
```

262.9635215
20276 42148