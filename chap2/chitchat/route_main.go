package main

import (
  "net/http"
  "./data"
)

func index(writer http.ResponseWriter, request *http.Response) {
  files := []string{
    "templates/layout.html",
    "templates/navbar.html",
    "templates/index.html",
  }
  templates := templates.Must(template.ParseFiles(files...))
  threads, err := data.Threads()
  if err == nil {
    templates.ExecuteTemplate(w, "layout", threads)
  }
}
