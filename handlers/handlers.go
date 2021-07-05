package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/form/v4"
	"gopkg.in/go-playground/validator.v8"

	res "github.com/ec965/todo-api/handlers/response"
	"github.com/ec965/todo-api/models"
)

var validate *validator.Validate
var formDecoder *form.Decoder

func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
	formDecoder = form.NewDecoder()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newUser := struct {
		FirstName string `form:"firstName" validate:"required,max=64"`
		LastName  string `form:"lastName" validate:"required,max=64"`
		Username  string `form:"username" validate:"required,max=36"`
		Password  string `form:"password" validate:"required,max=36,min=6"`
		Email     string `form:"email" validate:"required,email"`
		Role      string `form:"role" validate:"required,eq=user|eq=admin"` // maybe validate this against the db
	}{}
	err := formDecoder.Decode(&newUser, r.Form)

	if err != nil {
		errJson := res.Error("invalid form")
		res.Status(http.StatusBadRequest).Json(errJson).Send(w)
		return
	}
	fmt.Println(newUser)

	err = validate.Struct(newUser)

	if err != nil {
		fmt.Println(err)
		res.Status(http.StatusBadRequest).Json(err).Send(w)
		return
	}

	role := models.FindRoleByName(newUser.Role)
	if role == (models.Role{}) {
		errJson := res.Error("invalid role")
		res.Status(http.StatusBadRequest).Json(errJson).Send(w)
		return
	}

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

func Ping(w http.ResponseWriter, r *http.Request) {
	res.Status(http.StatusOK).Json(res.Message("pong")).Send(w)
}
