package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/ec965/todo-api/config"
	"github.com/ec965/todo-api/handlers/validator"
	"github.com/ec965/todo-api/models"
)

// create a new user
// returns status OK
func Signup(w http.ResponseWriter, r *http.Request) {
	// parse form
	newUser := struct {
		FirstName string `form:"firstName" validate:"required,max=64"`
		LastName  string `form:"lastName" validate:"required,max=64"`
		Email     string `form:"email" validate:"required,email"`
		Role      string `form:"role" validate:"required,eq=user|eq=admin"` // maybe validate this against the db
		Password  string `form:"password" validate:"required,max=36,min=6"`
		Username  string `form:"username" validate:"required,max=36"`
	}{}
	if errMap, err := validator.IsValid(r, &newUser); err != nil {
		sendStatus(w, http.StatusBadRequest)
		sendJson(w, errMap)
		return
	}

	// check if username or email already exists
	user := models.User{}
	hasUsername, hasEmail, err := user.CheckUsernameEmail(newUser.Username, newUser.Email)
	if err != nil {
		sendError(w, err)
		return
	}
	if hasUsername || hasEmail {
		errJson := map[string]string{"error": "invalid field"}
		if hasEmail {
			errJson["email"] = "exists"
		}
		if hasUsername {
			errJson["username"] = "exists"
		}
		sendStatus(w, http.StatusBadRequest)
		sendJson(w, errJson)
		return
	}

	// check the users role
	role := models.Role{}
	err = role.SelectByName(newUser.Role)
	if err != nil {
		sendError(w, err)
		return
	}
	if role == (models.Role{}) {
		sendStatus(w, http.StatusBadRequest)
		sendJsonErr(w, "invalid role")
		return
	}

	// create the new user
	user = models.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Username:  newUser.Username,
		Password:  newUser.Password,
		Email:     newUser.Email,
		RoleId:    role.ID,
	}

	_, err = user.Insert()
	if err != nil {
		sendError(w, err)
		return
	}
	sendJsonMsg(w, "signup successful")
}

// login the user
// returns the JWToken
func Login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	l := struct {
		Password string `form:"password" validate:"required,max=36,min=6"`
		Username string `form:"username" validate:"required,max=36"`
	}{}
	// parse form and validate fields
	if errMap, err := validator.IsValid(r, &l); err != nil {
		sendStatus(w, http.StatusBadRequest)
		sendJson(w, errMap)
		return
	}

	// find user
	user := models.User{}
	err := user.SelectByUsername(l.Username)
	if err != nil {
		sendError(w, err)
		return
	}
	// check if user in db
	if user == (models.User{}) {
		sendStatus(w, http.StatusNotFound)
		sendJsonErr(w, "user not found")
		return
	}
	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password))
	if err != nil {
		sendStatus(w, http.StatusBadRequest)
		sendJsonErr(w, "password incorrect")
		return
	}

	ju := Jwt{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + config.TokenDuration,
			Issuer:    config.TokenIssuer,
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, ju.GetMap()).SignedString([]byte(config.Secret))

	sendJson(w, map[string]string{"token": token})
}
