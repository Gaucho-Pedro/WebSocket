package main

import (
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		log.Println(r.Header.Get("Upgrade"))
		if err != nil {
			log.Println(err)
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				log.Printf("Msg:%s", string(msg))
				if err != nil {
					log.Println(err)
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}))
}
