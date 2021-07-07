package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ec965/todo-api/config"
)

const (
	ContentTypeJSON = "application/json"
	ContentTypeText = "text/plain; charset=utf-8"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value(config.CtxUser))
	sendText(w, "pong")
}

// helper functions for response writer

func sendStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func sendJson(w http.ResponseWriter, j interface{}) {
	msg, _ := json.Marshal(j)
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.Write(msg)
}

func sendJsonErr(w http.ResponseWriter, err string) {
	e := map[string]string{"error": err}
	sendJson(w, e)
}

func sendJsonMsg(w http.ResponseWriter, msg string) {
	m := map[string]string{"message": msg}
	sendJson(w, m)
}

func sendText(w http.ResponseWriter, t string) {
	w.Header().Set("Content-Type", ContentTypeText)
	w.Write([]byte(t))
}

func sendError(w http.ResponseWriter, err error){
	http.Error(w, err.Error(), http.StatusInternalServerError)
}