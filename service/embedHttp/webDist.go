package embedHttp

import (
	"embed"
	"flag"
	"log"
	"net/http"
)

// go:embed webDist
// var webDist embed.FS

//go:embed webDist/*
var f embed.FS

func webDistServe() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()
	log.Println("ListenAndServe:", *addr)

	// err := http.ListenAndServe(*addr, http.FileServer(http.FS(webDist)))

	http.Handle("/static/", http.StripPrefix("/webDist/", http.FileServer(http.FS(f))))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
