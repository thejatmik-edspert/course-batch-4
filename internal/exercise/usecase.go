package exercise

import (
	"context"
	"strconv"
	"strings"
	"time"

	"course/internal/domain"

	"github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

type ExerciseService struct {
	db *gorm.DB
}

func NewExerciseService(dbConn *gorm.DB) *ExerciseService {
	return &ExerciseService{
		db: dbConn,
	}
}

func (es ExerciseService) CreateExercise(c *gin.Context) {
	var exerciseRequest domain.ExerciseRequest
	err := c.ShouldBindJSON(&exerciseRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx := es.db.Begin()
	defer tx.Rollback()
	query := `
	INSERT INTO ` + "`exercises`" + `(title, description)
	VALUES (?, ?)
	`
	result, err := tx.ConnPool.ExecContext(ctx, query, exerciseRequest.Title, exerciseRequest.Description)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "unable to create Exercise",
		})
		return
	}
	tx.Commit()
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(201, gin.H{
		"id":          id,
		"title":       exerciseRequest.Title,
		"description": exerciseRequest.Description,
	})
}

func (es ExerciseService) GetExerciseScoreByID(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = es.db.Where("id = ?", id).Preload("Question").Take(&exercise).Error
	if err != nil || len(exercise.Question) == 0 {
		c.JSON(404, gin.H{
			"message": "exercise not found",
		})
		return
	}

	var score int = 0
	var answers []domain.Answer
	ctx := c.Request.Context()
	userID := int(ctx.Value("user_id").(float64))

	err = es.db.Where("exercise_id = ? AND user_id = ?", exercise.ID, userID).Find(&answers).Error
	if err != nil {
		c.JSON(200, gin.H{
			"score": score,
		})
		return
	}

	// loop on each questions ?
	mapQA := make(map[int]string)
	for _, answer := range answers {
		mapQA[answer.QuestionID] = answer.Answer
	}

	for _, question := range exercise.Question {
		answer := mapQA[question.ID]
		correctAnswer := question.CorrectAnswer
		if strings.EqualFold(answer, correctAnswer) {
			score = score + question.Score
		}
	}

	c.JSON(200, gin.H{
		"score": score,
	})
}

func (es ExerciseService) GetExerciseByID(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid execise id",
		})
		return
	}

	var exercise domain.Exercise
	// err = es.db.Where("id = ?", id).Take(&exercise).Error
	err = es.db.Where("id = ?", id).Preload("Question").Take(&exercise).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "exercise not found",
		})
		return
	}

	c.JSON(200, exercise)
}
