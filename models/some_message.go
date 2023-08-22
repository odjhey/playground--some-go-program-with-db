package models

import (
	"errors"
	"net/http"
)

type SomeMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func (s *SomeMessage) Bind(r *http.Request) error {
	if s.Message == "" {
		return errors.New("message is required")
	}

	return nil
}

func (s *SomeMessage) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
