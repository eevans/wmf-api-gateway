// A REST API that simply echos request data as JSON

package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	portNo = flag.Int("port", 8080, "Port to listen on")
)

type jsonResponse struct {
	Path    string              `json:"path"`
	Method  string              `json:"method"`
	Body    string              `json:"body"`
	Form    map[string][]string `json:"form"`
	Headers map[string][]string `json:"headers"`
}

func newJSONResponse(req *http.Request) jsonResponse {
	if err := req.ParseForm(); err != nil {
		log.Println(err)
	}

	var encoded string

	if body, err := ioutil.ReadAll(req.Body); err != nil {
		encoded = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Error base64 encoding body: %s", err)))
	} else {
		encoded = base64.StdEncoding.EncodeToString(body)
	}

	return jsonResponse{req.URL.Path, req.Method, encoded, req.Form, req.Header}
}


func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		response := newJSONResponse(req)

		b, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, fmt.Sprintf("%s", err))
			log.Printf("%d %s %s\n", 500, req.RemoteAddr, req.URL.Path)
			return
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		log.Printf("%d %s %s\n", 200, req.RemoteAddr, req.URL.Path)

		io.WriteString(w, fmt.Sprintf("%v", string(b)))
	})

	listen := fmt.Sprintf(":%d", *portNo)
	log.Printf("Starting up %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}
