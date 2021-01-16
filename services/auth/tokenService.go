package auth

import (
	"crypto/rand"
	"mmbt/models"

	"github.com/go-pg/pg/v10"
)

// TokenService manages auth tokens
type TokenService struct {
	DB *pg.DB
}

// CreateToken creates a token for a user
func (svc *TokenService) CreateToken(user *models.User) (result *models.UserLoginToken, err error) {
	bytes := make([]byte, 64)

	if _, err = rand.Read(bytes); err != nil {
		return nil, err
	}

	allowedLetters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVYXYZ1234567890")
	for i := 0; i < len(bytes); i++ {
		bytes[i] = allowedLetters[int(bytes[i])%len(allowedLetters)]
	}

	tokenStr := string(bytes)

	result = &models.UserLoginToken{
		Token:  tokenStr,
		UserID: user.ID,
	}

	_, err = svc.DB.Model(result).Insert()

	return
}
