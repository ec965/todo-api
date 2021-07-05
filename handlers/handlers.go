package handlers

import (
	"encoding/json"
	"net/http"

	// "gopkg.in/go-playground/validator.v8"

	res "github.com/ec965/todo-api/handlers/response"
	"github.com/ec965/todo-api/models"
)

type NewUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var nu NewUser
	err := json.NewDecoder(r.Body).Decode(&nu)

	if err != nil {
		res.Status(http.StatusBadRequest).Json(res.Error{Error: "invalid json"}).Send(w)
		return
	}

	role := models.FindRoleByName(nu.Role)
	if role == (models.Role{}) {
		res.Status(http.StatusBadRequest).Json(res.Error{Error: "invalid role"}).Send(w)
		return
	}

	user := models.User{
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
		Username:  nu.Username,
		Password:  nu.Password,
		Email:     nu.Email,
		Role:      role,
	}
	models.Db.Create(&user)

	res.Status(http.StatusOK).Send(w)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	res.Status(http.StatusOK).Json(res.Message{Message: "pong"}).Send(w)
}
