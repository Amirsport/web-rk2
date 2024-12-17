package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetQuestions(e echo.Context) error {
	questions, err := s.uc.ListQuestions()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, questions)
}

func (s *Server) CreateAnswer(e echo.Context) error {
	return s.uc.CreateAnswer()
}
