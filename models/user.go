package models

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	RoleId int `json:"roleId"`
}

func hashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash), err
}

// type User struct {
// 	gorm.Model
// 	FirstName string
// 	LastName  string
// 	Username  string `gorm:"uniqueIndex:idx_username"`
// 	Password  string
// 	Email     string `gorm:"index:uniqueIndex:idx_email"`
// 	RoleID    uint
// 	Role      Role
// }

// func CreateUserIfNotExist(u User) User {
// 	existingUser := User{}
// 	Db.Where("username = ?", u.Username).Find(existingUser)
// 	if existingUser == (User{}) {
// 		Db.Create(&u)
// 		return u
// 	}
// 	return existingUser
// }


// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.Password, err = hashPassword(u.Password)
// 	return
// }

// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
// 	if tx.Statement.Changed("Password") {
// 		u.Password, err = hashPassword(u.Password)
// 	}
// 	return
// }
