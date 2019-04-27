package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

func main() {

}

var db *mgo.Session

// MongoDBとの接続
func dialdb() error {
	var err error
	log.Println("MongoDBにダイアル中: localhost")
	db, err = mgo.Dial("localhost")
	return err
}

// MongoDBとの接続を解除
func closedb() {
	db.Close()
	log.Println("データベース接続を閉じられました")
}
