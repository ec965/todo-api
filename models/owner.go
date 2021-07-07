package models

import "time"

const ownerTable = "owners"

type Owner struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId    int64     `json:"user_id"`
	TaskId    int64     `json:"task_id"`
}
