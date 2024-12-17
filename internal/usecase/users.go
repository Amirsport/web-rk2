package usecase

import (
	"github.com/Amrisport/web-rk2/internal/entities"

	"github.com/labstack/echo/v4"
)

func (u *Usecase) CreateAnswer(c echo.Context) int {
	var input struct {
		Ans1 int `json:"ans1"`
		Ans2 int `json:"ans2"`
		Ans3 int `json:"ans3"`
		Ans4 int `json:"ans4"`
		Ans5 int `json:"ans5"`
	}

	points := [5]int{3, 3, 2, 3, 2}
	var correct int
	correct = 0
	if points[0] == input.Ans1 {
		correct++
	}
	if points[1] == input.Ans2 {
		correct++
	}
	if points[2] == input.Ans3 {
		correct++
	}
	if points[3] == input.Ans4 {
		correct++
	}
	if points[4] == input.Ans5 {
		correct++
	}

	return correct
}
func (u *Usecase) ListQuestions() ([]*entities.Question, error) {
	questions, err := u.p.GetQuestion()
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (u *Usecase) GetScore() (*entities.Score, error) {
	score, err := u.p.GetScore()
	if err != nil {
		return nil, err
	}
	return score, nil
}
