package domain

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var SignatureKey = []byte("leleyeye")

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"     binding:"required"`
	Email     string    `json:"email"    binding:"required,email"`
	Password  string    `json:"password" binding:"required"`
	NoHP      string    `json:"no_hp"    binding:"required,numeric"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (u *User) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iss":     "lele",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SignatureKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *User) ComparePassword(input string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(u.Password), []byte(input),
	)
}
