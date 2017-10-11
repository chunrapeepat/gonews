package app

import "net/http"

func Mount(mux *http.ServeMux) {
  mux.HandleFunc("/", indexHandler)
}
