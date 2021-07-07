package models

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const usersTable = "users"

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	RoleId    int64     `json:"roleId"`
}

func hashPassword(pw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password")
	}
	return string(hash)
}

func (u *User) Insert() int {
	ut := table{usersTable, []string{"username", "email", "password", "first_name", "last_name", "role_id"}}

	u.Password = hashPassword(u.Password)
	var id int
	err := db.QueryRow(
		ut.insertStr(),
		u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.RoleId,
	).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (u *User) Update(hashPw bool) error {
	ut := table{usersTable, []string{"username", "email", "password", "first_name", "last_name", "role_id"}}
	if u.ID == 0 {
		return errors.New("User ID is null, cannot update user")
	}
	if hashPw {
		u.Password = hashPassword(u.Password)
	}
	_, err := db.Exec(ut.updateStr(), u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.RoleId, u.ID)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (u *User) SelectById(id int) {
	ut := table{usersTable, []string{"id", "created_at", "updated_at", "username", "email", "password", "first_name", "last_name", "role_id"}}
	err := db.QueryRow(ut.selectByStr("id"), id).Scan(
		u.ID, u.CreatedAt, u.UpdatedAt, u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.RoleId,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (u *User) SelectByUsername(username string) {
	ut := table{usersTable, []string{"id", "created_at", "updated_at", "username", "email", "password", "first_name", "last_name", "role_id"}}
	err := db.QueryRow(ut.selectByStr("username"), username).Scan(
		u.ID, u.CreatedAt, u.UpdatedAt, u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.RoleId,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (u User) CheckUsernameEmail(qUsername string, qEmail string) (bool, bool) {
	rows, err := db.Query("SELECT username, email WHERE username = $1, email = $2", qUsername, qEmail)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	hasUsername := false
	hasEmail := false
	for rows.Next() {
		var (
			username string
			email    string
		)
		if err := rows.Scan(&username, &email); err != nil {
			log.Fatal(err)
		}
		if username == qUsername {
			hasUsername = true
		}
		if email == qEmail {
			hasEmail = true
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return hasUsername, hasEmail

	// u.ID, u.CreatedAt, u.UpdatedAt, u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.RoleId,
}
