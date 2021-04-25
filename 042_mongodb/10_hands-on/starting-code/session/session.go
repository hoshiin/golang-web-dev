package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hoshiin/golang-web-dev/042_mongodb/10_hands-on/starting-code/models"
)

const sessionLength int = 30

var sessions = map[string]models.Session{} // session ID, session
var lastCleaned time.Time = time.Now()

func Get(w http.ResponseWriter, req *http.Request) (id string) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return c.Value
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := sessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		sessions[c.Value] = s
	}
	_, ok = users[s.UserName]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

func Clean() {
	// clean up dbSessions
	if time.Since(lastCleaned) > (time.Second * 30) {
		go clean()
	}
}

func clean() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	Show()                      // for demonstration purposes
	for k, v := range sessions {
		if time.Since(v.LastActivity) > (time.Second * 30) {
			delete(sessions, k)
		}
	}
	lastCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	Show()                     // for demonstration purposes
}

// Create .
func Create(w http.ResponseWriter, req *http.Request, key string) {
	// create session
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	sessions[c.Value] = models.Session{key, time.Now()}
}

// Delete .
func Delete(w http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("session")
	// delete the session
	delete(sessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	Clean()
}

// for demonstration purposes
func Show() {
	fmt.Println("********")
	for k, v := range sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
