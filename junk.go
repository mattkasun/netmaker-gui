package main

import (
	"log"
	"net/http"
	//	. "github.com/maragudk/gomponents/html"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(imagehandler)))
}

func imagehandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./images/netmaker2.png")
}
