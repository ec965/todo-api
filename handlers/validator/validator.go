package validator

import (
	"net/http"
	"strings"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator"
)

type FormError struct {
	Error string `json:"error"`
	Fields map[string]string `json:"fields"`
}

var validate *validator.Validate
var formDecoder *form.Decoder

func init() {
	validate = validator.New()
	formDecoder = form.NewDecoder()
}

func pascalToCamelCase(in string) string{
	f := strings.ToLower(string(in[0]))
	r := string(in[1:])
	return f+r
}

// parse form data and validate it
// return the bad fields as a FormError
func IsValid(r *http.Request, data interface{}) (FormError,error) {
	r.ParseForm()
	err := formDecoder.Decode(data, r.Form)
	if err != nil {
		return FormError{}, err
	}
	err = validate.Struct(data)
	// collect errors into a map
	if err != nil {
		errMap := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			key := pascalToCamelCase(fieldErr.Field())
			errMap[key] = fieldErr.ActualTag()
		}
		errJson := FormError{
			Error: "validation error",
			Fields: errMap,
		}
		return errJson, err
	}
	return FormError{}, nil
}
