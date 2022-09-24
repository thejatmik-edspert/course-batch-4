package domain

import (
	"time"
)

type ExerciseRequest struct {
	Title       string `json:"title"       binding:"required"`
	Description string `json:"description" binding:"required"`
}
type Exercise struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Question    []Question `json:"questions"`
}

type QuestionRequest struct {
	ExerciseID    int    `json:"exercise_id"`
	Body          string `json:"body"           binding:"required"`
	OptionA       string `json:"option_a"       binding:"required"`
	OptionB       string `json:"option_b"       binding:"required"`
	OptionC       string `json:"option_c"       binding:"required"`
	OptionD       string `json:"option_d"       binding:"required"`
	CorrectAnswer string `json:"correct_answer" binding:"required"`
}
type Question struct {
	ID            int       `json:"id"`
	ExerciseID    int       `json:"-"`
	Body          string    `json:"body"`
	OptionA       string    `json:"option_a"`
	OptionB       string    `json:"option_b"`
	OptionC       string    `json:"option_c"`
	OptionD       string    `json:"option_d"`
	CorrectAnswer string    `json:"-"`
	Score         int       `json:"score"`
	CreatorID     int       `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AnswerRequest struct {
	Answer string `json:"answer" binding:"required"`
}
type Answer struct {
	ID         int       `json:"id"`
	ExerciseID int       `json:"exercise_id"`
	QuestionID int       `json:"question_id"`
	UserID     int       `json:"user_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
