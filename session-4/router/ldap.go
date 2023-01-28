package router

import (
	"Hacktiv8project/session-4/service"
	"fmt"
	"net/http"
	"text/template"
)

type RouterHandler struct {
	loginService *service.LoginSvc
}

func NewRouterHandler(loginService *service.LoginSvc) *RouterHandler {
	return &RouterHandler{loginService: loginService}
}
func (rh *RouterHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("login.html")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rh *RouterHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//aunthenticate via ldap
	data, err := rh.loginService.Authenticate(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//greet user on success
	message := fmt.Sprintf("welcome %s", data.FullName)
	w.Write([]byte(message))
}
