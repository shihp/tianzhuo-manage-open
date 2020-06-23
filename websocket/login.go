package main

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

var requestBody = make(map[string]string)

func Login(conn net.Conn, op ws.OpCode, params []string) {
	for _, param := range params {
		log.Info(param)
		kv := strings.Split(string(param), "=")
		requestBody[kv[0]] = kv[1]
	}

	//todo map默认值
	if requestBody["token"] != "" {
		connMap[requestBody["token"]] = conn
	}

	err := wsutil.WriteServerMessage(conn, op, []byte("ok"))
	if err != nil {
		log.Error(err)
	}
}

func Logout(conn net.Conn) {

}

func Hi(net.Conn, ws.OpCode, []string) {

}
