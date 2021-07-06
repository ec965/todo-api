package models

import (
	"context"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
// Turn a Pascal case string into snake case
func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// create the INSERT database string and the INSERT args array
func makeInsert(model interface{}) (string, []interface{}) {
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)
	// example: convert "User" to "users"
	name := t.Name()
	name = strings.ToLower(string(name[0])) + string(name[1:]) + "s"

	// collect field values
	var values []interface{}
	// collect field names
	var fields []string
	// create the ( $1, $2 ... ) thing
	var dollar []string
	count := 1
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		omit := f.Tag.Get(dbomitTag)
		if omit == "" || !strings.Contains("insert", omit) {
			fields = append(fields, toSnakeCase(f.Name))
			dollar = append(dollar, "$"+strconv.Itoa(count))
			values = append(values, val)
			count += 1
		}
	}
	// join the array into a comma delimited string
	fieldsStr := strings.Join(fields[:], ", ")
	dollarStr := strings.Join(dollar[:], ", ")

	// create the sql statement
	s := "INSERT INTO " + name + " ( " + fieldsStr + " ) VALUES ( " + dollarStr + " ) RETURNING id"
	return s, values
}

// Insert with context
func InsertContext(ctx context.Context, model interface{}) (int64, error) {
	s, v := makeInsert(model)
	var id int64
	err := db.QueryRowContext(ctx, s, v...).Scan(&id)
	return id, err
}

// Insert a model into the database
func Insert(model interface{}) (int64, error) {
	s, v := makeInsert(model)
	var id int64
	err := db.QueryRow(s, v...).Scan(&id)
	return id, err
}
