package main

/*
ストリーミングAPIの接続関連を管理する関数群
*/
import (
	"io"
	"net"
	"time"
)

var conn net.Conn        //コネクション
var reader io.ReadCloser // Read & Closeメソッドのグループ

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

// コネクション切断用の関数
// ctrl + Cで終了する際に利用
func closeConn() {
	if conn != nil {
		conn.Close()
	}
	if reader != nil {
		reader.Close()
	}
}
