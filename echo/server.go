package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
	"net/url"
)



func handler(rw http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			http.Error(rw, err.Error(), 500)
		}
	}()

	defer req.Body.Close()

	var body []byte

	u := req.URL.String()
	u, _ =  url.QueryUnescape(u)
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("handler body read error: %s\n", err.Error())
		return
	}

	tpl, err := template.New("res").Parse(`
			-------------
			req url:
			-------------
			{{.url}}

			-------------
			req body:
			-------------
			{{.body}}

			time: {{.time}}
		`)
	if err != nil {
		panic("")
	}
	brw := bufio.NewWriter(rw)
	tpl.Execute(brw, map[string]interface{}{
		"url":    u,
		"body":   string(body),
		"time":   time.Now(),
	})
	brw.Flush()
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
