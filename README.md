# GoDistributedSystemPractice
分散システムと柔軟なデータ処理の練習用
# 使用にあたってのメモ
## 起動前に打つコマンド
**nsqdのインスタンス確認**
```
nsqlookupd
```

**nsqdを起動**
```
nsqd --lookupd-tcp-address=localhost:4160
```

**mongodbを起動**
```
mongod --dbpath .db
```
