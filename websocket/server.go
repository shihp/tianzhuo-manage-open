package main

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"strings"
)

var connMap = make(map[string]net.Conn)

func main() {

	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			//defer closeConn(conn)
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Fatal(err)
				}
				message := string(msg)

				messageBody := strings.Split(message, "?")
				if len(messageBody) != 2 {
					err = wsutil.WriteServerMessage(conn, op, msg)
					if err != nil {
						log.Fatal(err)
					}
					closeConn(conn)
				}

				uri := messageBody[0]
				params := strings.Split(messageBody[1], "&")
				switch uri {
				case "login":
					Login(conn, op, params)
				//case "logout":
				//	Logout(conn)
				//case "hello":
				//	Hi(conn, op, params)
				default:
					err = wsutil.WriteServerMessage(conn, op, []byte("？"))
					if err != nil {
						log.Error(err)
					}
				}
			}
		}()
	}))

	if err != nil {
		log.Fatal(err)
	}
}

func closeConn(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}

}
