package httpServer

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["param"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Golang!")
}

type serverHandler struct{}

func (h *serverHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	sayHello(res, req)
}

func init() {
	// http.HandleFunc("/", sayHello)
	http.HandleFunc("/", nil)

	fmt.Println("Http server listening at http://localhost:9090")
	err := http.ListenAndServe(":9090", &serverHandler{})
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
