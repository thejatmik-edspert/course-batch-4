package answer

import (
	"strconv"
	"strings"

	"course/internal/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AnswerService struct {
	db *gorm.DB
}

func NewAnswerServie(dbConn *gorm.DB) *AnswerService {
	return &AnswerService{
		db: dbConn,
	}
}

func (as AnswerService) CreateAnswer(c *gin.Context) {
	paramID := c.Param("id")
	paramQID := c.Param("qid")

	exerciseID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid exercise id",
		})
		return
	}
	var exists bool
	err = as.db.Model(&domain.Exercise{}).Select(
		"COUNT(*) > 0",
	).Where("id = ?", exerciseID).Find(&exists).Error
	if err != nil || !exists {
		c.JSON(404, gin.H{
			"message": "exercise not found",
		})
		return
	}

	questionID, err := strconv.Atoi(paramQID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid question id",
		})
		return
	}
	err = as.db.Model(&domain.Question{}).Select(
		"COUNT(*) > 0",
	).Where("id = ?", questionID).Find(&exists).Error
	if err != nil || !exists {
		c.JSON(404, gin.H{
			"message": "question not found",
		})
		return
	}

	userID := int(c.Request.Context().Value("user_id").(float64))

	var answerRequest domain.AnswerRequest
	err = c.ShouldBindJSON(&answerRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}
	switch strings.ToLower(answerRequest.Answer) {
	case "a", "b", "c", "d": // empty string as 'delete answer'
		answerRequest.Answer = strings.ToLower(answerRequest.Answer)
	default:
		c.JSON(400, gin.H{
			"message": "invalid answer",
		})
		return
	}

	answer := domain.Answer{
		ExerciseID: exerciseID,
		QuestionID: questionID,
		UserID:     userID,
		Answer:     answerRequest.Answer,
	}

	result := as.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_id"}, {Name: "question_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"answer"}),
	}).Create(&answer)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "unable to create Answer",
		})
		return
	}
	c.JSON(201, gin.H{
		"answer_id": answer.ID,
	})
}
