package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/ec965/todo-api/config"
	jwtUser "github.com/ec965/todo-api/handlers/jwt"
	res "github.com/ec965/todo-api/handlers/response"
	"github.com/ec965/todo-api/handlers/validator"
	"github.com/ec965/todo-api/models"
)

type login struct {
	Password string `form:"password" validate:"required,max=36,min=6"`
	Username string `form:"username" validate:"required,max=36"`
}

// create a new user
// returns status OK
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// parse form
	newUser := struct {
		FirstName string `form:"firstName" validate:"required,max=64"`
		LastName  string `form:"lastName" validate:"required,max=64"`
		Email     string `form:"email" validate:"required,email"`
		Role      string `form:"role" validate:"required,eq=user|eq=admin"` // maybe validate this against the db
		login
	}{}
	if errMap, err := validator.IsValid(r, &newUser); err != nil {
		res.Status(http.StatusBadRequest).Json(errMap).Send(w)
		return
	}

	fmt.Println("validated")

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

// login the user
// returns the JWToken
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	l := login{}
	// parse form and validate fields
	if errMap, err := validator.IsValid(r, &l); err != nil {
		res.Status(http.StatusBadRequest).Json(errMap).Send(w)
		return
	}

	// find user
	user := models.User{}
	models.Db.Where("username = ?", l.Username).Preload("Role").First(&user)
	// check if user in db
	if user == (models.User{}) {
		res.Status(http.StatusNotFound).Json(res.Error("user not found")).Send(w)
		return
	}
	// check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password))
	if err != nil {
		res.Status(http.StatusNotFound).Json(res.Error("password is incorrect")).Send(w)
		return
	}

	ju := jwtUser.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role.Name,
		RoleId:    user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + config.TokenDuration,
			Issuer:    config.TokenIssuer,
		},
	}
	juMap := ju.GetMap()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, juMap)

	tokenStr, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		panic(err)
	}

	res.Status(http.StatusOK).Text(tokenStr).Send(w)
}
