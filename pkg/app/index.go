package app

import (
	"net/http"

	"../model"
	"../view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	list := model.ListNews()
	view.Index(w, &view.IndexData{
		List: list,
	})
}
