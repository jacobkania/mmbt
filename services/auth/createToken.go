package auth

import (
	"crypto/rand"
	"mmbt/db"
	"mmbt/models"
)

// CreateToken creates a token for a user
func CreateToken(user *models.User) (token *models.UserLoginToken, err error) {
	bytes := make([]byte, 64)

	if _, err = rand.Read(bytes); err != nil {
		return nil, err
	}

	allowedLetters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVYXYZ1234567890")
	for i := 0; i < len(bytes); i++ {
		bytes[i] = allowedLetters[int(bytes[i])%len(allowedLetters)]
	}

	tokenStr := string(bytes)

	newToken := &models.UserLoginToken{
		Token:  tokenStr,
		UserID: user.ID,
	}

	_, err = db.DB.Model(newToken).Insert()

	return newToken, err
}
