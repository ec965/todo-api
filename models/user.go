package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username" gorm:"uniqueIndex:idx_username"`
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"index:uniqueIndex:idx_email"`
	RoleID    uint
	Role      Role `json:"role"`
}

func CreateUser(u *User) {
	result := db.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func UpdateUser(u *User) {
	result := db.Save(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func DeleteUser(u *User) {
	result := db.Delete(u)
	if result.Error != nil {
		panic(result.Error)
	}
}

func CreateUserIfNotExist(u User) User{
	existingUser := User{}
	db.Where("username = ?", u.Username).Find(existingUser)
	if existingUser == (User{}) {
		CreateUser(&u);
		return u;
	}
	return existingUser;
}