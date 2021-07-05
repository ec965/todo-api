// package models

// import (
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

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

// func hashPassword(pw string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
// 	return string(hash), err
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
