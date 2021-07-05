package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string `gorm:"uniqueIndex:idx_username"`
	Password  string
	Email     string `gorm:"index:uniqueIndex:idx_email"`
	RoleID    uint
	Role      Role
}

func CreateUserIfNotExist(u User) User {
	existingUser := User{}
	Db.Where("username = ?", u.Username).Find(existingUser)
	if existingUser == (User{}) {
		Db.Create(&u)
		return u
	}
	return existingUser
}
