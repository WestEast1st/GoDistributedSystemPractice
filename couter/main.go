package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	nsq "github.com/bitly/go-nsq"
	mgo "gopkg.in/mgo.v2"
)

var fatalErr error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

func main() {
	var countsLook sync.Mutex
	var counts map[string]int
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()
	log.Println("データベースに接続します...")
	db, err := mgo.Dial("localhost")
	if err != nil {
		fatal(err)
		return
	}
	defer func() {
		log.Println("データベースの接続を閉じます...")
		db.Close()
	}()
	pollData := db.DB("ballots").C("polls")
	//NSQに接続
	log.Println("NSQに接続ます")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())
	if err != nil {
		fatal(err)
		return
	}

	//NSQからメッセージを受け取った時の処理
	q.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLook.Lock()
		defer countsLook.Unlock()
		if counts == nil {
			counts = make(map[string]int)
		}
		vote := string(m.Body)
		counts[vote]++
		return nil
	}))
	if err := q.ConnectToNSQLookupd("localhost:4161"); err != nil {
		fatal(err)
		return
	}
}
