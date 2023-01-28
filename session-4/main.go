package main

import (
	"fmt"
	"net/http"

	"Hacktiv8project/session-4/config"
	"Hacktiv8project/session-4/repository"
	"Hacktiv8project/session-4/router"
	"Hacktiv8project/session-4/service"

	"github.com/go-ldap/ldap"
	"github.com/gorilla/mux"
)

func main() {
	//connect to LDAP Server
	ldapConn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", config.LdapServer, config.LdapPort))
	if err != nil {
		panic(err)
	}
	defer ldapConn.Close()
	//bind to LDAP Directory
	if err := ldapConn.Bind(config.LdapBindDN, config.LdapPassword); err != nil {
		panic(err)
	}

	ldapRepo := repository.NewLDAPRepo(ldapConn)
	loginService := service.NewLoginService(ldapRepo)
	r := mux.NewRouter()
	routerHandler := router.NewRouterHandler(loginService)
	r.HandleFunc("/", routerHandler.IndexHandler)
	r.HandleFunc("/login", routerHandler.LoginHandler)
	portString := fmt.Sprintf(":%d", config.WebServerPort)
	fmt.Println("server started at", portString)
	http.ListenAndServe(portString, r)
}
