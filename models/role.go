package models

import (
	"log"
	"time"
)

const rolesTable = "roles"

type Role struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

func (r *Role) Insert() int {
	rt := table{rolesTable, []string{"name"}}
	var id int
	if err := db.QueryRow(rt.insertStr(), r.Name).Scan(&id); err != nil {
		log.Fatal(err)
	}
	return id
}

func (r *Role) Update() {
	rt := table{rolesTable, []string{"name"}}
	if _, err := db.Exec(rt.updateStr(), r.Name); err != nil {
		log.Fatal(err)
	}
}

func (r *Role) SelectById(id int) {
	rt := table{rolesTable, []string{"id", "created_at", "updated_at", "name"}}
	err := db.QueryRow(rt.selectByStr("id"), id).Scan(r.ID, r.CreatedAt, r.UpdatedAt, r.Name)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Role) SelectByName(name string) {
	rt := table{rolesTable, []string{"id", "created_at", "updated_at", "name"}}
	err := db.QueryRow(rt.selectByStr("name"), name).Scan(r.ID, r.CreatedAt, r.UpdatedAt, r.Name)
	if err != nil {
		log.Fatal(err)
	}
}
