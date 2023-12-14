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
