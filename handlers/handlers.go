package handlers

import (
	"fmt"
	"net/http"

	"github.com/ec965/todo-api/config"
	res "github.com/ec965/todo-api/handlers/response"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value(config.CtxUser))
	res.Status(http.StatusOK).Json(res.Message("pong")).Send(w)
}
