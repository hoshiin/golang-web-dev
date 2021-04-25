package controllers

import (
	"html/template"
	"net/http"

	"github.com/hoshiin/golang-web-dev/042_mongodb/10_hands-on/starting-code/models"
	"github.com/hoshiin/golang-web-dev/042_mongodb/10_hands-on/starting-code/session"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	tpl *template.Template
}

func NewUserController(tpl *template.Template) *UserController {
	return &UserController{tpl: tpl}
}

func (uc UserController) Index(w http.ResponseWriter, req *http.Request) {
	id := session.Get(w, req)
	u := session.GetUser(id)
	session.Show() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (uc UserController) Bar(w http.ResponseWriter, req *http.Request) {
	id := session.Get(w, req)
	u := session.GetUser(id)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	session.Show() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func (uc UserController) Signup(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		_, ok := session.User(un)
		if ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		session.Create(w, req, un)

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{
			UserName: un,
			Password: bs,
			First:    f,
			Last:     l,
			Role:     r,
		}
		session.CreateUser(u)

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.Show() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func (uc UserController) Login(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		var ok bool
		u, ok = session.User(un)
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		session.Create(w, req, un)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.Show() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (uc UserController) Logout(w http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	session.Delete(w, req)
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
