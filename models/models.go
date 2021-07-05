package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ec965/todo-api/config"
)

type model interface {
	Insert() interface{}
	Select() interface{}
	Update() interface{}
	Delete() interface{}
}

var db *sql.DB
var ctx context.Context

func Init() {
	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

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

	r := &Role{Name: "user"}
	result, err := r.Insert(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
