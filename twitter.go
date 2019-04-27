package main

/*
ストリーミングAPIの接続関連を管理する関数群
*/
import (
	"net"
	"time"
)

var conn net.Conn

//ストリーミングAPIダイアル用の関数
func dial(netw, addr string) (net.Conn, error) {
	//コネクトが切れているか確認
	if conn != nil {
		conn.Close()
		conn = nil
	}
	// APIのエラー処理
	netc, err := net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil {
		return nil, err
	}
	conn = netc
	return netc, nil
}
