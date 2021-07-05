package handlers

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/ec965/todo-api/config"
	res "github.com/ec965/todo-api/handlers/response"
)

// get the jwt and parse it into a user
func JwtMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tStr := r.Header.Get("Authorization")
		tStr = strings.Split(tStr, "Bearer ")[1]

		// token , err := jwt.Parse(tStr, func(token *jwt.Token) (interface{}, error){
		// 	return []byte(config.Secret), nil
		// })

		// if err != nil {
		// 	res.Status(http.StatusBadRequest).Json(res.Error("invalid token")).Send(w)
		// 	return
		// }

		// TODO: parse out the JWT into a struct

		next.ServeHTTP(w,r)
	})
}