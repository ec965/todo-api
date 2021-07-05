package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ec965/todo-api/config"
)

var Db *gorm.DB

func resetDB(dst ...interface{}) {
	Db.Migrator().DropTable(dst...)
	Db.AutoMigrate(dst...)
}

func Init() {
	var err error
	Db, err = gorm.Open(sqlite.Open(config.DbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	resetDB(&User{}, &Role{})

	// create default roles
	roles := createRoles()
	// create admin user
	CreateUserIfNotExist(User{
		FirstName: "admin",
		LastName:  "user",
		Username:  config.AdminUser,
		Password:  config.AdminPass,
		Email:     "adminEmail",
		Role:      roles.Admin,
	})
}
