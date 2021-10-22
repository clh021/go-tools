package websocketGo

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	} // use default options
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func Main() {
	addr := flag.String("addr", "0.0.0.0:8080", "http service address")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome for you!\n"))
	})
	http.HandleFunc("/ws", ws)
	fmt.Println(*addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
