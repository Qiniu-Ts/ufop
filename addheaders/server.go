package main

import (
	"bufio"
	"log"
	"strings"
	"net/http"
	"os"
	"fmt"
)


func handler(rw http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			http.Error(rw, err.Error(), 500)
		}
	}()

	defer req.Body.Close()

	u := req.URL.Query().Get("url")
	cmd := req.URL.Query().Get("cmd")

	fname := "demo"
	path := strings.Split(cmd, "/")
	if len(path) > 1 {
		fname = path[1]
	}


	res, err := http.Get(u)
	if err != nil {
		return
	}
	defer res.Body.Close()

	fname = fmt.Sprintf("inline;filename=%s", fname)
	rw.Header().Add("Content-Disposition", fname)
	rw.Header().Add("Cdddddn", fname)
	bw := bufio.NewWriter(rw)
	if err := res.Write(bw); err != nil {
		return
	}

	bw.Flush()
}

func health(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("ok"))
}

func main() {
	port := os.Getenv("PORT_HTTP")
	if port == "" {
		port = "9100"
	}
	http.HandleFunc("/handler", handler)
	http.HandleFunc("/health", health)
	log.Fatalln(http.ListenAndServe("0.0.0.0:"+port, nil))
}
