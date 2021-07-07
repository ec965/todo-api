package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"errors"

	"github.com/ec965/todo-api/config"
	"github.com/ec965/todo-api/models"
)

const (
	ContentTypeJSON = "application/json"
	ContentTypeText = "text/plain; charset=utf-8"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value(config.CtxUser))
	_, err:=getUser(r)
	if err != nil {
		sendError(w, err)
		return
	}
	sendText(w, "pong")
}

// helper to get context items
// this only works on routes that require a token
// otherwise an empty user will be returned
func getUser(r *http.Request) (models.User, error) {
	ctxUser := r.Context().Value(config.CtxUser)
	user, ok := ctxUser.(models.User)
	if !ok {
		return models.User{}, errors.New("couldn't find the jwt user in request context")
	}
	return user, nil
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

func sendError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
