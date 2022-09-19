package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var signatureKey = []byte("myPrivateKey")

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	NoHp      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreatePassword(password string) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(encryptedPassword)
	return nil
}

func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	})

	tokenString, err := token.SignedString(signatureKey)
	return tokenString, err
}

func (u *User) DecryptJwt(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return signatureKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return map[string]interface{}{}, errors.New("auth invalid")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}

func (u User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
