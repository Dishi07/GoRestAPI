package models

import (
	"golang.org/x/crypto/bcrypt"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (params *LoginParams) EncodePassword() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
