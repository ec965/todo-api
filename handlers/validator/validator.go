package validator

import (
	"net/http"
	"fmt"

	"github.com/go-playground/form/v4"
	"gopkg.in/go-playground/validator.v8"

	res "github.com/ec965/todo-api/handlers/response"
)

var validate *validator.Validate
var formDecoder *form.Decoder

func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
	formDecoder = form.NewDecoder()
}

func isValidForm(w http.ResponseWriter, r *http.Request, form interface{}) bool {
	r.ParseForm()
	err := formDecoder.Decode(form, r.Form)
	if err != nil {
		errJson := res.Error("invalid form")
		res.Status(http.StatusBadRequest).Json(errJson).Send(w)
		return false
	}
	return true
}

func isValidData(w http.ResponseWriter, r * http.Request, data interface{}) bool {
	err := validate.Struct(data)
	if err != nil {
		// TODO: fix this error json
		fmt.Println(err)
		res.Status(http.StatusBadRequest).Json(err).Send(w)
		return false
	}
	return true
}

func IsValid(w http.ResponseWriter, r *http.Request, data interface{})bool {
	if(!isValidForm(w, r, &data)){
		return false
	}
	if(!isValidData(w, r, &data)){
		return false
	}
	return true
}