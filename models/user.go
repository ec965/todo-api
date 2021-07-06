package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id" db:"auto"`
	CreatedAt time.Time `json:"createdAt" db:"auto"`
	UpdatedAt time.Time `json:"updatedAt" db:"auto"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	RoleId    int       `json:"roleId"`
}

func hashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash), err
}

func (u *User) beforeInsert() {
	var err error
	u.Password, err = hashPassword(u.Password)
	if err != nil {
		log.Fatal("Failed to hash password")
	}
}

func (u *User) InsertContext(ctx context.Context) (int64, error) {
	u.beforeInsert()
	id, err := InsertContext(ctx, u)
	return id, err
}

func (u *User) Insert(ctx context.Context) (int64, error) {
	u.beforeInsert()
	id, err := Insert(u)
	return id, err
}

func (u *User) SelectById(id int64) error {
	rows, err := db.Query("SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println(rows.Columns)
		fmt.Println(rows.ColumnTypes)
	}
	err = rows.Err()
	return err
}

// func (u *User) Update(ctx context.Context) (sql.Result, error){
// 	result, err := db.ExecContext(
// 		ctx,
// 		"UPDATE users (username, email, password)"
// 	)
// }

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
