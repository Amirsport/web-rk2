package api

import "github.com/ValeryBMSTU/web-rk2/internal/entities"

type Usecase interface {
	CreateAnswer() error
	GetQuestions() ([]*entities.Question, error)
	GetScore() (*entities.Score, error)
}
