package user

import (
	"errors"
	"fmt"

	"course/internal/domain"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(dbConn *gorm.DB) *UserService {
	return &UserService{
		db: dbConn,
	}
}

func (us UserService) PostRegister(c *gin.Context) {
	var user domain.User
	// check input
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println(len(user.Password), "pass len")
	if len(user.Password) < 6 {
		c.JSON(400, gin.H{
			"message": "password length must be >6",
		})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	err = us.db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "registration failed",
		})
		return
	}

	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error on generating token",
		})
		return
	}
	c.JSON(201, gin.H{
		"token": token,
	})
}

type login struct {
	Email    string
	Password string
}

func (us UserService) PostLogin(c *gin.Context) {
	var userRequest login
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}
	if userRequest.Email == "" || userRequest.Password == "" {
		c.JSON(400, gin.H{
			"message": "wrong email/password",
		})
		return
	}

	var user domain.User
	err = us.db.Where("email = ?", userRequest.Email).Take(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "wrong email/password",
		})
		return
	}
	err = user.ComparePassword(userRequest.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "wrong email/password",
		})
		return
	}
	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error on generating token",
		})
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func (us UserService) DecriptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return domain.SignatureKey, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}
	if !parsedToken.Valid {
		return data, errors.New("invalid token")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
