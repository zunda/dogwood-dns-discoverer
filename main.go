package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func usage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, `
<html><title>dogwood-dns-discoverer</title></html>
<h1>dogwood-dns-discoverer</h1>
<p>Try <a href="/dig/example.com">/dig/example.com</a>.</p>
`)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := httprouter.New()
	router.GET("/", usage)

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}
