package repository

import "task-manager/models"

var tasks []models.Task
var nextID = 1

func CreateTask(t models.Task) models.Task{
	t.ID = nextID
	nextID++
	tasks = append(tasks, t)
	return t
}



func GetAllTasks() []models.Task{
	return tasks
}

func GetTaskByID(id int) *models.Task{
	for i:= range tasks {
		if tasks[i].ID == id{
			return &tasks[i]
		}
	}
	return nil
}

func UpdateTask(id int,updated models.Task) *models.Task{
	for i:= range tasks{
		if tasks[i].ID == id{
			tasks[i].Title = updated.Title
			tasks[i].Description = updated.Description
			tasks[i].Completed = updated.Completed
			return &tasks[i]
		}
	}
	return nil
}
func DeleteTask(id int) bool{
	for i:= range tasks{
		if tasks[i].ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
