package handlers

import (
	"net/http"
	"encoding/json"
	"time"
)

type User struct {
	ID int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ApiHandler(w http.ResponseWriter, r * http.Request){
	user := User{1,"Enoch", "Chau", "ec965", "pwhash", "enoch965@gmail.com", time.Now(), time.Now()}
	b, _:= json.Marshal(user)
	w.Write(b)
}
