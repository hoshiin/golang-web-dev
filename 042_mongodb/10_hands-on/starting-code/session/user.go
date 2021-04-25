package session

import (
	"time"

	"github.com/hoshiin/golang-web-dev/042_mongodb/10_hands-on/starting-code/models"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]models.User{} // user ID, user

// if the user exists already, get user .
func GetUser(sessionID string) models.User {
	var u models.User
	if s, ok := sessions[sessionID]; ok {
		s.LastActivity = time.Now()
		sessions[sessionID] = s
		u = users[s.UserName]
	}
	return u
}

// CreateUser .
func CreateUser(u models.User) (models.User, error) {
	// store user in dbUsers
	bs, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.MinCost)
	if err != nil {
		return u, err
	}
	user := models.User{
		UserName: u.UserName,
		Password: bs,
		First:    u.First,
		Last:     u.Last,
		Role:     u.Role,
	}
	users[u.UserName] = user
	return user, nil
}

// check if there is a user .
func User(userName string) (models.User, bool) {
	u, ok := users[userName]
	return u, ok
}
