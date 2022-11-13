package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8181", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func main() {
	fmt.Println("The server started")
	flag.Parse()
	hub := newHub()
	go hub.run()
	go castIoTMessage(hub)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/message-center/public", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get a ws request")
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
