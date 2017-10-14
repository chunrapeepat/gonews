package view

import (
	"net/http"

	"../model"
)

// IndexData structure type
type IndexData struct {
	List []*model.News
}

// Index renders index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// New render new view
func New(w http.ResponseWriter, data *model.News) {
	render(tpNew, w, data)
}

// AdminLogin renders admin login view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminList renders admin list view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// AdminCreate renders admin create view
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit renders admin edit view
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminEdit, w, data)
}
