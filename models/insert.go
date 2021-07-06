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
	// call Elem() b/c we only want this function to work on references
	t := reflect.TypeOf(model).Elem()
	v := reflect.ValueOf(model).Elem()
	// example: convert "User" to "users"
	name := t.Name()
	name = strings.ToLower(string(name[0])) + string(name[1:]) + "s"

	// collect field values
	var values []interface{}
	// collect field names
	var fields []string
	// create the ( $1, $2 ... ) thing
	var dollar []string
	// collect fields that are auto added
	var auto []string
	count := 1
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		tag := f.Tag.Get(dbTag)
		if tag == "" || !strings.Contains("auto", tag) {
			fields = append(fields, toSnakeCase(f.Name))
			dollar = append(dollar, "$"+strconv.Itoa(count))
			values = append(values, val)
			count += 1
		} else if strings.Contains("auto", tag) {
			auto = append(auto, toSnakeCase(f.Name))
		}
	}
	// join the array into a comma delimited string
	fieldsStr := strings.Join(fields[:], ", ")
	dollarStr := strings.Join(dollar[:], ", ")
	autoStr := strings.Join(auto[:], ", ")

	// create the sql statement
	s := "INSERT INTO " + name + " ( " + fieldsStr + " ) VALUES ( " + dollarStr + " ) RETURNING " + autoStr
	return s, values
}

// Insert with context
// model should be a reference
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
	rows, err := db.Query(s, v...)
	// 1. create a new reflection of the model
	// 2. parse the fields and Column Names and match up the snake case titles
	// 3. scan in values from the db
	// 4. apply the scanned values to the new reflection
	// 5. assign the new reflection to the original model reference
	for rows.Next() {
		for i := range rows.Columns() {
			valuePrts[i] = &values[i]
		}
	}
	return id, err
}
