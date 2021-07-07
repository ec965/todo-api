package config

import (
	"context"
	"net/http"
)

type ContextKey string

func (c ContextKey) String() string {
	return "todo-api ctx key " + string(c)
}

const (
	CtxUser = ContextKey("user")
)

func AddRequestContext(r *http.Request, key ContextKey, value interface{}) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), key, value))
}
