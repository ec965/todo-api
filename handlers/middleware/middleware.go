package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/ec965/todo-api/config"
	userToken "github.com/ec965/todo-api/handlers/jwt"
	res "github.com/ec965/todo-api/handlers/response"
)

// get the jwt from Authorization Header and parse it into a user
// user is added to request context
func Jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		headerSplit := strings.Split(header, "Bearer ")
		if len(headerSplit) < 2 {
			res.Status(http.StatusBadRequest).Json(res.Error("invalid token")).Send(w)
			return
		}

		tokenStr := headerSplit[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})

		if err != nil {
			res.Status(http.StatusBadRequest).Json(res.Error("invalid token")).Send(w)
			return
		}

		// parse claims into user struct
		user := userToken.GetFromClaims(token.Claims)

		// add user struct to context
		r = config.AddRequestContext(r, config.CtxUser, user)

		next.ServeHTTP(w, r)
	})
}
