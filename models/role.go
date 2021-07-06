package models

import (
	"context"
	"fmt"
	"time"
)

type Role struct {
	ID        int64     `json:"id" db:"auto"`
	CreatedAt time.Time `json:"createdAt" db:"auto"`
	UpdatedAt time.Time `json:"updatedAt" db:"auto"`
	Name      string    `json:"name"`
}

func (r *Role) Insert() error {
	_, err := Insert(r)
	return err
}

func (r *Role) InsertContext(ctx context.Context) error {
	_, err := InsertContext(ctx, r)
	return err
}

func (r *Role) SelectById(id int64) error {
	rows, err := db.Query("SELECT * FROM roles WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println(rows.Columns())
		fmt.Println(rows.ColumnTypes())
		var one interface{}
		var two interface{}
		var three interface{}
		var four interface{}
		rows.Scan(&one, &two, &three, &four)
		fmt.Println(one, two, three, four)
	}
	err = rows.Err()
	return err
}
