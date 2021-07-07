package handlers

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"

	"github.com/ec965/todo-api/models"
)

type Jwt struct {
	models.User
	jwt.StandardClaims
}

func (j *Jwt) GetMap() jwt.MapClaims {
	var m jwt.MapClaims
	jByte, _ := json.Marshal(j)
	json.Unmarshal(jByte, &m)
	return m
}

func GetFromClaims(c jwt.Claims) Jwt {
	u := Jwt{}
	mByte, _ := json.Marshal(c)
	json.Unmarshal(mByte, &u)
	return u
}
