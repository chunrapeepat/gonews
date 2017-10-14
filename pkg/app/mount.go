package app

import "net/http"

// Mount mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/edit", adminEdit)

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
}

// Admin middleware
func onlyAdmin(h http.Handler) http.Handler {
	return h
}
