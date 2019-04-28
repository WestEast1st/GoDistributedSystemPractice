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
## mongo
mongoDBのコンソールを起動
```
mongo
```

ballotsを選択
```
use ballots
```
調査項目をinsert
```
db.polls.insert({"title":"今日の気分は?","options":["happy","sad","fail","win"]}}
```
## nsqのメッセージを出力
```
nsq_tail --topic="votes" --lookupd-http-address=localhost:4161
```
