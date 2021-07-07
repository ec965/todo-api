package models

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ec965/todo-api/config"
)

var db *sql.DB

type table struct {
	name string
	cols []string
}

func Init() {
	var err error
	db, err = sql.Open("pgx", config.DatabaseUrl)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to ping the Database:", err)
	}
}

func (t *table) insertStr() string {
	cStr := ""
	dStr := ""
	for i, c := range t.cols {
		cStr += c
		dStr += "$" + strconv.Itoa(i+1)
		if i != len(t.cols)-1 {
			cStr += ", "
			dStr += ", "
		}
	}

	return "INSERT INTO " + t.name + " ( " + cStr + " ) VALUES ( " + dStr + " ) RETURNING id"
}

func (t *table) updateStr() string {
	str := ""
	for i, c := range t.cols {
		str += c + " = $" + strconv.Itoa(i+1)
		if i != len(t.cols)-1 {
			str += ", "
		}
	}

	return "UPDATE " + t.name + " SET " + str + " WHERE id = $" + strconv.Itoa(len(t.cols)+1)
}

func (t *table) selectByStr(which string) string {
	cStr := ""
	for i, c := range t.cols {
		cStr += c
		if i != len(t.cols)-1 {
			cStr += ", "
		}
	}
	return "SELECT " + cStr + " FROM " + t.name + " WHERE " + which + " = $1"
}
