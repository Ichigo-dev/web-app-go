package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "golang.org/x/net/http2"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func main() {
  mux := httprouter.New()
  mux.GET("/hello/:name", hello)
  server := http.Server{
    Addr: "127.0.0.1:8080",
    Handler: mux,
  }

  http2.ConfigureServer(&server, &http2.Server{})
  server.ListenAndServe()
}
