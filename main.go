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
	t := `
<html><title>dogwood-dns-discoverer</title></html>
<body>
<h1>dogwood-dns-discoverer</h1>
<p>Try:</p>
<ul>
<li>Heroku: <a href="/lookup/www.heroku.com">/lookup/www.heroku.com</a></li>
<li>Domain name reserved for documentation purposes: <a href="/lookup/example.com">/lookup/example.com</a></li>
`
	h := os.Getenv("HEROKU_DNS_DYNO_NAME")
	if h != "" {
		t += "<li>This dyno: <a href=\"/lookup/" + h + "\">/lookup/" + h + "</a></li>\n"
	}

	fmt.Fprintf(w, t + "</ul>\n</body>\n")
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
