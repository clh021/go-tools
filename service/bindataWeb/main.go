package bindataWeb

import (
	"fmt"
	"log"
	"net/http"
)

func newHandler() http.Handler {
	// return http.FileServer(AssetFile())

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		byte, err := Asset("index.html")
		if err != nil {
			log.Println("Read WebTest Error:", err)
		}
		w.Write(byte)
	})
}

func Main() {
	httpPort := 8081
	log.Printf("Web/Api Server listen on :%d\n", httpPort)
	mux := http.NewServeMux()
	mux.Handle("/", newHandler())
	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}
