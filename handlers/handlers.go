package handlers

import (
	"net/http"

	res "github.com/ec965/todo-api/handlers/response"
	"github.com/ec965/todo-api/handlers/validator"
	"github.com/ec965/todo-api/models"
)

type login struct {
	Password string `form:"password" validate:"required,max=36,min=6"`
	Username string `form:"username" validate:"required,max=36"`
}

// create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// parse form
	newUser := struct {
		FirstName string `form:"firstName" validate:"required,max=64"`
		LastName  string `form:"lastName" validate:"required,max=64"`
		Email     string `form:"email" validate:"required,email"`
		Role      string `form:"role" validate:"required,eq=user|eq=admin"` // maybe validate this against the db
		login
	}{}
	if !validator.IsValid(w, r, &newUser) {
		return
	}

	// get the user's role
	role := models.FindRoleByName(newUser.Role)
	if role == (models.Role{}) {
		errJson := res.Error("invalid role")
		res.Status(http.StatusBadRequest).Json(errJson).Send(w)
		return
	}

	// create the new user
	user := models.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Username:  newUser.Username,
		Password:  newUser.Password,
		Email:     newUser.Email,
		Role:      role,
	}
	models.Db.Create(&user)

	res.Status(http.StatusOK).Send(w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	l := login{}
	if !validator.IsValid(w, r, &l) {
		return
	}

	user := models.User{}
	models.Db.Where("username = ?", l.Username).First(&user)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	res.Status(http.StatusOK).Json(res.Message("pong")).Send(w)
}
