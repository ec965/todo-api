package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_name"`
}

type roles struct {
	Admin Role
	User  Role
}

func FindRoleByName(n string) Role {
	role := Role{}
	Db.Where("name = ?", n).Find(&role)
	return role
}

func createRoleIfNotExist(n string) Role {
	r := FindRoleByName(n)
	if r == (Role{}) {
		r = Role{Name: n}
		Db.Create(&r)
	}
	return r
}

func createRoles() roles {
	admin := createRoleIfNotExist("admin")
	user := createRoleIfNotExist("user")
	return roles{Admin: admin, User: user}
}
