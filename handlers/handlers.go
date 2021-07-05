package handlers

import (
	"net/http"

	res "github.com/ec965/todo-api/handlers/response"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	res.Status(http.StatusOK).Json(res.Message("pong")).Send(w)
	w.Write([]byte("hi"))
}
