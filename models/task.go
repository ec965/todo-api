package models

import (
	"context"
	"errors"
	"time"
)

const taskTable = "tasks"

type Task struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Complete  bool      `json:"complete"`
}

func (t *Task) CreateForUser(userId int64) (int64, error) {
	if t.Title == "" || t.Text == "" {
		return 0, errors.New("title or text is missing from task")
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	var taskId int64
	err = tx.QueryRow("INSERT INTO tasks (title, text) VALUES ( $1, $2 ) RETURNING id", t.Title, t.Text).Scan(&taskId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec("INSERT INTO owners (user_id, task_id) VALUES ( $1, $2 )", userId, taskId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return taskId, nil
}

func (t *Task) SelectById(id int64) error {
	return db.QueryRow("SELECT id, created_at, updated_at, title, text, complete FROM "+taskTable+" WHERE id = $1", id).Scan(
		&t.ID, &t.CreatedAt, &t.UpdatedAt, &t.Title, &t.Text, &t.Complete,
	)
}

func SelectTasksForUser(userId int64) ([]Task, error) {
	rows, err := db.Query("SELECT id, created_at, updated_at, title, text, complete FROM "+taskTable+" WHERE id IN (SELECT task_id FROM owners WHERE user_id = $1)", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []Task

	for rows.Next() {
		t := Task{}
		if err := rows.Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt, &t.Title, &t.Text, &t.Complete); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
