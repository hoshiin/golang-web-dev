package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/hoshiin/golang-web-dev/042_mongodb/10_hands-on/starting-code/controllers"
)

var tpl *template.Template
var dbSessionsCleaned time.Time

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	uc := controllers.NewUserController(tpl)
	http.HandleFunc("/", uc.Index)
	http.HandleFunc("/bar", uc.Bar)
	http.HandleFunc("/signup", uc.Signup)
	http.HandleFunc("/login", uc.Login)
	http.HandleFunc("/logout", uc.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
