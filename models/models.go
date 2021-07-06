package models

import (
	"database/sql"
	"log"


	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ec965/todo-api/config"
)

type model interface {
	Insert() (sql.Result, error)
	InsertContext()
	SelectById()
	// Update()
	// UpdateContext()
	// Delete()
	// DeleteContext() 
}

// used to omit a field from it's respective database actions
// actions: insert, update
const dbomitTag = "dbomit"

var db *sql.DB
// var ctx context.Context

func Init() {
	// ctx = context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	// defer cancel()

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

	r := Role{}
	r.SelectById(1)
}
