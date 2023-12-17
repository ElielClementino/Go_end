package handler

import "fmt"

func errParamIsRequired(name, category string) error {
	return fmt.Errorf("field: %s, type: %s is required", name, category)
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueTo       string `json:"dueTo"`
}

func (t *CreateTaskRequest) Validate() error {
	if t.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if t.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if t.DueTo == "" {
		return errParamIsRequired("dueTo", "string")
	}
	return nil
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueTo       string `json:"dueTo"`
}

func (t *UpdateTaskRequest) Validate() error {
	if t.Title != "" || t.Description != "" || t.DueTo != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
