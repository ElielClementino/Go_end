package schemas

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	title       string
	description string
	due_to      string
}
