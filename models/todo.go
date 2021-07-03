package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID    uint
	Title string
	Done  bool
}

func CreateTodo(title string) (*Todo, error) {
	todo := Todo{Title: title, Done: false}
	result := db.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func GetTodoById(id uint) (*Todo, error) {
	todo := Todo{ID: id}
	result := db.First(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func GetAllTodo() ([]*Todo, error) {
	todos := []*Todo{}
	db.Find(&todos)
	return todos, nil
}

func DeleteTodo(id uint) error {
	todo := Todo{ID: id}
	result := db.Delete(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateTodo(id uint, title string, done bool) (*Todo, error) {
	todo := Todo{ID: id, Title: title, Done: done}
	result := db.Save(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}
