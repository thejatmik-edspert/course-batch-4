package main

import (
	"course/internal/answer"
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/question"
	"course/internal/user"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"message": "hello world",
		})
	})
	db := database.CreateConn()

	// user endpoint
	user := user.NewUserService(db)
	r.POST("/register", user.PostRegister)
	r.POST("/login", user.PostLogin)

	// exercise endpoint
	exercise := exercise.NewExerciseService(db)
	r.GET("/exercises/:id", middleware.UseClaims(user), exercise.GetExerciseByID)
	r.GET("/exercises/:id/scores", middleware.UseClaims(user), exercise.GetExerciseScoreByID)
	r.POST("/exercises", middleware.UseClaims(user), exercise.CreateExercise)

	// question endpoint
	question := question.NewQuestionService(db)
	r.POST("/exercises/:id/questions", middleware.UseClaims(user), question.CreateQuestion)

	// answer endpoint
	answer := answer.NewAnswerServie(db)
	r.POST("/exercises/:id/questions/:qid/answer", middleware.UseClaims(user), answer.CreateAnswer)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "1234"
	}
	r.Run(fmt.Sprintf(":%s", appPort))
}
