package app

import (
	"net/http"

	"../model"
	"../view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
	}
	view.AdminLogin(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {
	view.AdminList(w, nil)
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		n := model.News{
			Topic:       r.FormValue("topic"),
			Description: r.FormValue("description"),
		}
		err := model.CreateNews(&n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/create", http.StatusSeeOther)
		return
	}
	view.AdminCreate(w, nil)
}

func adminEdit(w http.ResponseWriter, r *http.Request) {
	view.AdminEdit(w, nil)
}
