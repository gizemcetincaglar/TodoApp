package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/gizemcetincaglar/todo_app/models"
)

var stepList []models.Step

func GetStepsByTodoID(todoID string) []models.Step {
	var steps []models.Step
	for _, step := range stepList {
		if step.TodoID == todoID && step.DeletedAt == nil {
			steps = append(steps, step)
		}
	}
	return steps
}

func CreateStep(todoID string, content string) models.Step {
	step := models.Step{
		ID:        uuid.New().String(),
		TodoID:    todoID,
		Content:   content,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	stepList = append(stepList, step)
	UpdateCompletionPercentage(todoID)
	return step
}

func MarkStepCompleted(stepID string) bool {
	for i, step := range stepList {
		if step.ID == stepID && step.DeletedAt == nil {
			stepList[i].Completed = true
			stepList[i].UpdatedAt = time.Now()
			UpdateCompletionPercentage(step.TodoID)
			return true
		}
	}
	return false
}

// Adımı soft-delete ile sil
func DeleteStep(stepID string) bool {
	for i, step := range stepList {
		if step.ID == stepID && step.DeletedAt == nil {
			now := time.Now()
			stepList[i].DeletedAt = &now
			UpdateCompletionPercentage(step.TodoID)
			return true
		}
	}
	return false
}

// Listenin tamamlanma yüzdesini güncelle
func UpdateCompletionPercentage(todoID string) {
	var total, completed int
	for _, step := range stepList {
		if step.TodoID == todoID && step.DeletedAt == nil {
			total++
			if step.Completed {
				completed++
			}
		}
	}
	for i, todo := range todoList {
		if todo.ID == todoID {
			percent := 0.0
			if total > 0 {
				percent = (float64(completed) / float64(total)) * 100
			}
			todoList[i].Completion = percent
			todoList[i].UpdatedAt = time.Now()
		}
	}
}
