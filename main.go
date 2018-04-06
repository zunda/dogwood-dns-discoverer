package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func usage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, `
<html><title>dogwood-dns-discoverer</title></html>
<h1>dogwood-dns-discoverer</h1>
<p>Try <a href="/lookup/example.com">/lookup/example.com</a>.</p>
`)
}

func lookup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	addrs, err := net.LookupHost(p.ByName("hostname"))
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(500)
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	log.Println("Looked up " + p.ByName("hostname"))
	for _, addr := range addrs {
		fmt.Fprintf(w, addr + "\n")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := httprouter.New()
	router.GET("/", usage)
	router.GET("/lookup/:hostname", lookup)

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}
