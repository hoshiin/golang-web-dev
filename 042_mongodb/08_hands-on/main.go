package main

import (
	"net/http"

	"github.com/hoshiin/golang-web-dev/042_mongodb/08_hands-on/controllers"
	"github.com/hoshiin/golang-web-dev/042_mongodb/08_hands-on/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return models.LoadUser()
}
