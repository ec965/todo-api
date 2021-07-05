package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"uniqueIndex:idx_name"`
}

func FindRoleByName(n string) Role {
	role := Role{}
	Db.Where("name = ?", n).Find(&role)
	return role
}

func CreateRoleIfNotExist(n string) Role {
	r := FindRoleByName(n)
	if r == (Role{}) {
		r = Role{Name: n}
		Db.Create(&r)
	}
	return r
}
