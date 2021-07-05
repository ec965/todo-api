package validator

import (
	"net/http"
	"strings"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator"
)

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
// return the bad fields as a map
func IsValid(r *http.Request, data interface{}) (map[string]string,error) {
	r.ParseForm()
	err := formDecoder.Decode(data, r.Form)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(data)
	// collect errors into a map
	if err != nil {
		errMap := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			key := pascalToCamelCase(fieldErr.Field())
			errMap[key] = fieldErr.ActualTag()
		}
		return errMap, err
	}
	return nil, nil
}
