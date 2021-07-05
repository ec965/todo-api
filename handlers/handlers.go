package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/go-playground/validator.v8"

	res "github.com/ec965/todo-api/handlers/response"
	"github.com/ec965/todo-api/models"
)

var validate *validator.Validate

func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
}

type NewUser struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Role      string `json:"role" validate:"required"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var nu NewUser
	err := json.NewDecoder(r.Body).Decode(&nu)

	if err != nil {
		res.Status(http.StatusBadRequest).Json(res.Error{Error: "invalid json"}).Send(w)
		return
	}
	fmt.Println(nu)

	errs := validate.Struct(nu)

	if errs != nil {
		res.Status(http.StatusBadRequest).Json(errs).Send(w)
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
