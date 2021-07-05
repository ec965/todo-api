package models

import (
	"context"
	"time"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Role struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (r *Role) Insert(ctx context.Context) (sql.Result, error) {
	result, err := db.ExecContext(ctx, "INSERT INTO roles (name) VALUES ( $1 )", r.Name)
	return result, err
}
