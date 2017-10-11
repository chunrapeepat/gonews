package main

import (
  "net/http"
  "../../pkg/app"
)

const port = ":8080"

func main() {
  mux := http.NewServeMux()
  app.Mount(mux)
  http.ListenAndServe(port, mux)
}
