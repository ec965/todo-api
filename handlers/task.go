package handlers

import (
	"net/http"

	"github.com/ec965/todo-api/handlers/validator"
	"github.com/ec965/todo-api/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		sendError(w, err)
		return
	}
	data := struct {
		Title string `form:"title" validate:"required,max=64"`
		Text  string `form:"text" validate:"required,max=1024"`
	}{}
	if errMap, err := validator.IsValid(r, &data); err != nil {
		sendStatus(w, http.StatusBadRequest)
		sendJson(w, errMap)
		return
	}

	task := models.Task{Title: data.Title, Text: data.Text}
	id, err := task.CreateForUser(user.ID)
	if err != nil {
		sendError(w, err)
		return
	}
	task.SelectById(id)
	sendJson(w, task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		sendError(w, err)
		return
	}
	tasks, err := models.SelectTasksForUser(user.ID)
	if err != nil {
		sendError(w, err)
		return
	}
	sendJson(w, tasks)
}

// func UpdateTasks(w http.ResponseWriter, r *http.Request) {
// 	user, err := getUser(r)
// 	if err != nil {
// 		sendError(w, err)
// 		return
// 	}

// }
