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