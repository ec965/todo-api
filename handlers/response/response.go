package response

import (
	"encoding/json"
	"net/http"
)

const (
	ContentTypeJSON = "application/json"
	ContentTypeText = "text/plain; charset=utf-8"
)

type response struct {
	Status      int
	Body        []byte
	ContentType string
}

type ErrorJson struct {
	Error string `json:"error"`
}

type MessageJson struct {
	Message string `json:"message"`
}

func (res *response) Json(input interface{}) *response {
	msg, _ := json.Marshal(input)
	res.Body = msg
	res.ContentType = ContentTypeJSON
	return res
}

func (res *response) Text(input string) *response {
	res.ContentType = ContentTypeText
	res.Body = []byte(input)
	return res
}

func (res *response) Send(w http.ResponseWriter) {
	if res.ContentType != "" {
		w.Header().Set("Content-Type", res.ContentType)
	}
	if res.Body != nil {
		w.Write(res.Body)
	}
	if res.Status != 0 && res.Status != http.StatusOK {
		w.WriteHeader(res.Status)
	}
}

func Status(status int) *response {
	return &response{Status: status}
}

func Error(err string) ErrorJson {
	return ErrorJson{Error: err}
}

func Message(msg string) MessageJson {
	return MessageJson{Message: msg}
}
