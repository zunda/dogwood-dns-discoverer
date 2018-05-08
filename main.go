package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func usage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := `
<html><title>dogwood-dns-discoverer</title></html>
<body>
<h1>dogwood-dns-discoverer</h1>
<p>Try:</p>
<ul>
<li>This host and port: <a href="/self">/self</a></li>
<li>Heroku: <a href="/lookup/www.heroku.com">/lookup/www.heroku.com</a></li>
<li>Domain name reserved for documentation purposes: <a href="/lookup/example.com">/lookup/example.com</a></li>
`
	h := os.Getenv("HEROKU_DNS_DYNO_NAME")
	if h != "" {
		t += "<li>This dyno: <a href=\"/lookup/" + h + "\">/lookup/" + h + "</a></li>\n"
	}

	fmt.Fprintf(w, t+"</ul>\n</body>\n")
}

func lookupAndRespond(w http.ResponseWriter, hostname string, port string) {
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		log.Printf("%s\n", err)
		w.WriteHeader(500)
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	log.Println("Looked up " + hostname)
	for _, addr := range addrs {
		s := addr
		if net.ParseIP(addr).To4() == nil {
			s = "[" + addr + "]"
		}
		if port == "" {
			fmt.Fprintf(w, s+"\n")
		} else {
			fmt.Fprintf(w, s+":"+port+"\n")
		}
	}
}

func self(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hn := os.Getenv("HEROKU_PRIVATE_IP")
	pt := port()
	if hn == "" {
		hn, pt, _ = net.SplitHostPort(r.Host)
	}
	if hn == "" {
		hn = r.Host
		pt = ""
	}
	if hn == "" {
		hn = "localhost"
	}
	lookupAndRespond(w, hn, pt)
}

func lookup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	hn := p.ByName("hostname")
	lookupAndRespond(w, hn, "")
}

func main() {
	port := port()

	router := httprouter.New()
	router.GET("/", usage)
	router.GET("/self", self)
	router.GET("/lookup/:hostname", lookup)

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, router)
	log.Fatal(err)
}
