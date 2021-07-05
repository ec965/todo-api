package jwt

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id        uint
	CreatedAt time.Time
	Username  string
	FirstName string
	LastName  string
	Email     string
	Role      string
	RoleId    string
}

func (j *User) GetMap() map[string]string {
	var m map[string]string
	jByte, _ := json.Marshal(j)
	json.Unmarshal(jByte, &m)
	return m
}

func GetFromClaims(c jwt.Claims) User {
	u := User{}
	mByte, _ := json.Marshal(c)
	json.Unmarshal(mByte, &u)
	return u
}
