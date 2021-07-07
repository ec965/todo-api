package handlers

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/ec965/todo-api/config"
)

// get the jwt from Authorization Header and parse it into a user
// user is added to request context
func JwtMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		headerSplit := strings.Split(header, "Bearer ")
		if len(headerSplit) < 2 {
			sendStatus(w, http.StatusBadRequest)
			sendJsonErr(w, "invalid token")
			return
		}

		tokenStr := headerSplit[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})

		if err != nil {
			sendStatus(w, http.StatusBadRequest)
			sendJsonErr(w, "invalid token")
			return
		}

		// parse claims into user struct
		user := GetFromClaims(token.Claims).User

		// add user struct to context
		r = config.AddRequestContext(r, config.CtxUser, user)

		next.ServeHTTP(w, r)
	})
}
