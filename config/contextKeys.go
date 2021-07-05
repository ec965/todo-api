package config

import (
	"context"
	"net/http"
)

type contextKey string

func (c contextKey) String() string {
	return "todo-api ctx key " + string(c)
}

const (
	CtxUser = contextKey("user")
)

func AddRequestContext(r *http.Request, key contextKey, value interface{}) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), key, value))
}
