package question

import (
	"strconv"
	"strings"

	"course/internal/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionService struct {
	db *gorm.DB
}

func NewQuestionService(dbConn *gorm.DB) *QuestionService {
	return &QuestionService{
		db: dbConn,
	}
}

// create question from exercise ID
func (qs QuestionService) CreateQuestion(c *gin.Context) {
	paramID := c.Param("id")
	userID := int(c.Request.Context().Value("user_id").(float64))
	exerciseID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid exercise id",
		})
		return
	}

	var exists bool
	err = qs.db.Model(&domain.Exercise{}).Select(
		"COUNT(*) > 0",
	).Where("id = ?", exerciseID).Find(&exists).Error
	if err != nil || !exists {
		c.JSON(404, gin.H{
			"message": "exercise not found",
		})
		return
	}

	var questionRequest domain.QuestionRequest
	err = c.ShouldBindJSON(&questionRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	switch strings.ToLower(questionRequest.CorrectAnswer) {
	case "a", "b", "c", "d":
		questionRequest.CorrectAnswer = strings.ToLower(questionRequest.CorrectAnswer)
	default:
		c.JSON(400, gin.H{
			"message": "invalid correct answer",
		})
		return
	}

	question := domain.Question{
		ExerciseID:    exerciseID,
		CreatorID:     userID,
		Body:          questionRequest.Body,
		OptionA:       questionRequest.OptionA,
		OptionB:       questionRequest.OptionB,
		OptionC:       questionRequest.OptionC,
		OptionD:       questionRequest.OptionD,
		CorrectAnswer: questionRequest.CorrectAnswer,
		Score:         5,
	}

	result := qs.db.Create(&question)
	if result.Error != nil || question.ID == 0 {
		c.JSON(400, gin.H{
			"message": "unable to create Question",
		})
		return
	}
	c.JSON(201, gin.H{
		"question_id": question.ID,
	})
}
