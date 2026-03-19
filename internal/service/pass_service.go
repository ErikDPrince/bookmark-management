package service

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const (
	charset        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	passwordLength = 10
)

type passwordService struct {
}

// Password  interface represents the password services interface
//go:generate mockery --name Password --filename pass_service.go

type Password interface {
	GeneratePassword() (string, error)
}

// NewPassWordService creates a new password service
func NewPassWordService() Password {
	return &passwordService{}
}

func (s *passwordService) GeneratePassword() (string, error) {
	var strBuilder bytes.Buffer

	for i := 0; i < passwordLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		strBuilder.WriteByte(charset[randomIndex.Int64()])
	}
	return strBuilder.String(), nil

}
