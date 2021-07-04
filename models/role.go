package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"role" gorm:"uniqueIndex:idx_name"`
}

func CreateRole(r *Role) {
	result := db.Create(r)
	if result.Error != nil {
		panic(result.Error)
	}
}

func DeleteRole(r *Role) {
	result := db.Delete(r)
	if result.Error != nil {
		panic(result.Error)
	}
}

func FindRoleByName(n string) Role {
	role := Role{}
	result := db.Where("name = ?", n).Find(&role)
	if result.Error != nil {
		panic(result.Error)
	}
	return role
}

func CreateRoleIfNotExist(n string) Role {
	r := FindRoleByName(n)
	if r == (Role{}) {
		r = Role{Name: n}
		CreateRole(&r)
	}
	return r
}
